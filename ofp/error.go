package ofp

import (
	"encoding/binary"
	"errors"
)

// Values for 'Type' in Error.  These values are immutable: they
// will not change in future versions of the protocol (although new values may
// be added).
const (
	OFPET_HELLO_FAILED    = iota // Hello protocol failed.
	OFPET_BAD_REQUEST            // Request was not understood.
	OFPET_BAD_ACTION             // Error in action description.
	OFPET_FLOW_MOD_FAILED        // Problem modifying flow entry.
	OFPET_PORT_MOD_FAILED        // Port mod request failed.
	OFPET_QUEUE_OP_FAILED        // Queue operation failed.
)

// Error 'Code' values for OFPET_HELLO_FAILED. 'Data' contains an
// ASCII text string that may give failure details.
const (
	OFPHFC_INCOMPATIBLE = iota // No compatible version.
	OFPHFC_EPERM               // Permissions error.
)

// Error 'Code' values for OFPET_BAD_REQUEST. 'Data' contains at least
// the first 64 bytes of the failed request.
const (
	OFPBRC_BAD_VERSION    = iota // Header.Version not supported.
	OFPBRC_BAD_TYPE              // Header.Type not supported.
	OFPBRC_BAD_STAT              // StatsRequest.Type not supported.
	OFPBRC_BAD_VENDOR            // Vendor not supported
	OFPBRC_BAD_SUBTYPE           // Vendor subtype not supported.
	OFPBRC_EPERM                 // Permissions error.
	OFPBRC_BAD_LEN               // Wrong request length for type.
	OFPBRC_BUFFER_EMPTY          // Specified buffer has already been used.
	OFPBRC_BUFFER_UNKNOWN        // Specified buffer does not exist.
)

// Error 'Code' values for OFPET_BAD_ACTION.  ’Data’ contains at least
// the first 64 bytes of the failed request.
const (
	OFPBAC_BAD_TYPE        = iota // Unknown action type.
	OFPBAC_BAD_LEN                // Length problem in actions.
	OFPBAC_BAD_VENDOR             // Unknown vendor id specified.
	OFPBAC_BAD_VENDOR_TYPE        // Unknown action type for vendor id.
	OFPBAC_BAD_OUT_PORT           // Problem validating output action.
	OFPBAC_BAD_ARGUMENT           // Bad action argument.
	OFPBAC_EPERM                  // Permissions error.
	OFPBAC_TOO_MANY               // Can’t handle this many actions.
	OFPBAC_BAD_QUEUE              // Problem validating output queue.
)

// Error 'Code' values for OFPET_FLOW_MOD_FAILED.  'Data' contains at least the first 64 bytes of the failed request.
const (
	OFPFMFC_ALL_TABLES_FULL   = iota // Flow not added because of full tables.
	OFPFMFC_OVERLAP                  // Attempted to add overlapping flow with CHECK_OVERLAP flag set.
	OFPFMFC_EPERM                    //  Permissions error.
	OFPFMFC_BAD_EMERG_TIMEOUT        // Flow not added because of non-zero idle/hard timeout.
	OFPFMFC_BAD_COMMAND              // Unknown command.
	OFPFMFC_UNSUPPORTED              // Unsupported action list - cannot process in the order specified.
)

// Error 'Code' values for OFPET_PORT_MOD_FAILED. 'Data' contains at least the first 64 bytes of the failed request.
const (
	OFPPMFC_BAD_PORT    = iota // Specified port does not exist.
	OFPPMFC_BAD_HW_ADDR        // Specified hardware address is wrong.
)

// Error 'Code' alues for OFPET_QUEUE_OP_FAILED. 'Data' contains at least the first 64 bytes of the failed request.
const (
	OFPQOFC_BAD_PORT = iota // Invalid port (or port does not exist).
	FPQOFC_BAD_QUEUE        // Queue does not exist.
	OFPQOFC_EPERM           // Permissions error.
)

// Error is openflow error message, openflow switch -> controller.
type Error struct {
	Header
	Type   uint16 // High-level type of error, value is one of OFPET_*.
	Code   uint16 // Interpreted based on the type.
	// Variable-length data. Interpreted based on the type and code, in most
	// cases this is the message that caused the problem.
	Data []byte
}

func (msg *Error) Marshal(buf []byte) (n int, err error) {
	if len(buf) < msg.Len() {
		return 0, errors.New("buffer is too short")
	}
	n, err = msg.Header.Marshal(buf)
	if err != nil {
		return
	}
	binary.BigEndian.PutUint16(buf[n:], msg.Type)
	n += 2
	binary.BigEndian.PutUint16(buf[n:], msg.Code)
	n += 2
	copy(buf[n:], msg.Data)
	return n, nil
}

func (msg *Error) Unmarshal(buf []byte) (n int, err error) {
	if len(buf) < msg.Len() {
		return 0, errors.New("buffer is too short")
	}
	n, err = msg.Header.Unmarshal(buf)
	if err != nil {
		return
	}
	msg.Type = binary.BigEndian.Uint16(buf[n:])
	n += 2
	msg.Code = binary.BigEndian.Uint16(buf[n:])
	n += 2
	copy(msg.Data, buf[n:msg.Len()-n])
	return msg.Len(), nil
}

func (msg *Error) Len() int {
	return int(msg.Header.Length)
}
