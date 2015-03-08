package ofp10
import (
    "encoding/binary"

    "github.com/kuun/ofgo/ofp"
)

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

func (self *Match)Len() int {
    return 40
}

func (self *Match)Marshal(buff []byte) (n int, err error) {
    if (len(buff) < self.Len()) {
        return 0, ofp.NewNoBuffError()
    }
    n = 0
    binary.BigEndian.PutUint32(buff, self.Wildcards)
    n += 4
    binary.BigEndian.PutUint16(buff[n:], self.InPort)
    n += 2
    copy(buff[n:], self.EthSrc[:])
    n += OFP_ETH_ALAN
    copy(buff[n:], self.EthDst[:])
    n += OFP_ETH_ALAN
    binary.BigEndian.PutUint16(buff[n:], self.VlanId)
    n += 2
    buff[n] = self.VlanPcp
    n += 2  // plus one padding byte
    binary.BigEndian.PutUint16(buff[n:], self.EthType)
    n += 2
    buff[n] = self.NwTos
    n++
    buff[n] = self.NwProto
    n += 3 // plus 2 padding byte
    binary.BigEndian.PutUint32(buff[n:], self.NwSrc)
    n += 4
    binary.BigEndian.PutUint32(buff[n:], self.NwDst)
    n += 4
    binary.BigEndian.PutUint16(buff[n:], self.TpSrc)
    n += 2
    binary.BigEndian.PutUint16(buff[n:], self.TpDst)
    n += 2

    return n, nil
}

func (self *Match)Unmarshal(buff []byte) (n int, err error) {
    if (len(buff) < self.Len()) {
        return 0, ofp.NewNoBuffError()
    }
    n = 0
    self.Wildcards = binary.BigEndian.Uint32(buff)
    n += 4
    self.InPort = binary.BigEndian.Uint16(buff[n:])
    n += 2
    copy(self.EthSrc[:], buff[n:])
    n += OFP_ETH_ALAN
    copy(self.EthDst[:], buff[n:])
    n += OFP_ETH_ALAN
    self.VlanId = binary.BigEndian.Uint16(buff[n:])
    n += 2
    self.VlanPcp = buff[n]
    n += 2
    self.EthType = binary.BigEndian.Uint16(buff[n:])
    n += 2
    self.NwTos = buff[n]
    n++
    self.NwProto = buff[n]
    n += 3
    self.NwSrc = binary.BigEndian.Uint32(buff[n:])
    n += 4
    self.NwDst = binary.BigEndian.Uint32(buff[n:])
    n += 4
    self.TpSrc = binary.BigEndian.Uint16(buff[n:])
    n += 2
    self.TpDst = binary.BigEndian.Uint16(buff[n:])

    return n, nil
}