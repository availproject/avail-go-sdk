package sdk

import (
	"encoding/json"
	"errors"

	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/primitives"
)

type transactionRPC struct {
	client *Client
}

func (this *transactionRPC) State(txHash primitives.H256, finalized bool) ([]metadata.TransactionState, error) {
	params := RPCParams{}
	params.Add(txHash.ToRpcParam())
	params.AddBool(finalized)

	value, err := this.client.RequestWithRetry("transaction_state", params.Build())
	if err != nil {
		return []metadata.TransactionState{}, err
	}

	var jsonData []interface{}
	if err := json.Unmarshal([]byte(value), &jsonData); err != nil {
		return []metadata.TransactionState{}, newError(err, ErrorCode002)
	}

	result := []metadata.TransactionState{}
	for _, elem := range jsonData {
		elemJson := elem.(map[string]interface{})

		if elemJson["block_hash"] == nil {
			return []metadata.TransactionState{}, errors.New("Failed to find block_hash")
		}

		if elemJson["block_height"] == nil {
			return []metadata.TransactionState{}, errors.New("Failed to find block_height")
		}

		if elemJson["pallet_index"] == nil {
			return []metadata.TransactionState{}, errors.New("Failed to find pallet_index")
		}

		if elemJson["call_index"] == nil {
			return []metadata.TransactionState{}, errors.New("Failed to find call_index")
		}

		if elemJson["is_finalized"] == nil {
			return []metadata.TransactionState{}, errors.New("Failed to find is_finalized")
		}

		if elemJson["tx_success"] == nil {
			return []metadata.TransactionState{}, errors.New("Failed to find tx_success")
		}

		if elemJson["tx_hash"] == nil {
			return []metadata.TransactionState{}, errors.New("Failed to find tx_hash")
		}

		if elemJson["tx_index"] == nil {
			return []metadata.TransactionState{}, errors.New("Failed to find tx_index")
		}

		value := metadata.TransactionState{}
		blockHash, err := primitives.NewH256FromByteSlice(primitives.Hex.FromHex(elemJson["block_hash"].(string)))
		if err != nil {
			return []metadata.TransactionState{}, err
		}
		value.BlockHash = blockHash

		txHash, err := primitives.NewH256FromByteSlice(primitives.Hex.FromHex(elemJson["tx_hash"].(string)))
		if err != nil {
			return []metadata.TransactionState{}, err
		}
		value.TxHash = txHash
		value.BlockHeight = uint32(elemJson["block_height"].(float64))
		value.TxIndex = uint32(elemJson["tx_index"].(float64))
		value.PalletIndex = uint8(elemJson["pallet_index"].(float64))
		value.CallIndex = uint8(elemJson["call_index"].(float64))
		value.IsFinalized = elemJson["is_finalized"].(bool)
		value.TxSuccess = elemJson["tx_success"].(bool)

		result = append(result, value)
	}

	return result, nil
}
