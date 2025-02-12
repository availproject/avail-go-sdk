package main

import (
	"github.com/availproject/avail-go-sdk/examples"
	prim "github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func main() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	examples.PanicOnError(err)

	{
		// system_transaction_state
		hash, err := prim.NewBlockHashFromHexString("0x92cdb77314063a01930b093516d19a453399710cc8ae635ff5ab6cf76b26f218")
		examples.PanicOnError(err)

		values, err := sdk.Client.Rpc.System.TransactionState(hash, false)
		examples.PanicOnError(err)
		for _, val := range values {
			println(val.BlockHash.ToHuman())
			println(val.BlockHeight)
			println(val.TxHash.ToHuman())
			println(val.TxIndex)
			println(val.TxSuccess)
			println(val.PalletIndex)
			println(val.CallIndex)
			println(val.IsFinalized)
		}
	}
}
