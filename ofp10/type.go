package ofp10

import (
	"github.com/kuun/ofgo/ofp"
)

// openflow 1.0 message type
const (
	// immutable messages, symmetric messages.
	OFPT_HELLO = iota
	OFPT_ERROR
	OFPT_ECHO_REQUEST
	OFPT_ECHO_REPLY
	OFPT_VENDOR

	// switch configuration message.
	OFPT_FEATURES_REQUEST
	OFPT_FEATURES_REPLY
	OFPT_GET_CONFIG_REQUEST
	OFPT_GET_CONFIG_REPLY
	OFPT_SET_CONFIG

	// asynchronous messages.
	OFPT_PACKET_IN
	OFPT_FLOW_REMOVED
	OFPT_PORT_STATUS

	// controller command messages.
	OFPT_PACKET_OUT
	OFPT_FLOW_MOD
	OFPT_PORT_MOD

	// statistics messages.
	OFPT_STATS_REQUEST
	OFPT_STATS_REPLY

	// barrier messages.
	OFPT_BARRIER_REQUEST
	OFPT_BARRIER_REPLY

	// queue configuration messages.
	OFPT_QUEUE_GET_CONFIG_REQUEST
	OFPT_QUEUE_GET_CONFIG_REPLYA
)

func NewHello() *ofp.Hello {
	return &ofp.Hello{
		Header: ofp.Header{
			Version: ofp.OFP10_VERSION,
			Type: OFPT_HELLO,
			Length: ofp.HeaderLength,
		},
	}
}

func NewEchoRequest() *ofp.EchoRequest {
	return &ofp.EchoRequest{
		Header: ofp.Header{
			Version: ofp.OFP10_VERSION,
			Type:   OFPT_ECHO_REQUEST,
			Length: ofp.HeaderLength,
		},
	}
}

