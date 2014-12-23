package ofp

import (
	"encoding/binary"
	"errors"
)

// Header is the header of all openflow messages.
type Header struct {
	Version uint8  // An OpenFlow version number, e.g. OFP10_VERSION.
	Type    uint8  // An OpenFlow version number, e.g. OFP10_VERSION.
	Length  uint16 // Length of the message including this Header.
	// Transaction id associated with this packet.
	// Replies use the same id as was in the request
	// to facilitate pairing.
	Xid uint32
}

const HeaderLength = 8

func (h *Header) Marshal(b []byte) (n int, err error) {
	if len(b) < h.Len() {
		return 0, errors.New("buffer is too short")
	}
	b[0] = h.Version
	b[1] = h.Type
	n = 2
	binary.BigEndian.PutUint16(b[n:], h.Length)
	n += 2
	binary.BigEndian.PutUint32(b[n:], h.Xid)
	n += 4
	return n, nil
}

func (h *Header) Unmarshal(b []byte) (n int, err error) {
	if len(b) < h.Len() {
		return 0, errors.New("buffer is too short")
	}
	h.Version = b[0]
	h.Type = b[1]
	n = 2
	h.Length = binary.BigEndian.Uint16(b[n:])
	n += 2
	h.Xid = binary.BigEndian.Uint32(b[n:])
	n += 4
	return n, nil
}

func (h *Header) Len() int {
	return HeaderLength
}
