package examples

import (
	"fmt"
	"math/rand/v2"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func Run_data_submission() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	// Use SDK.Account.NewKeyPair("Your key") to use a different account than Alice
	acc, err := SDK.Account.Alice()
	if err != nil {
		panic(err)
	}

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

	events := res.Events.Unwrap()
	event := SDK.EventFindFirst(events, daPallet.EventApplicationKeyCreated{}).Unwrap()

	appId := event.Id
	println(fmt.Sprintf(`Owner: %v, Key: %v, AppId: %v`, event.Owner.ToHuman(), string(event.Key), appId))

	tx = sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	res, err = tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(appId))
	if err != nil {
		panic(err)
	}

	// Transaction Details
	println(fmt.Sprintf(`Block Hash: %v, Block Index: %v, Tx Hash: %v, Tx Index: %v`, res.BlockHash.ToHexWith0x(), res.BlockNumber, res.TxHash.ToHexWith0x(), res.TxIndex))

	events = res.Events.Unwrap()
	event2 := SDK.EventFindFirst(events, daPallet.EventDataSubmitted{}).Unwrap()

	println(fmt.Sprintf(`Who: %v, Datahash: %v`, event2.Who.ToHuman(), event2.DataHash.ToHexWith0x()))
}
