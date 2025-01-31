package examples

import (
	"fmt"
	"github.com/availproject/avail-go-sdk/metadata/pallets"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

type CustomTransaction struct {
	Value []byte
}

func (this CustomTransaction) PalletName() string {
	return "DataAvailability"
}

func (this CustomTransaction) PalletIndex() uint8 {
	return 29
}

func (this CustomTransaction) CallName() string {
	return "submit_data"
}

func (this CustomTransaction) CallIndex() uint8 {
	return 1
}

func RunTransactionCustom() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Creating custom transaction
	customTx := CustomTransaction{Value: []byte("Hello World")}
	tx := SDK.NewTransaction(sdk.Client, pallets.ToPayload(customTx))

	// Executing Custom transaction
	txDetails, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), SDK.NewTransactionOptions())
	PanicOnError(err)

	// Returns None if there was no way to determine the
	// success status of a transaction. Otherwise it returns
	// true or false.
	AssertTrue(txDetails.IsSuccessful().Unwrap(), "Transaction is supposed to succeed")

	// Printout Transaction Details
	fmt.Println(fmt.Sprintf(`Block Hash: %v, Block Index: %v, Tx Hash: %v, Tx Index: %v`, txDetails.BlockHash, txDetails.BlockNumber, txDetails.TxHash, txDetails.TxIndex))

	// Printout Transaction Events
	AssertTrue(txDetails.Events.IsSome(), "We should be able to find events")
	txEvents := txDetails.Events.Unwrap()
	for _, ev := range txEvents {
		fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position))
	}

	// Converts from generic transaction to a specific one
	event := SDK.EventFindFirst(txEvents, daPallet.EventDataSubmitted{}).UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, DataHash: %v, Who: %v`, event.PalletName(), event.EventName(), event.DataHash, event.Who.ToHuman()))

	fmt.Println("RunTransactionCustom finished correctly.")
}
