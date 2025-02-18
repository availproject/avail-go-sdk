package examples

import (
	"fmt"

	prim "github.com/availproject/avail-go-sdk/primitives"

	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockEvents() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := prim.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	// All Block Events
	blockEvents := block.Events().UnsafeUnwrap()
	AssertEq(len(blockEvents), 53, "Block event count must be 53")

	// Printout All Block Events
	for _, ev := range blockEvents {
		fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v, Tx Index: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position, ev.TxIndex()))
	}

	// Find Transfer event
	baEvents := SDK.EventFind(blockEvents, baPallet.EventTransfer{})
	PanicOnError(err)
	AssertEq(len(baEvents), 2, "Event Transfer event count is not 2")

	for _, ev := range baEvents {
		fmt.Println(fmt.Sprintf(`From: %v, To: %v, Amount: %v`, ev.From.ToHuman(), ev.To.ToHuman(), ev.Amount))
	}

	// Find ApplicationKeyCreated event
	daEventMyb := SDK.EventFindFirst(blockEvents, daPallet.EventApplicationKeyCreated{})
	daEvent := daEventMyb.UnsafeUnwrap().UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, Id: %v, Key: %v, Owner: %v`, daEvent.PalletName(), daEvent.EventName(), daEvent.Id, string(daEvent.Key), daEvent.Owner.ToHuman()))

	// Check
	AssertEq(len(SDK.EventFind(blockEvents, daPallet.EventDataSubmitted{})), 4, "Incorrect count of Data Submitted Event")
	AssertEq(len(SDK.EventFind(blockEvents, daPallet.EventApplicationKeyCreated{})), 1, "Incorrect count of Application Key Created Event")

	// Events for Specific Transaction
	txIndex := uint32(0)
	txEvents := block.EventsForTransaction(txIndex).UnsafeUnwrap()
	AssertEq(len(txEvents), 1, "Tx event count is not 1")

	// Printout All Tx Events
	for _, ev := range txEvents {
		AssertEq(ev.TxIndex(), prim.Some(txIndex), "Tx Index is not the same")
		fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v, Tx Index: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position, ev.TxIndex()))
	}

	// Find ExtrinsicSuccess event
	syEventMyb := SDK.EventFindFirst(txEvents, syPallet.EventExtrinsicSuccess{})
	syEvent := syEventMyb.UnsafeUnwrap().UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, Class: %v`, syEvent.PalletName(), syEvent.EventName(), syEvent.DispatchInfo.Class))

	// Check
	tx2 := block.Transactions(SDK.Filter{}.WTxIndex(txIndex))
	AssertEq(len(tx2), 1, "")
	tx2Events := tx2[0].Events()
	AssertTrue(tx2Events.IsSome(), "")
	AssertEq(len(tx2Events.Unwrap()), len(txEvents), "")

	fmt.Println("RunBlockEvents finished correctly.")
}
