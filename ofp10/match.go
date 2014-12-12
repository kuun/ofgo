package ofp10

const OFP_ETH_ALAN = 6

// Flow wildcards
const (
	OFPFW_IN_PORT      = 1 << iota //
	OFPFW_DL_VLAN                  //
	OFPFW_DL_SRC                   //
	OFPFW_DL_DST                   //
	OFPFW_DL_TYPE                  //
	OFPFW_NW_PROTO                 //
	OFPFW_TP_SRC                   //
	OFPFW_TP_DST                   //

	// IP source address wildcard bit count.  0 is exact match, 1 ignores the
    // LSB, 2 ignores the 2 least-significant bits, ..., 32 and higher wildcard
    // the entire field.  This is the *opposite* of the usual convention where
    // e.g. /24 indicates that 8 bits (not 24 bits) are wildcarded. 
	OFPFW_NW_SRC_SHIFT = 8
	OFPFW_NW_SRC_BITS  = 6
	OFPFW_NW_SRC_MASK  = (1<<OFPFW_NW_SRC_BITS - 1) << OFPFW_NW_SRC_SHIFT
	OFPFW_NW_SRC_ALL   = 32 << OFPFW_NW_SRC_SHIFT

	// IP destination address wildcard bit count.  Same format as source.
	OFPFW_NW_DST_SHIFT = 14
	OFPFW_NW_DST_BITS  = 6
	OFPFW_NW_DST_MASK  = (1<<OFPFW_NW_DST_BITS - 1) << OFPFW_NW_DST_SHIFT
	OFPFW_NW_DST_ALL   = 32 << OFPFW_NW_DST_SHIFT

	OFPFW_DL_VLAN_PCP = 1 << 20   // VLAN priority.
	OFPFW_NW_TOS      = 1 << 21   // IP ToS (DSCP field, 6 bits).
	OFPFW_ALL         = 1<<22 - 1 // Wildcard all fields.
)

// Match is used to describe a flow entry, fileds to match against flows.
type Match struct {
	Wildcards uint32             // Wildcard fields.
	InPort    uint16             // Input switch port.
	EthSrc    [OFP_ETH_ALAN]byte // Ethernet source address.
	EthDst    [OFP_ETH_ALAN]byte // Ethernet destination address.
	VlanId    uint16             // Input VLAN id.
	VlanPcp   uint8              // Input VLAN priority.
	pad1      byte               // Align to 64-bits
	EthType   uint16             // Ethernet frame type.
	NwTos     byte               // IP ToS (actually DSCP field, 6 bits).
	NwProto   byte               // IP protocol or lower 8 bits of ARP opcode.
	pad2      [2]byte            // Align to 64-bits
	NwSrc     uint32             // IP source address.
	NwDst     uint32             // IP destination address.
	TpSrc     uint16             // TCP/UDP source port.
	TpDst     uint16             // TCP/UDP destination port.
}
