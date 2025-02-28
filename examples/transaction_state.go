package examples

import (
	"fmt"
	"time"

	"github.com/availproject/avail-go-sdk/metadata"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionState() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Transaction will be signed, and sent.
	//
	// There is no guarantee that the transaction was executed at all. It might have been
	// dropped or discarded for various reasons. The caller is responsible for querying future
	// blocks in order to determine the execution status of that transaction.
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	txHash, err := tx.Execute(SDK.Account.Alice(), SDK.NewTransactionOptions().WithAppId(1))
	PanicOnError(err)
	fmt.Println("Tx Hash:", txHash)

	details := []metadata.TransactionState{}
	for {
		details, err = sdk.Client.TransactionState(txHash, false)
		PanicOnError(err)
		if len(details) != 0 {
			break
		}

		time.Sleep(time.Second)
	}
	AssertEq(len(details), 1, "")
	detail := details[0]
	fmt.Println(fmt.Sprintf(`Block Hash: %v, Block Height: %v, Tx Hash: %v, Tx Index: %v`, detail.BlockHash.ToHuman(), detail.BlockHeight, detail.TxHash.ToHuman(), detail.TxIndex))
	fmt.Println(fmt.Sprintf(`Pallet Index: %v, Call Index: %v, Tx Success: %v, Is Finalized: %v`, detail.PalletIndex, detail.CallIndex, detail.TxSuccess, detail.IsFinalized))

	fmt.Println("RunTransactionState finished correctly.")
}
