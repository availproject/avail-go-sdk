package sdk

import (
	"time"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type Watcher struct {
	client             *Client
	txHash             prim.H256
	waitFor            uint8
	blockCountTimeout  prim.Option[uint32]
	blockHeightTimeout prim.Option[uint32]
	blockFetchInterval uint32
	logger             CustomLogger
}

func NewWatcher(client *Client, txHash prim.H256) Watcher {
	return Watcher{
		client:             client,
		txHash:             txHash,
		waitFor:            Inclusion,
		blockCountTimeout:  prim.None[uint32](),
		blockHeightTimeout: prim.None[uint32](),
		blockFetchInterval: 3_000,
		logger:             NewCustomLogger(prim.H256{}, false),
	}
}

func (this Watcher) Logger(value CustomLogger) Watcher {
	this.logger = value
	return this
}

func (this Watcher) WaitFor(value uint8) Watcher {
	this.waitFor = value
	return this
}

func (this Watcher) TxHash(value prim.H256) Watcher {
	this.txHash = value
	return this
}

func (this Watcher) BlockCountTimeout(value uint32) Watcher {
	this.blockCountTimeout = prim.Some(value)
	return this
}

func (this Watcher) BlockHeightTimeout(value uint32) Watcher {
	this.blockHeightTimeout = prim.Some(value)
	return this
}

// In milliseconds
func (this Watcher) BlockFetchInterval(value uint32) Watcher {
	this.blockFetchInterval = value
	return this
}

func (this *Watcher) getBlockHash() (prim.H256, error) {
	if this.waitFor == Finalization {
		return this.client.FinalizedBlockHash()
	} else {
		return this.client.BestBlockHash()
	}
}

func (this *Watcher) getTxEvents(ext *prim.DecodedExtrinsic, blockHash prim.H256) prim.Option[EventRecords] {
	blockEvents, err := this.client.EventsAt(prim.Some(blockHash))
	if err != nil {
		return prim.None[EventRecords]()
	}

	return prim.Some(EventFilterByTxIndex(blockEvents, ext.TxIndex))
}

func (this *Watcher) Run() (prim.Option[TransactionDetails], error) {
	currentBlockHash := prim.None[prim.H256]()
	blockHeightTimeout, err := this.calculateBlockHeightTimeout()
	if err != nil {
		return prim.None[TransactionDetails](), err
	}
	this.logger.LogWatcherRun(this.waitFor, blockHeightTimeout)

	for {
		block, err := this.fetchNextBlock(&currentBlockHash)
		if err != nil {
			return prim.None[TransactionDetails](), err
		}
		blockHash := currentBlockHash.Unwrap()
		this.logger.LogWatcherNewBlock(&block, blockHash)

		if txDetails := this.findTransaction(&block, blockHash); txDetails.IsSome() {
			details := txDetails.Unwrap()
			this.logger.LogWatcherTxFound(&details)
			return prim.Some(details), nil
		}

		if block.Header.Number >= blockHeightTimeout {
			this.logger.LogWatcherStop()
			return prim.None[TransactionDetails](), nil
		}
	}
}

func (this *Watcher) fetchNextBlock(currentBlockHash *prim.Option[prim.H256]) (RPCBlock, error) {
	for {
		blockHash, err := this.getBlockHash()
		if err != nil {
			return RPCBlock{}, err
		}

		if currentBlockHash.IsSome() && currentBlockHash.Unwrap() == blockHash {
			time.Sleep(time.Millisecond * time.Duration(this.blockFetchInterval))
			continue
		}
		*currentBlockHash = prim.Some(blockHash)

		return this.client.RPCBlockAt(prim.Some(blockHash))
	}

}

func (this *Watcher) calculateBlockHeightTimeout() (uint32, error) {
	if this.blockHeightTimeout.IsSome() {
		return this.blockHeightTimeout.Unwrap(), nil
	}

	count := this.blockCountTimeout.UnwrapOr(5)
	current_height, err := this.client.BestBlockNumber()

	return current_height + count, err
}

func (this *Watcher) findTransaction(block *RPCBlock, blockHash prim.H256) prim.Option[TransactionDetails] {
	blockNumber := block.Header.Number

	extrinsics := block.Extrinsics
	for i := 0; i < len(extrinsics); i++ {
		if extrinsics[i].TxHash != this.txHash {
			continue
		}

		txEvents := this.getTxEvents(&extrinsics[i], blockHash)
		res := TransactionDetails{client: this.client, TxHash: extrinsics[i].TxHash, TxIndex: extrinsics[i].TxIndex, BlockHash: blockHash, BlockNumber: blockNumber, Events: txEvents}
		return prim.Some(res)
	}

	return prim.None[TransactionDetails]()
}
