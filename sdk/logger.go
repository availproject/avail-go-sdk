package sdk

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
	"github.com/sirupsen/logrus"
	"github.com/vedhavyas/go-subkey/v2"
)

type CustomLogger struct {
	txHash  prim.H256
	marker  string
	enabled bool
}

func NewCustomLogger(txHash prim.H256, enabled bool) CustomLogger {
	marker := txHash.ToString()[0:10]
	return CustomLogger{
		txHash:  txHash,
		marker:  marker,
		enabled: enabled,
	}
}

func NewCustomLoggerEmpty(marker string, enabled bool) CustomLogger {
	return CustomLogger{
		txHash:  prim.H256{},
		marker:  marker,
		enabled: enabled,
	}
}

func (l *CustomLogger) Enabled(value bool) {
	l.enabled = value
}

func (l *CustomLogger) LogWatcherRun(waitFor uint8, blockHeightTimeout uint32) {
	if !logrus.IsLevelEnabled(logrus.DebugLevel) || !l.enabled {
		return
	}

	logrus.Debug(fmt.Sprintf("%v: Watching for Tx Hash: %v. Waiting for: %v, Block height timeout: %v", l.marker, l.txHash, waitFor, blockHeightTimeout))
}

func (l *CustomLogger) LogWatcherNewBlock(block *RPCBlock, blockHash prim.H256) {
	if !logrus.IsLevelEnabled(logrus.DebugLevel) || !l.enabled {
		return
	}

	logrus.Debug(fmt.Sprintf("%v: New block fetched. Hash: %v, Number: %v", l.marker, blockHash, block.Header.Number))
}

func (l *CustomLogger) LogWatcherTxFound(details *TransactionDetails) {
	if !logrus.IsLevelEnabled(logrus.DebugLevel) || !l.enabled {
		return
	}

	logrus.Debug(fmt.Sprintf("%v: Transaction was found. Tx Hash: %v, Tx Index: %v, Block Hash: %v, Block Number: %v", l.marker, details.TxHash, details.TxIndex, details.BlockHash, details.BlockNumber))
}

func (l *CustomLogger) LogWatcherStop() {
	if !logrus.IsLevelEnabled(logrus.DebugLevel) || !l.enabled {
		return
	}

	logrus.Debug(fmt.Sprintf("%v: No more blocks to search. Failed to find transaction. Tx Hash: %v", l.marker, l.txHash))
}

func (l *CustomLogger) LogTxSubmitted(keypair *subkey.KeyPair, mortality uint64) {
	if !logrus.IsLevelEnabled(logrus.DebugLevel) || !l.enabled {
		return
	}

	address := (*keypair).SS58Address(42)
	logrus.Debug(fmt.Sprintf("%v: Transaction was submitted. Account: %v, TxHash: %v, Mortality Period: %v", l.marker, address, l.txHash, mortality))
}

func (l *CustomLogger) LogTxSubmitting(keypair *subkey.KeyPair, payload *metadata.Payload, nonce uint32, appId uint32) {
	if !logrus.IsLevelEnabled(logrus.DebugLevel) || !l.enabled {
		return
	}

	address := (*keypair).SS58Address(42)
	logrus.Debug(fmt.Sprintf("Signing and submitting new transaction. Account: %v, Nonce: %v, Pallet Name: %v, Call Name: %v, App Id: %v", address, nonce, payload.PalletName(), payload.CallName(), appId))
}

func (l *CustomLogger) LogTxRetryAbort() {
	logrus.Warn(fmt.Sprintf("%v: Failed to submit and find transaction. Aborting. Tx Hash: %v", l.marker, l.txHash))
}

func (l *CustomLogger) LogTxRetry() {
	if !logrus.IsLevelEnabled(logrus.DebugLevel) || !l.enabled {
		return
	}

	logrus.Debug(fmt.Sprintf("%v: Failed to submit and find transaction. Retrying.", l.marker))
}

func (l *CustomLogger) LogRpcRetry(method string) {
	if !logrus.IsLevelEnabled(logrus.DebugLevel) || !l.enabled {
		return
	}

	logrus.Debug(fmt.Sprintf("%v: Failed to get concrete value from RPC call. Method: %v, Getting `null` as result. Retrying.", l.marker, method))
}

func (l *CustomLogger) LogRpcRetryAbort(method string) {
	logrus.Warn(fmt.Sprintf("%v: Failed to get concrete value from RPC call. Method: %v, Getting `null` as result. Aborting.", l.marker, method))
}
