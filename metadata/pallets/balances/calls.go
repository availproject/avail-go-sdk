package balances

import (
	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"

	"github.com/itering/scale.go/utiles/uint128"
)

// Transfer some liquid free balance to another account.
//
// `transfer_allow_death` will set the `FreeBalance` of the sender and receiver.
// If the sender's account is below the existential deposit as a result
// of the transfer, the account will be reaped.
//
// The dispatch origin for this call must be `Signed` by the transactor.
type CallTransferAlowDeath struct {
	Dest  prim.MultiAddress
	Value uint128.Uint128 `scale:"compact"`
}

func (ctad CallTransferAlowDeath) PalletIndex() uint8 {
	return PalletIndex
}

func (ctad CallTransferAlowDeath) PalletName() string {
	return PalletName
}

func (ctad CallTransferAlowDeath) CallIndex() uint8 {
	return 0
}

func (ctad CallTransferAlowDeath) CallName() string {
	return "transfer_allow_death"
}

// Exactly as `TransferAlowDeath`, except the origin must be root and the source account
// may be specified.
type CallForceTransfer struct {
	Source prim.MultiAddress
	Dest   prim.MultiAddress
	Value  metadata.Balance `scale:"compact"`
}

func (cft CallForceTransfer) PalletIndex() uint8 {
	return PalletIndex
}

func (cft CallForceTransfer) PalletName() string {
	return PalletName
}

func (cft CallForceTransfer) CallIndex() uint8 {
	return 2
}

func (cft CallForceTransfer) CallName() string {
	return "force_transfer"
}

// Same as the `TransferAlowDeath` call, but with a check that the transfer will not
// kill the origin account.
type CallTransferKeepAlive struct {
	Dest  prim.MultiAddress
	Value metadata.Balance `scale:"compact"`
}

func (ctka CallTransferKeepAlive) PalletIndex() uint8 {
	return PalletIndex
}

func (ctka CallTransferKeepAlive) PalletName() string {
	return PalletName
}

func (ctka CallTransferKeepAlive) CallIndex() uint8 {
	return 3
}

func (ctka CallTransferKeepAlive) CallName() string {
	return "transfer_keep_alive"
}

// Transfer the entire transferable balance from the caller account.
//
// NOTE: This function only attempts to transfer _transferable_ balances. This means that
// any locked, reserved, or existential deposits (when `keep_alive` is `true`), will not be
// transferred by this function.
type CallTransferAll struct {
	Dest      prim.MultiAddress
	KeepAlive bool
}

func (cta CallTransferAll) PalletIndex() uint8 {
	return PalletIndex
}

func (cta CallTransferAll) PalletName() string {
	return PalletName
}

func (cta CallTransferAll) CallIndex() uint8 {
	return 4
}

func (cta CallTransferAll) CallName() string {
	return "transfer_all"
}
