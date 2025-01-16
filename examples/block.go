package examples

import (
	"fmt"

	"go-sdk/metadata"
	balancesPallet "go-sdk/metadata/pallets/balances"
	SDK "go-sdk/sdk"
)

func run_block() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	acc, err := SDK.Account.Alice()
	if err != nil {
		panic(err)
	}

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(1))
	if err != nil {
		panic(err)
	}

	// Fetching
	// Fetching Best Block
	_, _ = SDK.NewBestBlock(sdk.Client)

	// Fetching Last Finalized Block
	_, _ = SDK.NewFinalizedBlock(sdk.Client)

	// Fetching block with Block Hash
	block, err := SDK.NewBlock(sdk.Client, res.BlockHash)
	if err != nil {
		panic(err)
	}

	// Transactions
	//
	// Available methods:
	// TransactionAll, TransactionBySigner, TransactionByIndex, TransactionByHash, TransactionByAppId
	allTxs := block.TransactionAll()
	println("Transaction Count: ", len(allTxs))

	// Data Submission
	//
	// Available methods:
	// DataSubmissionAll, DataSubmissionBySigner, DataSubmissionByIndex, DataSubmissionByHash, DataSubmissionByAppId
	allDS := block.DataSubmissionAll()
	println("Data Submission Count: ", len(allDS))

	ds := allDS[0]
	println(fmt.Sprintf(`Tx Hash: %v, Tx Index: %v, Tx Data: %v, Tx AppId: %v`, ds.TxHash.ToHexWith0x(), ds.TxIndex, string(ds.Data), ds.AppId))

	accountId := metadata.AccountId{Value: ds.TxSigner.Id.Unwrap()}
	println(fmt.Sprintf(`Tx Signer: %v`, accountId.ToHuman()))

	// Events
	//
	events := block.Events().Unwrap()

	// Find First Event
	withdrawEvent := SDK.EventFindFirst(events, balancesPallet.EventWithdraw{}).Unwrap()
	println("Who:", withdrawEvent.Who.ToHuman())
	println("Amount: ", withdrawEvent.Amount.ToHuman())

	// Find First Event - Checked
	withdrawEvent2, err := SDK.EventFindFirstChecked(events, balancesPallet.EventWithdraw{})
	if err != nil {
		// Failed to decode events
		panic(err)
	}
	if withdrawEvent2.IsNone() {
		panic("Deposit Event was not found")
	}
	withdrawEvent = withdrawEvent2.Unwrap()
	println("Who:", withdrawEvent.Who.ToHuman())
	println("Amount: ", withdrawEvent.Amount.ToHuman())
}
