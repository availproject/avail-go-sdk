package examples

import (
	"fmt"

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

func RunTransactionCustom() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	customTx := CustomTransaction{Value: []byte("Hello World")}
	tx := SDK.NewTransaction(sdk.Client, pallets.ToPayload(customTx))

	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), SDK.NewTransactionOptions())
	PanicOnError(err)

	isSuc, err := res.IsSuccessful()
	PanicOnError(err)
	AssertEq(isSuc, true, "Transaction needs to be successful")

	block, err := SDK.NewBlock(sdk.Client, res.BlockHash)
	PanicOnError(err)

	genTx := block.TransactionByIndex(res.TxIndex).UnsafeUnwrap()
	foundTx := CustomTransaction{}
	isOk := pallets.Decode(&foundTx, genTx.Extrinsic)
	AssertEq(isOk, true, "Transaction needs to be decodable")
	fmt.Println("Value:", string(foundTx.Value))

	fmt.Println("RunCustomTransaction finished correctly.")
}
