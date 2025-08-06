package proxy

import (
	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/primitives"
)

// A proxy was executed correctly, with the given.
type EventProxyExecuted struct {
	Result metadata.DispatchResult
}

func (epe EventProxyExecuted) PalletIndex() uint8 {
	return PalletIndex
}

func (epe EventProxyExecuted) PalletName() string {
	return PalletName
}

func (epe EventProxyExecuted) EventIndex() uint8 {
	return 0
}

func (epe EventProxyExecuted) EventName() string {
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

func (epc EventPureCreated) PalletIndex() uint8 {
	return PalletIndex
}

func (epc EventPureCreated) PalletName() string {
	return PalletName
}

func (epc EventPureCreated) EventIndex() uint8 {
	return 1
}

func (epc EventPureCreated) EventName() string {
	return "PureCreated"
}

// An announcement was placed to make a call in the future.
type EventAnnounced struct {
	Real     primitives.AccountId
	Proxy    primitives.AccountId
	CallHash primitives.H256
}

func (ea EventAnnounced) PalletIndex() uint8 {
	return PalletIndex
}

func (ea EventAnnounced) PalletName() string {
	return PalletName
}

func (ea EventAnnounced) EventIndex() uint8 {
	return 2
}

func (ea EventAnnounced) EventName() string {
	return "Announced"
}

// A proxy was added.
type EventProxyAdded struct {
	Delegator primitives.AccountId
	Delegatee primitives.AccountId
	ProxyTpe  metadata.ProxyType
	Delay     uint32
}

func (epa EventProxyAdded) PalletIndex() uint8 {
	return PalletIndex
}

func (epa EventProxyAdded) PalletName() string {
	return PalletName
}

func (epa EventProxyAdded) EventIndex() uint8 {
	return 3
}

func (epa EventProxyAdded) EventName() string {
	return "ProxyAdded"
}

// A proxy was removed.
type EventProxyRemoved struct {
	Delegator primitives.AccountId
	Delegatee primitives.AccountId
	ProxyTpe  metadata.ProxyType
	Delay     uint32
}

func (epr EventProxyRemoved) PalletIndex() uint8 {
	return PalletIndex
}

func (epr EventProxyRemoved) PalletName() string {
	return PalletName
}

func (epr EventProxyRemoved) EventIndex() uint8 {
	return 4
}

func (epr EventProxyRemoved) EventName() string {
	return "ProxyRemoved"
}
