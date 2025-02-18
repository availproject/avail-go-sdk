package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata/pallets"
	"github.com/availproject/avail-go-sdk/primitives"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockTransactionAll() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := primitives.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	// All Transactions
	blockTxs := block.Transactions(SDK.Filter{})
	fmt.Println("Transaction Count: ", len(blockTxs))
	AssertEq(len(blockTxs), 9, "Transaction count is not 9")

	// Printout Block Transactions
	for _, tx := range blockTxs {
		fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Call Name: %v, Call Index: %v, Tx Hash: %v, Tx Index: %v`, tx.PalletName(), tx.PalletIndex(), tx.CallName(), tx.CallIndex(), tx.TxHash(), tx.TxIndex()))
		fmt.Println(fmt.Sprintf(`Tx Signer: %v, App Id: %v, Tip: %v, Mortality: %v, Nonce: %v`, tx.SS58Address(), tx.AppId(), tx.Tip(), tx.Mortality(), tx.Nonce()))
	}

	// Convert from Block Transaction to Specific Transaction
	daTx := daPallet.CallSubmitData{}
	isOk := pallets.Decode(&daTx, blockTxs[2].Extrinsic)
	AssertEq(isOk, true, "Transaction number 3 was not of type Call Submit Data")
	fmt.Println(fmt.Sprintf(`Data: %v,`, string(daTx.Data)))

	// Printout all Transaction Events
	txEvents := blockTxs[2].Events().UnsafeUnwrap()
	AssertEq(len(txEvents), 7, "Events count is not 7")

	for _, ev := range txEvents {
		fmt.Println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v, Tx Index: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position, ev.TxIndex()))
	}

	// Find DataSubmitted event
	eventMyb := SDK.EventFindFirst(txEvents, daPallet.EventDataSubmitted{})
	event := eventMyb.UnsafeUnwrap().UnsafeUnwrap()
	fmt.Println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, Who: %v, Data Hash: %v`, event.PalletName(), event.EventName(), event.Who.ToHuman(), event.DataHash))

	fmt.Println("RunBlockTransactionAll finished correctly.")
}
