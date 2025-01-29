package examples

import (
	"fmt"

	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunTransactionPayment() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
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
		fmt.Println("Adjusted Weight Fee:", InclusionFee.AdjustedWeightFee.ToHuman())
		fmt.Println("Len Fee:", InclusionFee.LenFee.ToHuman())
		fmt.Println("Base Fee:", InclusionFee.BaseFee.ToHuman())
	}

	// PaymentQueryFeeDetails
	feeDetails2, err := tx.PaymentQueryFeeDetails(acc, options)
	PanicOnError(err)

	AssertEq(feeDetails1.InclusionFee.IsSome(), true, "InclusionFee Must Exist")
	if feeDetails2.InclusionFee.IsSome() {
		InclusionFee := feeDetails2.InclusionFee.Unwrap()
		fmt.Println("Adjusted Weight Fee:", InclusionFee.AdjustedWeightFee.ToHuman())
		fmt.Println("Len Fee:", InclusionFee.LenFee.ToHuman())
		fmt.Println("Base Fee:", InclusionFee.BaseFee.ToHuman())
	}

	// PaymentQueryCallFeeInfo
	feeInfo1, err := tx.PaymentQueryCallFeeInfo()
	PanicOnError(err)

	fmt.Println("ProofSize:", feeInfo1.Weight.ProofSize)
	fmt.Println("RefTime:", feeInfo1.Weight.RefTime)
	fmt.Println("Class:", feeInfo1.Class.ToHuman())
	fmt.Println("Partial Fee:", feeInfo1.PartialFee.ToHuman())

	// PaymentQueryFeeInfo
	feeInfo, err := tx.PaymentQueryFeeInfo(acc, options)
	PanicOnError(err)

	fmt.Println("ProofSize:", feeInfo.Weight.ProofSize)
	fmt.Println("RefTime:", feeInfo.Weight.RefTime)
	fmt.Println("Class:", feeInfo.Class.ToHuman())
	fmt.Println("Partial Fee:", feeInfo.PartialFee.ToHuman())

	fmt.Println("RunTransactionPayment finished correctly.")

}
