package ofp10

import (
	"encoding/binary"
	"github.com/kuun/ofgo/ofp"
)

// Defines openflow 1.0 action type.
const (
	OFPAT_OUTPUT       = iota // Output to switch port.
	OFPAT_SET_VLAN_VID        // Set the 802.1q VLAN id.
	OFPAT_SET_VLAN_PCP        // Set the 802.1q priority.
	OFPAT_STRIP_VLAN          // Strip the 802.1q header.
	OFPAT_SET_DL_SRC          // Ethernet source address.
	OFPAT_SET_DL_DST          // Ethernet destination address.
	OFPAT_SET_NW_SRC          // IP source address.
	OFPAT_SET_NW_DST          // IP destination address.
	OFPAT_SET_NW_TOS          // IP ToS (DSCP field, 6 bits).
	OFPAT_SET_TP_SRC          // TCP/UDP source port.
	OFPAT_SET_TP_DST          // TCP/UDP destination port.
	OFPAT_ENQUEUE             // Output to queue.
	OFPAT_VENDOR       = 0xffff
)

type ActionType uint16

type Action interface {
	Type() ActionType
}

// Actionheader is common to all actions. The length includes the
// header and any padding used to make the action 64-bit aligned.
// NB: The length of an action *must* always be a multiple of eight.
type ActionHeader struct {
	// One of OFPAT_*.
	Type ActionType
	// Length of action, including this header. This is the length of action,
	// including any padding to make it 64-bit aligned.
	Length uint16
}

func (self *ActionHeader) Len() int {
	return 4
}

func (self *ActionHeader) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	n = 0
	binary.BigEndian.PutUint16(buff, uint16(self.Type))
	n += 2
	binary.BigEndian.PutUint16(buff[n:], self.Length)
	n += 2
	return n, nil
}

func (self *ActionHeader) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	n = 0
	self.Type = ActionType(binary.BigEndian.Uint16(buff))
	n += 2
	self.Length = binary.BigEndian.Uint16(buff[n:])
	n += 2
	return n, nil
}

// Action structure for OFPAT_OUTPUT, which sends packets out ’port’.
// When the ’port’ is the OFPP_CONTROLLER, ’max_len’ indicates the max
// number of bytes to send. A ’max_len’ of zero means no bytes of the
// packet should be sent
type ActionOutput struct {
	ActionHeader        // type must be OFPAT_OUTPUT.
	Port         uint16 // Output port.
	MaxLen       uint16 // Max length to send to controller.
}

func NewActionOutput() *ActionOutput {
	return &ActionOutput{
		ActionHeader: ActionHeader{Type: OFPAT_OUTPUT, Length: 8},
	}
}

func (self *ActionOutput) Type() ActionType {
	return OFPAT_OUTPUT
}

func (self *ActionOutput) Len() int {
	return int(self.Length)
}

func (self *ActionOutput) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Marshal(buff); err != nil {
		return n, err
	}
	binary.BigEndian.PutUint16(buff[n:], self.Port)
	n += 2
	binary.BigEndian.PutUint16(buff[n:], self.MaxLen)
	n += 2
	return n, nil
}

func (self *ActionOutput) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Unmarshal(buff); err != nil {
		return n, err
	}
	self.Port = binary.BigEndian.Uint16(buff[n:])
	n += 2
	self.MaxLen = binary.BigEndian.Uint16(buff[n:])
	n += 2
	return n, nil
}

type ActionEnqueue struct {
	ActionHeader
	Port    uint16
	Pad     [6]byte
	QueueId uint32
}

func NewActionEnqueue() *ActionEnqueue {
	return &ActionEnqueue{
		ActionHeader: ActionHeader{Type: OFPAT_ENQUEUE, Length: 16},
	}
}

func (self *ActionEnqueue) Type() ActionType {
	return self.ActionHeader.Type
}

func (self *ActionEnqueue) Len() int {
	return int(self.ActionHeader.Length)
}

func (self *ActionEnqueue) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Marshal(buff); err != nil {
		return n, err
	}
	binary.BigEndian.PutUint16(buff[n:], self.Port)
	n += 8
	binary.BigEndian.PutUint32(buff[n:], self.QueueId)
	n += 4
	return n, nil
}

func (self *ActionEnqueue) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Unmarshal(buff); err != nil {
		return n, err
	}
	self.Port = binary.BigEndian.Uint16(buff[n:])
	n += 8
	self.QueueId = binary.BigEndian.Uint32(buff[n:])
	return n, nil
}

// ActionVlanVid is action for OFPAT_SET_VLAN_VID
type ActionVlanVid struct {
	ActionHeader
	VlanVid uint16
	pad     [2]byte
}

func NewActionVlanVid() *ActionVlanVid {
	return &ActionVlanVid{
		ActionHeader: ActionHeader{Type: OFPAT_SET_VLAN_VID, Length: 8},
	}
}

func (self *ActionVlanVid) Type() ActionType {
	return self.ActionHeader.Type
}

func (self *ActionVlanVid) Len() int {
	return int(self.ActionHeader.Length)
}

func (self *ActionVlanVid) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Marshal(buff); err != nil {
		return n, err
	}
	binary.BigEndian.PutUint16(buff[n:], self.VlanVid)
	n += 4
	return n, nil
}

func (self *ActionVlanVid) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Unmarshal(buff); err != nil {
		return n, err
	}
	self.VlanVid = binary.BigEndian.Uint16(buff[n:])
	n += 4
	return n, nil
}

// ActionVlanPcp is action for OFPAT_SET_VLAN_PCP
type ActionVlanPcp struct {
	ActionHeader
	VlanPcp uint8
	Pad     [3]byte
}

func NewActionVlanPcp() *ActionVlanPcp {
	return &ActionVlanPcp{
		ActionHeader: ActionHeader{Type: OFPAT_SET_VLAN_PCP, Length: 8},
	}
}

func (self *ActionVlanPcp) Type() ActionType {
	return self.ActionHeader.Type
}

func (self *ActionVlanPcp) Len() int {
	return int(self.ActionHeader.Length)
}

func (self *ActionVlanPcp) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Marshal(buff); err != nil {
		return n, err
	}
	buff[n] = self.VlanPcp
	n += 4
	return n, nil
}

func (self *ActionVlanPcp) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Unmarshal(buff); err != nil {
		return n, err
	}
	self.VlanPcp = buff[n]
	n += 4
	return n, nil
}

// ActionDlAddr is action for OFPAT_SET_DL_SRC/DST
type ActionDlAddr struct {
	ActionHeader
	Addr [OFP_ETH_ALAN]byte
	Pad  [6]byte
}

func newActionDlAddr(actionType ActionType) *ActionDlAddr {
	return &ActionDlAddr{
		ActionHeader: ActionHeader{Type: actionType, Length: 16},
	}
}

func NewActionDlSrc() *ActionDlAddr {
	return newActionDlAddr(OFPAT_SET_DL_SRC)
}

func NewActionDlDst() *ActionDlAddr {
	return newActionDlAddr(OFPAT_SET_DL_DST)
}

func (self *ActionDlAddr) Type() ActionType {
	return self.ActionHeader.Type
}

func (self *ActionDlAddr) Len() int {
	return int(self.ActionHeader.Length)
}

func (self *ActionDlAddr) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Marshal(buff); err != nil {
		return n, err
	}
	copy(buff[n:], self.Addr[:])
	n += OFP_ETH_ALAN + 6
	return n, nil
}

func (self *ActionDlAddr) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Unmarshal(buff); err != nil {
		return n, err
	}
	copy(self.Addr[:], buff[n:])
	n += OFP_ETH_ALAN + 6
	return n, nil
}

// ActionNwAddr is action for OFPAT_SET_NW_SRC/DST.
type ActionNwAddr struct {
	ActionHeader
	Addr uint32
}

func newActionNwAddr(actionType ActionType) *ActionNwAddr {
	return &ActionNwAddr{
		ActionHeader: ActionHeader{Type: actionType, Length: 8},
	}
}

func NewActionNwSrc() *ActionNwAddr {
	return newActionNwAddr(OFPAT_SET_NW_SRC)
}

func NewActionNwDst() *ActionNwAddr {
	return newActionNwAddr(OFPAT_SET_NW_DST)
}

func (self *ActionNwAddr) Type() ActionType {
	return self.ActionHeader.Type
}

func (self *ActionNwAddr) Len() int {
	return int(self.ActionHeader.Length)
}

func (self *ActionNwAddr) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Marshal(buff); err != nil {
		return n, err
	}
	binary.BigEndian.PutUint32(buff[n:], self.Addr)
	n += 4
	return n, nil
}

func (self *ActionNwAddr) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Unmarshal(buff); err != nil {
		return n, err
	}
	self.Addr = binary.BigEndian.Uint32(buff[n:])
	n += 4
	return n, nil
}

// ActionNwTos is action for OFPAT_SET_NW_TOS.
type ActionNwTos struct {
	ActionHeader
	Tos byte // IP ToS (DSCP field, 6 bits).
	Pad [3]byte
}

func NewActionNwTos() *ActionNwTos {
	return &ActionNwTos{
		ActionHeader: ActionHeader{Type: OFPAT_SET_NW_TOS, Length: 8},
	}
}

func (self *ActionNwTos) Type() ActionType {
	return self.ActionHeader.Type
}

func (self *ActionNwTos) Len() int {
	return int(self.ActionHeader.Length)
}

func (self *ActionNwTos) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Marshal(buff); err != nil {
		return n, err
	}
	buff[n] = self.Tos
	n += 4
	return n, nil
}

func (self *ActionNwTos) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Unmarshal(buff); err != nil {
		return n, err
	}
	self.Tos = buff[n]
	n += 4
	return n, nil
}

type ActionTpPort struct {
	ActionHeader
	Port uint16
	Pad  [2]byte
}

func newActionTpPort(actionType ActionType) *ActionTpPort {
	return &ActionTpPort{
		ActionHeader: ActionHeader{Type: actionType, Length: 8},
	}
}

func NewActionTpSrc() *ActionTpPort {
	return newActionTpPort(OFPAT_SET_TP_SRC)
}

func NewActionTpDst() *ActionTpPort {
	return newActionTpPort(OFPAT_SET_TP_DST)
}

func (self *ActionTpPort) Type() ActionType {
	return self.ActionHeader.Type
}

func (self *ActionTpPort) Len() int {
	return int(self.ActionHeader.Length)
}

func (self *ActionTpPort) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Marshal(buff); err != nil {
		return n, err
	}
	binary.BigEndian.PutUint16(buff[n:], self.Port)
	n += 4
	return n, nil
}

func (self *ActionTpPort) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Unmarshal(buff); err != nil {
		return n, err
	}
	self.Port = binary.BigEndian.Uint16(buff[n:])
	n += 4
	return n, nil
}

type ActionVendorHeader struct {
	ActionHeader
	Vendor uint32
}

func (self *ActionVendorHeader) Len() int {
	return int(self.ActionHeader.Length)
}

func (self *ActionVendorHeader) Marshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Marshal(buff); err != nil {
		return n, err
	}
	binary.BigEndian.PutUint32(buff[n:], self.Vendor)
	n += 4
	return n, nil
}

func (self *ActionVendorHeader) Unmarshal(buff []byte) (n int, err error) {
	if len(buff) < self.Len() {
		return 0, ofp.NewNoBuffError()
	}
	if n, err = self.ActionHeader.Unmarshal(buff); err != nil {
		return n, err
	}
	self.Vendor = binary.BigEndian.Uint32(buff[n:])
	n += 4
	return n, nil
}
