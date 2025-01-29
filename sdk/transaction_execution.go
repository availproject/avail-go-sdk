package sdk

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vedhavyas/go-subkey/v2"

	"errors"

	"github.com/availproject/avail-go-sdk/metadata"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

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
