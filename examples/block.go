package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	balancesPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlock() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	if err != nil {
		panic(err)
	}

	acc := SDK.Account.Alice()

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(1))
	if err != nil {
		panic(err)
	}

	// Fetching
	// Fetching Best Block
	_, err = SDK.NewBestBlock(sdk.Client)
	if err != nil {
		panic(err)
	}

	// Fetching Last Finalized Block
	_, err = SDK.NewFinalizedBlock(sdk.Client)
	if err != nil {
		panic(err)
	}

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
	blobs := block.DataSubmissionAll()
	println("Data Submission Count: ", len(blobs))

	blob := blobs[0]
	println(fmt.Sprintf(`Tx Hash: %v, Tx Index: %v, Tx Data: %v, Tx AppId: %v`, blob.TxHash.ToHexWith0x(), blob.TxIndex, string(blob.Data), blob.AppId))

	accountId := metadata.AccountId{Value: blob.TxSigner.Id.Unwrap()}
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

	println("RunBlock finished correctly.")
}
