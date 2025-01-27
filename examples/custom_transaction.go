package examples

import (
	"github.com/availproject/avail-go-sdk/metadata/pallets"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

type CustomTransaction struct {
	Value []byte
}

func (this CustomTransaction) PalletName() string {
	return "DataAvailability"
}

func (this CustomTransaction) PalletIndex() uint8 {
	return 29
}

func (this CustomTransaction) CallName() string {
	return "submit_data"
}

func (this CustomTransaction) CallIndex() uint8 {
	return 1
}

func RunCustomTransaction() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	if err != nil {
		panic(err)
	}

	customTx := CustomTransaction{Value: []byte("Hello World")}
	tx := SDK.NewTransaction(sdk.Client, pallets.ToPayload(customTx))

	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), SDK.NewTransactionOptions())
	if err != nil {
		panic(err)
	}
	if isSuc, err := res.IsSuccessful(); err != nil {
		panic(err)
	} else if !isSuc {
		println("The transaction was unsuccessful")
	}

	block, err := SDK.NewBlock(sdk.Client, res.BlockHash)
	if err != nil {
		panic(err)
	}

	genTx := block.TransactionByIndex(res.TxIndex).SafeUnwrap()

	foundTx := CustomTransaction{}
	if isOk := pallets.Decode(&foundTx, genTx.Extrinsic); !isOk {
		panic("Failed to Decode Custom Transaction")
	}
	println("Value:", string(foundTx.Value))

	println("RunCustomTransaction finished correctly.")
}
