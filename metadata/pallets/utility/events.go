package utility

import (
	"github.com/availproject/avail-go-sdk/metadata"
)

// Batch of dispatches did not complete fully. Index of first failing dispatch given, as well as the error
type EventBatchInterrupted struct {
	Index uint32
	Error metadata.DispatchError
}

func (this EventBatchInterrupted) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventBatchInterrupted) PalletName() string {
	return PalletName
}

func (this EventBatchInterrupted) EventIndex() uint8 {
	return 0
}

func (this EventBatchInterrupted) EventName() string {
	return "BatchInterrupted"
}

// Batch of dispatches completed fully with no error.
type EventBatchCompleted struct{}

func (this EventBatchCompleted) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventBatchCompleted) PalletName() string {
	return PalletName
}

func (this EventBatchCompleted) EventIndex() uint8 {
	return 1
}

func (this EventBatchCompleted) EventName() string {
	return "BatchCompleted"
}

// Batch of dispatches completed but has errors.
type EventBatchCompletedWithErrors struct{}

func (this EventBatchCompletedWithErrors) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventBatchCompletedWithErrors) PalletName() string {
	return PalletName
}

func (this EventBatchCompletedWithErrors) EventIndex() uint8 {
	return 2
}

func (this EventBatchCompletedWithErrors) EventName() string {
	return "BatchCompletedWithErrors"
}

// A single item within a Batch of dispatches has completed with no error.
type EventItemCompleted struct{}

func (this EventItemCompleted) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventItemCompleted) PalletName() string {
	return PalletName
}

func (this EventItemCompleted) EventIndex() uint8 {
	return 3
}

func (this EventItemCompleted) EventName() string {
	return "ItemCompleted"
}

// A single item within a Batch of dispatches has completed with error.
type EventItemFailed struct {
	Error metadata.DispatchError
}

func (this EventItemFailed) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventItemFailed) PalletName() string {
	return PalletName
}

func (this EventItemFailed) EventIndex() uint8 {
	return 4
}

func (this EventItemFailed) EventName() string {
	return "ItemFailed"
}

// A call was dispatched.
type EventDispatchedAs struct {
	Result metadata.DispatchResult
}

func (this EventDispatchedAs) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventDispatchedAs) PalletName() string {
	return PalletName
}

func (this EventDispatchedAs) EventIndex() uint8 {
	return 5
}

func (this EventDispatchedAs) EventName() string {
	return "DispatchedAs"
}
