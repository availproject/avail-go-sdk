package examples

import (
	"sync"
	"time"

	"github.com/availproject/avail-go-sdk/primitives"
	"github.com/availproject/avail-go-sdk/sdk"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunIndexer() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	indexer := Indexer{sdk: sdk}
	indexer.Init()
	go indexer.Run()

	sub := indexer.Subscribe()
	sub.Fetch()
	sub.Fetch()
	sub.Fetch()

	indexer.Close()

}

type Indexer struct {
	sdk      SDK.SDK
	shutdown bool
	block    IndexedBlock
	lock     sync.Mutex
}

type IndexedBlock struct {
	hash   primitives.H256
	number uint32
	block  sdk.Block
}

func (this *Indexer) Run() {
	for {
		if this.shutdown {
			return
		}

		hash, err := this.sdk.Client.FinalizedBlockHash()
		PanicOnError(err)
		if this.block.hash == hash {
			time.Sleep(1 * time.Second)
			continue
		}

		block, err := SDK.NewBlock(this.sdk.Client, hash)
		PanicOnError(err)

		number, err := this.sdk.Client.BlockNumber(hash)
		PanicOnError(err)

		this.lock.Lock()
		this.block.block = block
		this.block.hash = hash
		this.block.number = number
		this.lock.Unlock()
	}
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
	this.block.number = number
	this.lock.Unlock()

}

func (this *Indexer) Close() {
	this.shutdown = true
}

func (this *Indexer) GetBlock() (uint32, primitives.H256, SDK.Block) {
	this.lock.Lock()
	number := this.block.number
	hash := this.block.hash
	block := this.block.block
	this.lock.Unlock()

	return number, hash, block
}

func (this *Indexer) Callback(cb func(uint32, primitives.H256, SDK.Block)) {
	go func() {
		sub := this.Subscribe()
		for {
			if this.shutdown {
				return
			}

			blockNumber, blockHash, block := sub.Fetch()
			cb(blockNumber, blockHash, block)
		}

	}()
	return
}

func (this *Indexer) GetHash() primitives.H256 {
	this.lock.Lock()
	hash := this.block.hash
	this.lock.Unlock()
	return hash
}

func (this *Indexer) GetNumber() uint32 {
	this.lock.Lock()
	number := this.block.number
	this.lock.Unlock()
	return number
}

func (this *Indexer) Subscribe() BlockSubscription {
	return BlockSubscription{number: this.block.number, sdk: this.sdk, indexer: this}
}

type BlockSubscription struct {
	number   uint32
	sdk      SDK.SDK
	indexer  *Indexer
	shutdown bool
}

func (this *BlockSubscription) Fetch() (uint32, primitives.H256, SDK.Block) {
	for {
		blockNumber, blockHash, block := this.indexer.GetBlock()
		if this.number > blockNumber {
			time.Sleep(1 * time.Second)
			continue
		}

		if this.number == blockNumber {
			this.number += 1
			return blockNumber, blockHash, block
		}

		hash, err := this.sdk.Client.BlockHash(this.number)
		PanicOnError(err)

		oldBlock, err := SDK.NewBlock(this.sdk.Client, hash)
		PanicOnError(err)

		number, err := this.sdk.Client.BlockNumber(hash)
		PanicOnError(err)

		this.number += 1
		return number, hash, oldBlock
	}
}

func (this *BlockSubscription) Close() {
	this.shutdown = true
}
