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
		blockFetchInterval: 5_000,
		logger:             NewCustomLogger(prim.H256{}, false),
	}
}

func (w Watcher) Logger(value CustomLogger) Watcher {
	w.logger = value
	return w
}

func (w Watcher) WaitFor(value uint8) Watcher {
	w.waitFor = value
	return w
}

func (w Watcher) TxHash(value prim.H256) Watcher {
	w.txHash = value
	return w
}

func (w Watcher) BlockCountTimeout(value uint32) Watcher {
	w.blockCountTimeout = prim.Some(value)
	return w
}

func (w Watcher) BlockHeightTimeout(value uint32) Watcher {
	w.blockHeightTimeout = prim.Some(value)
	return w
}

// In milliseconds
func (w Watcher) BlockFetchInterval(value uint32) Watcher {
	w.blockFetchInterval = value
	return w
}

func (w *Watcher) getBlockHash(waitFor uint8) (prim.H256, error) {
	if waitFor == Finalization {
		return w.client.FinalizedBlockHash()
	} else {
		return w.client.BestBlockHash()
	}
}

func (w *Watcher) getBlockHeight(waitFor uint8) (uint32, error) {
	if waitFor == Finalization {
		return w.client.FinalizedBlockNumber()
	} else {
		return w.client.BestBlockNumber()
	}
}

func (w *Watcher) getTxEvents(ext *prim.DecodedExtrinsic, blockHash prim.H256) prim.Option[EventRecords] {
	blockEvents, err := w.client.EventsAt(prim.Some(blockHash))
	if err != nil {
		return prim.None[EventRecords]()
	}

	return prim.Some(EventFilterByTxIndex(blockEvents, ext.TxIndex))
}

func (w *Watcher) Run() (prim.Option[TransactionDetails], error) {
	blockHeightTimeout, err := w.calculateBlockHeightTimeout()
	if err != nil {
		return prim.None[TransactionDetails](), err
	}
	w.logger.LogWatcherRun(w.waitFor, blockHeightTimeout)

	if w.waitFor == Finalization {
		return w.runFinalized(blockHeightTimeout)

	} else {
		return w.runIncluded(blockHeightTimeout)
	}
}

func (w *Watcher) runFinalized(blockHeightTimeout uint32) (prim.Option[TransactionDetails], error) {
	nextBlockHeight, err := w.getBlockHeight(Finalization)
	if err != nil {
		return prim.None[TransactionDetails](), err
	}

	for {
		block, blockHash, err := w.fetchNextBlockFinalized(nextBlockHeight)
		if err != nil {
			return prim.None[TransactionDetails](), err
		}
		w.logger.LogWatcherNewBlock(&block, blockHash)

		if txDetails := w.findTransaction(&block, blockHash); txDetails.IsSome() {
			details := txDetails.Unwrap()
			w.logger.LogWatcherTxFound(&details)
			return prim.Some(details), nil
		}

		if block.Header.Number >= blockHeightTimeout {
			w.logger.LogWatcherStop()
			return prim.None[TransactionDetails](), nil
		}

		nextBlockHeight += 1
	}
}

func (w *Watcher) fetchNextBlockFinalized(nextBlockHeight uint32) (RPCBlock, prim.H256, error) {
	for {
		blockHeight, err := w.getBlockHeight(Finalization)
		if err != nil {
			return RPCBlock{}, prim.H256{}, err
		}

		if nextBlockHeight > blockHeight {
			time.Sleep(time.Millisecond * time.Duration(w.blockFetchInterval))
			continue
		}

		blockHash, err := w.client.BlockHash(nextBlockHeight)
		if err != nil {
			return RPCBlock{}, prim.H256{}, err
		}

		block, err := w.client.RPCBlockAt(prim.Some(blockHash))
		if err != nil {
			return RPCBlock{}, prim.H256{}, err
		}

		return block, blockHash, nil
	}
}

func (w *Watcher) runIncluded(blockHeightTimeout uint32) (prim.Option[TransactionDetails], error) {
	currentBlockHash := prim.None[prim.H256]()

	for {
		block, err := w.fetchNextBlockIncluded(&currentBlockHash)
		if err != nil {
			return prim.None[TransactionDetails](), err
		}
		blockHash := currentBlockHash.Unwrap()
		w.logger.LogWatcherNewBlock(&block, blockHash)

		if txDetails := w.findTransaction(&block, blockHash); txDetails.IsSome() {
			details := txDetails.Unwrap()
			w.logger.LogWatcherTxFound(&details)
			return prim.Some(details), nil
		}

		if block.Header.Number >= blockHeightTimeout {
			w.logger.LogWatcherStop()
			return prim.None[TransactionDetails](), nil
		}
	}
}

func (w *Watcher) fetchNextBlockIncluded(currentBlockHash *prim.Option[prim.H256]) (RPCBlock, error) {
	for {
		blockHash, err := w.getBlockHash(Inclusion)
		if err != nil {
			return RPCBlock{}, err
		}

		if currentBlockHash.IsSome() && currentBlockHash.Unwrap() == blockHash {
			time.Sleep(time.Millisecond * time.Duration(w.blockFetchInterval))
			continue
		}
		*currentBlockHash = prim.Some(blockHash)

		return w.client.RPCBlockAt(prim.Some(blockHash))
	}
}

func (w *Watcher) calculateBlockHeightTimeout() (uint32, error) {
	if w.blockHeightTimeout.IsSome() {
		return w.blockHeightTimeout.Unwrap(), nil
	}

	count := w.blockCountTimeout.UnwrapOr(32)
	current_height, err := w.client.BestBlockNumber()

	return current_height + count, err
}

func (w *Watcher) findTransaction(block *RPCBlock, blockHash prim.H256) prim.Option[TransactionDetails] {
	blockNumber := block.Header.Number

	extrinsics := block.Extrinsics
	for i := 0; i < len(extrinsics); i++ {
		if extrinsics[i].TxHash != w.txHash {
			continue
		}

		txEvents := w.getTxEvents(&extrinsics[i], blockHash)
		res := TransactionDetails{client: w.client, TxHash: extrinsics[i].TxHash, TxIndex: extrinsics[i].TxIndex, BlockHash: blockHash, BlockNumber: blockNumber, Events: txEvents}
		return prim.Some(res)
	}

	return prim.None[TransactionDetails]()
}
