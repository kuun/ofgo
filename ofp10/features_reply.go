package ofp10

import (
	"encoding/binary"
	"errors"

	"github.com/kuun/ofgo/ofp"
)

// Capabilities supported by the datapath.
const (
	OFPC_FLOW_STATS   = 1 << iota // Flow statistics.
	OFPC_TABLE_STATS              // Table statistics.
	OFPC_PORT_STATS               // Port statistics.
	OFPC_STP                      // 802.1d spanning tree.
	OFPC_RESERVED                 // Reserved, must be zero.
	OFPC_IP_REASM                 // Can reassemble IP fragments.
	OFPC_QUEUE_STATS              // Queue statistics.
	OFPC_ARP_MATCH_IP             // Match IP addresses in ARP pkts.
)

type FeaturesReply struct {
	ofp.Header
	Dpid         uint64  // Datapath unique id, the lower 48-bits are for a MAC address, while the upper 16-bits are implementer-defined.
	NBuffers     uint32  // Max packets buffered at once.
	pad           [3]byte // Align to 64-bits.
	Capabilities uint32  // Datapath capabilities, bitmap of OFPC_*
	Actions      uint32  // Supported actions, bitmap of OFPAT_*
	Ports        []Port // Port definitions.  The number of ports is inferred from the length field in the header.
}

func (msg *FeaturesReply) Len() int {
	return int(msg.Length)
}

func (msg *FeaturesReply) Marshal(buf []byte) (n int, err error) {
	panic("not implement!")
}

func (msg *FeaturesReply) Unmarshal(buf []byte) (n int, err error) {
	if n, err = msg.Header.Unmarshal(buf); err != nil {
		return
	}
	if len(buf) < msg.Len() {
		return 0, errors.New("buffer is too short")
	}
	n += msg.Header.Len()
	msg.Dpid = binary.BigEndian.Uint64(buf[n:])
	n += 8
	msg.NBuffers = binary.BigEndian.Uint32(buf[n:])
	n += 4
	n += 3 // 3 bytes pad
	msg.Capabilities = binary.BigEndian.Uint32(buf[n:])
	n += 4
	msg.Actions = binary.BigEndian.Uint32(buf[n:])
	n += 4
	leftSize := int(msg.Len()) - n
	portNum := leftSize / portSize
	for i := 0; i < portNum; i++ {
		port := Port{}
		var m int
		if m, err = port.Unmarshal(buf[n:]); err != nil {
			return n + m, err
		}
		msg.Ports = append(msg.Ports, port)
		n += m
	}
	return
}
