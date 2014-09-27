package ofp

import (
	"encoding/binary"
	"errors"
)

// VendorHeader is openflow vendor message header
type VendorHeader struct {
	Header
	VendorId uint32
}

func (h *VendorHeader) Marshal(buf []byte) (n int, err error) {
	if len(buf) < h.Len() {
		return 0, errors.New("buffer is too short")
	}
	if n, err = h.Header.Marshal(buf); err != nil {
		return n, err
	}
	binary.BigEndian.PutUint32(buf[n:], h.VendorId)
	n += 4
	return n, nil
}

func (h *VendorHeader) Unmarshal(buf []byte) (n int, err error) {
	if len(buf) < h.Len() {
		return 0, errors.New("buffer is too short")
	}
	if n, err = h.Header.Unmarshal(buf); err != nil {
		return n, err
	}
	h.VendorId = binary.BigEndian.Uint32(buf[n:])
	n += 4
	return n, nil
}

func (h *VendorHeader) Len() int {
	return h.Header.Len() + 4
}
