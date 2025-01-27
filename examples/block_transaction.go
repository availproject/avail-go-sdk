package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata/pallets"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockTransaction() {
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

	for _, tx := range allTxs {
		println(fmt.Sprintf(`Pallet Name: %v, Pallet Index: %v, Call Name: %v, Call Index: %v`, tx.PalletName(), tx.PalletIndex(), tx.CallName(), tx.CallIndex()))
		println(fmt.Sprintf(`Tx Hash: %v, Tx Index: %v`, tx.TxHash().ToHuman(), tx.TxIndex()))

		// Converting generic block tx to a DataAvailability SubmitData Call Data
		call := daPallet.CallSubmitData{}
		if !pallets.Decode(&call, tx.Extrinsic) {
			continue
		}

		println(fmt.Sprintf(`Found Submit Data call. Length: %v`, len(call.Data)))

	}
}
