package sdk

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vedhavyas/go-subkey/v2"

	"errors"
	"strconv"
	"time"

	"github.com/availproject/avail-go-sdk/metadata"
	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

const Finalization = uint8(0)
const Inclusion = uint8(1)

type Transaction struct {
	client  *Client
	Payload metadata.Payload
}

func NewTransaction(client *Client, payload metadata.Payload) Transaction {
	return Transaction{
		client:  client,
		Payload: payload,
	}
}

func (this *Transaction) Execute(account subkey.KeyPair, options TransactionOptions) (prim.H256, error) {
	return TransactionSignAndSend(this.client, account, this.Payload, options)
}

func (this *Transaction) ExecuteAndWatch(account subkey.KeyPair, waitFor uint8, blockTimeout uint32, options TransactionOptions) (TransactionDetails, error) {
	return TransactionSignSendWatch(this.client, account, this.Payload, waitFor, blockTimeout, options)
}

func (this *Transaction) ExecuteAndWatchFinalization(account subkey.KeyPair, options TransactionOptions) (TransactionDetails, error) {
	return TransactionSignSendWatch(this.client, account, this.Payload, Finalization, 5, options)
}

func (this *Transaction) ExecuteAndWatchInclusion(account subkey.KeyPair, options TransactionOptions) (TransactionDetails, error) {
	return TransactionSignSendWatch(this.client, account, this.Payload, Inclusion, 3, options)
}

func (this *Transaction) PaymentQueryFeeDetails(account subkey.KeyPair, options TransactionOptions) (metadata.InclusionFee, error) {
	extra, additional, err := options.ToPrimitive(this.client, account.SS58Address(42))
	if err != nil {
		return metadata.InclusionFee{}, err
	}
	tx, err := prim.CreateSigned(this.Payload.Call, extra, additional, account)
	if err != nil {
		return metadata.InclusionFee{}, err
	}

	return this.client.Rpc.Payment.QueryFeeDetails(tx.Value, prim.NewNone[prim.H256]())
}

func TransactionSignAndSend(client *Client, account subkey.KeyPair, payload metadata.Payload, options TransactionOptions) (prim.H256, error) {
	extra, additional, err := options.ToPrimitive(client, account.SS58Address(42))
	if err != nil {
		return prim.H256{}, err
	}
	tx, err := prim.CreateSigned(payload.Call, extra, additional, account)
	if err != nil {
		return prim.H256{}, err
	}

	return client.Send(tx)
}

func TransactionSignSendWatch(client *Client, account subkey.KeyPair, payload metadata.Payload, waitFor uint8, blockTimeout uint32, options TransactionOptions) (TransactionDetails, error) {
	extra, additional, err := options.ToPrimitive(client, account.SS58Address(42))
	if err != nil {
		return TransactionDetails{}, err
	}

	var retryCount = 3
	for {
		tx, err := prim.CreateSigned(payload.Call, extra, additional, account)
		if err != nil {
			return TransactionDetails{}, err
		}

		txHash, err := client.Send(tx)
		if err != nil {
			return TransactionDetails{}, err
		}
		iden := txHash.ToString()[0:10]
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.Debug(fmt.Sprintf("%v: Transaction was submitted. Account: %v, TxHash: %v", iden, account.SS58Address(42), txHash.ToHexWith0x()))
		}
		maybeDetails, err := TransactionWatch(client, txHash, waitFor, blockTimeout)
		if err != nil {
			return TransactionDetails{}, err
		}
		if maybeDetails.IsSome() {
			return maybeDetails.Unwrap(), nil
		}

		if retryCount == 0 {
			if logrus.IsLevelEnabled(logrus.DebugLevel) {
				logrus.Debug(fmt.Sprintf("%v: Failed to find transaction. TxHash: %v. Aborting", iden, txHash.ToHexWith0x()))
			}
			break
		}

		RegenerateEra(client, &extra, &additional)

		retryCount -= 1
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.Debug(fmt.Sprintf("%v: Failed to find transaction. TxHash: %v, Retry Count: %v/3. Trying again", iden, txHash.ToHexWith0x(), retryCount))
		}
	}

	return TransactionDetails{}, errors.New("Failed to submit transaction. Tried 3 times.")
}

type TransactionDetails struct {
	Client      *Client
	TxHash      prim.H256
	TxIndex     uint32
	BlockHash   prim.H256
	BlockNumber uint32
	Events      prim.Option[EventRecords]
}

func TransactionWatch(client *Client, txHash prim.H256, waitFor uint8, blockTimeout uint32) (prim.Option[TransactionDetails], error) {
	shouldSleep := false
	currentBlockHash := prim.NewNone[prim.H256]()
	timeoutBlockNumber := prim.NewNone[uint32]()
	iden := txHash.ToString()[0:10]
	var err error
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		if waitFor == Finalization {
			logrus.Debug(iden + ": Watching for Tx Hash: " + txHash.ToHexWith0x() + ", Waiting for finalization")
		} else {
			logrus.Debug(iden + ": Watching for Tx Hash: " + txHash.ToHexWith0x() + ", Waiting for inclusion")
		}
	}

	for {
		if shouldSleep {
			time.Sleep(time.Second * 3)
		}
		if !shouldSleep {
			shouldSleep = true
		}

		blockHash := prim.H256{}
		if waitFor == Finalization {
			blockHash, err = client.Rpc.Chain.GetFinalizedHead()
			if err != nil {
				return prim.NewNone[TransactionDetails](), err
			}
		} else {
			blockHash, err = client.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
			if err != nil {
				return prim.NewNone[TransactionDetails](), err
			}
		}

		if currentBlockHash.IsSome() {
			if currentBlockHash.Unwrap() == blockHash {
				continue
			}
		}
		currentBlockHash = prim.NewSome(blockHash)

		block, err := client.RPCBlockAt(prim.NewSome(blockHash))
		if err != nil {
			return prim.NewNone[TransactionDetails](), err
		}

		blockNumber := block.Header.Number
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.Debug(iden + ": New Block fetched. Hash: " + blockHash.ToHexWith0x() + ", Number: " + strconv.FormatUint(uint64(blockNumber), 10))
		}

		for _, element := range block.Extrinsics {
			if element.TxHash.ToHexWith0x() == txHash.ToHexWith0x() {

				// Get Events
				blockEvents, err := client.EventsAt(prim.NewSome(blockHash))
				events := prim.NewNone[EventRecords]()
				if err != nil {
					logrus.Error(err.Error())
				} else {
					events.Set(EventFilterByTxIndex(blockEvents, element.TxIndex))
				}

				details := TransactionDetails{
					TxHash:      txHash,
					TxIndex:     element.TxIndex,
					BlockHash:   blockHash,
					BlockNumber: blockNumber,
					Events:      events,
				}

				if logrus.IsLevelEnabled(logrus.DebugLevel) {
					logrus.Debug(fmt.Sprintf("%v: Transaction was found. Tx Hash: %v, Tx Index: %v, Block Hash: %v, Block Number: %v", iden, details.TxHash.ToHexWith0x(), details.TxIndex, details.BlockHash.ToHexWith0x(), details.BlockNumber))
				}

				return prim.NewSome(details), nil
			}
		}

		if timeoutBlockNumber.IsNone() {
			timeoutBlockNumber = prim.NewSome(blockNumber + blockTimeout)
			if logrus.IsLevelEnabled(logrus.DebugLevel) {
				logrus.Debug(iden + ": Current Block Number: " + strconv.FormatUint(uint64(blockNumber), 10) + ", Timeout Block Number: " + strconv.FormatUint(uint64(blockNumber+blockTimeout+1), 10))
			}
		}

		if timeoutBlockNumber.IsSome() {
			timeoutBlock := timeoutBlockNumber.Unwrap()
			if timeoutBlock < blockNumber {
				break
			}
		}
	}

	return prim.NewNone[TransactionDetails](), nil
}

func (this *TransactionDetails) IsSuccessful() (bool, error) {
	if this.Events.IsNone() {
		return false, errors.New("No events were decoded.")
	}
	events := this.Events.Unwrap()

	maybeFound, err := EventFindFirstChecked(events, syPallet.EventExtrinsicFailed{})
	if err != nil {
		return false, err
	}

	if maybeFound.IsNone() {
		return true, nil
	}

	return false, nil
}
