package ofp

import (
	"errors"
)

// EchoResponse is openflow echo response message.
type EchoResponse struct {
	Header
	data []byte		// An arbitrary-length user data, zero size is valid.
}

func (msg *EchoResponse) Marshal(buf []byte) (n int, err error) {
	if len(buf) < msg.Len() {
		return 0, errors.New("buffer is too short")
	}
	n, err = msg.Header.Marshal(buf)
	if err != nil {
		return n, err
	}
	if msg.data != nil {
		copy(buf[n:msg.Len()], msg.data)
	}
	return msg.Len(), nil
}

func (msg *EchoResponse) Unmarshal(buf []byte) (n int, err error) {
	if len(buf) < msg.Len() {
		return 0, errors.New("buffer is too short")
	}
	if n, err = msg.Header.Unmarshal(buf); err != nil {
		return n, err
	}
	dataLen := msg.Len() - n
	if dataLen > 0 {
		msg.data = make([]byte, dataLen, dataLen)
		copy(msg.data, buf[n:])
	}
	return msg.Len(), nil
}

func (msg *EchoResponse) Len() int {
	return int(msg.Header.Length)
}

// SetData sets the message's payload data. the message will own the 'data'.
func (msg *EchoResponse) SetData(data []byte, len int) *EchoResponse {
	msg.data = data
	
	return msg
}

// Data gets the message's payload data.
func (msg *EchoResponse) Data() []byte {
	return msg.data
}

