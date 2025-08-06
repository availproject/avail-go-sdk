package examples

import (
	"fmt"
	"sync"
	"time"

	"github.com/availproject/avail-go-sdk/primitives"
	"github.com/availproject/avail-go-sdk/sdk"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunIndexer() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	indexer := Indexer{sdk: sdk}
	// Initializing indexer with default values
	indexer.Init()
	// Running indexer in the background
	go indexer.Run()

	// Fetching blocks in procedural way
	sub := indexer.Subscribe()
	for i := 0; i < 3; i++ {
		block := sub.Fetch()
		fmt.Println(fmt.Sprintf("Current: Block Height: %v, Block hash: %v", block.height, block.hash))
	}

	// Fetching historical blocks
	sub.Height = sub.Height - 100
	for i := 0; i < 3; i++ {
		block := sub.Fetch()
		fmt.Println(fmt.Sprintf("Historical: Block Height: %v, Block hash: %v", block.height, block.hash))
	}

	// Using Callbacks
	callBack := func(block IndexedBlock) {
		fmt.Println(fmt.Sprintf("Callback: Block Height: %v, Block hash: %v", block.height, block.hash))
	}
	sub2 := indexer.Callback(callBack)
	time.Sleep(25 * time.Second)

	sub2.Shutdown = true
	indexer.Shutdown()

	time.Sleep(3 * time.Second)
	fmt.Println("RunIndexer finished correctly.")
}

type IndexedBlock struct {
	hash   primitives.H256
	height uint32
	block  sdk.Block
}

type Indexer struct {
	sdk      SDK.SDK
	shutdown bool
	block    IndexedBlock
	lock     sync.Mutex
}

func (indexer *Indexer) Init() {
	hash, err := indexer.sdk.Client.FinalizedBlockHash()
	PanicOnError(err)

	block, err := SDK.NewBlock(indexer.sdk.Client, hash)
	PanicOnError(err)

	number, err := indexer.sdk.Client.BlockNumber(hash)
	PanicOnError(err)

	indexer.lock.Lock()
	indexer.block.block = block
	indexer.block.hash = hash
	indexer.block.height = number
	indexer.lock.Unlock()

}

func (indexer *Indexer) Run() {
	for {
		block, shutdown := indexer.fetchBlock()
		if shutdown {
			return
		}

		indexer.lock.Lock()
		indexer.block.block = block.block
		indexer.block.hash = block.hash
		indexer.block.height = block.height
		indexer.lock.Unlock()
	}
}

func (indexer *Indexer) fetchBlock() (IndexedBlock, bool) {
	for {
		if indexer.shutdown {
			return IndexedBlock{}, true
		}

		hash, err := indexer.sdk.Client.FinalizedBlockHash()
		PanicOnError(err)
		if indexer.block.hash == hash {
			time.Sleep(15 * time.Second)
			continue
		}

		block, err := SDK.NewBlock(indexer.sdk.Client, hash)
		PanicOnError(err)

		number, err := indexer.sdk.Client.BlockNumber(hash)
		PanicOnError(err)

		return IndexedBlock{hash: hash, height: number, block: block}, false
	}
}

func (indexer *Indexer) GetBlock(blockNumber uint32) IndexedBlock {
	for {
		if indexer.shutdown {
			return IndexedBlock{}
		}

		indexer.lock.Lock()
		block := indexer.block
		indexer.lock.Unlock()

		if blockNumber > indexer.block.height {
			time.Sleep(15 * time.Second)
			continue
		}

		if blockNumber == indexer.block.height {
			return block
		}

		hash, err := indexer.sdk.Client.BlockHash(blockNumber)
		PanicOnError(err)

		oldBlock, err := SDK.NewBlock(indexer.sdk.Client, hash)
		PanicOnError(err)

		number, err := indexer.sdk.Client.BlockNumber(hash)
		PanicOnError(err)

		return IndexedBlock{hash: hash, height: number, block: oldBlock}
	}
}

func (indexer *Indexer) Shutdown() {
	indexer.shutdown = true
}

func (indexer *Indexer) Callback(cb func(IndexedBlock)) *BlockSubscription {
	sub := indexer.Subscribe()
	go func() {
		for {
			block := sub.Fetch()
			if indexer.shutdown || sub.Shutdown {
				return
			}

			cb(block)
		}

	}()
	return &sub
}

func (indexer *Indexer) Subscribe() BlockSubscription {
	return BlockSubscription{Height: indexer.block.height, indexer: indexer}
}

type BlockSubscription struct {
	Height   uint32
	indexer  *Indexer
	Shutdown bool
}

func (indexer *BlockSubscription) Fetch() IndexedBlock {
	if indexer.Shutdown {
		return IndexedBlock{}
	}

	block := indexer.indexer.GetBlock(indexer.Height)
	indexer.Height += 1

	return block
}
