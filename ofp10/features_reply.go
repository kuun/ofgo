package ofp10

import (
	"encoding/binary"
	"errors"

	"github.com/kuun/ofgo/ofp"
)

type FeaturesReply struct {
	ofp.Header
	Dpid uint64
	NBuffers uint32
	pad [3]byte
	Capabilities uint32
	Actions uint32
	Ports []Port
}

func (msg *FeaturesReply) Len() int {
	return int(msg.Length)
}

func (msg *FeaturesReply) Marshal(buf []byte) (n int , err error) {
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
