package ofp4sw

import (
	"github.com/hkwi/gopenflow/ofp4"
	"time"
)

type packetDrop struct{}

func (_ packetDrop) Error() string {
	return "drop packet"
}

const BASE_INTERVAL = 2.0

type meter struct {
	flagPkts    bool
	flagBurst   bool
	flagStats   bool
	commands    chan func()
	created     time.Time
	packetCount uint64
	byteCount   uint64
	bands       []band
	highestBand band

	meterTime time.Time
	rate      float64
}

func (m *meter) process(data *frame) error {
	ch := make(chan error)
	m.commands <- func() {
		ch <- func() error {
			length := 0
			if buf, err := data.data(); err != nil {
				return err
			} else {
				length = len(buf)
			}

			now := time.Now()
			meterInterval := float64(now.Sub(m.meterTime)) / float64(time.Second)

			if m.flagStats {
				m.packetCount++
				m.byteCount += uint64(length)
			}

			inc := float64(length*8) / 1000.0 // kilobits
			if m.flagPkts {
				inc = 1
			}

			rate := (m.rate*BASE_INTERVAL + inc) / (BASE_INTERVAL + meterInterval)

			if m.flagBurst {
				for _, bi := range m.bands {
					drain := meterInterval * float64(bi.getRate()) * 8 / 1000.0
					if m.flagPkts {
						drain = meterInterval * float64(bi.getRate())
					}
					switch b := bi.(type) {
					case *bandDrop:
						b.bucket += inc - drain
						if b.bucket < 0 {
							b.bucket = 0
						}
					case *bandDscpRemark:
						b.bucket += inc - drain
						if b.bucket < 0 {
							b.bucket = 0
						}
					case *bandExperimenter:
						b.bucket += inc - drain
						if b.bucket < 0 {
							b.bucket = 0
						}
					default:
						panic("Unexpected band")
					}
				}
				for _, bi := range m.bands {
					switch b := bi.(type) {
					case *bandDrop:
						if b.bucket > float64(b.burstSize) {
							if m.flagStats {
								b.packetCount++
								b.byteCount += uint64(length)
							}
							return &packetDrop{}
						}
					case *bandDscpRemark:
						if b.bucket > float64(b.burstSize) {
							// do dscp remark
							m.rate = rate
							if m.flagStats {
								b.packetCount++
								b.byteCount += uint64(length)
							}
							return nil
						}
					case *bandExperimenter:
						if b.bucket > float64(b.burstSize) {
							// do nothing
							m.rate = rate
							if m.flagStats {
								b.packetCount++
								b.byteCount += uint64(length)
							}
							return nil
						}
					default:
						panic("Unexpected band")
					}
				}
			}
			if m.highestBand != nil && rate > float64(m.highestBand.getRate()) {
				switch b := m.highestBand.(type) {
				case *bandDrop:
					if m.flagStats {
						b.packetCount++
						b.byteCount += uint64(length)
					}
					return &packetDrop{}
				case *bandDscpRemark:
					// do dscp remark
					m.rate = rate
					if m.flagStats {
						b.packetCount++
						b.byteCount += uint64(length)
					}
					return nil
				case *bandExperimenter:
					// do nothing
					m.rate = rate
					if m.flagStats {
						b.packetCount++
						b.byteCount += uint64(length)
					}
					return nil
				default:
					panic("Unexpected band")
				}
			}
			m.meterTime = now
			m.rate = rate
			return nil
		}()
		close(ch)
	}
	return <-ch
}

type band interface {
	getRate() uint32
}

type bandCommon struct {
	rate      uint32
	burstSize uint32

	bucket      float64
	packetCount uint64
	byteCount   uint64
}

func (b bandCommon) getRate() uint32 {
	return b.rate
}

type bandDrop struct {
	bandCommon
}

type bandDscpRemark struct {
	bandCommon
	precLevel uint8
}

type bandExperimenter struct {
	bandCommon
	experimenter uint32
	data         []byte
}

func newMeter(msg ofp4.MeterMod) *meter {
	var highestBand band
	var bands []band
	for _, bi := range msg.Bands {
		var b band
		switch msgBand := bi.(type) {
		case *ofp4.MeterBandDrop:
			b = &bandDrop{
				bandCommon: bandCommon{
					rate:      msgBand.Rate,
					burstSize: msgBand.BurstSize,
				},
			}
		case *ofp4.MeterBandDscpRemark:
			b = &bandDscpRemark{
				bandCommon: bandCommon{
					rate:      msgBand.Rate,
					burstSize: msgBand.BurstSize,
				},
				precLevel: msgBand.PrecLevel,
			}
		case *ofp4.MeterBandExperimenter:
			b = &bandExperimenter{
				bandCommon: bandCommon{
					rate:      msgBand.Rate,
					burstSize: msgBand.BurstSize,
				},
				experimenter: msgBand.Experimenter,
				data:         msgBand.Data,
			}
		}
		if b != nil {
			bands = append(bands, b)
			if highestBand == nil || highestBand.getRate() < b.getRate() {
				highestBand = b
			}
		}
	}

	m := meter{
		commands:    make(chan func(), 16),
		created:     time.Now(),
		bands:       bands,
		highestBand: highestBand,
	}
	if msg.Flags&ofp4.OFPMF_PKTPS != 0 {
		m.flagPkts = true
	}
	if msg.Flags&ofp4.OFPMF_BURST != 0 {
		m.flagBurst = true
	}
	if msg.Flags&ofp4.OFPMF_STATS != 0 {
		m.flagStats = true
	}
	go func() {
		for cmd := range m.commands {
			if cmd != nil {
				cmd()
			} else {
				break
			}
		}
	}()
	return &m
}