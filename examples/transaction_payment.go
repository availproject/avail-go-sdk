package examples

import (
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionPayment() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	acc := SDK.Account.Alice()
	options := SDK.NewTransactionOptions().WithAppId(1)
	tx := sdk.Tx.DataAvailability.SubmitData([]byte("Hello World"))

	// PaymentQueryFeeDetails
	feeDetails, err := tx.PaymentQueryFeeDetails(acc, options)
	PanicOnError(err)

	println("Adjusted Weight Fee:", feeDetails.AdjustedWeightFee.ToHuman())
	println("Len Fee:", feeDetails.LenFee.ToHuman())
	println("Base Fee:", feeDetails.BaseFee.ToHuman())

	// PaymentQueryFeeInfo
	feeInfo, err := tx.PaymentQueryFeeInfo(acc, options)
	PanicOnError(err)

	println("ProofSize:", feeInfo.Weight.ProofSize)
	println("RefTime:", feeInfo.Weight.RefTime)
	println("Class:", feeInfo.Class.ToHuman())
	println("Partial Fee:", feeInfo.PartialFee.ToHuman())

	println("RunTransactionPayment finished correctly.")
}
