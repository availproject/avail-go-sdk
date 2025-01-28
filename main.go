package main

import (
	"fmt"
	"math/rand/v2"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func main() {
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
	println("Executing...")
	res, err := tx.ExecuteAndWatchFinalization(acc, SDK.NewTransactionOptions())
	if err != nil {
		// Failed to submit transaction
		panic(err)
	}
	if isSuc, err := res.IsSuccessful(); err != nil {
		panic(err)
	} else if !isSuc {
		println("The transaction was unsuccessful")
	}

	events := res.Events.Unwrap()
	event := SDK.EventFindFirst(events, daPallet.EventApplicationKeyCreated{}).Unwrap()

	appId := event.Id
	println(fmt.Sprintf(`Owner: %v, Key: %v, AppId: %v`, event.Owner.ToHuman(), string(event.Key), appId))

	tx = sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	res, err = tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(appId))
	if err != nil {
		panic(err)
	}

	if isSuc, err := res.IsSuccessful(); err != nil {
		panic(err)
	} else if !isSuc {
		println("The transaction was unsuccessful")
	}

	// Transaction Details
	println(fmt.Sprintf(`Block Hash: %v, Block Index: %v, Tx Hash: %v, Tx Index: %v`, res.BlockHash.ToHexWith0x(), res.BlockNumber, res.TxHash.ToHexWith0x(), res.TxIndex))

	events = res.Events.Unwrap()
	event2 := SDK.EventFindFirst(events, daPallet.EventDataSubmitted{}).Unwrap()

	println(fmt.Sprintf(`Who: %v, Datahash: %v`, event2.Who.ToHuman(), event2.DataHash.ToHexWith0x()))

	println("RunDataSubmission finished correctly.")
}
