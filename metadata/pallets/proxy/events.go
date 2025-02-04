package proxy

import (
	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/primitives"
)

// A proxy was executed correctly, with the given.
type EventProxyExecuted struct {
	Result metadata.DispatchResult
}

func (this EventProxyExecuted) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventProxyExecuted) PalletName() string {
	return PalletName
}

func (this EventProxyExecuted) EventIndex() uint8 {
	return 0
}

func (this EventProxyExecuted) EventName() string {
	return "ProxyExecuted"
}

// A pure account has been created by new proxy with given
// disambiguation index and proxy type.
type EventPureCreated struct {
	Pure                primitives.AccountId
	Who                 primitives.AccountId
	ProxyTpe            metadata.ProxyType
	DisambiguationIndex uint16
}

func (this EventPureCreated) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventPureCreated) PalletName() string {
	return PalletName
}

func (this EventPureCreated) EventIndex() uint8 {
	return 1
}

func (this EventPureCreated) EventName() string {
	return "PureCreated"
}

// An announcement was placed to make a call in the future.
type EventAnnounced struct {
	Real     primitives.AccountId
	Proxy    primitives.AccountId
	CallHash primitives.H256
}

func (this EventAnnounced) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventAnnounced) PalletName() string {
	return PalletName
}

func (this EventAnnounced) EventIndex() uint8 {
	return 2
}

func (this EventAnnounced) EventName() string {
	return "Announced"
}

// A proxy was added.
type EventProxyAdded struct {
	Delegator primitives.AccountId
	Delegatee primitives.AccountId
	ProxyTpe  metadata.ProxyType
	Delay     uint32
}

func (this EventProxyAdded) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventProxyAdded) PalletName() string {
	return PalletName
}

func (this EventProxyAdded) EventIndex() uint8 {
	return 3
}

func (this EventProxyAdded) EventName() string {
	return "ProxyAdded"
}

// A proxy was removed.
type EventProxyRemoved struct {
	Delegator primitives.AccountId
	Delegatee primitives.AccountId
	ProxyTpe  metadata.ProxyType
	Delay     uint32
}

func (this EventProxyRemoved) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventProxyRemoved) PalletName() string {
	return PalletName
}

func (this EventProxyRemoved) EventIndex() uint8 {
	return 4
}

func (this EventProxyRemoved) EventName() string {
	return "ProxyRemoved"
}
