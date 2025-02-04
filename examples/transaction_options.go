package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionOptions() {
	RunTransactionOptionsAppId()
	RunTransactionOptionsNonce()
	RunTransactionOptionsMortality()
	RunTransactionOptionsTip()

	fmt.Println("RunTransactionOptions finished correctly.")
}

func RunTransactionOptionsAppId() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Setting AppId
	appId := uint32(5)
	options := SDK.NewTransactionOptions().WithAppId(appId)

	// Executing Transaction
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))
	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), options)
	PanicOnError(err)

	AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

	block, err := SDK.NewBlock(sdk.Client, res.BlockHash)
	PanicOnError(err)

	// Checking is the App Id was used correctly
	blockTxs := block.Transactions(SDK.Filter{}.WTxHash(res.TxHash))
	AssertEq(len(blockTxs), 1, "")
	foundAppId := blockTxs[0].AppId().UnsafeUnwrap()
	AssertEq(appId, foundAppId, "App Ids are not the same")

	fmt.Println("RunTransactionOptionsAppId finished correctly.")
}

func RunTransactionOptionsNonce() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Getting Nonce
	acc := SDK.Account.Alice()
	currentNonce, err := SDK.Account.Nonce(sdk.Client, primitives.NewAccountIdFromKeyPair(acc))
	PanicOnError(err)

	// Executing Transaction
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))
	options := SDK.NewTransactionOptions().WithNonce(currentNonce).WithAppId(5)
	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), options)
	PanicOnError(err)

	AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

	block, err := SDK.NewBlock(sdk.Client, res.BlockHash)
	PanicOnError(err)

	// Checking is the Nonce was used correctly
	blockTxs := block.Transactions(SDK.Filter{}.WTxHash(res.TxHash))
	AssertEq(len(blockTxs), 1, "")
	foundNonce := blockTxs[0].Nonce().UnsafeUnwrap()
	AssertEq(foundNonce, currentNonce, "Nonces are not the same")

	newNonce, err := SDK.Account.Nonce(sdk.Client, primitives.NewAccountIdFromKeyPair(acc))
	PanicOnError(err)
	AssertEq(newNonce, currentNonce+1, "New nonce and old nonce + 1 are not the same.")

	fmt.Println("RunTransactionOptionsNonce finished correctly.")
}

func RunTransactionOptionsMortality() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Setting Mortality
	mortality := uint32(16)
	options := SDK.NewTransactionOptions().WithMortality(mortality).WithAppId(1)

	// Executing Transaction
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))
	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), options)
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

	block, err := SDK.NewBlock(sdk.Client, res.BlockHash)
	PanicOnError(err)

	// Checking if the Mortality is the same as the one expected
	blockTxs := block.Transactions(SDK.Filter{}.WTxHash(res.TxHash))
	AssertEq(len(blockTxs), 1, "")
	actualMortality := uint32(blockTxs[0].Mortality().UnsafeUnwrap().Period)
	AssertEq(actualMortality, mortality, "Mortalities are not the same.")

	fmt.Println("RunTransactionOptionsMortality finished correctly.")
}

func RunTransactionOptionsTip() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Setting Tip
	tip := SDK.OneAvail()
	options := SDK.NewTransactionOptions().WithTip(tip).WithAppId(1)

	// Executing Transaction
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))
	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), options)
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().UnsafeUnwrap(), "Transaction is supposed to succeed")

	block, err := SDK.NewBlock(sdk.Client, res.BlockHash)
	PanicOnError(err)

	// Checking if the Tip is the same as the one expected
	blockTxs := block.Transactions(SDK.Filter{}.WTxHash(res.TxHash))
	AssertEq(len(blockTxs), 1, "")
	actualTip := blockTxs[0].Tip().UnsafeUnwrap()
	AssertEq(actualTip, tip, "Tips are not the same.")

	fmt.Println("RunTransactionOptionsTip finished correctly.")
}
