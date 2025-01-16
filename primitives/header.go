package primitives

import (
	"encoding/json"
	"errors"
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

func NewHeaderFromJson(rawJson string) (Header, error) {
	var mappedData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		return Header{}, err
	}

	if mappedData["extrinsicsRoot"] == nil {
		return Header{}, errors.New("Header is missing extrinsicsRoot")
	}
	if mappedData["number"] == nil {
		return Header{}, errors.New("Header is missing number")
	}
	if mappedData["parentHash"] == nil {
		return Header{}, errors.New("Header is missing parentHash")
	}
	if mappedData["stateRoot"] == nil {
		return Header{}, errors.New("Header is missing stateRoot")
	}

	//
	extrinsicsRoot, err := NewH256FromHexString(mappedData["extrinsicsRoot"].(string))
	if err != nil {
		return Header{}, err
	}
	parentHash, err := NewH256FromHexString(mappedData["parentHash"].(string))
	if err != nil {
		return Header{}, err
	}
	stateRoot, err := NewH256FromHexString(mappedData["stateRoot"].(string))
	if err != nil {
		return Header{}, err
	}
	number, err := strconv.ParseUint(mappedData["number"].(string)[2:], 16, 32)
	if err != nil {
		return Header{}, err
	}

	return Header{
		ExtrinsicsRoot: extrinsicsRoot,
		Number:         uint32(number),
		ParentHash:     parentHash,
		StateRoot:      stateRoot,
	}, nil
}
