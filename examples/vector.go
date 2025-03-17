package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	sudoPallet "github.com/availproject/avail-go-sdk/metadata/pallets/sudo"
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunVector() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	alice := SDK.Account.Alice()

	// Whitelisting 0
	{
		call := sdk.Tx.Vector.SetWhitelistedDomains([]uint32{0})
		sudoTx := sdk.Tx.Sudo.Sudo(call.Payload.Call)
		res, err := sudoTx.ExecuteAndWatchInclusion(alice, SDK.NewTransactionOptions())
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction must be successful")
		events := res.Events.UnsafeUnwrap()
		event := SDK.EventFindFirst(events, sudoPallet.EventSudid{}).UnsafeUnwrap().UnsafeUnwrap()
		AssertEq(event.SudoResult.VariantIndex, 0, "Sudo Failed")
	}

	// Enabling mock
	{
		call := sdk.Tx.Vector.EnableMock(true)
		sudoTx := sdk.Tx.Sudo.Sudo(call.Payload.Call)
		res, err := sudoTx.ExecuteAndWatchInclusion(alice, SDK.NewTransactionOptions())
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction must be successful")
		events := res.Events.UnsafeUnwrap()
		event := SDK.EventFindFirst(events, sudoPallet.EventSudid{}).UnsafeUnwrap().UnsafeUnwrap()
		AssertEq(event.SudoResult.VariantIndex, 0, "Sudo Failed")
	}

	// Setting updater
	{
		updater, err := primitives.NewH256FromByteSlice(alice.AccountID())
		PanicOnError(err)

		call := sdk.Tx.Vector.SetUpdater(updater)
		sudoTx := sdk.Tx.Sudo.Sudo(call.Payload.Call)
		res, err := sudoTx.ExecuteAndWatchInclusion(alice, SDK.NewTransactionOptions())
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction must be successful")
		events := res.Events.UnsafeUnwrap()
		event := SDK.EventFindFirst(events, sudoPallet.EventSudid{}).UnsafeUnwrap().UnsafeUnwrap()
		AssertEq(event.SudoResult.VariantIndex, 0, "Sudo Failed")
	}

	// Mock Fulfill
	{
		publicValues := []byte{0, 1, 2}
		tx := sdk.Tx.Vector.MockFulfill(publicValues)
		res, err := tx.ExecuteAndWatchInclusion(alice, SDK.NewTransactionOptions())
		PanicOnError(err)
		AssertEq(res.IsSuccessful().UnsafeUnwrap(), false, "Should fail as public values are not correct")
	}

	// Send Message
	{
		message := metadata.VectorMessage{VariantIndex: 0, ArbitraryMessage: primitives.Some([]byte{0, 1, 2})}
		tx := sdk.Tx.Vector.SendMessage(message, primitives.H256{}, 0)
		res, err := tx.ExecuteAndWatchInclusion(alice, SDK.NewTransactionOptions())
		PanicOnError(err)
		AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction must be successful")
	}

	fmt.Println("RunVector finished correctly.")

}
