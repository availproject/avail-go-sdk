package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata/pallets"
	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	utPallet "github.com/availproject/avail-go-sdk/metadata/pallets/utility"
	"github.com/availproject/avail-go-sdk/primitives"
	prim "github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBatch() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Use SDK.Account.NewKeyPair("Your key") to use a different account than Alice
	acc := SDK.Account.Alice()

	callsToExecute := []prim.Call{}

	// One way to create a suitable call for the batch transaction is to manually create the desired call and then convert it to a generic call
	{
		destBob, err := primitives.NewAccountIdFromAddress("5FHneW46xGXgs5mUiveU4sbTyGBzmstUspZC92UhjJM694ty")
		PanicOnError(err)

		call := baPallet.CallTransferKeepAlive{Dest: destBob.ToMultiAddress(), Value: SDK.OneAvail()}
		callsToExecute = append(callsToExecute, pallets.ToCall(call))
	}

	// The other was it to create a transaction using the sdk api and then use the `call` field member
	{
		destCharlie, err := primitives.NewAccountIdFromAddress("5FLSigC9HGRKVhB9FiEo4Y3koPsNmBmLJbpXg2mp1hXcS59Y")
		PanicOnError(err)

		tx := sdk.Tx.Balances.TransferKeepAlive(destCharlie.ToMultiAddress(), SDK.OneAvail())
		callsToExecute = append(callsToExecute, tx.Payload.Call)
	}

	//
	// Happy Path
	//

	// Batch call
	{
		tx := sdk.Tx.Utility.Batch(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

		events := res.Events.UnsafeUnwrap()

		event := SDK.EventFindFirst(events, utPallet.EventBatchCompleted{})
		AssertTrue(event.IsSome(), "BatchCompleted event must be present.")

		event_count := len(SDK.EventFind(events, utPallet.EventItemCompleted{}))
		AssertEq(event_count, 2, "ItemCompleted events must be produced twice")

		fmt.Println("Batch call done")
	}

	// Batch All call
	{
		tx := sdk.Tx.Utility.BatchAll(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

		events := res.Events.UnsafeUnwrap()

		event := SDK.EventFindFirst(events, utPallet.EventBatchCompleted{})
		AssertTrue(event.IsSome(), "BatchCompleted event must be present.")

		event_count := len(SDK.EventFind(events, utPallet.EventItemCompleted{}))
		AssertEq(event_count, 2, "ItemCompleted events must be produced twice")

		fmt.Println("Batch All call done")
	}

	// Force Batch call
	{
		tx := sdk.Tx.Utility.ForceBatch(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

		events := res.Events.UnsafeUnwrap()

		event := SDK.EventFindFirst(events, utPallet.EventBatchCompleted{})
		AssertTrue(event.IsSome(), "BatchCompleted event must be present.")

		event_count := len(SDK.EventFind(events, utPallet.EventItemCompleted{}))
		AssertEq(event_count, 2, "ItemCompleted events must be produced twice")

		fmt.Println("Force Batch call done")
	}

	//
	//	Things differ when we introduce a call that will fail
	//

	// The 3. is poisoned with a too high transfer amount
	{
		destEve, err := primitives.NewAccountIdFromAddress("5HGjWAeFDfFCWPsjFQdVV2Msvz2XtMktvgocEZcCj68kUMaw")
		PanicOnError(err)

		tx := sdk.Tx.Balances.TransferKeepAlive(destEve.ToMultiAddress(), SDK.OneAvail().Mul64(uint64(1_000_000_000)))
		callsToExecute = append(callsToExecute, tx.Payload.Call)
	}

	// The 4. call is a normal one
	{
		destDave, err := primitives.NewAccountIdFromAddress("5DAAnrj7VHTznn2AWBemMuyBwZWs6FNFjdyVXUeYum3PTXFy")
		PanicOnError(err)

		tx := sdk.Tx.Balances.TransferKeepAlive(destDave.ToMultiAddress(), SDK.OneAvail())
		callsToExecute = append(callsToExecute, tx.Payload.Call)
	}

	// Batch call
	{
		tx := sdk.Tx.Utility.Batch(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

		events := res.Events.UnsafeUnwrap()

		event := SDK.EventFindFirst(events, utPallet.EventBatchInterrupted{})
		AssertTrue(event.IsSome(), "BatchInterrupted event must be present.")

		event2 := SDK.EventFindFirst(events, utPallet.EventBatchCompleted{})
		AssertTrue(event2.IsNone(), "BatchCompleted event must NOT be present.")

		event_count := len(SDK.EventFind(events, utPallet.EventItemCompleted{}))
		AssertEq(event_count, 2, "ItemCompleted events must be produced twice")

		fmt.Println("Batch call done")
	}

	// Batch All call
	{
		tx := sdk.Tx.Utility.BatchAll(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertEq(res.IsSuccessful(), prim.Some(false), "Transaction is supposed to fail")

		fmt.Println("Batch All call done")
	}

	// Force Batch call
	{
		tx := sdk.Tx.Utility.ForceBatch(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

		events := res.Events.UnsafeUnwrap()

		event := SDK.EventFindFirst(events, utPallet.EventBatchCompletedWithErrors{})
		AssertTrue(event.IsSome(), "BatchCompletedWithErrors event must be present.")

		event_count := len(SDK.EventFind(events, utPallet.EventItemCompleted{}))
		AssertEq(event_count, 3, "ItemCompleted events must be produced thrice")

		event_count2 := len(SDK.EventFind(events, utPallet.EventItemFailed{}))
		AssertEq(event_count2, 1, "ItemFailed events must be produced once")

		fmt.Println("Force Batch call done")
	}

	fmt.Println("RunBatch finished correctly.")
}
