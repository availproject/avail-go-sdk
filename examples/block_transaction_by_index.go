package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata/pallets"
	"github.com/availproject/avail-go-sdk/primitives"

	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockTransactionByIndex() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := primitives.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	// Transaction filtered by Transaction index
	tx := block.TransactionByIndex(1).UnsafeUnwrap()

	// Printout Block Transaction
	println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Call Name: %v, Call Index: %v, Tx Hash: %v, Tx Index: %v, Tx Signer: %v, App Id: %v`, tx.PalletName(), tx.PalletIndex(), tx.CallName(), tx.CallIndex(), tx.TxHash().ToHuman(), tx.TxIndex(), tx.Signer(), tx.AppId()))

	// Convert from Block Transaction to Specific Transaction
	baTx := baPallet.CallTransferKeepAlive{}
	isOk := pallets.Decode(&baTx, tx.Extrinsic)
	AssertEq(isOk, true, "Transaction was not of type Transfer Keep Alive")
	println(fmt.Sprintf(`Destination: %v, Value: %v`, baTx.Dest.Id.UnsafeUnwrap().ToHuman(), baTx.Value.ToHuman()))

	// Printout all Transaction Events
	txEvents := tx.Events().UnsafeUnwrap()
	AssertEq(len(txEvents), 9, "Events count is not 9")

	for _, ev := range txEvents {
		println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position))
	}

	// Convert from Block Transaction Event to Specific Transaction Event
	event := SDK.EventFindFirst(txEvents, syPallet.EventNewAccount{}).UnsafeUnwrap()
	println(fmt.Sprintf(`Account: %v`, event.Account.ToHuman()))

	println("RunBlockTransactionByIndex finished correctly.")
}
