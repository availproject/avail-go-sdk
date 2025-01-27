package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata/pallets"
	"github.com/availproject/avail-go-sdk/primitives"

	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockTransactionByHash() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := primitives.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	// Transaction filtered by Transaction hash
	txHash, err := primitives.NewH256FromHexString("0x19c486e107c926ff4af3fa9b1d95aaba130cb0bc89515d0f5b523ef6bac06338")
	PanicOnError(err)
	tx := block.TransactionByHash(txHash).UnsafeUnwrap()

	// Printout Block Transaction
	println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Call Name: %v, Call Index: %v, Tx Hash: %v, Tx Index: %v, Tx Signer: %v, App Id: %v`, tx.PalletName(), tx.PalletIndex(), tx.CallName(), tx.CallIndex(), tx.TxHash(), tx.TxIndex(), tx.Signer(), tx.AppId()))

	// Convert from Block Transaction to Specific Transaction
	baTx := baPallet.CallTransferKeepAlive{}
	isOk := pallets.Decode(&baTx, tx.Extrinsic)
	AssertEq(isOk, true, "Transaction was not of type Transfer Keep Alive")
	println(fmt.Sprintf(`Destination: %v, Value: %v`, baTx.Dest.Id.UnsafeUnwrap().ToHuman(), baTx.Value.ToHuman()))

	// Printout all Transaction Events
	txEvents := tx.Events().UnsafeUnwrap()
	AssertEq(len(txEvents), 7, "Events count is not 7")

	for _, ev := range txEvents {
		println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position))
	}

	// Convert from Block Transaction Event to Specific Transaction Event
	event := SDK.EventFindFirst(txEvents, baPallet.EventTransfer{}).UnsafeUnwrap()
	println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, From: %v, To: %v, Amount: %v`, event.PalletName(), event.EventName(), event.From.ToHuman(), event.To.ToHuman(), event.Amount))

	println("RunBlockTransactionByHash finished correctly.")
}
