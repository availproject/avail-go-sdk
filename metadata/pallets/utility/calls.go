package utility

import (
	"github.com/availproject/avail-go-sdk/metadata"
	. "github.com/availproject/avail-go-sdk/metadata/pallets"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

// Send a batch of dispatch calls.
//
// May be called from any origin except `None`.
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

func (this CallAsDerivate) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallAsDerivate) PalletName() string {
	return PalletName
}

func (this CallAsDerivate) CallIndex() uint8 {
	return 1
}

func (this CallAsDerivate) CallName() string {
	return "AsDerivate"
}

func (this *CallAsDerivate) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallAsDerivate) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallAsDerivate) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}

// Send a batch of dispatch calls and atomically execute them.
// The whole transaction will rollback and fail if any of the calls failed.
//
// May be called from any origin except `None`.
type CallBatchAll struct {
	Calls []prim.Call
}

func (this CallBatchAll) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallBatchAll) PalletName() string {
	return PalletName
}

func (this CallBatchAll) CallIndex() uint8 {
	return 2
}

func (this CallBatchAll) CallName() string {
	return "BatchAll"
}

func (this *CallBatchAll) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallBatchAll) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallBatchAll) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}

// Send a batch of dispatch calls.
// Unlike `batch`, it allows errors and won't interrupt.
//
// May be called from any origin except `None`.
type CallForceBatch struct {
	Calls []prim.Call
}

func (this CallForceBatch) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallForceBatch) PalletName() string {
	return PalletName
}

func (this CallForceBatch) CallIndex() uint8 {
	return 4
}

func (this CallForceBatch) CallName() string {
	return "ForceBatch"
}

func (this *CallForceBatch) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallForceBatch) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallForceBatch) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}
