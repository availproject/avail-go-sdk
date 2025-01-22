package utility

import (
	"go-sdk/metadata"
	. "go-sdk/metadata/pallets"
	prim "go-sdk/primitives"
)

// Do not add, remove or change any of the field members.
type CallBatch struct {
	Calls []prim.Call
}

func (this CallBatch) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallBatch) PalletName() string {
	return PalletName
}

func (this CallBatch) CallIndex() uint8 {
	return 0
}

func (this CallBatch) CallName() string {
	return "batch"
}

func (this *CallBatch) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallBatch) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallBatch) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}

func (this *CallBatch) AddCall(value prim.Call) {
	this.Calls = append(this.Calls, value)
}

func (this *CallBatch) AddPayload(value metadata.Payload) {
	this.Calls = append(this.Calls, value.Call)
}
