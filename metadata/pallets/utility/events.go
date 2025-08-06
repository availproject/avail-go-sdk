package utility

import (
	"github.com/availproject/avail-go-sdk/metadata"
)

// Batch of dispatches did not complete fully. Index of first failing dispatch given, as well as the error
type EventBatchInterrupted struct {
	Index uint32
	Error metadata.DispatchError
}

func (ebi EventBatchInterrupted) PalletIndex() uint8 {
	return PalletIndex
}

func (ebi EventBatchInterrupted) PalletName() string {
	return PalletName
}

func (ebi EventBatchInterrupted) EventIndex() uint8 {
	return 0
}

func (ebi EventBatchInterrupted) EventName() string {
	return "BatchInterrupted"
}

// Batch of dispatches completed fully with no error.
type EventBatchCompleted struct{}

func (ebc EventBatchCompleted) PalletIndex() uint8 {
	return PalletIndex
}

func (ebc EventBatchCompleted) PalletName() string {
	return PalletName
}

func (ebc EventBatchCompleted) EventIndex() uint8 {
	return 1
}

func (ebc EventBatchCompleted) EventName() string {
	return "BatchCompleted"
}

// Batch of dispatches completed but has errors.
type EventBatchCompletedWithErrors struct{}

func (ebcwe EventBatchCompletedWithErrors) PalletIndex() uint8 {
	return PalletIndex
}

func (ebcwe EventBatchCompletedWithErrors) PalletName() string {
	return PalletName
}

func (ebcwe EventBatchCompletedWithErrors) EventIndex() uint8 {
	return 2
}

func (ebcwe EventBatchCompletedWithErrors) EventName() string {
	return "BatchCompletedWithErrors"
}

// A single item within a Batch of dispatches has completed with no error.
type EventItemCompleted struct{}

func (eic EventItemCompleted) PalletIndex() uint8 {
	return PalletIndex
}

func (eic EventItemCompleted) PalletName() string {
	return PalletName
}

func (eic EventItemCompleted) EventIndex() uint8 {
	return 3
}

func (eic EventItemCompleted) EventName() string {
	return "ItemCompleted"
}

// A single item within a Batch of dispatches has completed with error.
type EventItemFailed struct {
	Error metadata.DispatchError
}

func (eif EventItemFailed) PalletIndex() uint8 {
	return PalletIndex
}

func (eif EventItemFailed) PalletName() string {
	return PalletName
}

func (eif EventItemFailed) EventIndex() uint8 {
	return 4
}

func (eif EventItemFailed) EventName() string {
	return "ItemFailed"
}

// A call was dispatched.
type EventDispatchedAs struct {
	Result metadata.DispatchResult
}

func (eda EventDispatchedAs) PalletIndex() uint8 {
	return PalletIndex
}

func (eda EventDispatchedAs) PalletName() string {
	return PalletName
}

func (eda EventDispatchedAs) EventIndex() uint8 {
	return 5
}

func (eda EventDispatchedAs) EventName() string {
	return "DispatchedAs"
}
