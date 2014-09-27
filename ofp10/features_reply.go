package ofp10

import (
	"encoding/binary"
	"errors"

	"ofp"
)

type FeaturesReply struct {
	ofp.Header
	Dpid uint64
	NBuffers
	pad [3]byte
	Capabilities uint32
	Actions uint32
	Ports []
	
}
