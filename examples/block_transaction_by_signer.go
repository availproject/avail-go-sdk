package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/metadata/pallets"
	"github.com/availproject/avail-go-sdk/primitives"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockTransactionBySigner() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := primitives.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	accountId, err := metadata.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	PanicOnError(err)

	// All Transaction filtered by Signer
	blockTxs := block.TransactionBySigner(accountId)
	println("Transaction Count: ", len(blockTxs))
	AssertEq(len(blockTxs), 5, "Transaction count is not 5")

	// Printout Block Transactions made by Signer
	for _, tx := range blockTxs {
		println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Call Name: %v, Call Index: %v, Tx Hash: %v, Tx Index: %v, Tx Signer: %v, App Id: %v`, tx.PalletName(), tx.PalletIndex(), tx.CallName(), tx.CallIndex(), tx.TxHash(), tx.TxIndex(), tx.Signer(), tx.AppId()))
		AssertEq(tx.Signer().UnsafeUnwrap(), accountId.ToHuman(), "Signer is not the correct one")
	}

	// Convert from Block Transaction to Specific Transaction
	daTx := daPallet.CallCreateApplicationKey{}
	isOk := pallets.Decode(&daTx, blockTxs[0].Extrinsic)
	AssertEq(isOk, true, "Transaction was not of type Create Application Key")
	println(fmt.Sprintf(`Key: %v`, string(daTx.Key)))

	// Printout all Transaction Events
	txEvents := blockTxs[0].Events().UnsafeUnwrap()
	AssertEq(len(txEvents), 7, "Events count is not 7")

	for _, ev := range txEvents {
		println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Event Name: %v, Event Index: %v, Event Position: %v`, ev.PalletName, ev.PalletIndex, ev.EventName, ev.EventIndex, ev.Position))
	}

	// Convert from Block Transaction Event to Specific Transaction Event
	event := SDK.EventFindFirst(txEvents, daPallet.EventApplicationKeyCreated{}).UnsafeUnwrap()
	println(fmt.Sprintf(`Pallet Name: %v, Event Name: %v, Owner: %v, Key: %v, AppId: %v`, event.PalletName(), event.EventName(), event.Owner.ToHuman(), string(event.Key), event.Id))

	println("RunBlockTransactionBySigner finished correctly.")
}
