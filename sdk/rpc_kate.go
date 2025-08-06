package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type kateRPC struct {
	client *Client
}

func (k *kateRPC) BlockLength(blockHash prim.Option[prim.H256]) (metadata.BlockLength, error) {
	var params = &RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	rawJson, err := k.client.RequestWithRetry("kate_blockLength", params.Build())
	if err != nil {
		return metadata.BlockLength{}, err
	}

	var mappedData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		return metadata.BlockLength{}, newError(err, ErrorCode002)
	}

	if mappedData["chunkSize"] == nil {
		return metadata.BlockLength{}, errors.New("Block Length is missing chunkSize")
	}
	if mappedData["cols"] == nil {
		return metadata.BlockLength{}, errors.New("Block Length is missing cols")
	}
	if mappedData["max"] == nil {
		return metadata.BlockLength{}, errors.New("Block Length is missing max")
	}
	if mappedData["rows"] == nil {
		return metadata.BlockLength{}, errors.New("Block Length is missing rows")
	}

	res := metadata.BlockLength{}

	res.ChunkSize = uint32(mappedData["chunkSize"].(float64))
	res.Cols = uint32(mappedData["cols"].(float64))
	res.Rows = uint32(mappedData["rows"].(float64))

	arrT := mappedData["max"].([]interface{})

	res.Max.Normal = uint32(arrT[0].(float64))
	res.Max.Operational = uint32(arrT[1].(float64))
	res.Max.Mandatory = uint32(arrT[2].(float64))

	return res, nil
}

func (k *kateRPC) QueryDataProof(transactionIndex uint32, blockHash prim.Option[prim.H256]) (metadata.ProofResponse, error) {
	res := metadata.ProofResponse{}
	var params = &RPCParams{}
	params.AddUint32(transactionIndex)
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	rawJson, err := k.client.RequestWithRetry("kate_queryDataProof", params.Build())
	if err != nil {
		return res, err
	}

	fmt.Println(rawJson)

	var mappedData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		return res, newError(err, ErrorCode002)
	}

	if mappedData["dataProof"] == nil {
		return res, errors.New("QueryDataProof is missing dataProof")
	}

	dataProofMap := mappedData["dataProof"].(map[string]interface{})

	if dataProofMap["leaf"] == nil {
		return res, errors.New("QueryDataProof is missing leaf")
	}
	if dataProofMap["leafIndex"] == nil {
		return res, errors.New("QueryDataProof is missing leafIndex")
	}
	if dataProofMap["numberOfLeaves"] == nil {
		return res, errors.New("QueryDataProof is missing numberOfLeaves")
	}
	if dataProofMap["proof"] == nil {
		return res, errors.New("QueryDataProof is missing proof")
	}
	if dataProofMap["roots"] == nil {
		return res, errors.New("QueryDataProof is missing roots")
	}

	res.DataProof.Leaf, err = prim.NewH256FromHexString(dataProofMap["leaf"].(string))
	if err != nil {
		return metadata.ProofResponse{}, err
	}
	res.DataProof.LeafIndex = uint32(dataProofMap["leafIndex"].(float64))
	res.DataProof.NumberOfLeaves = uint32(dataProofMap["numberOfLeaves"].(float64))

	rootsMap := dataProofMap["roots"].(map[string]interface{})
	if rootsMap["blobRoot"] == nil {
		return res, errors.New("QueryDataProof is missing blobRoot")
	}
	if rootsMap["bridgeRoot"] == nil {
		return res, errors.New("QueryDataProof is missing bridgeRoot")
	}
	if rootsMap["dataRoot"] == nil {
		return res, errors.New("QueryDataProof is missing dataRoot")
	}

	res.DataProof.Roots.BlobRoot, err = prim.NewH256FromHexString(rootsMap["blobRoot"].(string))
	if err != nil {
		return metadata.ProofResponse{}, err
	}
	res.DataProof.Roots.BridgeRoot, err = prim.NewH256FromHexString(rootsMap["bridgeRoot"].(string))
	if err != nil {
		return metadata.ProofResponse{}, err
	}
	res.DataProof.Roots.DataRoot, err = prim.NewH256FromHexString(rootsMap["dataRoot"].(string))
	if err != nil {
		return metadata.ProofResponse{}, err
	}

	proofMap := dataProofMap["proof"].([]interface{})
	for i := range proofMap {
		val, err := prim.NewH256FromHexString(proofMap[i].(string))
		if err != nil {
			return metadata.ProofResponse{}, err
		}
		res.DataProof.Proof = append(res.DataProof.Proof, val)
	}

	if mappedData["message"] == nil {
		res.Message.Unset()
		return res, nil
	}

	addressedMsgMap := mappedData["message"].(map[string]interface{})
	if addressedMsgMap["destinationDomain"] == nil {
		return res, errors.New("QueryDataProof is missing destinationDomain")
	}
	if addressedMsgMap["originDomain"] == nil {
		return res, errors.New("QueryDataProof is missing originDomain")
	}
	if addressedMsgMap["from"] == nil {
		return res, errors.New("QueryDataProof is missing from")
	}
	if addressedMsgMap["to"] == nil {
		return res, errors.New("QueryDataProof is missing to")
	}
	if addressedMsgMap["id"] == nil {
		return res, errors.New("QueryDataProof is missing id")
	}
	if addressedMsgMap["message"] == nil {
		return res, errors.New("QueryDataProof is missing message")
	}

	msg := metadata.AddressedMessage{}
	msg.Id = uint64(addressedMsgMap["id"].(float64))
	msg.DestinationDomain = uint32(addressedMsgMap["destinationDomain"].(float64))
	msg.OriginDomain = uint32(addressedMsgMap["originDomain"].(float64))

	msg.From, err = prim.NewH256FromHexString(addressedMsgMap["from"].(string))
	if err != nil {
		return metadata.ProofResponse{}, err
	}
	msg.To, err = prim.NewH256FromHexString(addressedMsgMap["to"].(string))
	if err != nil {
		return metadata.ProofResponse{}, err
	}

	msg2Map := addressedMsgMap["message"].(map[string]interface{})

	if msg2Map["fungibleToken"] != nil {
		msg.Message.VariantIndex = 0
		fungMap := msg2Map["fungibleToken"].(map[string]interface{})
		if fungMap["asset_id"] == nil {
			return res, errors.New("QueryDataProof is missing AssetId")
		}
		if fungMap["amount"] == nil {
			return res, errors.New("QueryDataProof is missing Amount")
		}

		t := metadata.MessageFungibleToken{}

		t.AssetId, err = prim.NewH256FromHexString(fungMap["asset_id"].(string))
		if err != nil {
			return metadata.ProofResponse{}, err
		}

		amountF := fungMap["amount"].(float64)
		amount, err := metadata.NewBalanceFromString(strconv.FormatFloat(amountF, 'f', -1, 64))
		if err != nil {
			return res, err
		}
		t.Amount = amount

		msg.Message.FungibleToken.Set(t)

	} else if msg2Map["arbitraryMessage"] != nil {
		msg.Message.VariantIndex = 1
		// TODO
		panic("TODO")

	} else {
		panic("Something went wrong with mapping message")
	}

	res.Message.Set(msg)

	return res, nil
}

func (k *kateRPC) QueryProof(cells []KateCell, blockHash prim.Option[prim.H256]) ([]GDataProof, error) {
	var params = &RPCParams{}
	res := []GDataProof{}

	cellsEnc := "["
	for i := range cells {
		cellsEnc += "[" + fmt.Sprintf("%v", cells[i].Row) + "," + fmt.Sprintf("%v", cells[i].Col) + "]"

		if i < (len(cells) - 1) {
			cellsEnc += ","
		}
	}
	cellsEnc += "]"
	params.Add(cellsEnc)
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	rawJson, err := k.client.RequestWithRetry("kate_queryProof", params.Build())
	if err != nil {
		return res, err
	}

	var mappedData []interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		return res, newError(err, ErrorCode002)
	}

	for i := range mappedData {
		mappedData2 := mappedData[i].([]interface{})
		gProofArray := mappedData2[1].([]interface{})

		gProof := [48]byte{}
		if len(gProofArray) != 48 {
			return res, errors.New("GProof is not 48 bytes long")
		}
		for i := range gProofArray {
			gProof[i] = byte(gProofArray[i].(float64))
		}

		res = append(res, metadata.NewTuple2(mappedData2[0].(string), gProof))

	}

	return res, nil
}

type KateCell struct {
	Row int32 `scale:"compact"`
	Col int32 `scale:"compact"`
}

type GDataProof = metadata.Tuple2[string, [48]byte]

func (k *kateRPC) QueryRows(rows []uint32, blockHash prim.Option[prim.H256]) ([][]string, error) {
	var params = &RPCParams{}
	res := [][]string{}

	rowsEnc := "["
	for i := range rows {
		rowsEnc += fmt.Sprintf("%v", rows[i])

		if i < (len(rows) - 1) {
			rowsEnc += ","
		}
	}
	rowsEnc += "]"
	params.Add(rowsEnc)
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	rawJson, err := k.client.RequestWithRetry("kate_queryRows", params.Build())
	if err != nil {
		return res, err
	}

	var outerArrays []interface{}
	if err := json.Unmarshal([]byte(rawJson), &outerArrays); err != nil {
		return res, newError(err, ErrorCode002)
	}

	for i := range outerArrays {
		res = append(res, []string{})
		innerArray := outerArrays[i].([]interface{})
		for j := range innerArray {
			res[i] = append(res[i], innerArray[j].(string))
		}
	}

	return res, nil
}
