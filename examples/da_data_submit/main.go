package main

import (
	"avail-go-sdk/src/config"
	"avail-go-sdk/src/sdk"
	"avail-go-sdk/src/sdk/tx"

	"fmt"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("cannot load config:%v", err)
	}
	api, err := sdk.NewSDK(config.ApiURL)
	if err != nil {
		fmt.Printf("cannot create api:%v", err)
	}

	appID := 0

	// if app id is greater than 0 then it must be created before submitting data
	if config.AppID != 0 {
		appID = config.AppID
	}
	fmt.Println("Submitting data ...")
	WaitFor := sdk.BlockInclusion

	// submit data
	BlockHash, txHash, err := tx.SubmitData(api, config.Seed, appID, "my happy data", WaitFor)
	if err != nil {
		fmt.Printf("cannot submit data:%v", err)
	}
	fmt.Printf("Data submitted successfully with block hash: %v\n and ext hash:%v", BlockHash.Hex(), txHash.Hex())
}
