package examples

import (
	"fmt"
	"math/rand/v2"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunDataSubmission() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Use SDK.Account.NewKeyPair("Your key") to use a different account than Alice
	acc := SDK.Account.Alice()

	key := fmt.Sprintf("MyKey%v", rand.Uint32())
	// Transactions can be found under sdk.Tx.*
	tx := sdk.Tx.DataAvailability.CreateApplicationKey([]byte(key))

	// Transaction Execution
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions())
	// Err means that we failed to submit and execute the transaction.
	PanicOnError(err)

	isOk := res.IsSuccessful()
	// If the return value from `IsSuccessful` is None, it means that we cannot
	// determine if the transaction was successful or not.
	//
	// In this case we assume that we were able to determine it.
	AssertTrue(isOk.IsSome(), "Failed to determine transaction status.")

	// If the value of `IsSuccessful()` is Some(false) then the transaction has failed.
	AssertTrue(isOk.Unwrap(), "The transaction failed.")

	// We might have failed to decode the events so res.Events could be None.
	//
	// In the case we assume that we were able to decode then.
	AssertTrue(res.Events.IsSome(), "Failed to decode events.")
	events := res.Events.Unwrap()
	event := SDK.EventFindFirst(events, daPallet.EventApplicationKeyCreated{}).Unwrap()

	// Printing out all the values of the event ApplicationKeyCreated
	appId := event.Id
	fmt.Println(fmt.Sprintf(`Owner: %v, Key: %v, AppId: %v`, event.Owner.ToHuman(), string(event.Key), appId))

	tx = sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	res, err = tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(appId))
	PanicOnError(err)

	isOk = res.IsSuccessful()
	AssertTrue(isOk.Unwrap(), "The transaction failed or event decoding failed.")

	// Transaction Details
	fmt.Println(fmt.Sprintf(`Block Hash: %v, Block Index: %v, Tx Hash: %v, Tx Index: %v`, res.BlockHash.ToHexWith0x(), res.BlockNumber, res.TxHash.ToHexWith0x(), res.TxIndex))

	AssertTrue(res.Events.IsSome(), "Failed to decode events.")
	events = res.Events.Unwrap()
	event2 := SDK.EventFindFirst(events, daPallet.EventDataSubmitted{}).Unwrap()

	fmt.Println(fmt.Sprintf(`Who: %v, Datahash: %v`, event2.Who.ToHuman(), event2.DataHash.ToHexWith0x()))

	fmt.Println("RunDataSubmission finished correctly.")
}
