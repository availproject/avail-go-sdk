package sdk

import (
	"errors"
	"strconv"
)

type systemRPC struct {
	client *Client
}

func (this *systemRPC) AccountNextIndex(accountId string) (uint32, error) {
	if len(accountId) < 1 {
		return uint32(0), errors.New("AccountId needs to have a length of > 0")
	}

	if accountId[0] != '"' {
		accountId = "\"" + accountId
	}

	if accountId[len(accountId)-1] != '"' {
		accountId += "\""
	}

	params := RPCParams{}
	params.Add(accountId)

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
