package examples

import (
	"fmt"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionExecute() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Transaction will be signed, and sent.
	//
	// There is no guarantee that the transaction was executed at all. It might have been
	// dropped or discarded for various reasons. The caller is responsible for querying future
	// blocks in order to determine the execution status of that transaction.
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	txHash, err := tx.Execute(SDK.Account.Alice(), SDK.NewTransactionOptions().WithAppId(1))
	PanicOnError(err)
	fmt.Println("Tx Hash:", txHash)

	// Checking if the transaction was included
	//
	// It's not necessary to use the builtin watcher. A custom watcher
	// might yield better results in some cases.
	watcher := SDK.NewWatcher(sdk.Client, txHash, SDK.Inclusion, 3, 3)
	mybTxDetails, err := watcher.Run()
	PanicOnError(err)
	AssertEq(mybTxDetails.IsSome(), true, "Watcher must have found the status for our transaction")

	// Printout Transaction Details
	txDetails := mybTxDetails.UnsafeUnwrap()
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

	fmt.Println("RunTransactionExecute finished correctly.")
}
