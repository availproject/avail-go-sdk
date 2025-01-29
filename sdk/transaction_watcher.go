package sdk

import (
	"github.com/sirupsen/logrus"

	"strconv"
	"time"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type Watcher struct {
	client             *Client
	txHash             prim.H256
	waitFor            uint8
	blockTimeout       uint32
	timeoutBlockNumber prim.Option[uint32]
	sleepDuration      uint32
}

func newWatcher(client *Client, txHash prim.H256, waitFor uint8, blockTimeout uint32, sleepDuration uint32) Watcher {
	return Watcher{
		client:             client,
		txHash:             txHash,
		waitFor:            waitFor,
		blockTimeout:       blockTimeout,
		sleepDuration:      sleepDuration,
		timeoutBlockNumber: prim.NewNone[uint32](),
	}
}

func (this *Watcher) getBlockHash() (prim.H256, error) {
	if this.waitFor == Finalization {
		return this.client.Rpc.Chain.GetFinalizedHead()
	} else {
		return this.client.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
	}
}

func (this *Watcher) getTxEvents(ext *prim.DecodedExtrinsic, blockHash prim.H256) prim.Option[EventRecords] {
	blockEvents, err := this.client.EventsAt(prim.NewSome(blockHash))
	if err != nil {
		logrus.Error(err.Error())
		return prim.NewNone[EventRecords]()
	}

	return prim.NewSome(EventFilterByTxIndex(blockEvents, ext.TxIndex))
}

func (this *Watcher) timeout(iden string, blockNumber uint32) bool {
	if this.timeoutBlockNumber.IsNone() {
		this.timeoutBlockNumber = prim.NewSome(blockNumber + this.blockTimeout)
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.Debug(iden + ": Current Block Number: " + strconv.FormatUint(uint64(blockNumber), 10) + ", Timeout Block Number: " + strconv.FormatUint(uint64(blockNumber+this.blockTimeout+1), 10))
		}
	}

	timeoutBlock := this.timeoutBlockNumber.Unwrap()
	if timeoutBlock < blockNumber {
		return true
	}

	return false
}

func (this *Watcher) runDebugMessage(iden string) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		if this.waitFor == Finalization {
			logrus.Debug(iden + ": Watching for Tx Hash: " + this.txHash.ToHexWith0x() + ", Waiting for finalization")
		} else {
			logrus.Debug(iden + ": Watching for Tx Hash: " + this.txHash.ToHexWith0x() + ", Waiting for inclusion")
		}
	}
}

func (this *Watcher) findExtrinsic(extrinsics []prim.DecodedExtrinsic) int {
	for i := range extrinsics {
		if extrinsics[i].TxHash != this.txHash {
			continue
		}
		return i
	}
	return -1
}

func (this *Watcher) wait() {
	time.Sleep(time.Second * time.Duration(this.sleepDuration))
}

func (this *Watcher) Run() (prim.Option[TransactionDetails], error) {
	iden := this.txHash.ToString()[0:10]
	currentBlockHash := prim.NewNone[prim.H256]()

	this.runDebugMessage(iden)
	for {
		blockHash, err := this.getBlockHash()
		if err != nil {
			return prim.NewNone[TransactionDetails](), err
		}

		if currentBlockHash.IsSome() && currentBlockHash.Unwrap() == blockHash {
			this.wait()
			continue
		}
		currentBlockHash = prim.NewSome(blockHash)

		block, err := this.client.RPCBlockAt(prim.NewSome(blockHash))
		if err != nil {
			return prim.NewNone[TransactionDetails](), err
		}
		blockNumber := block.Header.Number

		extrinsics := block.Extrinsics
		if i := this.findExtrinsic(extrinsics); i != -1 {
			txEvents := this.getTxEvents(&extrinsics[i], blockHash)
			res := TransactionDetails{client: this.client, TxHash: extrinsics[i].TxHash, TxIndex: extrinsics[i].TxIndex, BlockHash: blockHash, BlockNumber: blockNumber, Events: txEvents}
			return prim.NewSome(res), nil
		}

		if this.timeout(iden, blockNumber) {
			break
		}

		this.wait()
	}

	return prim.NewNone[TransactionDetails](), nil
}
