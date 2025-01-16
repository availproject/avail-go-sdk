package main

import (
	// Complex "go-sdk/complex"
	// "go-sdk/metadata"
	// Balances "go-sdk/metadata/pallets/balances"
	"go-sdk/complex"
	"go-sdk/primitives"
	"sync"
	"time"
	// "github.com/vedhavyas/go-subkey/v2"
	// "github.com/vedhavyas/go-subkey/v2/sr25519"
)

func Job(targetBlockNumber uint32, wg *sync.WaitGroup) {
	sdk := NewSDK2("https://mainnet-rpc.avail.so/rpc")
	targetBlockHash := sdk.Client.Rpc.Chain.GetBlockHash(primitives.NewSome(targetBlockNumber))
	println("Block Number:", targetBlockNumber, "Block Hash: ", targetBlockHash.ToHexWith0x())

	targetBlockHash2 := targetBlockHash
	/*
	   	Turing
	   if targetBlockNumber == 297282 {
	   		targetBlockHash2 = sdk.Client.Rpc.Chain.GetBlockHash(primitives.NewSome(targetBlockNumber - 1))
	   	}
	*/
	if err := sdk.Client.InitMetadata(primitives.NewSome(targetBlockHash2)); err != nil {
		panic(err)
	}

	// Decoding Extrinsics and Events
	complex.NewBlock(sdk.Client, targetBlockHash)
	wg.Done()
}

func main() {
	sdk := NewSDK2("https://mainnet-rpc.avail.so/rpc")
	println("Genesis: ", sdk.Client.Rpc.Chain.GetBlockHash(primitives.NewSome(uint32(0))).ToHexWith0x())

	targetBlockNumber := uint32(246585)
	var wg sync.WaitGroup
	for {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go Job(targetBlockNumber, &wg)
			targetBlockNumber += 1
		}

		wg.Wait()
		time.Sleep(time.Millisecond * 50)
	}

	/*
		 	// Go through blocks and see which one is not decodable

			add, err := metadata.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
			if err != nil {
				panic(err)
			}
			println(add.ToHuman())

			uri := "bottom drive obey lake curtain smoke basket hold race lonely fit walk//Alice"
			scheme := new(sr25519.Scheme)
			account, _ := subkey.DeriveKeyPair(scheme, uri)

			var tx = sdk.Tx.DataAvailability.SubmitData([]byte("aabbcc"))
			var options = Complex.NewTransactionOptions().WithAppId(uint32(1))

			var details = tx.ExecuteAndWatch(account, false, options)
			println("Block Hash: ", details.BlockHash.ToHexWith0x())
			println("Block Number: ", details.BlockNumber)
			println("TX Events ", len(details.Events.Unwrap()))

			var eves = details.Events.Unwrap()

			var res = Complex.FindFirst(eves, Balances.EventWithdraw{})
			println("Who:", res.Unwrap().Who.ToHuman())
			println("Amount: ", res.Unwrap().Amount.ToHuman())
	*/
}
