package vector

import (
	"github.com/availproject/avail-go-sdk/metadata"
	. "github.com/availproject/avail-go-sdk/metadata/pallets"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

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

func (this *CallSendMessage) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallSendMessage) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallSendMessage) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}
