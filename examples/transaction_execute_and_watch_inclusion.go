package examples

import (
	"fmt"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionExecuteAndWatchInclusion() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Transaction will be signed, sent, and watched
	// If the transaction was dropped or never executed, the system will retry it
	// for 2 more times using the same nonce and app id.
	//
	// Waits for transaction inclusion. Most of the time you would want to call `ExecuteAndWatchFinalization` as
	// inclusion doesn't mean that the transaction will be in the canonical chain.
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	txDetails, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), SDK.NewTransactionOptions().WithAppId(1))
	PanicOnError(err)

	// Returns None if there was no way to determine the
	// success status of a transaction. Otherwise it returns
	// true or false.
	AssertTrue(txDetails.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

	// Printout Transaction Details
	fmt.Println(fmt.Sprintf(`Block Hash: %v, Block Index: %v, Tx Hash: %v, Tx Index: %v`, txDetails.BlockHash, txDetails.BlockNumber, txDetails.TxHash, txDetails.TxIndex))

	// Printout Transaction Events
	txEvents := txDetails.Events.UnsafeUnwrap()
	for _, ev := range txEvents {
		fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v, Tx Index: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position, ev.TxIndex()))
	}

	// Converts generic event to a specific one
	eventMyb := SDK.EventFindFirst(txEvents, daPallet.EventDataSubmitted{})
	event := eventMyb.UnsafeUnwrap().UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, DataHash: %v, Who: %v`, event.PalletName(), event.EventName(), event.DataHash, event.Who.ToHuman()))

	fmt.Println("RunTransactionExecuteAndWatchInclusion finished correctly.")
}
