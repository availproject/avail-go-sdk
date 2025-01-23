package main

import (
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func main() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	sdk.Client.Rpc.Kate.BlockLength(primitives.NewNone[primitives.H256]())

}
