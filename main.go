package main

import (
	// "go-sdk/examples"

	sesPall "go-sdk/metadata/pallets/session"
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

	{
		storage := sesPall.StorageValidators{}
		value, err := storage.Fetch(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value[0].ToHuman())
	}
}
