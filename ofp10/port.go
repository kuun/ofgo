package ofp10

import (
	"encoding/binary"
	"errors"
)

const OFP_MAX_PORT_NAME_LEN = 16

// port config flags, can be used to configure the port's behavior.
const (
	OFPPC_PORT_DOWN    = 1 << iota // port is administratively down
	OFPPC_NO_STP                   // diable 802.1D spanning tree on port
	OFPPC_NO_RECV                  // drop all packets except 802.1D spanning tree packets
	OFPPC_NO_RECV_STP              // drop received 802.1D STP packets
	OFPPC_NO_FLOOD                 // don't include this port when flooding.
	OFPPC_NO_FWD                   // drop packets forwarded to port
	OFPPC_NO_PACKET_IN             // do not send packet-in message for port
)

// port state, these are not configurable from the controller.
const (
	OFPPS_LINK_DOWN = 1 << 0 // No physical link present.

	// The OFPPS_STP_* bits have no effect on switch operation.  The
	// controller must adjust OFPPC_NO_RECV, OFPPC_NO_FWD, and
	// OFPPC_NO_PACKET_IN appropriately to fully implement an 802.1D spanning
	// tree.
	OFPPS_STP_LISTEN  = 0 << 8 // Not learning or relaying frames.
	OFPPS_STP_LEARN   = 1 << 8 // Learning but not relaying frames.
	OFPPS_STP_FORWARD = 2 << 8 // Learning and relaying frames.
	OFPPS_STP_BLOCK   = 3 << 8 // Not part of spanning tree.
	OFPPS_STP_MASK    = 3 << 8 // Bit mask for OFPPS_STP_* values.
)

// Reserved openflow port number.
const (
	OFPP_MAX = 0xff00 // Maximum number of physical switch ports.

	// Fake output "ports".
	OFPP_IN_PORT = 0xfff8 /* Send the packet out the input port.  This virtual
	* port must be explicitly used in order to send back out of the input */
	OFPP_TABLE = 0xfff9 /* Perform actions in flow table. NB: This can only
	* be the destination port for packet-out messages */
	OFPP_NORMAL     = 0xfffa // Process with normal L2/L3 switching.
	OFPP_FLOOD      = 0xfffb // All physical ports except input port and those disabled by STP.
	OFPP_ALL        = 0xfffc // All physical ports except input port.
	OFPP_CONTROLLER = 0xfffd // Send to controller.
	OFPP_LOCAL      = 0xfffe // Local openflow "port".
	OFPP_NONE       = 0xffff // Not associated with a physical port.
)

// Features of physical ports available in a datapath.
const (
	OFPPF_10MB_HD    = 1 << iota // 10 Mb half-duplex rate support.
	OFPPF_10MB_FD                // 10 Mb full-duplex rate support.
	OFPPF_100MB_HD               // 100 Mb half-duplex rate support.
	OFPPF_100MB_FD               // 100 Mb full-duplex rate support.
	OFPPF_1GB_HD                 // 1 Gb half-duplex rate support
	OFPPF_1GB_FD                 // 1 Gb full-duplex rate support.
	OFPPF_10GB_FD                // 10 Gb full-duplex rate support.
	OFPPF_COPPER                 // Copper medium.
	OFPPF_FIBER                  // Fiber medium.
	OFPPF_AUTONEG                // Auto-negotiation.
	OFPPF_PAUSE                  // Pause.
	OFPPF_PAUSE_ASYM             // Asymmetric pause.
)

// port binary size, in byte
const portSize = 48

// Port is an openflow physical port.
type Port struct {
	PortNo uint16
	HwAddr [6]byte
	Name   [OFP_MAX_PORT_NAME_LEN]byte // Null-terminated

	Config uint32 // Bitmap of OFPPC_* flags.
	State  uint32 // Bitmap of OFPPS_* flags.

	// Bitmaps of OFPPF_* that describe features.  All bits zeroed if unsupported or unavailable.
	CurrFeatures       uint32 // Current features.
	AdvertisedFeatures uint32 // Features being advertised by the port.
	SupportedFeatures  uint32 // Features supported by the port.
	PeerFeatures       uint32 // Features advertised by peer.
}

func (p *Port) Unmarshal(buf []byte) (n int, err error) {
	if len(buf) < p.Len() {
		return 0, errors.New("buffer is too short")
	}
	n = 0
	p.PortNo = binary.BigEndian.Uint16(buf)
	n += 2
	copy(p.HwAddr[:], buf[n:])
	n += 6
	copy(p.Name[:], buf[n:])
	n += OFP_MAX_PORT_NAME_LEN
	p.Config = binary.BigEndian.Uint32(buf[n:])
	n += 4
	p.State = binary.BigEndian.Uint32(buf[n:])
	n += 4
	p.CurrFeatures = binary.BigEndian.Uint32(buf[n:])
	n += 4
	p.AdvertisedFeatures = binary.BigEndian.Uint32(buf[n:])
	n += 4
	p.SupportedFeatures = binary.BigEndian.Uint32(buf[n:])
	n += 4
	p.PeerFeatures = binary.BigEndian.Uint32(buf[n:])
	n += 4
	return n, nil
}

func (p *Port) Marshal(buf []byte) (n int, err error) {
	if len(buf) < p.Len() {
		return n, errors.New("buffer is too short")
	}
	n = 0
	binary.BigEndian.PutUint16(buf, p.PortNo)
	n += 2
	copy(buf[n:], p.HwAddr[:])
	n += 6
	copy(buf[n:], p.Name[:])
	n += OFP_MAX_PORT_NAME_LEN
	binary.BigEndian.PutUint32(buf[n:], p.Config)
	n += 4
	binary.BigEndian.PutUint32(buf[n:], p.State)
	n += 4
	binary.BigEndian.PutUint32(buf[n:], p.CurrFeatures)
	n += 4
	binary.BigEndian.PutUint32(buf[n:], p.AdvertisedFeatures)
	n += 4
	binary.BigEndian.PutUint32(buf[n:], p.SupportedFeatures)
	n += 4
	binary.BigEndian.PutUint32(buf[n:], p.PeerFeatures)
	n += 4
	return n, nil
}

func (p *Port) Len() int {
	return portSize
}
