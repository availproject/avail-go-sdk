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
	txIndex := uint32(1)
	txs := block.Transactions(SDK.Filter{}.WTxIndex(txIndex))
	AssertEq(len(txs), 1, "")
	tx := &txs[0]

	// Printout Block Transaction filtered by Tx Index
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Call Name: %v, Call Index: %v, Tx Hash: %v, Tx Index: %v`, tx.PalletName(), tx.PalletIndex(), tx.CallName(), tx.CallIndex(), tx.TxHash(), tx.TxIndex()))
	fmt.Println(fmt.Sprintf(`Tx Signer: %v, App Id: %v, Tip: %v, Mortality: %v, Nonce: %v`, tx.SS58Address(), tx.AppId(), tx.Tip(), tx.Mortality(), tx.Nonce()))

	// Convert from Block Transaction to Specific Transaction
	baTx := baPallet.CallTransferKeepAlive{}
	isOk := pallets.Decode(&baTx, tx.Extrinsic)
	AssertEq(isOk, true, "Transaction was not of type Transfer Keep Alive")
	fmt.Println(fmt.Sprintf(`Destination: %v, Value: %v`, baTx.Dest.Id.UnsafeUnwrap().ToHuman(), baTx.Value))

	// Printout all Transaction Events
	txEvents := tx.Events().UnsafeUnwrap()
	AssertEq(len(txEvents), 9, "Events count is not 9")

	for _, ev := range txEvents {
		fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v, Tx Index: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position, ev.TxIndex()))
	}

	// Find NewAccount event
	eventMyb := SDK.EventFindFirst(txEvents, syPallet.EventNewAccount{})
	event := eventMyb.UnsafeUnwrap().UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, Account: %v`, event.PalletName(), event.EventName(), event.Account.ToHuman()))

	fmt.Println("RunBlockTransactionByIndex finished correctly.")
}
