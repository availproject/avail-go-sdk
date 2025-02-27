package system

import (
	meta "github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/primitives"
)

// An extrinsic completed successfully.
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

// An extrinsic failed.
type EventExtrinsicFailed struct {
	DispatchError meta.DispatchError
	DispatchInfo  meta.DispatchInfo
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

// A new account was created.
type EventNewAccount struct {
	Account primitives.AccountId
}

func (this EventNewAccount) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventNewAccount) PalletName() string {
	return PalletName
}

func (this EventNewAccount) EventIndex() uint8 {
	return 3
}

func (this EventNewAccount) EventName() string {
	return "NewAccount"
}

// An account was reaped.
type EventKilledAccount struct {
	Account primitives.AccountId
}

func (this EventKilledAccount) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventKilledAccount) PalletName() string {
	return PalletName
}

func (this EventKilledAccount) EventIndex() uint8 {
	return 4
}

func (this EventKilledAccount) EventName() string {
	return "KilledAccount"
}

// On on-chain remark happened
type EventRemarked struct {
	Sender primitives.AccountId
	Hash   primitives.H256
}

func (this EventRemarked) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventRemarked) PalletName() string {
	return PalletName
}

func (this EventRemarked) EventIndex() uint8 {
	return 5
}

func (this EventRemarked) EventName() string {
	return "Remarked"
}
