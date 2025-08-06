package multisig

import (
	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

// Immediately dispatch a multi-signature call using a single approval from the caller.
//
// The dispatch origin for this call must be _Signed_.
//
// - `other_signatories`: The accounts (other than the sender) who are part of the
// multi-signature, but do not participate in the approval process.
// - `call`: The call to be executed.
//
// Result is equivalent to the dispatched result.
type CallAsMultiThreshold1 struct {
	OtherSignatories []prim.AccountId
	Call             prim.Call
}

func (camt CallAsMultiThreshold1) PalletIndex() uint8 {
	return PalletIndex
}

func (camt CallAsMultiThreshold1) PalletName() string {
	return PalletName
}

func (camt CallAsMultiThreshold1) CallIndex() uint8 {
	return 0
}

func (camt CallAsMultiThreshold1) CallName() string {
	return "as_multi_threshold_1"
}

// Register approval for a dispatch to be made from a deterministic composite account if
// approved by a total of `threshold - 1` of `other_signatories`.
//
// If there are enough, then dispatch the call.
//
// Payment: `DepositBase` will be reserved if this is the first approval, plus
// `threshold` times `DepositFactor`. It is returned once this dispatch happens or
// is cancelled.
//
// The dispatch origin for this call must be _Signed_.
//
// - `threshold`: The total number of approvals for this dispatch before it is executed.
// - `other_signatories`: The accounts (other than the sender) who can approve this
// dispatch. May not be empty.
// - `maybe_timepoint`: If this is the first approval, then this must be `None`. If it is
// not the first approval, then it must be `Some`, with the timepoint (block number and
// transaction index) of the first approval transaction.
// - `call`: The call to be executed.
//
// NOTE: Unless this is the final approval, you will generally want to use
// `approve_as_multi` instead, since it only requires a hash of the call.
//
// Result is equivalent to the dispatched result if `threshold` is exactly `1`. Otherwise
// on success, result is `Ok` and the result from the interior call, if it was executed,
// may be found in the deposited `MultisigExecuted` event.
type CallAsMulti struct {
	Threshold        uint16
	OtherSignatories []prim.AccountId
	MaybeTimepoint   prim.Option[metadata.TimepointBlockNumber]
	Call             prim.Call
	MaxWeight        metadata.Weight
}

func (cam CallAsMulti) PalletIndex() uint8 {
	return PalletIndex
}

func (cam CallAsMulti) PalletName() string {
	return PalletName
}

func (cam CallAsMulti) CallIndex() uint8 {
	return 1
}

func (cam CallAsMulti) CallName() string {
	return "as_multi"
}

// Register approval for a dispatch to be made from a deterministic composite account if
// approved by a total of `threshold - 1` of `other_signatories`.
//
// Payment: `DepositBase` will be reserved if this is the first approval, plus
// `threshold` times `DepositFactor`. It is returned once this dispatch happens or
// is cancelled.
//
// The dispatch origin for this call must be _Signed_.
//
// - `threshold`: The total number of approvals for this dispatch before it is executed.
// - `other_signatories`: The accounts (other than the sender) who can approve this
// dispatch. May not be empty.
// - `maybe_timepoint`: If this is the first approval, then this must be `None`. If it is
// not the first approval, then it must be `Some`, with the timepoint (block number and
// transaction index) of the first approval transaction.
// - `call_hash`: The hash of the call to be executed.
//
// NOTE: If this is the final approval, you will want to use `as_multi` instead.
type CallApproveAsMulti struct {
	Threshold        uint16
	OtherSignatories []prim.AccountId
	MaybeTimepoint   prim.Option[metadata.TimepointBlockNumber]
	CallHash         prim.H256
	MaxWeight        metadata.Weight
}

func (caam CallApproveAsMulti) PalletIndex() uint8 {
	return PalletIndex
}

func (caam CallApproveAsMulti) PalletName() string {
	return PalletName
}

func (caam CallApproveAsMulti) CallIndex() uint8 {
	return 2
}

func (caam CallApproveAsMulti) CallName() string {
	return "approve_as_multi"
}

// Cancel a pre-existing, on-going multisig transaction. Any deposit reserved previously
// for this operation will be unreserved on success.
//
// The dispatch origin for this call must be _Signed_.
//
// - `threshold`: The total number of approvals for this dispatch before it is executed.
// - `other_signatories`: The accounts (other than the sender) who can approve this
// dispatch. May not be empty.
// - `timepoint`: The timepoint (block number and transaction index) of the first approval
// transaction for this dispatch.
// - `call_hash`: The hash of the call to be executed.
type CallCancelAsMulti struct {
	Threshold        uint16
	OtherSignatories []prim.AccountId
	MaybeTimepoint   metadata.TimepointBlockNumber
	CallHash         prim.H256
}

func (ccam CallCancelAsMulti) PalletIndex() uint8 {
	return PalletIndex
}

func (ccam CallCancelAsMulti) PalletName() string {
	return PalletName
}

func (ccam CallCancelAsMulti) CallIndex() uint8 {
	return 3
}

func (ccam CallCancelAsMulti) CallName() string {
	return "cancel_as_multi"
}
