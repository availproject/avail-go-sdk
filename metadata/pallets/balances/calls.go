package balances

import (
	"go-sdk/metadata"
	. "go-sdk/metadata/pallets"
	prim "go-sdk/primitives"

	"github.com/itering/scale.go/utiles/uint128"
)

// Do not add, remove or change any of the field members.
type CallTransferAlowDeath struct {
	Dest  prim.MultiAddress
	Value uint128.Uint128 `scale:"compact"`
}

func (this CallTransferAlowDeath) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallTransferAlowDeath) PalletName() string {
	return PalletName
}

func (this CallTransferAlowDeath) CallIndex() uint8 {
	return 0
}

func (this CallTransferAlowDeath) CallName() string {
	return "transfer_allow_death"
}

func (this *CallTransferAlowDeath) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallTransferAlowDeath) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallTransferAlowDeath) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	if this.PalletIndex() != tx.Call.PalletIndex {
		return false
	}

	if this.CallIndex() != tx.Call.CallIndex {
		return false
	}

	var bytes = tx.Call.Fields.ToBytes()
	var decoder = prim.NewDecoder(bytes, 0)
	decoder.Decode(this)
	return true
}

// Do not add, remove or change any of the field members.
type CallForceTransfer struct {
	Source prim.MultiAddress
	Dest   prim.MultiAddress
	Value  uint128.Uint128 `scale:"compact"`
}

func (this CallForceTransfer) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallForceTransfer) PalletName() string {
	return PalletName
}

func (this CallForceTransfer) CallIndex() uint8 {
	return 2
}

func (this CallForceTransfer) CallName() string {
	return "force_transfer"
}

func (this *CallForceTransfer) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallForceTransfer) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallForceTransfer) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
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
type CallTransferKeepAlive struct {
	Dest  prim.MultiAddress
	Value uint128.Uint128 `scale:"compact"`
}

func (this CallTransferKeepAlive) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallTransferKeepAlive) PalletName() string {
	return PalletName
}

func (this CallTransferKeepAlive) CallIndex() uint8 {
	return 3
}

func (this CallTransferKeepAlive) CallName() string {
	return "transfer_keep_alive"
}

func (this *CallTransferKeepAlive) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallTransferKeepAlive) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallTransferKeepAlive) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}

// Do not add, remove or change any of the field members.
type CallTransferAll struct {
	Dest      prim.MultiAddress
	KeepAlive bool
}

func (this CallTransferAll) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallTransferAll) PalletName() string {
	return PalletName
}

func (this CallTransferAll) CallIndex() uint8 {
	return 4
}

func (this CallTransferAll) CallName() string {
	return "transfer_all"
}

func (this *CallTransferAll) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallTransferAll) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallTransferAll) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
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
