package sdk

import (
	"encoding/json"
	"fmt"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type stateRPC struct {
	client *Client
}

func (s *stateRPC) GetRuntimeVersion(blockHash prim.Option[prim.H256]) (prim.RuntimeVersion, error) {
	var params = &RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	value, err := s.client.RequestWithRetry("state_getRuntimeVersion", params.Build())
	if err != nil {
		return prim.RuntimeVersion{}, err
	}

	return prim.NewRuntimeVersionFromJson(value)
}

func (s *stateRPC) GetStorage(key string, at prim.Option[prim.H256]) (prim.Option[string], error) {
	params := RPCParams{}
	params.Add("\"" + key + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	return s.client.Request("state_getStorage", params.Build())
}

func (s *stateRPC) GetKeys(key string, at prim.Option[prim.H256]) ([]string, error) {
	params := RPCParams{}
	params.Add("\"" + key + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	value, err := s.client.RequestWithRetry("state_getKeys", params.Build())
	if err != nil {
		return nil, err
	}

	res := []string{}
	if err := json.Unmarshal([]byte(value), &res); err != nil {
		return nil, newError(err, ErrorCode002)
	}

	return res, nil
}

func (s *stateRPC) GetKeysPaged(key string, count uint32, startKey prim.Option[string], at prim.Option[prim.H256]) ([]string, error) {
	params := RPCParams{}
	params.Add("\"" + key + "\"")
	params.Add(fmt.Sprintf(`%v`, count))
	if startKey.IsSome() {
		params.Add("\"" + startKey.Unwrap() + "\"")
	} else {
		params.Add("null")
	}
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	value, err := s.client.RequestWithRetry("state_getKeysPaged", params.Build())
	if err != nil {
		return nil, err
	}

	res := []string{}
	if err := json.Unmarshal([]byte(value), &res); err != nil {
		return nil, newError(err, ErrorCode002)
	}

	return res, nil
}

func (s *stateRPC) GetMetadata(at prim.Option[prim.H256]) (string, error) {
	params := RPCParams{}
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	return s.client.RequestWithRetry("state_getMetadata", params.Build())
}

func (s *stateRPC) GetEvents(at prim.Option[prim.H256]) (string, error) {
	params := RPCParams{}
	params.Add("\"" + "0x26aa394eea5630e07c48ae0c9558cef780d41e5e16056765bc8461851072c9d7" + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	return s.client.RequestWithRetry("state_getStorage", params.Build())
}

func (s *stateRPC) Call(method string, data string, at prim.Option[prim.H256]) (string, error) {
	params := RPCParams{}
	params.Add("\"" + method + "\"")
	params.Add("\"" + data + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	return s.client.RequestWithRetry("state_call", params.Build())
}
