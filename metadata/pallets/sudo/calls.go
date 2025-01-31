package sudo

import (
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
