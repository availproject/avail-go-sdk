package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	utPallet "github.com/availproject/avail-go-sdk/metadata/pallets/utility"
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
		destBob, err := metadata.NewAccountIdFromAddress("5FHneW46xGXgs5mUiveU4sbTyGBzmstUspZC92UhjJM694ty")
		PanicOnError(err)

		call := baPallet.CallTransferKeepAlive{Dest: destBob.ToMultiAddress(), Value: SDK.OneAvail()}
		callsToExecute = append(callsToExecute, call.ToCall())
	}

	// The other was it to create a transaction using the sdk api and then use the `call` field member
	{
		destCharlie, err := metadata.NewAccountIdFromAddress("5FLSigC9HGRKVhB9FiEo4Y3koPsNmBmLJbpXg2mp1hXcS59Y")
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
		AssertTrue(res.IsSuccessful().Unwrap(), "Transaction is supposed to succeed")

		events := res.Events.Unwrap()

		if SDK.EventFindFirst(events, utPallet.EventBatchCompleted{}).IsSome() {
			fmt.Println("Batch was successfully completed")
		} else {
			panic("Batch call failed")
		}

		if len(SDK.EventFindAll(events, utPallet.EventItemCompleted{})) == 2 {
			fmt.Println("All batch items completed")
		} else {
			panic("No all items were completed")
		}

		fmt.Println("Batch call done")
	}

	// Batch All call
	{
		tx := sdk.Tx.Utility.BatchAll(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().Unwrap(), "Transaction is supposed to succeed")

		events := res.Events.Unwrap()

		if SDK.EventFindFirst(events, utPallet.EventBatchCompleted{}).IsSome() {
			fmt.Println("Batch was successfully completed")
		} else {
			panic("Batch All call failed")
		}

		if len(SDK.EventFindAll(events, utPallet.EventItemCompleted{})) == 2 {
			fmt.Println("All batch items completed")
		} else {
			panic("No all items were completed")
		}

		fmt.Println("Batch All call done")
	}

	// Force Batch call
	{
		tx := sdk.Tx.Utility.ForceBatch(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().Unwrap(), "Transaction is supposed to succeed")

		events := res.Events.Unwrap()

		if SDK.EventFindFirst(events, utPallet.EventBatchCompleted{}).IsSome() {
			fmt.Println("Batch was successfully completed")
		} else {
			panic("Batch All call failed")
		}

		if len(SDK.EventFindAll(events, utPallet.EventItemCompleted{})) == 2 {
			fmt.Println("All batch items completed")
		} else {
			panic("No all items were completed")
		}

		fmt.Println("Force Batch call done")
	}

	//
	//	Things differ when we introduce a call that will fail
	//

	// The 3. is poisoned with a too high transfer amount
	{
		destEve, err := metadata.NewAccountIdFromAddress("5HGjWAeFDfFCWPsjFQdVV2Msvz2XtMktvgocEZcCj68kUMaw")
		PanicOnError(err)

		tx := sdk.Tx.Balances.TransferKeepAlive(destEve.ToMultiAddress(), SDK.OneAvail().Mul64(uint64(1_000_000_000)))
		callsToExecute = append(callsToExecute, tx.Payload.Call)
	}

	// The 4. call is a normal one
	{
		destDave, err := metadata.NewAccountIdFromAddress("5DAAnrj7VHTznn2AWBemMuyBwZWs6FNFjdyVXUeYum3PTXFy")
		PanicOnError(err)

		tx := sdk.Tx.Balances.TransferKeepAlive(destDave.ToMultiAddress(), SDK.OneAvail())
		callsToExecute = append(callsToExecute, tx.Payload.Call)
	}

	// Batch call
	{
		tx := sdk.Tx.Utility.Batch(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().Unwrap(), "Transaction is supposed to succeed")

		events := res.Events.Unwrap()

		if event := SDK.EventFindFirst(events, utPallet.EventBatchInterrupted{}); event.IsSome() {
			ev := event.Unwrap()
			fmt.Println("Batch was interrupted. Reason: ", ev.Error.ToHuman())
			fmt.Println("Tx Index that caused failure: ", ev.Index)
		} else {
			panic("Failed to find EventBatchInterrupted event.")
		}

		if len(SDK.EventFindAll(events, utPallet.EventItemCompleted{})) == 2 {
			fmt.Println("Some batch items completed")
		} else {
			panic("Cannot be more than 2")
		}

		fmt.Println("Batch call done")
	}

	// Batch All call
	{
		tx := sdk.Tx.Utility.BatchAll(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)

		isOk := res.IsSuccessful()
		AssertTrue(isOk.IsSome(), "It should be possible to decode transaction events.")
		AssertEq(isOk.Unwrap(), false, "Transaction is supposed to fail")

		events := res.Events.Unwrap()

		if event := SDK.EventFindFirst(events, syPallet.EventExtrinsicFailed{}); event.IsSome() {
			fmt.Println("Batch was interrupted. Reason: ", event.Unwrap().DispatchError.ToHuman())
		} else {
			panic("Failed to find EventExtrinsicFailed event.")
		}

		fmt.Println("Batch All call done")
	}

	// Force Batch call
	{
		tx := sdk.Tx.Utility.ForceBatch(callsToExecute)
		res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().Unwrap(), "Transaction is supposed to succeed")

		events := res.Events.Unwrap()

		if SDK.EventFindFirst(events, utPallet.EventBatchCompletedWithErrors{}).IsSome() {
			fmt.Println("Batch completed with errors")
		} else {
			panic("Failed to find EventBatchCompletedWithErrors")
		}

		if len(SDK.EventFindAll(events, utPallet.EventItemCompleted{})) == 3 {
			fmt.Println("3 of out 4 items completed")
		} else {
			panic("3 items must be completed")
		}

		if event := SDK.EventFindFirst(events, utPallet.EventItemFailed{}); event.IsSome() {
			fmt.Println("Item failed. Reason: ", event.Unwrap().Error.ToHuman())
		} else {
			panic("Failed to find EventItemFailed")
		}

		fmt.Println("Force Batch call done")
	}

	fmt.Println("RunBatch finished correctly.")
}
