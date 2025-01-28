package sdk

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vedhavyas/go-subkey/v2"

	"errors"

	"github.com/availproject/avail-go-sdk/metadata"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
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

func (this *Transaction) ExecuteAndWatch(account subkey.KeyPair, waitFor uint8, options TransactionOptions, blockTimeout uint32) (TransactionDetails, error) {
	return TransactionSignSendWatch(this.client, account, this.Payload, waitFor, options, blockTimeout, 3)
}

func (this *Transaction) ExecuteAndWatchFinalization(account subkey.KeyPair, options TransactionOptions) (TransactionDetails, error) {
	return TransactionSignSendWatch(this.client, account, this.Payload, Finalization, options, 5, 3)
}

func (this *Transaction) ExecuteAndWatchInclusion(account subkey.KeyPair, options TransactionOptions) (TransactionDetails, error) {
	return TransactionSignSendWatch(this.client, account, this.Payload, Inclusion, options, 3, 3)
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

func (this *Transaction) PaymentQueryFeeInfo(account subkey.KeyPair, options TransactionOptions) (metadata.FeeInfo, error) {
	extra, additional, err := options.ToPrimitive(this.client, account.SS58Address(42))
	if err != nil {
		return metadata.FeeInfo{}, err
	}
	tx, err := prim.CreateSigned(this.Payload.Call, extra, additional, account)
	if err != nil {
		return metadata.FeeInfo{}, err
	}

	encodedTxLen := len(tx.HexToBytes())
	encoded := tx.ToHexWith0x() + prim.Encoder.Encode(uint32(encodedTxLen))

	val, err := this.client.Rpc.State.Call("TransactionPaymentApi_query_info", encoded, prim.NewNone[prim.H256]())
	if err != nil {
		return metadata.FeeInfo{}, err
	}

	res := metadata.FeeInfo{}
	decoder := prim.NewDecoder(prim.Hex.FromHex(val), 0)
	err = decoder.Decode(&res)

	return res, newError(err, ErrorCode004)
}

func TransactionSignAndSend(client *Client, account subkey.KeyPair, payload metadata.Payload, options TransactionOptions) (prim.H256, error) {
	if !checkPayloadAndOptionsValidity(&payload, &options) {
		return prim.H256{}, errors.New("Transaction is not compatible with non-zero AppIds")
	}

	extra, additional, err := options.ToPrimitive(client, account.SS58Address(42))
	if err != nil {
		return prim.H256{}, err
	}

	return signAndSend(client, account, payload, extra, additional)
}

func TransactionSignSendWatch(client *Client, account subkey.KeyPair, payload metadata.Payload, waitFor uint8, options TransactionOptions, blockTimeout uint32, retryCount uint32) (TransactionDetails, error) {
	if !checkPayloadAndOptionsValidity(&payload, &options) {
		return TransactionDetails{}, errors.New("Transaction is not compatible with non-zero AppIds")
	}

	extra, additional, err := options.ToPrimitive(client, account.SS58Address(42))
	if err != nil {
		return TransactionDetails{}, err
	}

	for {
		txHash, err := signAndSend(client, account, payload, extra, additional)
		if err != nil {
			return TransactionDetails{}, err
		}

		iden := txHash.ToString()[0:10]
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.Debug(fmt.Sprintf("%v: Transaction was submitted. Account: %v, TxHash: %v", iden, account.SS58Address(42), txHash.ToHexWith0x()))
		}

		watcher := newWatcher(client, txHash, waitFor, blockTimeout, 3)
		maybeDetails, err := watcher.Run()
		if err != nil {
			return TransactionDetails{}, err
		}

		if maybeDetails.IsSome() {
			val := maybeDetails.Unwrap()
			if logrus.IsLevelEnabled(logrus.DebugLevel) {
				logrus.Debug(fmt.Sprintf("%v: Transaction was found. Tx Hash: %v, Tx Index: %v, Block Hash: %v, Block Number: %v", iden, val.TxHash.ToHexWith0x(), val.TxIndex, val.BlockHash.ToHexWith0x(), val.BlockNumber))
			}
			return val, nil
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

	customErr := ErrorCode003
	customErr.Message = fmt.Sprintf(`Retry count: %v`, retryCount)
	return TransactionDetails{}, &customErr
}

func signAndSend(client *Client, account subkey.KeyPair, payload metadata.Payload, extra prim.Extra, additional prim.Additional) (prim.H256, error) {
	tx, err := prim.CreateSigned(payload.Call, extra, additional, account)
	if err != nil {
		return prim.H256{}, err
	}

	return client.Send(tx)
}

// Check that ID is zero for non-submitData payloads
func checkPayloadAndOptionsValidity(payload *metadata.Payload, options *TransactionOptions) bool {
	targetPalletIndex := daPallet.CallSubmitData{}.PalletIndex()
	targetCallIndex := daPallet.CallSubmitData{}.CallIndex()

	if payload.Call.PalletIndex == targetPalletIndex && payload.Call.CallIndex == targetCallIndex {
		return true
	}

	if options.AppId.Unwrap() != 0 {
		return false
	}

	return true
}

type TransactionDetails struct {
	client      *Client
	TxHash      prim.H256
	TxIndex     uint32
	BlockHash   prim.H256
	BlockNumber uint32
	Events      prim.Option[EventRecords]
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
