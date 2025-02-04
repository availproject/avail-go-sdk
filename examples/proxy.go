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

func RunProxy() {
	RunNormalProxy()
	RunPureProxy()

	fmt.Println("RunProxy finished correctly.")
}

func RunNormalProxy() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	proxyAccount := SDK.Account.Bob()
	proxyAccountMulti := primitives.NewAccountIdFromKeyPair(proxyAccount).ToMultiAddress()
	mainAccount := SDK.Account.Ferdie()
	mainAccountMulti := primitives.NewAccountIdFromKeyPair(mainAccount).ToMultiAddress()

	// Creating proxy
	proxyType := metadata.ProxyType{VariantIndex: 0} // Any Proxy
	tx := sdk.Tx.Proxy.AddProxy(proxyAccountMulti, proxyType, 0)
	res, err := tx.ExecuteAndWatchInclusion(mainAccount, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction has to succeed")

	// Finding the Proxy Added Event
	eventMyb := SDK.EventFindFirst(res.Events.UnsafeUnwrap(), pxPallet.EventProxyAdded{})
	event := eventMyb.UnsafeUnwrap().UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Delegatee: %v, Delegator %v, ProxyTpe %v, Delay: %v`, event.Delegatee.ToHuman(), event.Delegator.ToHuman(), event.ProxyTpe.ToHuman(), event.Delay))

	call := pallets.ToCall(baPallet.CallTransferKeepAlive{Dest: proxyAccountMulti, Value: SDK.OneAvail()})

	// Executing the Proxy.Proxy() call
	tx = sdk.Tx.Proxy.Proxy(mainAccountMulti, primitives.NewNone[metadata.ProxyType](), call)
	res, err = tx.ExecuteAndWatchInclusion(proxyAccount, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction has to succeed")

	// Checking for EventProxyExecuted event.
	event2Myb := SDK.EventFindFirst(res.Events.UnsafeUnwrap(), pxPallet.EventProxyExecuted{})
	event2 := event2Myb.UnsafeUnwrap().UnsafeUnwrap()
	AssertEq(event2.Result.VariantIndex, 0, "Event must OK")
	fmt.Println(fmt.Sprintf(`Dispatch Result: %v`, event2.Result.ToString()))

	// Removing Proxy
	tx = sdk.Tx.Proxy.RemoveProxy(proxyAccountMulti, proxyType, 0)
	res, err = tx.ExecuteAndWatchInclusion(mainAccount, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction has to succeed")

	// Checking for EventProxyExecuted event.
	event3Myb := SDK.EventFindFirst(res.Events.UnsafeUnwrap(), pxPallet.EventProxyRemoved{})
	event3 := event3Myb.UnsafeUnwrap().UnsafeUnwrap()
	AssertEq(event2.Result.VariantIndex, 0, "Event must OK")
	fmt.Println(fmt.Sprintf(`Delegatee: %v, Delegator %v, ProxyTpe %v, Delay: %v`, event3.Delegatee.ToHuman(), event3.Delegator.ToHuman(), event3.ProxyTpe.ToHuman(), event3.Delay))

	fmt.Println("RunNormalProxy finished correctly.")
}

func RunPureProxy() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	mainAccount := SDK.Account.Bob()

	// Creating Pure Proxy
	proxyType := metadata.ProxyType{VariantIndex: 0} // Any Proxy
	index := uint16(0)
	tx := sdk.Tx.Proxy.CreatePure(proxyType, 0, index)
	res, err := tx.ExecuteAndWatchInclusion(mainAccount, SDK.NewTransactionOptions())
	PanicOnError(err)

	AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction has to succeed")
	AssertTrue(res.Events.IsSome(), "Events need to be decodable")

	// Finding the Pure Created Event
	eventMyb := SDK.EventFindFirst(res.Events.UnsafeUnwrap(), pxPallet.EventPureCreated{})
	event := eventMyb.UnsafeUnwrap().UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pure: %v, Who %v, ProxyType %v, Index: %v`, event.Pure.ToHuman(), event.Who.ToHuman(), event.ProxyTpe.ToHuman(), event.DisambiguationIndex))
	pureProxy := event.Pure

	// Executing the Proxy.Proxy() call
	key := fmt.Sprintf("MyKey%v", rand.Uint32())
	call := pallets.ToCall(daPallet.CallCreateApplicationKey{Key: []byte(key)})
	tx = sdk.Tx.Proxy.Proxy(pureProxy.ToMultiAddress(), primitives.NewNone[metadata.ProxyType](), call)
	res, err = tx.ExecuteAndWatchInclusion(mainAccount, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction has to succeed")

	// Finding the Proxy Executed Event
	event2Myb := SDK.EventFindFirst(res.Events.UnsafeUnwrap(), pxPallet.EventProxyExecuted{})
	event2 := event2Myb.UnsafeUnwrap().UnsafeUnwrap()
	AssertEq(event2.Result.VariantIndex, 0, "Event must OK")
	fmt.Println(fmt.Sprintf(`Dispatch Result: %v`, event2.Result.ToString()))

	fmt.Println("RunPureProxy finished correctly.")
}
