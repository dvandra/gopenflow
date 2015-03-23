package ofp4

const (
	OFPFC_ADD = iota
	OFPFC_MODIFY
	OFPFC_MODIFY_STRICT
	OFPFC_DELETE
	OFPFC_DELETE_STRICT
)

const (
	OFPT_HELLO = iota
	OFPT_ERROR
	OFPT_ECHO_REQUEST
	OFPT_ECHO_REPLY
	OFPT_EXPERIMENTER

	OFPT_FEATURES_REQUEST
	OFPT_FEATURES_REPLY
	OFPT_GET_CONFIG_REQUEST
	OFPT_GET_CONFIG_REPLY
	OFPT_SET_CONFIG

	OFPT_PACKET_IN
	OFPT_FLOW_REMOVED
	OFPT_PORT_STATUS

	OFPT_PACKET_OUT
	OFPT_FLOW_MOD
	OFPT_GROUP_MOD
	OFPT_PORT_MOD
	OFPT_TABLE_MOD

	OFPT_MULTIPART_REQUEST
	OFPT_MULTIPART_REPLY

	OFPT_BARRIER_REQUEST
	OFPT_BARRIER_REPLY

	OFPT_QUEUE_GET_CONFIG_REQUEST
	OFPT_QUEUE_GET_CONFIG_REPLY

	OFPT_ROLE_REQUEST
	OFPT_ROLE_REPLY

	OFPT_GET_ASYNC_REQUEST
	OFPT_GET_ASYNC_REPLY
	OFPT_SET_ASYNC

	OFPT_METER_MOD
)

const (
	OFPP_MAX        = 0xffffff00
	OFPP_IN_PORT    = 0xfffffff8
	OFPP_TABLE      = 0xfffffff9
	OFPP_NORMAL     = 0xfffffffa
	OFPP_FLOOD      = 0xfffffffb
	OFPP_ALL        = 0xfffffffc
	OFPP_CONTROLLER = 0xfffffffd
	OFPP_LOCAL      = 0xfffffffe
	OFPP_ANY        = 0xffffffff
)

const (
	OFPPC_PORT_DOWN = 1 << iota
	_
	OFPPC_NO_RECV
	_
	_
	OFPPC_NO_FWD
	OFPPC_NO_PACKET_IN
)

const (
	OFPPS_LINK_DOWN = 1 << iota
	OFPPS_BLOCKED
	OFPPS_LIVE
)

const (
	OFPPF_10MB_HD = 1 << iota
	OFPPF_10MB_FD
	OFPPF_100MB_HD
	OFPPF_100MB_FD
	OFPPF_1GB_HD
	OFPPF_1GB_FD
	OFPPF_10GB_FD
	OFPPF_40GB_FD
	OFPPF_100GB_FD
	OFPPF_1TB_FD
	OFPPF_OTHER

	OFPPF_COPPER
	OFPPF_FIBER
	OFPPF_AUTONEG
	OFPPF_PAUSE
	OFPPF_PAUSE_ASYM
)

const (
	OFPQT_MIN_RATE     = 1
	OFPQT_MAX_RATE     = 2
	OFPQT_EXPERIMENTER = 0xffff
)

const (
	OFPMT_STANDARD = iota
	OFPMT_OXM
)

const (
	OFPXMC_NXM_0          = 0x0000
	OFPXMC_NXM_1          = 0x0001
	OFPXMC_OPENFLOW_BASIC = 0x8000
	OFPXMC_EXPERIMENTER   = 0xffff
)

const (
	OFPXMT_OFB_IN_PORT = iota
	OFPXMT_OFB_IN_PHY_PORT
	OFPXMT_OFB_METADATA
	OFPXMT_OFB_ETH_DST
	OFPXMT_OFB_ETH_SRC
	OFPXMT_OFB_ETH_TYPE
	OFPXMT_OFB_VLAN_VID
	OFPXMT_OFB_VLAN_PCP
	OFPXMT_OFB_IP_DSCP
	OFPXMT_OFB_IP_ECN
	OFPXMT_OFB_IP_PROTO
	OFPXMT_OFB_IPV4_SRC
	OFPXMT_OFB_IPV4_DST
	OFPXMT_OFB_TCP_SRC
	OFPXMT_OFB_TCP_DST
	OFPXMT_OFB_UDP_SRC
	OFPXMT_OFB_UDP_DST
	OFPXMT_OFB_SCTP_SRC
	OFPXMT_OFB_SCTP_DST
	OFPXMT_OFB_ICMPV4_TYPE
	OFPXMT_OFB_ICMPV4_CODE
	OFPXMT_OFB_ARP_OP
	OFPXMT_OFB_ARP_SPA
	OFPXMT_OFB_ARP_TPA
	OFPXMT_OFB_ARP_SHA
	OFPXMT_OFB_ARP_THA
	OFPXMT_OFB_IPV6_SRC
	OFPXMT_OFB_IPV6_DST
	OFPXMT_OFB_IPV6_FLABEL
	OFPXMT_OFB_ICMPV6_TYPE
	OFPXMT_OFB_ICMPV6_CODE
	OFPXMT_OFB_IPV6_ND_TARGET
	OFPXMT_OFB_IPV6_ND_SLL
	OFPXMT_OFB_IPV6_ND_TLL
	OFPXMT_OFB_MPLS_LABEL
	OFPXMT_OFB_MPLS_TC
	OFPXMT_OFB_MPLS_BOS
	OFPXMT_OFB_PBB_ISID
	OFPXMT_OFB_TUNNEL_ID
	OFPXMT_OFB_IPV6_EXTHDR
)

type OxmHeader uint32

const (
	OXM_CLASS_SHIFT   = 16
	OXM_FIELD_SHIFT   = 9
	OXM_HASMASK_SHIFT = 8
	OXM_TYPE_MASK     = 0xFFFFFE00
	OXM_CLASS_MASK    = 0xFFFF0000
	OXM_FIELD_MASK    = 0x0000FE00
	OXM_HASMASK_MASK  = 0x00000100
	OXM_LENGTH_MASK   = 0x000000FF
)

func (self OxmHeader) Type() uint32 {
	return uint32(self & OXM_TYPE_MASK)
}

func (self OxmHeader) Class() uint16 {
	return uint16((self & OXM_CLASS_MASK) >> OXM_CLASS_SHIFT)
}

func (self OxmHeader) Field() uint8 {
	return uint8((self & OXM_FIELD_MASK) >> OXM_FIELD_SHIFT)
}

func (self OxmHeader) HasMask() bool {
	return (self & OXM_HASMASK_MASK) != 0
}

// Length of OXM payload
func (self OxmHeader) Length() int {
	return int(self & OXM_LENGTH_MASK)
}

func (self *OxmHeader) SetLength(length int) {
	*self = (*self &^ OXM_LENGTH_MASK) | OxmHeader(length&OXM_LENGTH_MASK)
}

func (self *OxmHeader) SetMask(mask bool) {
	if mask {
		*self |= OXM_HASMASK_MASK
	} else {
		*self &^= OXM_HASMASK_MASK
	}
}

/*
OXM_OF_ constnts are defined as combinations of OFPXMC_OPENFLOW_BASIC and their field.
Following constants does not have oxm_hasmask and length.

Following constants are usually used as match key.
Experimenter oxm may have variable sized length or mask, and experimenter
match key would be mixed with basic oxm key.
So match key would better not have length and mask.
This is because these constants does not have length and mask included.
*/
const (
	OXM_OF_IN_PORT        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IN_PORT<<OXM_FIELD_SHIFT
	OXM_OF_IN_PHY_PORT    = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IN_PHY_PORT<<OXM_FIELD_SHIFT
	OXM_OF_METADATA       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_METADATA<<OXM_FIELD_SHIFT
	OXM_OF_ETH_DST        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ETH_DST<<OXM_FIELD_SHIFT
	OXM_OF_ETH_SRC        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ETH_SRC<<OXM_FIELD_SHIFT
	OXM_OF_ETH_TYPE       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ETH_TYPE<<OXM_FIELD_SHIFT
	OXM_OF_VLAN_VID       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_VLAN_VID<<OXM_FIELD_SHIFT
	OXM_OF_VLAN_PCP       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_VLAN_PCP<<OXM_FIELD_SHIFT
	OXM_OF_IP_DSCP        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IP_DSCP<<OXM_FIELD_SHIFT
	OXM_OF_IP_ECN         = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IP_ECN<<OXM_FIELD_SHIFT
	OXM_OF_IP_PROTO       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IP_PROTO<<OXM_FIELD_SHIFT
	OXM_OF_IPV4_SRC       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IPV4_SRC<<OXM_FIELD_SHIFT
	OXM_OF_IPV4_DST       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IPV4_DST<<OXM_FIELD_SHIFT
	OXM_OF_TCP_SRC        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_TCP_SRC<<OXM_FIELD_SHIFT
	OXM_OF_TCP_DST        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_TCP_DST<<OXM_FIELD_SHIFT
	OXM_OF_UDP_SRC        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_UDP_SRC<<OXM_FIELD_SHIFT
	OXM_OF_UDP_DST        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_UDP_DST<<OXM_FIELD_SHIFT
	OXM_OF_SCTP_SRC       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_SCTP_SRC<<OXM_FIELD_SHIFT
	OXM_OF_SCTP_DST       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_SCTP_DST<<OXM_FIELD_SHIFT
	OXM_OF_ICMPV4_TYPE    = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ICMPV4_TYPE<<OXM_FIELD_SHIFT
	OXM_OF_ICMPV4_CODE    = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ICMPV4_CODE<<OXM_FIELD_SHIFT
	OXM_OF_ARP_OP         = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ARP_OP<<OXM_FIELD_SHIFT
	OXM_OF_ARP_SPA        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ARP_SPA<<OXM_FIELD_SHIFT
	OXM_OF_ARP_TPA        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ARP_TPA<<OXM_FIELD_SHIFT
	OXM_OF_ARP_SHA        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ARP_SHA<<OXM_FIELD_SHIFT
	OXM_OF_ARP_THA        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ARP_THA<<OXM_FIELD_SHIFT
	OXM_OF_IPV6_SRC       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IPV6_SRC<<OXM_FIELD_SHIFT
	OXM_OF_IPV6_DST       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IPV6_DST<<OXM_FIELD_SHIFT
	OXM_OF_IPV6_FLABEL    = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IPV6_FLABEL<<OXM_FIELD_SHIFT
	OXM_OF_ICMPV6_TYPE    = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ICMPV6_TYPE<<OXM_FIELD_SHIFT
	OXM_OF_ICMPV6_CODE    = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_ICMPV6_CODE<<OXM_FIELD_SHIFT
	OXM_OF_IPV6_ND_TARGET = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IPV6_ND_TARGET<<OXM_FIELD_SHIFT
	OXM_OF_IPV6_ND_SLL    = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IPV6_ND_SLL<<OXM_FIELD_SHIFT
	OXM_OF_IPV6_ND_TLL    = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IPV6_ND_TLL<<OXM_FIELD_SHIFT
	OXM_OF_MPLS_LABEL     = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_MPLS_LABEL<<OXM_FIELD_SHIFT
	OXM_OF_MPLS_TC        = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_MPLS_TC<<OXM_FIELD_SHIFT
	OXM_OF_MPLS_BOS       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_MPLS_BOS<<OXM_FIELD_SHIFT
	OXM_OF_PBB_ISID       = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_PBB_ISID<<OXM_FIELD_SHIFT
	OXM_OF_TUNNEL_ID      = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_TUNNEL_ID<<OXM_FIELD_SHIFT
	OXM_OF_IPV6_EXTHDR    = OFPXMC_OPENFLOW_BASIC<<OXM_CLASS_SHIFT | OFPXMT_OFB_IPV6_EXTHDR<<OXM_FIELD_SHIFT
)

func OxmOfDefs(oxm uint32) (length int, mayMask bool) {
	switch OxmHeader(oxm).Type() {
	default:
		return 0, false
	case OXM_OF_IN_PORT:
		return 4, false
	case OXM_OF_IN_PHY_PORT:
		return 4, false
	case OXM_OF_METADATA:
		return 8, true
	case OXM_OF_ETH_DST:
		return 6, true
	case OXM_OF_ETH_SRC:
		return 6, true
	case OXM_OF_ETH_TYPE:
		return 2, false
	case OXM_OF_VLAN_VID:
		return 2, true
	case OXM_OF_VLAN_PCP:
		return 1, false
	case OXM_OF_IP_DSCP:
		return 1, false
	case OXM_OF_IP_ECN:
		return 1, false
	case OXM_OF_IP_PROTO:
		return 1, false
	case OXM_OF_IPV4_SRC:
		return 4, true
	case OXM_OF_IPV4_DST:
		return 4, true
	case OXM_OF_TCP_SRC:
		return 2, false
	case OXM_OF_TCP_DST:
		return 2, false
	case OXM_OF_UDP_SRC:
		return 2, false
	case OXM_OF_UDP_DST:
		return 2, false
	case OXM_OF_SCTP_SRC:
		return 2, false
	case OXM_OF_SCTP_DST:
		return 2, false
	case OXM_OF_ICMPV4_TYPE:
		return 1, false
	case OXM_OF_ICMPV4_CODE:
		return 1, false
	case OXM_OF_ARP_OP:
		return 2, false
	case OXM_OF_ARP_SPA:
		return 4, true
	case OXM_OF_ARP_TPA:
		return 4, true
	case OXM_OF_ARP_SHA:
		return 6, true
	case OXM_OF_ARP_THA:
		return 6, true
	case OXM_OF_IPV6_SRC:
		return 16, true
	case OXM_OF_IPV6_DST:
		return 16, true
	case OXM_OF_IPV6_FLABEL:
		return 4, true
	case OXM_OF_ICMPV6_TYPE:
		return 1, false
	case OXM_OF_ICMPV6_CODE:
		return 1, false
	case OXM_OF_IPV6_ND_TARGET:
		return 16, false
	case OXM_OF_IPV6_ND_SLL:
		return 6, false
	case OXM_OF_IPV6_ND_TLL:
		return 6, false
	case OXM_OF_MPLS_LABEL:
		return 4, true
	case OXM_OF_MPLS_TC:
		return 1, false
	case OXM_OF_MPLS_BOS:
		return 1, false
	case OXM_OF_PBB_ISID:
		return 3, true
	case OXM_OF_TUNNEL_ID:
		return 8, true
	case OXM_OF_IPV6_EXTHDR:
		return 2, true
	}
}

const (
	OFPVID_PRESENT = 0x1000
	OFPVID_NONE    = 0x0000
)

const (
	OFPIEH_NONEXT = 1 << iota
	OFPIEH_ESP
	OFPIEH_AUTH
	OFPIEH_DEST
	OFPIEH_FRAG
	OFPIEH_ROUTER
	OFPIEH_HOP
	OFPIEH_UNREP
	OFPIEH_UNSEQ
)

const (
	_ = iota
	OFPIT_GOTO_TABLE
	OFPIT_WRITE_METADATA
	OFPIT_WRITE_ACTIONS
	OFPIT_APPLY_ACTIONS
	OFPIT_CLEAR_ACTIONS
	OFPIT_METER
	OFPIT_EXPERIMENTER = 0xffff
)

const (
	OFPAT_OUTPUT       = 0
	OFPAT_COPY_TTL_OUT = 11
	OFPAT_COPY_TTL_IN  = 12
	OFPAT_SET_MPLS_TTL = 15
	OFPAT_DEC_MPLS_TTL = 16
	OFPAT_PUSH_VLAN    = 17
	OFPAT_POP_VLAN     = 18
	OFPAT_PUSH_MPLS    = 19
	OFPAT_POP_MPLS     = 20
	OFPAT_SET_QUEUE    = 21
	OFPAT_GROUP        = 22
	OFPAT_SET_NW_TTL   = 23
	OFPAT_DEC_NW_TTL   = 24
	OFPAT_SET_FIELD    = 25
	OFPAT_PUSH_PBB     = 26
	OFPAT_POP_PBB      = 27
	OFPAT_EXPERIMENTER = 0xffff
)

const (
	OFPCML_MAX       = 0xffe5
	OFPCML_NO_BUFFER = 0xffff
)

const (
	OFPC_FLOW_STATS = 1 << iota
	OFPC_TABLE_STATS
	OFPC_PORT_STATS
	OFPC_GROUP_STATS
	_
	OFPC_IP_REASM
	OFPC_QUEUE_STATS
	_
	OFPC_PORT_BLOCKED
)

const (
	OFPC_FRAG_NORMAL = iota
	OFPC_FRAG_DROP
	OFPC_FRAG_REASM
	OFPC_FRAG_MARK
)

const (
	OFPTT_MAX = 0xfe
	OFPTT_ALL = 0xff
)

const OFPTC_DEPRECATED_MASK = 3

const (
	OFPFF_SEND_FLOW_REM = 1 << iota
	OFPFF_CHECK_OVERLAP
	OFPFF_RESET_COUNTS
	OFPFF_NO_PKT_COUNTS
	OFPFF_NO_BYT_COUNTS
)

const (
	OFPGC_ADD = iota
	OFPGC_MODIFY
	OFPGC_DELETE
)

const (
	OFPGT_ALL = iota
	OFPGT_SELECT
	OFPGT_INDIRECT
	OFPGT_FF
)

const (
	OFPG_MAX = 0xffffff00
	OFPG_ALL = 0xfffffffc
	OFPG_ANY = 0xffffffff
)

const (
	OFPM_MAX        = 0xffff0000
	OFPM_SLOWPATH   = 0xfffffffd
	OFPM_CONTROLLER = 0xfffffffe
	OFPM_ALL        = 0xffffffff
)

const (
	OFPMC_ADD = iota
	OFPMC_MODIFY
	OFPMC_DELETE
)

const (
	OFPMF_KBPS = 1 << iota
	OFPMF_PKTPS
	OFPMF_BURST
	OFPMF_STATS
)

const (
	OFPMBT_DROP         = 1
	OFPMBT_DSCP_REMARK  = 2
	OFPMBT_EXPERIMENTER = 0xffff
)

const OFPMPF_REQ_MORE = 1 << 0

const OFPMPF_REPLY_MORE = 1 << 0

const (
	OFPMP_DESC = iota
	OFPMP_FLOW
	OFPMP_AGGREGATE
	OFPMP_TABLE
	OFPMP_PORT_STATS
	OFPMP_QUEUE
	OFPMP_GROUP
	OFPMP_GROUP_DESC
	OFPMP_GROUP_FEATURES
	OFPMP_METER
	OFPMP_METER_CONFIG
	OFPMP_METER_FEATURES
	OFPMP_TABLE_FEATURES
	OFPMP_PORT_DESC
	OFPMP_EXPERIMENTER = 0xffff
)

const (
	OFPTFPT_INSTRUCTIONS = iota
	OFPTFPT_INSTRUCTIONS_MISS
	OFPTFPT_NEXT_TABLES
	OFPTFPT_NEXT_TABLES_MISS
	OFPTFPT_WRITE_ACTIONS
	OFPTFPT_WRITE_ACTIONS_MISS
	OFPTFPT_APPLY_ACTIONS
	OFPTFPT_APPLY_ACTIONS_MISS
	OFPTFPT_MATCH
	_
	OFPTFPT_WILDCARDS
	_
	OFPTFPT_WRITE_SETFIELD
	OFPTFPT_WRITE_SETFIELD_MISS
	OFPTFPT_APPLY_SETFIELD
	OFPTFPT_APPLY_SETFIELD_MISS
	OFPTFPT_EXPERIMENTER      = 0xfffe
	OFPTFPT_EXPERIMENTER_MISS = 0xffff
)

const (
	OFPGFC_SELECT_WEIGHT = 1 << iota
	OFPGFC_SELECT_LIVENESS
	OFPGFC_CHAINING
	OFPGFC_CHAINING_CHECKS
)

const (
	OFPCR_ROLE_NOCHANGE = iota
	OFPCR_ROLE_EQUAL
	OFPCR_ROLE_MASTER
	OFPCR_ROLE_SLAVE
)

const (
	OFPR_NO_MATCH = iota
	OFPR_ACTION
	OFPR_INVALID_TTL
)

const (
	OFPRR_IDLE_TIMEOUT = iota
	OFPRR_HARD_TIMEOUT
	OFPRR_DELETE
	OFPRR_GROUP_DELETE
)

const (
	OFPPR_ADD = iota
	OFPPR_DELETE
	OFPPR_MODIFY
)

const (
	OFPET_HELLO_FAILED = iota
	OFPET_BAD_REQUEST
	OFPET_BAD_ACTION
	OFPET_BAD_INSTRUCTION
	OFPET_BAD_MATCH
	OFPET_FLOW_MOD_FAILED
	OFPET_GROUP_MOD_FAILED
	OFPET_PORT_MOD_FAILED
	OFPET_TABLE_MOD_FAILED
	OFPET_QUEUE_OP_FAILED
	OFPET_SWITCH_CONFIG_FAILED
	OFPET_ROLE_REQUEST_FAILED
	OFPET_METER_MOD_FAILED
	OFPET_TABLE_FEATURES_FAILED
	OFPET_EXPERIMENTER = 0xffff
)

const (
	OFPHFC_INCOMPATIBLE = 0
	OFPHFC_EPERM        = 1
)

const (
	OFPBRC_BAD_VERSION = iota
	OFPBRC_BAD_TYPE
	OFPBRC_BAD_MULTIPART
	OFPBRC_BAD_EXPERIMENTER
	OFPBRC_BAD_EXP_TYPE
	OFPBRC_EPERM
	OFPBRC_BAD_LEN
	OFPBRC_BUFFER_EMPTY
	OFPBRC_BUFFER_UNKNOWN
	OFPBRC_BAD_TABLE_ID
	OFPBRC_IS_SLAVE
	OFPBRC_BAD_PORT
	OFPBRC_BAD_PACKET
	OFPBRC_MULTIPART_BUFFER_OVERFLOW
)

const (
	OFPBAC_BAD_TYPE = iota
	OFPBAC_BAD_LEN
	OFPBAC_BAD_EXPERIMENTER
	OFPBAC_BAD_EXP_TYPE
	OFPBAC_BAD_OUT_PORT
	OFPBAC_BAD_ARGUMENT
	OFPBAC_EPERM
	OFPBAC_TOO_MANY
	OFPBAC_BAD_QUEUE
	OFPBAC_BAD_OUT_GROUP
	OFPBAC_MATCH_INCONSISTENT
	OFPBAC_UNSUPPORTED_ORDER
	OFPBAC_BAD_TAG
	OFPBAC_BAD_SET_TYPE
	OFPBAC_BAD_SET_LEN
	OFPBAC_BAD_SET_ARGUMENT
)

const (
	OFPBIC_UNKNOWN_INST = iota
	OFPBIC_UNSUP_INST
	OFPBIC_BAD_TABLE_ID
	OFPBIC_UNSUP_METADATA
	OFPBIC_UNSUP_METADATA_MASK
	OFPBIC_BAD_EXPERIMENTER
	OFPBIC_BAD_EXP_TYPE
	OFPBIC_BAD_LEN
	OFPBIC_EPERM
)

const (
	OFPBMC_BAD_TYPE = iota
	OFPBMC_BAD_LEN
	OFPBMC_BAD_TAG
	OFPBMC_BAD_DL_ADDR_MASK
	OFPBMC_BAD_NW_ADDR_MASK
	OFPBMC_BAD_WILDCARDS
	OFPBMC_BAD_FIELD
	OFPBMC_BAD_VALUE
	OFPBMC_BAD_MASK
	OFPBMC_BAD_PREREQ
	OFPBMC_DUP_FIELD
	OFPBMC_EPERM
)

const (
	OFPFMFC_UNKNOWN = iota
	OFPFMFC_TABLE_FULL
	OFPFMFC_BAD_TABLE_ID
	OFPFMFC_OVERLAP
	OFPFMFC_EPERM
	OFPFMFC_BAD_TIMEOUT
	OFPFMFC_BAD_COMMAND
	OFPFMFC_BAD_FLAGS
)

const (
	OFPGMFC_GROUP_EXISTS = iota
	OFPGMFC_INVALID_GROUP
	OFPGMFC_WEIGHT_UNSUPPORTED
	OFPGMFC_OUT_OF_GROUPS
	OFPGMFC_OUT_OF_BUCKETS
	OFPGMFC_CHAINING_UNSUPPORTED
	OFPGMFC_WATCH_UNSUPPORTED
	OFPGMFC_LOOP
	OFPGMFC_UNKNOWN_GROUP
	OFPGMFC_CHAINED_GROUP
	OFPGMFC_BAD_TYPE
	OFPGMFC_BAD_COMMAND
	OFPGMFC_BAD_BUCKET
	OFPGMFC_BAD_WATCH
	OFPGMFC_EPERM
)

const (
	OFPPMFC_BAD_PORT = iota
	OFPPMFC_BAD_HW_ADDR
	OFPPMFC_BAD_CONFIG
	OFPPMFC_BAD_ADVERTISE
	OFPPMFC_EPERM
)

const (
	OFPTMFC_BAD_TABLE = iota
	OFPTMFC_BAD_CONFIG
	OFPTMFC_EPERM
)

const (
	OFPQOFC_BAD_PORT = iota
	OFPQOFC_BAD_QUEUE
	OFPQOFC_EPERM
)

const (
	OFPSCFC_BAD_FLAGS = iota
	OFPSCFC_BAD_LEN
	OFPSCFC_EPERM
)

const (
	OFPRRFC_STALE = iota
	OFPRRFC_UNSUP
	OFPRRFC_BAD_ROLE
)

const (
	OFPMMFC_UNKNOWN = iota
	OFPMMFC_METER_EXISTS
	OFPMMFC_INVALID_METER
	OFPMMFC_UNKNOWN_METER
	OFPMMFC_BAD_COMMAND
	OFPMMFC_BAD_FLAGS
	OFPMMFC_BAD_RATE
	OFPMMFC_BAD_BURST
	OFPMMFC_BAD_BAND
	OFPMMFC_BAD_BAND_VALUE
	OFPMMFC_OUT_OF_METERS
	OFPMMFC_OUT_OF_BANDS
)

const (
	OFPTFFC_BAD_TABLE = iota
	OFPTFFC_BAD_METADATA
	OFPTFFC_BAD_TYPE
	OFPTFFC_BAD_LEN
	OFPTFFC_BAD_ARGUMENT
	OFPTFFC_EPERM
)

const OFPHET_VERSIONBITMAP = 1

const OFP_ETH_ALEN = 6
const OFP_MAX_TABLE_NAME_LEN = 32
const OFP_MAX_PORT_NAME_LEN = 16

const OFP_NO_BUFFER = 0xffffffff

const DESC_STR_LEN = 256
const SERIAL_NUM_LEN = 32

const (
	MSG_UNKNOWN = iota
	MSG_SYMMETRIC
	MSG_ASYNC
	MSG_REQUEST
	MSG_REPLY
)

func MessageType(oftype uint8) int {
	switch oftype {
	default:
		return MSG_UNKNOWN
	case OFPT_HELLO, OFPT_ERROR, OFPT_EXPERIMENTER:
		return MSG_SYMMETRIC
	case OFPT_PACKET_IN, OFPT_FLOW_REMOVED, OFPT_PORT_STATUS:
		return MSG_ASYNC
	case OFPT_ECHO_REQUEST,
		OFPT_FEATURES_REQUEST,
		OFPT_GET_CONFIG_REQUEST,
		OFPT_SET_CONFIG,
		OFPT_PACKET_OUT,
		OFPT_FLOW_MOD,
		OFPT_GROUP_MOD,
		OFPT_PORT_MOD,
		OFPT_TABLE_MOD,
		OFPT_MULTIPART_REQUEST,
		OFPT_BARRIER_REQUEST,
		OFPT_QUEUE_GET_CONFIG_REQUEST,
		OFPT_ROLE_REQUEST,
		OFPT_GET_ASYNC_REQUEST,
		OFPT_SET_ASYNC,
		OFPT_METER_MOD:
		return MSG_REQUEST
	case OFPT_ECHO_REPLY,
		OFPT_FEATURES_REPLY,
		OFPT_GET_CONFIG_REPLY,
		OFPT_MULTIPART_REPLY,
		OFPT_BARRIER_REPLY,
		OFPT_QUEUE_GET_CONFIG_REPLY,
		OFPT_ROLE_REPLY,
		OFPT_GET_ASYNC_REPLY:
		return MSG_REPLY
	}
}
