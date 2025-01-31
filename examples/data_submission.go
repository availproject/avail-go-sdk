package examples

import (
	"fmt"
	"math/rand/v2"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunDataSubmission() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	if err != nil {
		panic(err)
	}

	// Use SDK.Account.NewKeyPair("Your key") to use a different account than Alice
	acc := SDK.Account.Alice()

	key := fmt.Sprintf("MyKey%v", rand.Uint32())
	// Transactions can be found under sdk.Tx.*
	tx := sdk.Tx.DataAvailability.CreateApplicationKey([]byte(key))

	// Available execution modes
	// Execute, ExecuteAndWatch, ExecuteAndWatchInclusion, ExecuteAndWatchInclusion
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions())
	if err != nil {
		// Failed to submit transaction
		panic(err)
	}
	if isOk := res.IsSuccessful(); isOk.IsNone() {
		// Failed to determine if the transaction was successful or if it failed.
		panic(err)
	} else if isOk.Unwrap() == false {
		panic("The transaction was unsuccessful")
	}

	events := res.Events.Unwrap()
	event := SDK.EventFindFirst(events, daPallet.EventApplicationKeyCreated{}).Unwrap()

	appId := event.Id
	fmt.Println(fmt.Sprintf(`Owner: %v, Key: %v, AppId: %v`, event.Owner.ToHuman(), string(event.Key), appId))

	tx = sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	res, err = tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(appId))
	if err != nil {
		panic(err)
	}

	if isOk := res.IsSuccessful(); isOk.IsNone() {
		// Failed to determine if the transaction was successful or if it failed.
		panic(err)
	} else if isOk.Unwrap() == false {
		panic("The transaction was unsuccessful")
	}

	// Transaction Details
	fmt.Println(fmt.Sprintf(`Block Hash: %v, Block Index: %v, Tx Hash: %v, Tx Index: %v`, res.BlockHash.ToHexWith0x(), res.BlockNumber, res.TxHash.ToHexWith0x(), res.TxIndex))

	events = res.Events.Unwrap()
	event2 := SDK.EventFindFirst(events, daPallet.EventDataSubmitted{}).Unwrap()

	fmt.Println(fmt.Sprintf(`Who: %v, Datahash: %v`, event2.Who.ToHuman(), event2.DataHash.ToHexWith0x()))

	fmt.Println("RunDataSubmission finished correctly.")
}
