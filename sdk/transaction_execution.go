package sdk

import (
	"fmt"
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

func TransactionSignSendWatch(client *Client, account subkey.KeyPair, payload metadata.Payload, waitFor uint8, options TransactionOptions) (TransactionDetails, error) {
	if !checkPayloadAndOptionsValidity(&payload, &options) {
		return TransactionDetails{}, errors.New("Transaction is not compatible with non-zero AppIds")
	}

	retryCount := 2

	extra, additional, err := options.ToPrimitive(client, account.SS58Address(42))
	if err != nil {
		return TransactionDetails{}, err
	}

	for {
		txHash, err := signAndSend(client, account, payload, extra, additional)
		if err != nil {
			return TransactionDetails{}, err
		}

		logger := NewTxLogger(txHash, true)
		logger.LogTxSubmitted(&account, extra.Era.Period)

		watcher := NewWatcher(client, txHash).WaitFor(waitFor).Logger(logger)
		maybeDetails, err := watcher.Run()
		if err != nil {
			return TransactionDetails{}, err
		}

		if maybeDetails.IsSome() {
			val := maybeDetails.Unwrap()
			return val, nil
		}

		if retryCount == 0 {
			logger.LogTxRetryAbort()
			break
		}

		RegenerateEra(client, &extra, &additional)

		retryCount -= 1
		logger.LogTxRetry()
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

	logger := NewTxLogger(prim.H256{}, true)
	logger.LogTxSubmitting(&account, &payload, extra.Nonce, extra.AppId)

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
