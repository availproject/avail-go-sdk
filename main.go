package main

import (
	// "go-sdk/examples"

	npmPool "go-sdk/metadata/pallets/identity"
	"go-sdk/primitives"
	SDK "go-sdk/sdk"
	/*
		 	"go-sdk/metadata"
			"go-sdk/primitives"

	*/)

func main() {
	sdk := SDK.NewClient(SDK.TuringEndpoint)

	storageAt, err := sdk.StorageAt(primitives.NewNone[primitives.H256]())
	if err != nil {
		panic(err)
	}

	{

		storage := npmPool.StorageIdentityOf{}
		value, err := storage.FetchAll(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value[1].Key.ToHuman())
		println(value[1].Value.T0.Info.Display.ToString())
	}
}
