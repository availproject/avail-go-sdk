package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunAccountNonce() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	// Via RPC
	nonce, err := sdk.Client.Rpc.System.AccountNextIndex("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	PanicOnError(err)

	fmt.Println("RPC Nonce: ", nonce)

	// Via Abstraction
	accountId, err := primitives.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	PanicOnError(err)

	nonce2, err := SDK.Account.Nonce(sdk.Client, accountId)
	PanicOnError(err)

	fmt.Println("Abstraction Nonce: ", nonce2)

	fmt.Println("RunAccountNonce finished correctly.")
}
