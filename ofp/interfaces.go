package ofp

// DataBlock is some type data in an openflow message, it can be the message header
// or a message body and so on. it can marshal to binary data as part of a message or
// unmarshal from binary data.
type DataBlock interface {
	// Marshal marshals DataBlock to binary, stores bytes to buffer, it returns
	// the number of bytes written and any error encountered.
	Marshal(buffer []byte) (n int, err error)
	// Unmarshal unmarshals binary to DataBlock from buffer, returns the number
	// of bytes read and any error encounterd.
	Unmarshal(buffer []byte) (n int, err error)
	// Len gets DataBlock's binary length by byte.
	Len() int
}
