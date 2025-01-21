package sdk

import (
	"encoding/json"
	prim "go-sdk/primitives"
)

func (this *stateRPC) GetRuntimeVersion(blockHash prim.Option[prim.H256]) (prim.RuntimeVersion, error) {
	var params = &RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	value, err := this.client.Request("state_getRuntimeVersion", params.Build())
	if err != nil {
		return prim.RuntimeVersion{}, err
	}

	return prim.NewRuntimeVersionFromJson(value)
}

func (this *stateRPC) GetStorage(key string, at prim.Option[prim.H256]) (string, error) {
	params := RPCParams{}
	params.Add("\"" + key + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	return this.client.Request("state_getStorage", params.Build())
}

func (this *stateRPC) GetKeys(key string, at prim.Option[prim.H256]) ([]string, error) {
	params := RPCParams{}
	params.Add("\"" + key + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	value, err := this.client.Request("state_getKeys", params.Build())
	if err != nil {
		return nil, err
	}

	res := []string{}
	if err := json.Unmarshal([]byte(value), &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (this *stateRPC) GetMetadata(at prim.Option[prim.H256]) (string, error) {
	params := RPCParams{}
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	return this.client.Request("state_getMetadata", params.Build())
}

func (this *stateRPC) GetEvents(at prim.Option[prim.H256]) (string, error) {
	params := RPCParams{}
	params.Add("\"" + "0x26aa394eea5630e07c48ae0c9558cef780d41e5e16056765bc8461851072c9d7" + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	return this.client.Request("state_getStorage", params.Build())
}
