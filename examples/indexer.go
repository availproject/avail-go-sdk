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

func (this *Indexer) Init() {
	hash, err := this.sdk.Client.FinalizedBlockHash()
	PanicOnError(err)

	block, err := SDK.NewBlock(this.sdk.Client, hash)
	PanicOnError(err)

	number, err := this.sdk.Client.BlockNumber(hash)
	PanicOnError(err)

	this.lock.Lock()
	this.block.block = block
	this.block.hash = hash
	this.block.height = number
	this.lock.Unlock()

}

func (this *Indexer) Run() {
	for {
		block, shutdown := this.fetchBlock()
		if shutdown {
			return
		}

		this.lock.Lock()
		this.block.block = block.block
		this.block.hash = block.hash
		this.block.height = block.height
		this.lock.Unlock()
	}
}

func (this *Indexer) fetchBlock() (IndexedBlock, bool) {
	for {
		if this.shutdown {
			return IndexedBlock{}, true
		}

		hash, err := this.sdk.Client.FinalizedBlockHash()
		PanicOnError(err)
		if this.block.hash == hash {
			time.Sleep(15 * time.Second)
			continue
		}

		block, err := SDK.NewBlock(this.sdk.Client, hash)
		PanicOnError(err)

		number, err := this.sdk.Client.BlockNumber(hash)
		PanicOnError(err)

		return IndexedBlock{hash: hash, height: number, block: block}, false
	}
}

func (this *Indexer) GetBlock(blockNumber uint32) IndexedBlock {
	for {
		if this.shutdown {
			return IndexedBlock{}
		}

		this.lock.Lock()
		block := this.block
		this.lock.Unlock()

		if blockNumber > this.block.height {
			time.Sleep(15 * time.Second)
			continue
		}

		if blockNumber == this.block.height {
			return block
		}

		hash, err := this.sdk.Client.BlockHash(blockNumber)
		PanicOnError(err)

		oldBlock, err := SDK.NewBlock(this.sdk.Client, hash)
		PanicOnError(err)

		number, err := this.sdk.Client.BlockNumber(hash)
		PanicOnError(err)

		return IndexedBlock{hash: hash, height: number, block: oldBlock}
	}
}

func (this *Indexer) Shutdown() {
	this.shutdown = true
}

func (this *Indexer) Callback(cb func(IndexedBlock)) *BlockSubscription {
	sub := this.Subscribe()
	go func() {
		for {
			block := sub.Fetch()
			if this.shutdown || sub.Shutdown {
				return
			}

			cb(block)
		}

	}()
	return &sub
}

func (this *Indexer) Subscribe() BlockSubscription {
	return BlockSubscription{Height: this.block.height, indexer: this}
}

type BlockSubscription struct {
	Height   uint32
	indexer  *Indexer
	Shutdown bool
}

func (this *BlockSubscription) Fetch() IndexedBlock {
	if this.Shutdown {
		return IndexedBlock{}
	}

	block := this.indexer.GetBlock(this.Height)
	this.Height += 1

	return block
}
