package primitives

import (
	"encoding/json"
	"strconv"
)

type Header struct {
	// Missing
	// digest
	// extension
	ExtrinsicsRoot H256
	Number         uint32
	ParentHash     H256
	StateRoot      H256
}

func NewHeaderFromJson(rawJson string) Header {
	var mappedData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		panic(err)
	}

	if mappedData["extrinsicsRoot"] == nil {
		panic("Header is missing extrinsicsRoot")
	}
	if mappedData["number"] == nil {
		panic("Header is missing number")
	}
	if mappedData["parentHash"] == nil {
		panic("Header is missing parentHash")
	}
	if mappedData["stateRoot"] == nil {
		panic("Header is missing stateRoot")
	}

	//
	extrinsicsRoot := NewH256FromHexString(mappedData["extrinsicsRoot"].(string))
	parentHash := NewH256FromHexString(mappedData["parentHash"].(string))
	stateRoot := NewH256FromHexString(mappedData["stateRoot"].(string))
	number, err := strconv.ParseUint(mappedData["number"].(string)[2:], 16, 32)
	if err != nil {
		panic(err)
	}

	return Header{
		ExtrinsicsRoot: extrinsicsRoot,
		Number:         uint32(number),
		ParentHash:     parentHash,
		StateRoot:      stateRoot,
	}
}
