package utility

import (
	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

// Send a batch of dispatch calls.
//
// May be called from any origin except `None`.
type CallBatch struct {
	Calls []prim.Call
}

func (c CallBatch) PalletIndex() uint8 {
	return PalletIndex
}

func (c CallBatch) PalletName() string {
	return PalletName
}

func (c CallBatch) CallIndex() uint8 {
	return 0
}

func (c CallBatch) CallName() string {
	return "batch"
}

func (c *CallBatch) AddCall(value prim.Call) {
	c.Calls = append(c.Calls, value)
}

func (c *CallBatch) AddPayload(value metadata.Payload) {
	c.Calls = append(c.Calls, value.Call)
}

// Send a call through an indexed pseudonym of the sender.
//
// Filter from origin are passed along. The call will be dispatched with an origin which
// use the same filter as the origin of this call.
//
// NOTE: If you need to ensure that any account-based filtering is not honored (i.e.
// because you expect `proxy` to have been used prior in the call stack and you do not want
// the call restrictions to apply to any sub-accounts), then use `as_multi_threshold_1`
// in the Multisig pallet instead.
type CallAsDerivate struct {
	Index uint16
	Call  prim.Call
}

func (c CallAsDerivate) PalletIndex() uint8 {
	return PalletIndex
}

func (c CallAsDerivate) PalletName() string {
	return PalletName
}

func (c CallAsDerivate) CallIndex() uint8 {
	return 1
}

func (c CallAsDerivate) CallName() string {
	return "AsDerivate"
}

// Send a batch of dispatch calls and atomically execute them.
// The whole transaction will rollback and fail if any of the calls failed.
//
// May be called from any origin except `None`.
type CallBatchAll struct {
	Calls []prim.Call
}

func (c CallBatchAll) PalletIndex() uint8 {
	return PalletIndex
}

func (c CallBatchAll) PalletName() string {
	return PalletName
}

func (c CallBatchAll) CallIndex() uint8 {
	return 2
}

func (c CallBatchAll) CallName() string {
	return "BatchAll"
}

// Send a batch of dispatch calls.
// Unlike `batch`, it allows errors and won't interrupt.
//
// May be called from any origin except `None`.
type CallForceBatch struct {
	Calls []prim.Call
}

func (c CallForceBatch) PalletIndex() uint8 {
	return PalletIndex
}

func (c CallForceBatch) PalletName() string {
	return PalletName
}

func (c CallForceBatch) CallIndex() uint8 {
	return 4
}

func (c CallForceBatch) CallName() string {
	return "ForceBatch"
}
