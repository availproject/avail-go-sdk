package primitives

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Block struct {
	Extrinsics []string
	Header     Header
}

func NewBlock(rawJson string) (Block, error) {
	var mappedData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		fmt.Println(fmt.Sprintf("Value: %v, Error: %v", rawJson, err))
		return Block{}, err
	}

	if mappedData["block"] == nil {
		fmt.Println(fmt.Sprintf("Value: %v", rawJson))
		return Block{}, errors.New("Block is missing block")
	}

	mappedData2 := mappedData["block"].(map[string]interface{})
	if mappedData2["extrinsics"] == nil {
		return Block{}, errors.New("Block is missing extrinsics")
	}

	if mappedData2["header"] == nil {
		return Block{}, errors.New("Block is missing header")
	}

	headerJson, err := json.Marshal(mappedData2["header"])
	if err != nil {
		fmt.Println(fmt.Sprintf("Value: %v, Error: %v", rawJson, err))
		return Block{}, err
	}
	header, err := NewHeaderFromJson(string(headerJson))
	if err != nil {
		fmt.Println(fmt.Sprintf("Value: %v, Error: %v", rawJson, err))
		return Block{}, err
	}

	extrinsicsRaw := mappedData2["extrinsics"].([]interface{})
	extrinsics := []string{}

	for i := 0; i < len(extrinsicsRaw); i++ {
		extrinsics = append(extrinsics, extrinsicsRaw[i].(string))
	}

	return Block{
		Extrinsics: extrinsics,
		Header:     header,
	}, nil
}
