package examples

import (
	"fmt"
	"math/rand/v2"

	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/metadata/pallets"
	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	pxPallet "github.com/availproject/avail-go-sdk/metadata/pallets/proxy"
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

// Understanding Proxy Account:
//
// From a simplified perspective, proxy account acts as a family member that we trust.
//
// If we designate the name "Child" to our main account, then te proxy account will
// be designated with the name "Parent". The "Parent" can do anything in the name of "Child"
// while the "Child" cannot do anything in the name of "Parent".
//
// This means that the "Parent", if he wants, can take all the funds from "Child" if he thinks
// that the "Child" has misbehaved, or the "Parent" can force the "Child" to work as a validator
// in order to gather some funds.
//
// The "Child" can, if needed, break this bond and force the "Parent" to become childless.
// The "Parent" can, if needed, break this bond and force the "Child" to become fatherless (an orphan).
//
// There are different types of Proxy account with different levels of authority.
//
// Pure Proxy accounts work in the opposite direction. If we designate the name "Parent" to our main account,
// then the proxy account will be designated with the name "Child". The "Child" cannot do anything on it's own
// and the "Parent" has full control over it. The "Parent" can forget that he even had the "Child" in the first
// place thus forcing the "Child" to become an orphan that no will ever see again.

func RunProxy() {
	RunNormalProxy()
	RunPureProxy()

	fmt.Println("RunProxy finished correctly.")
}

func RunNormalProxy() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Bob is the "Parent" account
	parent := SDK.Account.Bob()
	parentMulti := metadata.NewAccountIdFromKeyPair(parent).ToMultiAddress()
	// Ferdie is the "Child" account
	child := SDK.Account.Ferdie()

	// Let's create a proxy and make Bob the "Parent" and Ferdie the "Child"
	proxyType := metadata.ProxyType{VariantIndex: 0} // Any Proxy
	tx := sdk.Tx.Proxy.AddProxy(parentMulti, proxyType, 0)
	res, err := tx.ExecuteAndWatchInclusion(child, SDK.NewTransactionOptions())
	PanicOnError(err)

	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction has to succeed")
	AssertTrue(res.Events.IsSome(), "Events need to be decodable")

	// Finding the Proxy Added Event
	// In production you would use `EventFindFirstChecked`` instead of `EventFindFirst`.
	event := SDK.EventFindFirst(res.Events.Unwrap(), pxPallet.EventProxyAdded{}).UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Delegatee: %v, Delegator %v, ProxyTpe %v, Delay: %v`, event.Delegatee.ToHuman(), event.Delegator.ToHuman(), event.ProxyTpe.ToHuman(), event.Delay))

	// The "Child" account misbehaved so "Parent" decided to take his allowance. 1.0 Avail will be "taken"
	realAddress := metadata.NewAccountIdFromKeyPair(child).ToMultiAddress()
	call := pallets.ToCall(baPallet.CallTransferKeepAlive{Dest: parentMulti, Value: SDK.OneAvail()})

	tx = sdk.Tx.Proxy.Proxy(realAddress, primitives.NewNone[metadata.ProxyType](), call)
	res, err = tx.ExecuteAndWatchInclusion(parent, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction has to succeed")

	// Checking for EventProxyExecuted event.
	// In production you would use `EventFindFirstChecked`` instead of `EventFindFirst`.
	event2 := SDK.EventFindFirst(res.Events.Unwrap(), pxPallet.EventProxyExecuted{}).UnsafeUnwrap()
	AssertEq(event2.Result.VariantIndex, 0, "Event must OK")
	fmt.Println(fmt.Sprintf(`Dispatch Result: %v`, event2.Result.ToString()))

	// "Child" is angry at "Parent" so he decides to cut all the dies with him
	tx = sdk.Tx.Proxy.RemoveProxy(parentMulti, proxyType, 0)
	res, err = tx.ExecuteAndWatchInclusion(child, SDK.NewTransactionOptions())
	PanicOnError(err)

	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction has to succeed")
	AssertTrue(res.Events.IsSome(), "Events need to be decodable")

	// Checking for EventProxyExecuted event.
	// In production you would use `EventFindFirstChecked`` instead of `EventFindFirst`.
	event3 := SDK.EventFindFirst(res.Events.Unwrap(), pxPallet.EventProxyRemoved{}).UnsafeUnwrap()
	AssertEq(event2.Result.VariantIndex, 0, "Event must OK")
	fmt.Println(fmt.Sprintf(`Delegatee: %v, Delegator %v, ProxyTpe %v, Delay: %v`, event3.Delegatee.ToHuman(), event3.Delegator.ToHuman(), event3.ProxyTpe.ToHuman(), event3.Delay))

	fmt.Println("RunNormalProxy finished correctly.")
}

func RunPureProxy() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Bob is the "Parent" account
	parent := SDK.Account.Bob()

	// Creating Pure Proxy
	proxyType := metadata.ProxyType{VariantIndex: 0} // Any Proxy
	index := uint16(0)
	tx := sdk.Tx.Proxy.CreatePure(proxyType, 0, index)
	res, err := tx.ExecuteAndWatchInclusion(parent, SDK.NewTransactionOptions())
	PanicOnError(err)

	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction has to succeed")
	AssertTrue(res.Events.IsSome(), "Events need to be decodable")

	// Finding the Pure Created Event
	// In production you would use `EventFindFirstChecked`` instead of `EventFindFirst`.
	event := SDK.EventFindFirst(res.Events.Unwrap(), pxPallet.EventPureCreated{}).UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pure: %v, Who %v, ProxyType %v, Index: %v`, event.Pure.ToHuman(), event.Who.ToHuman(), event.ProxyTpe.ToHuman(), event.DisambiguationIndex))
	child := event.Pure

	// Forcing "Child" to create an application key
	key := fmt.Sprintf("MyKey%v", rand.Uint32())
	call := pallets.ToCall(daPallet.CallCreateApplicationKey{Key: []byte(key)})
	tx = sdk.Tx.Proxy.Proxy(child.ToMultiAddress(), primitives.NewNone[metadata.ProxyType](), call)
	res, err = tx.ExecuteAndWatchInclusion(parent, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction has to succeed")

	// Finding the Proxy Executed Event
	// In production you would use `EventFindFirstChecked`` instead of `EventFindFirst`.
	event2 := SDK.EventFindFirst(res.Events.Unwrap(), pxPallet.EventProxyExecuted{}).UnsafeUnwrap()
	AssertEq(event2.Result.VariantIndex, 0, "Event must OK")
	fmt.Println(fmt.Sprintf(`Dispatch Result: %v`, event2.Result.ToString()))

	fmt.Println("RunPureProxy finished correctly.")
}
