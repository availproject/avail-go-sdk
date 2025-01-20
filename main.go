package main

import (
	// "go-sdk/examples"

	"go-sdk/metadata"
	daPallet "go-sdk/metadata/pallets/data_availability"
	"go-sdk/primitives"
	SDK "go-sdk/sdk"
	/*
		 	"go-sdk/metadata"
			"go-sdk/primitives"

	*/)

func main() {
	sdk := SDK.NewClient(SDK.LocalEndpoint)

	storageAt, err := sdk.StorageAt(primitives.NewNone[primitives.H256]())
	if err != nil {
		panic(err)
	}

	acc, err := metadata.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	if err != nil {
		panic(err)
	}

	storage := daPallet.StorageAccount{}
	value, err := storage.Fetch(&storageAt, acc)
	if err != nil {
		panic(err)
	}
	println(value.IsSome())
	/*
		 	println(primitives.Hex.ToHex(value.Key))
			println(value.Value.AppId)
			println(value.Value.Owner.ToHuman())
	*/
}
