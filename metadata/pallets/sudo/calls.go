package sudo

import (
	"github.com/availproject/avail-go-sdk/metadata"
	. "github.com/availproject/avail-go-sdk/metadata/pallets"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

// Authenticates the sudo key and dispatches a function call with `Root` origin.
type CallSudo struct {
	Call prim.Call
}

func (this CallSudo) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSudo) PalletName() string {
	return PalletName
}

func (this CallSudo) CallIndex() uint8 {
	return 0
}

func (this CallSudo) CallName() string {
	return "sudo"
}

func (this *CallSudo) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallSudo) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallSudo) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}

// Authenticates the sudo key and dispatches a function call with `Root` origin.
// This function does not check the weight of the call, and instead allows the
// Sudo user to specify the weight of the call.
//
// The dispatch origin for this call must be _Signed_.
type CallSudoUncheckedWeight struct {
	Call prim.Call
}

func (this CallSudoUncheckedWeight) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSudoUncheckedWeight) PalletName() string {
	return PalletName
}

func (this CallSudoUncheckedWeight) CallIndex() uint8 {
	return 1
}

func (this CallSudoUncheckedWeight) CallName() string {
	return "sudo_unchecked_weight"
}

func (this *CallSudoUncheckedWeight) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallSudoUncheckedWeight) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallSudoUncheckedWeight) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}

// Authenticates the sudo key and dispatches a function call with `Signed` origin from
// a given account.
//
// The dispatch origin for this call must be _Signed_.
type CallSudoAs struct {
	Who  prim.MultiAddress
	Call prim.Call
}

func (this CallSudoAs) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSudoAs) PalletName() string {
	return PalletName
}

func (this CallSudoAs) CallIndex() uint8 {
	return 3
}

func (this CallSudoAs) CallName() string {
	return "sudo_as"
}

func (this *CallSudoAs) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallSudoAs) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallSudoAs) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}
