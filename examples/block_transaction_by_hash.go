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
	txs := block.Transactions(SDK.Filter{}.WTxHash(txHash))
	AssertEq(len(txs), 1, "")
	tx := &txs[0]

	// Printout Block Transaction filtered by Tx Hash
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Call Name: %v, Call Index: %v, Tx Hash: %v, Tx Index: %v`, tx.PalletName(), tx.PalletIndex(), tx.CallName(), tx.CallIndex(), tx.TxHash(), tx.TxIndex()))
	fmt.Println(fmt.Sprintf(`Tx Signer: %v, App Id: %v, Tip: %v, Mortality: %v, Nonce: %v`, tx.SS58Address(), tx.AppId(), tx.Tip(), tx.Mortality(), tx.Nonce()))

	// Convert from Block Transaction to Specific Transaction
	baTx := baPallet.CallTransferKeepAlive{}
	isOk := pallets.Decode(&baTx, tx.Extrinsic)
	AssertEq(isOk, true, "Transaction was not of type Transfer Keep Alive")
	fmt.Println(fmt.Sprintf(`Destination: %v, Value: %v`, baTx.Dest.Id.UnsafeUnwrap().ToHuman(), baTx.Value.ToHuman()))

	// Printout all Transaction Events
	txEvents := tx.Events().UnsafeUnwrap()
	AssertEq(len(txEvents), 7, "Events count is not 7")

	for _, ev := range txEvents {
		fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v, Tx Index: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position, ev.TxIndex()))
	}

	// Find Transfer event
	eventMyb := SDK.EventFindFirst(txEvents, baPallet.EventTransfer{})
	event := eventMyb.UnsafeUnwrap().UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, From: %v, To: %v, Amount: %v`, event.PalletName(), event.EventName(), event.From.ToHuman(), event.To.ToHuman(), event.Amount))

	fmt.Println("RunBlockTransactionByHash finished correctly.")
}
