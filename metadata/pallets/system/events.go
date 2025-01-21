package system

import (
	meta "github.com/availproject/avail-go-sdk/metadata"
)

// Do not add, remove or change any of the field members.
type EventExtrinsicSuccess struct {
	DispatchInfo meta.DispatchInfo
}

func (this EventExtrinsicSuccess) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventExtrinsicSuccess) PalletName() string {
	return PalletName
}

func (this EventExtrinsicSuccess) EventIndex() uint8 {
	return 0
}

func (this EventExtrinsicSuccess) EventName() string {
	return "ExtrinsicSuccess"
}

// Do not add, remove or change any of the field members.
type EventExtrinsicFailed struct {
	DispatchInfo meta.DispatchInfo
}

func (this EventExtrinsicFailed) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventExtrinsicFailed) PalletName() string {
	return PalletName
}

func (this EventExtrinsicFailed) EventIndex() uint8 {
	return 1
}

func (this EventExtrinsicFailed) EventName() string {
	return "ExtrinsicFailed"
}
