package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/primitives"

	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockEvents() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := primitives.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	// All Block Events
	blockEvents := block.Events().UnsafeUnwrap()
	AssertEq(len(blockEvents), 53, "Block event count is not 2")

	// Printout All Block Events
	for _, ev := range blockEvents {
		println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position))
	}

	// Convert from Block Transaction Event to Specific Transaction Event
	baEvents, err := SDK.EventFindAllChecked(blockEvents, baPallet.EventTransfer{})
	PanicOnError(err)
	AssertEq(len(baEvents), 2, "Event Transfer event count is not 2")

	for _, ev := range baEvents {
		println(fmt.Sprintf(`From: %v, To: %v, Amount: %v`, ev.From.ToHuman(), ev.To.ToHuman(), ev.Amount))
	}

	// Convert from Block Transaction Event to Specific Transaction Event
	daEventMyb, err := SDK.EventFindFirstChecked(blockEvents, daPallet.EventApplicationKeyCreated{})
	PanicOnError(err)
	daEvent := daEventMyb.UnsafeUnwrap()
	println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, Id: %v, Key: %v, Owner: %v`, daEvent.PalletName(), daEvent.EventName(), daEvent.Id, string(daEvent.Key), daEvent.Owner.ToHuman()))

	// Check
	AssertEq(len(SDK.EventFindAll(blockEvents, daPallet.EventDataSubmitted{})), 4, "Incorrect count of Data Submitted Event")
	AssertEq(len(SDK.EventFindAll(blockEvents, daPallet.EventApplicationKeyCreated{})), 1, "Incorrect count of Application Key Created Event")

	// Events for Specific Transaction
	txEvents := block.EventsForTransaction(0).UnsafeUnwrap()
	AssertEq(len(txEvents), 1, "Block event count is not 1")

	// Printout All Tx Events
	for _, ev := range txEvents {
		println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position))
	}

	// Convert from Block Transaction Event to Specific Transaction Event
	syEventMyb, err := SDK.EventFindFirstChecked(blockEvents, syPallet.EventExtrinsicSuccess{})
	PanicOnError(err)
	syEvent := syEventMyb.UnsafeUnwrap()
	println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, Class: %v`, syEvent.PalletName(), syEvent.EventName(), syEvent.DispatchInfo.Class))

	println("RunBlockEvents finished correctly.")
}
