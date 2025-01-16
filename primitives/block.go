package primitives

import "encoding/json"

type Block struct {
	Extrinsics []string
	Header     Header
}

func NewBlock(rawJson string) Block {
	var mappedData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		panic(err)
	}

	if mappedData["block"] == nil {
		panic("Block is missing block")
	}

	mappedData2 := mappedData["block"].(map[string]interface{})
	if mappedData2["extrinsics"] == nil {
		panic("Block is missing extrinsics")
	}

	if mappedData2["header"] == nil {
		panic("Block is missing header")
	}

	headerJson, err2 := json.Marshal(mappedData2["header"])
	if err2 != nil {
		panic(err2)
	}
	header := NewHeaderFromJson(string(headerJson))

	extrinsicsRaw := mappedData2["extrinsics"].([]interface{})
	extrinsics := []string{}

	for i := 0; i < len(extrinsicsRaw); i++ {
		extrinsics = append(extrinsics, extrinsicsRaw[i].(string))
	}

	return Block{
		Extrinsics: extrinsics,
		Header:     header,
	}
}
