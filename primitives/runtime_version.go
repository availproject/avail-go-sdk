package primitives

import (
	"encoding/json"
)

type RuntimeVersion struct {
	// Missing
	// apis
	SpecVersion      uint32
	TxVersion        uint32
	ImplVersion      uint32
	AuthoringVersion uint32
	StateVersion     uint32
	SpecName         string
	ImplName         string
}

func NewRuntimeVersionFromJson(rawJson string) RuntimeVersion {
	var mappedData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		panic(err)
	}

	if mappedData["specVersion"] == nil {
		panic("Header is missing specVersion")
	}
	if mappedData["transactionVersion"] == nil {
		panic("Header is missing transactionVersion")
	}
	if mappedData["implVersion"] == nil {
		panic("Header is missing implVersion")
	}
	if mappedData["authoringVersion"] == nil {
		panic("Header is missing authoringVersion")
	}
	if mappedData["stateVersion"] == nil {
		panic("Header is missing stateVersion")
	}
	if mappedData["specName"] == nil {
		panic("Header is missing specName")
	}
	if mappedData["implName"] == nil {
		panic("Header is missing implName")
	}

	specVersion := uint32(mappedData["specVersion"].(float64))
	txVersion := uint32(mappedData["transactionVersion"].(float64))
	implVersion := uint32(mappedData["implVersion"].(float64))
	authoringVersion := uint32(mappedData["authoringVersion"].(float64))
	stateVersion := uint32(mappedData["stateVersion"].(float64))
	specName := mappedData["specName"].(string)
	implName := mappedData["implName"].(string)

	return RuntimeVersion{
		SpecVersion:      specVersion,
		TxVersion:        txVersion,
		ImplVersion:      implVersion,
		AuthoringVersion: authoringVersion,
		StateVersion:     stateVersion,
		SpecName:         specName,
		ImplName:         implName,
	}
}
