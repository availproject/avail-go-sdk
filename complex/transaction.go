package complex

import (
	"strconv"
	"time"

	"github.com/vedhavyas/go-subkey/v2"

	"go-sdk/metadata"
	prim "go-sdk/primitives"
)

type Transaction struct {
	Client  *Client
	Payload metadata.Payload
}

func NewTransaction(client *Client, payload metadata.Payload) Transaction {
	return Transaction{
		Client:  client,
		Payload: payload,
	}
}

func (this *Transaction) Execute(account subkey.KeyPair, options TransactionOptions) prim.H256 {
	return TransactionSignAndSend(this.Client, account, this.Payload, options)
}

func (this *Transaction) ExecuteAndWatch(account subkey.KeyPair, waitForFin bool, options TransactionOptions) TransactionDetails {
	return TransactionSignSendWatch(this.Client, account, this.Payload, waitForFin, options)
}

func (this *Transaction) ExecuteAndWatchFinalization(account subkey.KeyPair, options TransactionOptions) TransactionDetails {
	return TransactionSignSendWatch(this.Client, account, this.Payload, true, options)
}

func (this *Transaction) ExecuteAndWatchInclusion(account subkey.KeyPair, options TransactionOptions) TransactionDetails {
	return TransactionSignSendWatch(this.Client, account, this.Payload, false, options)
}

func TransactionSignAndSend(client *Client, account subkey.KeyPair, payload metadata.Payload, options TransactionOptions) prim.H256 {
	var extra, additional = options.ToPrimitive(client, account.SS58Address(42))
	var tx = prim.CreateSigned(payload.Call, extra, additional, account)

	return client.Send(tx)
}

func TransactionSignSendWatch(client *Client, account subkey.KeyPair, payload metadata.Payload, waitForFin bool, options TransactionOptions) TransactionDetails {
	var extra, additional = options.ToPrimitive(client, account.SS58Address(42))
	var tx = prim.CreateSigned(payload.Call, extra, additional, account)

	var retryCount = 2
	for {
		var txHash = client.Send(tx)
		var maybeDetails = TransactionWatch(client, txHash, waitForFin)
		if maybeDetails.IsSome() {
			return maybeDetails.Unwrap()
		}

		retryCount -= 1
		if retryCount == 0 {
			break
		}
	}

	panic("Something went wrong :(")
}

type TransactionDetails struct {
	Client      *Client
	TxHash      prim.H256
	TxIndex     uint32
	BlockHash   prim.H256
	BlockNumber uint32
	Events      prim.Option[EventRecords]
}

func TransactionWatch(client *Client, txHash prim.H256, waitForFin bool) prim.Option[TransactionDetails] {
	shouldSleep := false
	currentBlockHash := prim.NewNone[prim.H256]()
	timeoutBlockNumber := prim.NewNone[uint32]()
	bTimeout := uint32(3)
	if waitForFin {
		bTimeout = uint32(5)
	}

	if waitForFin {
		println("Watching for Tx Hash: " + txHash.ToHexWith0x() + ", Waiting for finalization")
	} else {
		println("Watching for Tx Hash: " + txHash.ToHexWith0x() + ", Waiting for inclusion")
	}

	for {
		if shouldSleep {
			time.Sleep(time.Second * 3)
		}
		if !shouldSleep {
			shouldSleep = true
		}

		blockHash := prim.H256{}
		if waitForFin {
			blockHash = client.Rpc.Chain.GetFinalizedHead()
		} else {
			blockHash = client.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
		}

		if currentBlockHash.IsSome() {
			if currentBlockHash.Unwrap() == blockHash {
				continue
			}
		}
		currentBlockHash = prim.NewSome(blockHash)

		block := client.GetBlock(prim.NewSome(blockHash))
		blockNumber := block.Header.Number
		println("New Block fetched. Hash: " + blockHash.ToHexWith0x() + ", Number: " + strconv.FormatUint(uint64(blockNumber), 10))

		for _, element := range block.Extrinsics {
			if element.TxHash.ToHexWith0x() == txHash.ToHexWith0x() {
				// Get Events
				events, err := client.GetEvents(prim.NewSome(blockHash))
				if err != nil {
					panic(err)
				}
				events = FilterByTxIndex(events, element.TxIndex)

				details := TransactionDetails{
					TxHash:      txHash,
					TxIndex:     element.TxIndex,
					BlockHash:   blockHash,
					BlockNumber: blockNumber,
					Events:      prim.NewSome(events),
				}
				return prim.NewSome(details)
			}
		}

		if timeoutBlockNumber.IsNone() {
			timeoutBlockNumber = prim.NewSome(blockNumber + bTimeout)
			println("Current Block Number: " + strconv.FormatUint(uint64(blockNumber), 10) + ", Timeout Block Number: " + strconv.FormatUint(uint64(blockNumber+bTimeout+1), 10))
		}

		if timeoutBlockNumber.IsSome() {
			timeoutBlock := timeoutBlockNumber.Unwrap()
			if timeoutBlock < blockNumber {
				break
			}
		}
	}

	return prim.NewNone[TransactionDetails]()
}
