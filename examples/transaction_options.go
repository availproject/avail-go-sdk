package examples

import (
	"github.com/availproject/avail-go-sdk/metadata"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionOptions() {
	runAppId()
	runNonce()

	println("RunTransactionOptions finished correctly.")
}

func runAppId() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	if err != nil {
		panic(err)
	}

	appId := uint32(5)
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))
	options := SDK.NewTransactionOptions().WithAppId(appId)
	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), options)
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

	genTx := block.TransactionByHash(res.TxHash).SafeUnwrap()
	foundAppId := genTx.Signed().SafeUnwrap().AppId
	if appId != foundAppId {
		panic("Wrong appid was used.")
	}
}

func runNonce() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	if err != nil {
		panic(err)
	}

	acc := SDK.Account.Alice()
	currentNonce, err := SDK.Account.Nonce(sdk.Client, metadata.NewAccountIdFromKeyPair(acc))
	if err != nil {
		panic(err)
	}

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))
	options := SDK.NewTransactionOptions().WithNonce(currentNonce).WithAppId(5)
	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), options)
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

	genTx := block.TransactionByHash(res.TxHash).SafeUnwrap()
	foundNonce := genTx.Signed().SafeUnwrap().Nonce

	if foundNonce != currentNonce {
		panic("Wrong Nonce 1")
	}

	newNonce, err := SDK.Account.Nonce(sdk.Client, metadata.NewAccountIdFromKeyPair(acc))
	if err != nil {
		panic(err)
	}

	if newNonce != (currentNonce + 1) {
		panic("Wrong Nonce 2")
	}

}
