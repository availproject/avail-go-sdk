package proxy

import (
	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

// Dispatch the given `call` from an account that the sender is authorised for through
// `add_proxy`.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Real`: The account that the proxy will make a call on behalf of.
// - `ForceProxyType`: Specify the exact proxy type to be used and checked for this call.
// - `Call`: The call to be made by the `real` account.
type CallProxy struct {
	Real           prim.MultiAddress
	ForceProxyType prim.Option[metadata.ProxyType]
	Call           prim.Call
}

func (this CallProxy) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallProxy) PalletName() string {
	return PalletName
}

func (this CallProxy) CallIndex() uint8 {
	return 0
}

func (this CallProxy) CallName() string {
	return "proxy"
}

// Register a proxy account for the sender that is able to make calls on its behalf.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Delegate`: The account that the `caller` would like to make a proxy.
// - `ProxyType`: The permissions allowed for this proxy account.
// - `Delay`: The announcement period required of the initial proxy. Will generally be
// zero.
type CallAddProxy struct {
	Delegate  prim.MultiAddress
	ProxyType metadata.ProxyType
	Delay     uint32
}

func (this CallAddProxy) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallAddProxy) PalletName() string {
	return PalletName
}

func (this CallAddProxy) CallIndex() uint8 {
	return 1
}

func (this CallAddProxy) CallName() string {
	return "add_proxy"
}

// Unregister a proxy account for the sender.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Delegate`: The account that the `caller` would like to remove as a proxy.
// - `ProxyType`: The permissions currently enabled for the removed proxy account.
// - `Delay`:  Will generally be zero.
type CallRemoveProxy struct {
	Delegate  prim.MultiAddress
	ProxyType metadata.ProxyType
	Delay     uint32
}

func (this CallRemoveProxy) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallRemoveProxy) PalletName() string {
	return PalletName
}

func (this CallRemoveProxy) CallIndex() uint8 {
	return 2
}

func (this CallRemoveProxy) CallName() string {
	return "remove_proxy"
}

// Unregister all proxy accounts for the sender.
//
// The dispatch origin for this call must be _Signed_.
//
// WARNING: This may be called on accounts created by `pure`, however if done, then
// the unreserved fees will be inaccessible. **All access to this account will be lost.**
type CallRemoveProxies struct{}

func (this CallRemoveProxies) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallRemoveProxies) PalletName() string {
	return PalletName
}

func (this CallRemoveProxies) CallIndex() uint8 {
	return 3
}

func (this CallRemoveProxies) CallName() string {
	return "remove_proxies"
}

// Spawn a fresh new account that is guaranteed to be otherwise inaccessible, and
// initialize it with a proxy of `proxy_type` for `origin` sender.
//
// Requires a `Signed` origin.
//
// - `ProxyType`: The type of the proxy that the sender will be registered as over the
// new account. This will almost always be the most permissive `ProxyType` possible to
// allow for maximum flexibility.
// - `Index`: A disambiguation index, in case this is called multiple times in the same
// transaction (e.g. with `utility::batch`). Unless you're using `batch` you probably just
// want to use `0`.
// - `Delay`: The announcement period required of the initial proxy. Will generally be
// zero.
//
// Fails with `Duplicate` if this has already been called in this transaction, from the
// same sender, with the same parameters.
//
// Fails if there are insufficient funds to pay for deposit.
type CallCreatePure struct {
	ProxyType metadata.ProxyType
	Delay     uint32
	Index     uint16
}

func (this CallCreatePure) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallCreatePure) PalletName() string {
	return PalletName
}

func (this CallCreatePure) CallIndex() uint8 {
	return 4
}

func (this CallCreatePure) CallName() string {
	return "create_pure"
}

// Removes a previously spawned pure proxy.
//
// WARNING: **All access to this account will be lost.** Any funds held in it will be
// inaccessible.
//
// Requires a `Signed` origin, and the sender account must have been created by a call to
// `pure` with corresponding parameters.
//
// - `Spawner`: The account that originally called `pure` to create this account.
// - `Index`: The disambiguation index originally passed to `pure`. Probably `0`.
// - `ProxyType`: The proxy type originally passed to `pure`.
// - `Height`: The height of the chain when the call to `pure` was processed.
// - `ExtIndex`: The extrinsic index in which the call to `pure` was processed.
//
// Fails with `NoPermission` in case the caller is not a previously created pure
// account whose `pure` call has corresponding parameters.
type CallKillPure struct {
	Spawner   prim.MultiAddress
	ProxyType metadata.ProxyType
	Index     uint16
	Height    uint32 `scale:"compact"`
	ExtIndex  uint32 `scale:"compact"`
}

func (this CallKillPure) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallKillPure) PalletName() string {
	return PalletName
}

func (this CallKillPure) CallIndex() uint8 {
	return 5
}

func (this CallKillPure) CallName() string {
	return "kill_pure"
}

// Publish the hash of a proxy-call that will be made in the future.
//
// This must be called some number of blocks before the corresponding `proxy` is attempted
// if the delay associated with the proxy relationship is greater than zero.
//
// No more than `MaxPending` announcements may be made at any one time.
//
// This will take a deposit of `AnnouncementDepositFactor` as well as
// `AnnouncementDepositBase` if there are no other pending announcements.
//
// The dispatch origin for this call must be _Signed_ and a proxy of `real`.
//
// Parameters:
// - `Real`: The account that the proxy will make a call on behalf of.
// - `CallHash`: The hash of the call to be made by the `real` account.
type CallAnnounce struct {
	Real     prim.MultiAddress
	CallHash prim.H256
}

func (this CallAnnounce) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallAnnounce) PalletName() string {
	return PalletName
}

func (this CallAnnounce) CallIndex() uint8 {
	return 6
}

func (this CallAnnounce) CallName() string {
	return "announce"
}

// Remove a given announcement.
//
// May be called by a proxy account to remove a call they previously announced and return
// the deposit.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Real`: The account that the proxy will make a call on behalf of.
// - `CallHash`: The hash of the call to be made by the `real` account.
type CallRemoveAnnouncement struct {
	Real     prim.MultiAddress
	CallHash prim.H256
}

func (this CallRemoveAnnouncement) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallRemoveAnnouncement) PalletName() string {
	return PalletName
}

func (this CallRemoveAnnouncement) CallIndex() uint8 {
	return 7
}

func (this CallRemoveAnnouncement) CallName() string {
	return "remove_announcement"
}

// Remove the given announcement of a delegate.
//
// May be called by a target (proxied) account to remove a call that one of their delegates
// (`delegate`) has announced they want to execute. The deposit is returned.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Delegate`: The account that previously announced the call.
// - `CallHash`: The hash of the call to be made.
type CallRejectAnnouncement struct {
	Delegate prim.MultiAddress
	CallHash prim.H256
}

func (this CallRejectAnnouncement) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallRejectAnnouncement) PalletName() string {
	return PalletName
}

func (this CallRejectAnnouncement) CallIndex() uint8 {
	return 8
}

func (this CallRejectAnnouncement) CallName() string {
	return "reject_announcement"
}

// Dispatch the given `call` from an account that the sender is authorized for through
// `add_proxy`.
//
// Removes any corresponding announcement(s).
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Real`: The account that the proxy will make a call on behalf of.
// - `ForceProxyType`: Specify the exact proxy type to be used and checked for this call.
// - `Call`: The call to be made by the `real` account.
type CallProxyAnnounced struct {
	Delegate       prim.MultiAddress
	Real           prim.MultiAddress
	ForceProxyType prim.Option[metadata.ProxyType]
	Call           prim.Call
}

func (this CallProxyAnnounced) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallProxyAnnounced) PalletName() string {
	return PalletName
}

func (this CallProxyAnnounced) CallIndex() uint8 {
	return 9
}

func (this CallProxyAnnounced) CallName() string {
	return "proxy_announced"
}
