package main

import (
	"github.com/availproject/avail-go-sdk/src/config"
	"github.com/availproject/avail-go-sdk/src/sdk"
	"github.com/availproject/avail-go-sdk/src/sdk/tx"

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

	keyringPair, err := sdk.KeyringFromSeed(config.Seed)
	if err != nil {
		panic(fmt.Sprintf("cannot create KeyPair:%v", err))
	}

	// create extrinsic
	fmt.Println("Creating extrinsic ...")
	ext, err := sdk.CreateExtrinsic(api, "DataAvailability.submit_data", keyringPair, appID, "my happy data")
	if err != nil {
		fmt.Printf("cannot create extrinsic:%v", err)
	}
	fmt.Println("Extrinsic created successfully for method: DataAvailability.submit_data")

	// submit extrinsic
	fmt.Println("Submitting extrinsic ...")
	WaitFor := sdk.BlockInclusion
	BlockHash, txHash, err := tx.SubmitExtrinsic(api, ext, WaitFor)
	if err != nil {
		fmt.Printf("cannot submit extrinsic:%v", err)
	}
	fmt.Printf("\nExtrinsic submitted successfully with block hash: %v\n and ext hash:%v", BlockHash.Hex(), txHash.Hex())
}
