package examples

import (
	"fmt"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionExecuteAndWatch() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Transaction will be signed, sent, and watched
	// If the transaction was dropped or never executed, the system will retry it
	// for 2 more times using the same nonce and app id.
	//
	// Param `waitFor` can be either `SDK.Inclusion` or `SDK.Finalization`
	// Param `blockTimeout` defines how many blocks will the watcher explore before determining
	// that the transaction was not executed.
	//
	// Most likely you would want to call `ExecuteAndWatchFinalization` or `ExecuteAndWatchInclusion`
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	txDetails, err := tx.ExecuteAndWatch(SDK.Account.Alice(), SDK.Inclusion, SDK.NewTransactionOptions().WithAppId(1), 3)
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

	// Converts from generic transaction to a specific one
	eventMyb := SDK.EventFindFirst(txEvents, daPallet.EventDataSubmitted{})
	event := eventMyb.UnsafeUnwrap().UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, DataHash: %v, Who: %v`, event.PalletName(), event.EventName(), event.DataHash, event.Who.ToHuman()))

	fmt.Println("RunTransactionExecuteAndWatch finished correctly.")
}
