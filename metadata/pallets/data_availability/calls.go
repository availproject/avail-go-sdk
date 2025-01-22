package data_availability

import (
	"go-sdk/metadata"
	. "go-sdk/metadata/pallets"
	prim "go-sdk/primitives"
)

// Do not add, remove or change any of the field members.
type CallCreateApplicationKey struct {
	Key []uint8
}

func (this CallCreateApplicationKey) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallCreateApplicationKey) PalletName() string {
	return PalletName
}

func (this CallCreateApplicationKey) CallIndex() uint8 {
	return 0
}

func (this CallCreateApplicationKey) CallName() string {
	return "create_application_key"
}

func (this *CallCreateApplicationKey) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallCreateApplicationKey) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallCreateApplicationKey) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}

// Do not add, remove or change any of the field members.
type CallSubmitData struct {
	Data []uint8
}

func (this CallSubmitData) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSubmitData) PalletName() string {
	return PalletName
}

func (this CallSubmitData) CallIndex() uint8 {
	return 1
}

func (this CallSubmitData) CallName() string {
	return "submit_data"
}

func (this *CallSubmitData) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallSubmitData) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallSubmitData) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}
