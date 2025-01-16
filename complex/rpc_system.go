package complex

import (
	"strconv"
)

func (this *systemRPC) AccountNextIndex(accountId string) uint32 {
	if len(accountId) < 1 {
		panic("AccountId needs to have a length of > 0")
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
		panic(err)
	}
	parsedValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		panic(err)
	}

	return uint32(parsedValue)
}
