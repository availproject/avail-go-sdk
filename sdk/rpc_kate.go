package sdk

import (
	"encoding/json"
	"errors"

	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type kateRPC struct {
	client *Client
}

func (this *kateRPC) BlockLength(blockHash prim.Option[prim.H256]) (metadata.BlockLength, error) {
	var params = &RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	rawJson, err := this.client.Request("kate_blockLength", params.Build())
	if err != nil {
		return metadata.BlockLength{}, err
	}

	var mappedData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		return metadata.BlockLength{}, err
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
