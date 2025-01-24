package sdk

import (
	"encoding/json"
	"errors"
	"strconv"
)

type systemRPC struct {
	client *Client
}

func (this *systemRPC) AccountNextIndex(address string) (uint32, error) {
	if len(address) < 1 {
		return uint32(0), errors.New("Address needs to have a length of > 0")
	}

	if address[0] != '"' {
		address = "\"" + address
	}

	if address[len(address)-1] != '"' {
		address += "\""
	}

	params := RPCParams{}
	params.Add(address)

	value, err := this.client.Request("system_accountNextIndex", params.Build())
	if err != nil {
		return uint32(0), err
	}
	parsedValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return uint32(0), err
	}

	return uint32(parsedValue), nil
}

func (this *systemRPC) Chain() (string, error) {
	params := RPCParams{}
	return this.client.Request("system_chain", params.Build())
}

func (this *systemRPC) ChainType() (string, error) {
	params := RPCParams{}
	return this.client.Request("system_chainType", params.Build())
}

func (this *systemRPC) Health() (RpcSystemHealth, error) {
	params := RPCParams{}
	val, err := this.client.Request("system_health", params.Build())
	if err != nil {
		return RpcSystemHealth{}, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(val), &jsonData); err != nil {
		return RpcSystemHealth{}, err
	}

	if jsonData["peers"] == nil {
		return RpcSystemHealth{}, errors.New("Failed to find Peers")
	}
	if jsonData["isSyncing"] == nil {
		return RpcSystemHealth{}, errors.New("Failed to find isSyncing")
	}
	if jsonData["shouldHavePeers"] == nil {
		return RpcSystemHealth{}, errors.New("Failed to find shouldHavePeers")
	}

	res := RpcSystemHealth{}
	res.Peers = uint32(jsonData["peers"].(float64))
	res.IsSyncing = jsonData["isSyncing"].(bool)
	res.ShouldHavePeers = jsonData["shouldHavePeers"].(bool)

	return res, nil
}

type RpcSystemHealth struct {
	Peers           uint32
	IsSyncing       bool
	ShouldHavePeers bool
}

func (this *systemRPC) LocalPeerId() (string, error) {
	params := RPCParams{}
	return this.client.Request("system_localPeerId", params.Build())
}

func (this *systemRPC) Name() (string, error) {
	params := RPCParams{}
	return this.client.Request("system_name", params.Build())
}

func (this *systemRPC) NodeRoles() ([]string, error) {
	params := RPCParams{}
	val, err := this.client.Request("system_nodeRoles", params.Build())
	if err != nil {
		return []string{}, err
	}

	var jsonData []interface{}
	if err := json.Unmarshal([]byte(val), &jsonData); err != nil {
		return []string{}, err
	}

	res := []string{}
	for _, el := range jsonData {
		res = append(res, el.(string))
	}

	return res, nil
}

func (this *systemRPC) Properties() (RpcSystemChainProperties, error) {
	params := RPCParams{}
	val, err := this.client.Request("system_properties", params.Build())
	if err != nil {
		return RpcSystemChainProperties{}, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(val), &jsonData); err != nil {
		return RpcSystemChainProperties{}, err
	}

	if jsonData["ss58Format"] == nil {
		return RpcSystemChainProperties{}, errors.New("Failed to find ss58Format")
	}
	if jsonData["tokenDecimals"] == nil {
		return RpcSystemChainProperties{}, errors.New("Failed to find tokenDecimals")
	}
	if jsonData["tokenSymbol"] == nil {
		return RpcSystemChainProperties{}, errors.New("Failed to find tokenSymbol")
	}

	res := RpcSystemChainProperties{}
	if jsonData["isEthereum"] == nil {
		res.IsEthereum = false
	} else {
		res.IsEthereum = jsonData["isEthereum"].(bool)
	}

	res.Ss58Format = uint32(jsonData["ss58Format"].(float64))
	res.TokenDecimals = uint32(jsonData["tokenDecimals"].(float64))
	res.TokenSymbol = jsonData["tokenSymbol"].(string)

	return res, nil
}

type RpcSystemChainProperties struct {
	IsEthereum    bool
	Ss58Format    uint32
	TokenDecimals uint32
	TokenSymbol   string
}

func (this *systemRPC) SyncState() (RpcSystemSyncState, error) {
	params := RPCParams{}
	val, err := this.client.Request("system_syncState", params.Build())
	if err != nil {
		return RpcSystemSyncState{}, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(val), &jsonData); err != nil {
		return RpcSystemSyncState{}, err
	}

	if jsonData["startingBlock"] == nil {
		return RpcSystemSyncState{}, errors.New("Failed to find startingBlock")
	}
	if jsonData["currentBlock"] == nil {
		return RpcSystemSyncState{}, errors.New("Failed to find currentBlock")
	}
	if jsonData["highestBlock"] == nil {
		return RpcSystemSyncState{}, errors.New("Failed to find highestBlock")
	}

	res := RpcSystemSyncState{}
	res.StartingBlock = uint32(jsonData["startingBlock"].(float64))
	res.CurrentBlock = uint32(jsonData["currentBlock"].(float64))
	res.HighestBlock = uint32(jsonData["highestBlock"].(float64))

	return res, nil
}

type RpcSystemSyncState struct {
	StartingBlock uint32
	CurrentBlock  uint32
	HighestBlock  uint32
}

func (this *systemRPC) Version() (string, error) {
	params := RPCParams{}
	return this.client.Request("system_version", params.Build())
}
