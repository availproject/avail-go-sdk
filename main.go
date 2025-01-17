package main

import (
	// "go-sdk/examples"

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

	storage := daPallet.StorageAppKeys{}
	value, err := storage.FetchKeys(&storageAt)
	if err != nil {
		panic(err)
	}
	println(len(value))
	/*
		 	println(primitives.Hex.ToHex(value.Key))
			println(value.Value.AppId)
			println(value.Value.Owner.ToHuman())
	*/
}
