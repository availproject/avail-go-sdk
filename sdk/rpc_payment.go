package sdk

import (
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type paymentRPC struct {
	client *Client
}

func (this *paymentRPC) QueryFeeDetails(extrinsic string, at prim.Option[prim.H256]) (metadata.InclusionFee, error) {
	var params = &RPCParams{}

	if !strings.HasPrefix(extrinsic, "0x") {
		extrinsic = "0x" + extrinsic
	}

	params.Add("\"" + extrinsic + "\"")
	if at.IsSome() {
		params.AddH256(at.Unwrap())
	}

	rawJson, err := this.client.Request("payment_queryFeeDetails", params.Build())
	if err != nil {
		return metadata.InclusionFee{}, err
	}

	var mappedData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJson), &mappedData); err != nil {
		return metadata.InclusionFee{}, err
	}

	if mappedData["inclusionFee"] == nil {
		return metadata.InclusionFee{}, errors.New("Failed to find inclusionFee")
	}

	jsonData := mappedData["inclusionFee"].(map[string]interface{})
	if jsonData["adjustedWeightFee"] == nil {
		return metadata.InclusionFee{}, errors.New("Failed to find adjustedWeightFee")
	}

	if jsonData["baseFee"] == nil {
		return metadata.InclusionFee{}, errors.New("Failed to find baseFee")
	}

	if jsonData["lenFee"] == nil {
		return metadata.InclusionFee{}, errors.New("Failed to find lenFee")
	}

	res := metadata.InclusionFee{}

	{
		lenFeeString := jsonData["lenFee"].(string)
		v := big.Int{}
		if _, ok := v.SetString(lenFeeString[2:], 16); !ok {
			return metadata.InclusionFee{}, errors.New("Failed to convert lenFee")
		}
		res.LenFee = metadata.NewBalanceFromBigInt(&v)
	}

	{
		baseFeeString := jsonData["baseFee"].(string)
		v := big.Int{}
		if _, ok := v.SetString(baseFeeString[2:], 16); !ok {
			return metadata.InclusionFee{}, errors.New("Failed to convert baseFee")
		}
		res.BaseFee = metadata.NewBalanceFromBigInt(&v)
	}

	{
		adjustedWeightFeeString := jsonData["adjustedWeightFee"].(string)
		v := big.Int{}
		if _, ok := v.SetString(adjustedWeightFeeString[2:], 16); !ok {
			return metadata.InclusionFee{}, errors.New("Failed to convert adjustedWeightFee")
		}
		res.AdjustedWeightFee = metadata.NewBalanceFromBigInt(&v)
	}

	return res, err
}
