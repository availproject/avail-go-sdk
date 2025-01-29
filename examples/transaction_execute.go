package examples

import (
	"fmt"

	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionExecute() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	txHash, err := tx.Execute(SDK.Account.Alice(), SDK.NewTransactionOptions().WithAppId(1))
	PanicOnError(err)
	fmt.Println("Tx Hash:", txHash)

	fmt.Println("RunTransactionExecute finished correctly.")
}
