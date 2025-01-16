package complex

import (
	prim "go-sdk/primitives"
)

func (this *stateRPC) GetRuntimeVersion(blockHash prim.Option[prim.H256]) prim.RuntimeVersion {
	var params = &RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	var value, err = this.client.Request("state_getRuntimeVersion", params.Build())
	if err != nil {
		panic(err)
	}

	return prim.NewRuntimeVersionFromJson(value)
}

func (this *stateRPC) GetStorage(key string, at prim.Option[prim.H256]) string {
	params := RPCParams{}
	params.Add("\"" + key + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	value, err := this.client.Request("state_getStorage", params.Build())
	if err != nil {
		panic(err)
	}

	return value
}

func (this *stateRPC) GetMetadata(at prim.Option[prim.H256]) string {
	params := RPCParams{}
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	value, err := this.client.Request("state_getMetadata", params.Build())
	if err != nil {
		panic(err)
	}

	return value
}

func (this *stateRPC) GetEvents(at prim.Option[prim.H256]) string {
	params := RPCParams{}
	params.Add("\"" + "0x26aa394eea5630e07c48ae0c9558cef780d41e5e16056765bc8461851072c9d7" + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	value, err := this.client.Request("state_getStorage", params.Build())
	if err != nil {
		return ""
	}

	return value
}
