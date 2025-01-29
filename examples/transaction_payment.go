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

	// PaymentQueryCallFeeDetails
	feeDetails1, err := tx.PaymentQueryCallFeeDetails()
	PanicOnError(err)

	AssertEq(feeDetails1.InclusionFee.IsSome(), true, "InclusionFee Must Exist")
	if feeDetails1.InclusionFee.IsSome() {
		InclusionFee := feeDetails1.InclusionFee.Unwrap()
		println("Adjusted Weight Fee:", InclusionFee.AdjustedWeightFee.ToHuman())
		println("Len Fee:", InclusionFee.LenFee.ToHuman())
		println("Base Fee:", InclusionFee.BaseFee.ToHuman())
	}

	// PaymentQueryFeeDetails
	feeDetails2, err := tx.PaymentQueryFeeDetails(acc, options)
	PanicOnError(err)

	AssertEq(feeDetails1.InclusionFee.IsSome(), true, "InclusionFee Must Exist")
	if feeDetails2.InclusionFee.IsSome() {
		InclusionFee := feeDetails2.InclusionFee.Unwrap()
		println("Adjusted Weight Fee:", InclusionFee.AdjustedWeightFee.ToHuman())
		println("Len Fee:", InclusionFee.LenFee.ToHuman())
		println("Base Fee:", InclusionFee.BaseFee.ToHuman())
	}

	// PaymentQueryCallFeeInfo
	feeInfo1, err := tx.PaymentQueryCallFeeInfo()
	PanicOnError(err)

	println("ProofSize:", feeInfo1.Weight.ProofSize)
	println("RefTime:", feeInfo1.Weight.RefTime)
	println("Class:", feeInfo1.Class.ToHuman())
	println("Partial Fee:", feeInfo1.PartialFee.ToHuman())

	// PaymentQueryFeeInfo
	feeInfo, err := tx.PaymentQueryFeeInfo(acc, options)
	PanicOnError(err)

	println("ProofSize:", feeInfo.Weight.ProofSize)
	println("RefTime:", feeInfo.Weight.RefTime)
	println("Class:", feeInfo.Class.ToHuman())
	println("Partial Fee:", feeInfo.PartialFee.ToHuman())

	println("RunTransactionPayment finished correctly.")

}
