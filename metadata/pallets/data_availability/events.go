package data_availability

import (
	prim "github.com/availproject/avail-go-sdk/primitives"
)

// A new application key was created.
type EventApplicationKeyCreated struct {
	Key   []uint8
	Owner prim.AccountId
	Id    uint32 `scale:"compact"`
}

func (eakc EventApplicationKeyCreated) PalletIndex() uint8 {
	return PalletIndex
}

func (eakc EventApplicationKeyCreated) PalletName() string {
	return PalletName
}

func (eakc EventApplicationKeyCreated) EventIndex() uint8 {
	return 0
}

func (eakc EventApplicationKeyCreated) EventName() string {
	return "ApplicationKeyCreated"
}

// New Data was submitted.
type EventDataSubmitted struct {
	Who      prim.AccountId
	DataHash prim.H256
}

func (eds EventDataSubmitted) PalletIndex() uint8 {
	return PalletIndex
}

func (eds EventDataSubmitted) PalletName() string {
	return PalletName
}

func (eds EventDataSubmitted) EventIndex() uint8 {
	return 1
}

func (eds EventDataSubmitted) EventName() string {
	return "DataSubmitted"
}
