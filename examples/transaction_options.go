package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionOptions() {
	runAppId()
	runNonce()

	fmt.Println("RunTransactionOptions finished correctly.")
}

func runAppId() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	appId := uint32(5)
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))
	options := SDK.NewTransactionOptions().WithAppId(appId)
	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), options)
	PanicOnError(err)

	isSuc, err := res.IsSuccessful()
	PanicOnError(err)
	AssertEq(isSuc, true, "Transaction needs to be successful")

	block, err := SDK.NewBlock(sdk.Client, res.BlockHash)
	PanicOnError(err)

	genTx := block.TransactionByHash(res.TxHash).UnsafeUnwrap()
	foundAppId := genTx.Signed().UnsafeUnwrap().AppId
	AssertEq(appId, foundAppId, "App Ids are not the same")
}

func runNonce() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	acc := SDK.Account.Alice()
	currentNonce, err := SDK.Account.Nonce(sdk.Client, metadata.NewAccountIdFromKeyPair(acc))
	PanicOnError(err)

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))
	options := SDK.NewTransactionOptions().WithNonce(currentNonce).WithAppId(5)
	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), options)
	PanicOnError(err)

	isSuc, err := res.IsSuccessful()
	PanicOnError(err)
	AssertEq(isSuc, true, "Transaction needs to be successful")

	block, err := SDK.NewBlock(sdk.Client, res.BlockHash)
	PanicOnError(err)

	genTx := block.TransactionByHash(res.TxHash).UnsafeUnwrap()
	foundNonce := genTx.Signed().UnsafeUnwrap().Nonce
	AssertEq(foundNonce, currentNonce, "Nonces are not the same")

	newNonce, err := SDK.Account.Nonce(sdk.Client, metadata.NewAccountIdFromKeyPair(acc))
	PanicOnError(err)
	AssertEq(newNonce, currentNonce+1, "New nonce and old nonce + 1 are not the same.")
}
