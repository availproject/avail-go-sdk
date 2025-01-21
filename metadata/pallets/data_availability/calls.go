package data_availability

import (
	metadata "github.com/nmvalera/avail-go-sdk/metadata"
	prim "github.com/nmvalera/avail-go-sdk/primitives"
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

func (this *CallCreateApplicationKey) ToPayload() metadata.Payload {
	var call = prim.Call{
		PalletIndex: this.PalletIndex(),
		CallIndex:   this.CallIndex(),
		Fields:      prim.AlreadyEncoded{Value: prim.Encoder.Encode(this)},
	}

	return metadata.NewPayload(call, this.PalletName(), this.CallName())
}

func (this *CallCreateApplicationKey) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	if this.PalletIndex() != tx.Call.PalletIndex {
		return false
	}

	if this.CallIndex() != tx.Call.CallIndex {
		return false
	}

	var decoder = prim.NewDecoder(tx.Call.Fields.ToBytes(), 0)
	decoder.Decode(this)
	return true
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

func (this *CallSubmitData) ToPayload() metadata.Payload {
	var call = prim.Call{
		PalletIndex: this.PalletIndex(),
		CallIndex:   this.CallIndex(),
		Fields:      prim.AlreadyEncoded{Value: prim.Encoder.Encode(this)},
	}

	return metadata.NewPayload(call, this.PalletName(), this.CallName())
}

func (this *CallSubmitData) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	if this.PalletIndex() != tx.Call.PalletIndex {
		return false
	}

	if this.CallIndex() != tx.Call.CallIndex {
		return false
	}

	var decoder = prim.NewDecoder(tx.Call.Fields.ToBytes(), 0)
	decoder.Decode(this)
	return true
}
