package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata/pallets"
	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockTransactions() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	if err != nil {
		panic(err)
	}

	blockHash, err := primitives.NewH256FromHexString("0xd036cf19fc02229d8ca0b06b5d96fcd1feb7bd9b3609ff09920cac001e2ecd58")
	if err != nil {
		panic(err)
	}

	// Fetching block with Block Hash
	block, err := SDK.NewBlock(sdk.Client, blockHash)
	if err != nil {
		panic(err)
	}

	// Transactions
	//
	// Available methods:
	// TransactionAll, TransactionBySigner, TransactionByIndex, TransactionByHash, TransactionByAppId
	allTxs := block.TransactionAll()
	println("Transaction Count: ", len(allTxs))
	if len(allTxs) != 4 {
		panic("Transaction count needs to be 4")
	}

	decodedIndices := []uint32{}
	for _, tx := range allTxs {
		println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Call Name: %v, Call Index: %v`, tx.PalletName(), tx.PalletIndex(), tx.CallName(), tx.CallIndex()))
		println(fmt.Sprintf(`Tx Hash: %v, Tx Index: %v`, tx.TxHash().ToHuman(), tx.TxIndex()))

		// Converting generic block tx to specific one
		call := daPallet.CallSubmitData{}
		if !pallets.Decode(&call, tx.Extrinsic) {
			println("Failed to decode transaction. Skipping it.")
			println()
			continue
		}

		decodedIndices = append(decodedIndices, tx.TxIndex())
		println(fmt.Sprintf(`Found Submit Data call. Length: %v`, len(call.Data)))

		eventsMyb := tx.Events()
		if eventsMyb.IsNone() {
			panic("Events should be present for this transactions")
		}
		txEvents := eventsMyb.SafeUnwrap()

		if len(txEvents) != 7 {
			panic("There should be 7 events")
		}

		eventMyb, err := SDK.EventFindFirstChecked(txEvents, baPallet.EventWithdraw{})
		if err != nil {
			panic("Failed to decode events")
		}

		if eventMyb.IsNone() {
			panic("Failed to find  EventWithdraw event")
		}

		println("Withdraw amount:", eventMyb.Unwrap().Amount.ToHuman())

		println()
	}

	if len(decodedIndices) != 2 || decodedIndices[0] != 1 || decodedIndices[1] != 2 {
		panic("We we supposed to decode 2 transactions")
	}

	println("RunBlockTransactions finished correctly.")
}
