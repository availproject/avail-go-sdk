package main

import (
	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func main() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	acc, err := SDK.Account.Alice()
	if err != nil {
		panic(err)
	}

	val, _ := primitives.NewH256FromHexString("0xea552116539b130effd3404e409a3a9f99c55e47bcb855320bb70bc640b4eab3")

	fungt := metadata.MessageFungibleToken{}
	fungt.Amount = metadata.Balance{Value: SDK.OneAvail()}
	fungt.AssetId = val

	msg := metadata.VectorMessageKind{}
	msg.VariantIndex = 1
	msg.FungibleToken.Set(fungt)
	msg.ArbitraryMessage.Unset()
	{
		tx := sdk.Tx.Vector.SendMessage(msg, val, 0)
		_, err := tx.Execute(acc, SDK.NewTransactionOptions().WithAppId(0))
		if err != nil {
			panic(err)
		}
	}
	{
		tx := sdk.Tx.DataAvailability.SubmitData([]byte("aabbcc"))
		tx.Execute(acc, SDK.NewTransactionOptions().WithAppId(1))
		tx.Execute(acc, SDK.NewTransactionOptions().WithAppId(1))
	}

	_, err = sdk.Client.Rpc.Kate.QueryDataProof(4, primitives.NewSome(val))
	if err != nil {
		panic(err)
	}

}
