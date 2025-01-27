package examples

import (
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionPayment() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	if err != nil {
		panic(err)
	}

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))
	feeDetails, err := tx.PaymentQueryFeeDetails(SDK.Account.Alice(), SDK.NewTransactionOptions().WithAppId(1))
	if err != nil {
		panic(err)
	}
	println("Adjusted Weight Fee:", feeDetails.AdjustedWeightFee.ToHuman())
	println("Len Fee:", feeDetails.LenFee.ToHuman())
	println("Base Fee:", feeDetails.BaseFee.ToHuman())

	println("RunTransactionPayment finished correctly.")
}
