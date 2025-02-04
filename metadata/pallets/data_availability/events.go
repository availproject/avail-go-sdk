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

func (this EventApplicationKeyCreated) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventApplicationKeyCreated) PalletName() string {
	return PalletName
}

func (this EventApplicationKeyCreated) EventIndex() uint8 {
	return 0
}

func (this EventApplicationKeyCreated) EventName() string {
	return "ApplicationKeyCreated"
}

// New Data was submitted.
type EventDataSubmitted struct {
	Who      prim.AccountId
	DataHash prim.H256
}

func (this EventDataSubmitted) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventDataSubmitted) PalletName() string {
	return PalletName
}

func (this EventDataSubmitted) EventIndex() uint8 {
	return 1
}

func (this EventDataSubmitted) EventName() string {
	return "DataSubmitted"
}
