package system

import (
	meta "github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/primitives"
)

// An extrinsic completed successfully.
type EventExtrinsicSuccess struct {
	DispatchInfo meta.DispatchInfo
}

func (ees EventExtrinsicSuccess) PalletIndex() uint8 {
	return PalletIndex
}

func (ees EventExtrinsicSuccess) PalletName() string {
	return PalletName
}

func (ees EventExtrinsicSuccess) EventIndex() uint8 {
	return 0
}

func (ees EventExtrinsicSuccess) EventName() string {
	return "ExtrinsicSuccess"
}

// An extrinsic failed.
type EventExtrinsicFailed struct {
	DispatchError meta.DispatchError
	DispatchInfo  meta.DispatchInfo
}

func (eef EventExtrinsicFailed) PalletIndex() uint8 {
	return PalletIndex
}

func (eef EventExtrinsicFailed) PalletName() string {
	return PalletName
}

func (eef EventExtrinsicFailed) EventIndex() uint8 {
	return 1
}

func (eef EventExtrinsicFailed) EventName() string {
	return "ExtrinsicFailed"
}

// A new account was created.
type EventNewAccount struct {
	Account primitives.AccountId
}

func (ena EventNewAccount) PalletIndex() uint8 {
	return PalletIndex
}

func (ena EventNewAccount) PalletName() string {
	return PalletName
}

func (ena EventNewAccount) EventIndex() uint8 {
	return 3
}

func (ena EventNewAccount) EventName() string {
	return "NewAccount"
}

// An account was reaped.
type EventKilledAccount struct {
	Account primitives.AccountId
}

func (eka EventKilledAccount) PalletIndex() uint8 {
	return PalletIndex
}

func (eka EventKilledAccount) PalletName() string {
	return PalletName
}

func (eka EventKilledAccount) EventIndex() uint8 {
	return 4
}

func (eka EventKilledAccount) EventName() string {
	return "KilledAccount"
}

// On on-chain remark happened
type EventRemarked struct {
	Sender primitives.AccountId
	Hash   primitives.H256
}

func (er EventRemarked) PalletIndex() uint8 {
	return PalletIndex
}

func (er EventRemarked) PalletName() string {
	return PalletName
}

func (er EventRemarked) EventIndex() uint8 {
	return 5
}

func (er EventRemarked) EventName() string {
	return "Remarked"
}
