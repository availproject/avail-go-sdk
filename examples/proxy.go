package examples

import (
	"fmt"
	"math/rand/v2"

	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/metadata/pallets"
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
// There are different types of Proxy account and in not of all has the "Parent" the same control options.
//
// Pure Proxy accounts work in the opposite direction. If we designate the name "Parent" to our main account,
// then the proxy account will be designated with the name "Child". The "Child" cannot do anything on it's own
// and the "Parent" has full control over it. The "Parent" can forget that he even had the "Child" in the first
// place thus forcing the "Child" to become an orphan that no will ever see again.

func RunProxy() {
	RunPureProxy()
	fmt.Println("RunProxy finished correctly.")
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
	eventMyb, err := SDK.EventFindFirstChecked(res.Events.Unwrap(), pxPallet.EventPureCreated{})
	PanicOnError(err)
	AssertTrue(eventMyb.IsSome(), "Event must be found")

	// Finding "Child"
	event := eventMyb.Unwrap()
	fmt.Println(fmt.Sprintf(`Pure: %v, Who %v, ProxyType %v, Index: %v`, event.Pure.ToHuman(), event.Who.ToHuman(), event.ProxyTpe.ToHuman(), event.DisambiguationIndex))
	child := event.Pure
	/* 	childHeight := res.BlockNumber
	   	childExtIndex := res.TxIndex */

	// Forcing "Child" to create an application key
	key := fmt.Sprintf("MyKey%v", rand.Uint32())
	call := pallets.ToCall(daPallet.CallCreateApplicationKey{Key: []byte(key)})
	tx = sdk.Tx.Proxy.Proxy(child.ToMultiAddress(), primitives.NewNone[metadata.ProxyType](), call)
	res, err = tx.ExecuteAndWatchInclusion(parent, SDK.NewTransactionOptions())
	PanicOnError(err)

	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction has to succeed")
	AssertTrue(res.Events.IsSome(), "Events need to be decodable")

	// Finding the Proxy Executed Event
	eventMyb2, err := SDK.EventFindFirstChecked(res.Events.Unwrap(), pxPallet.EventProxyExecuted{})
	PanicOnError(err)
	AssertTrue(eventMyb2.IsSome(), "Event must be found")

	event2 := eventMyb2.Unwrap()
	AssertEq(event2.Result.VariantIndex, 0, "Event must OK")
	fmt.Println(fmt.Sprintf(`Dispatch Result: %v`, event2.Result.ToString()))

	/* 	// "Parent" decides that the "Child" should become fatherless.
	   	accountId := metadata.NewAccountIdFromKeyPair(parent)
	   	tx = sdk.Tx.Proxy.KillPure(accountId.ToMultiAddress(), proxyType, index, childHeight, childExtIndex)
	   	res, err = tx.ExecuteAndWatchInclusion(parent, SDK.NewTransactionOptions())
	   	PanicOnError(err)

	   	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction has to succeed")
	   	AssertTrue(res.Events.IsSome(), "Events need to be decodable") */

	fmt.Println("RunPureProxy finished correctly.")
}
