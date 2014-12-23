package ofp

import (
	"errors"
)

// EchoRequest is openflow echo request message.
type EchoRequest struct {
	Header
	data []byte // An arbitrary-length user data, zero size is valid.
}

func (msg *EchoRequest) Marshal(buf []byte) (n int, err error) {
	if len(buf) < msg.Len() {
		return 0, errors.New("buffer is too short")
	}
	if n, err = msg.Header.Marshal(buf); err != nil {
		return n, err
	}
	if msg.data != nil {
		copy(buf[n:msg.Len()], msg.data)
	}
	return msg.Len(), nil
}

func (msg *EchoRequest) Unmarshal(buf []byte) (n int, err error) {
	if n, err = msg.Header.Unmarshal(buf); err != nil {
		return n, err
	}
	if len(buf) < msg.Len() {
		return 0, errors.New("buffer is too short")
	}
	dataLen := msg.Len() - n
	if dataLen > 0 {
		msg.data = make([]byte, dataLen, dataLen)
		copy(msg.data, buf[n:])
	}
	return msg.Len(), nil
}

func (msg *EchoRequest) Len() int {
	return int(msg.Header.Length)
}

// SetData sets the message's payload data. the message will own the 'data'.
func (msg *EchoRequest) SetData(data []byte) *EchoRequest {
	msg.data = data
	msg.Header.Length = uint16(msg.Len() + len(data))

	return msg
}

// Data gets the message's payload data.
func (msg *EchoRequest) Data() []byte {
	return msg.data
}
