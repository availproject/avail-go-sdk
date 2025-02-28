package vector

import (
	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type CallFulfillCall struct {
	FunctionId prim.H256
	Input      []byte
	output     []byte
	proof      []byte
	slot       uint64 `scale:"compact"`
}

func (this CallFulfillCall) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallFulfillCall) PalletName() string {
	return PalletName
}

func (this CallFulfillCall) CallIndex() uint8 {
	return 0
}

func (this CallFulfillCall) CallName() string {
	return "fulfill_call"
}

// Send a batch of dispatch calls.
//
// May be called from any origin except `None`.
type CallSendMessage struct {
	Message metadata.VectorMessageKind
	To      prim.H256
	Domain  uint32 `scale:"compact"`
}

func (this CallSendMessage) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSendMessage) PalletName() string {
	return PalletName
}

func (this CallSendMessage) CallIndex() uint8 {
	return 3
}

func (this CallSendMessage) CallName() string {
	return "send_message"
}
