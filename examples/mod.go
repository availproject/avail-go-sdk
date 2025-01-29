package examples

import "fmt"

func Run() {
	RunAccountNonce()
	RunAccountBalance()
	RunBatch()
	RunBlock()
	RunDataSubmission()
	RunStorage()
	RunRpc()
	RunTransactionOptions()
	RunTransactionPayment()
	RunCustomTransaction()
}

// v1 is Actual value, v2 is Expected value
func AssertEq[T comparable](v1 T, v2 T, message string) {
	if v1 != v2 {
		panic(fmt.Sprintf("Failure. Message: %v, Actual: %v, Expected: %v", message, v1, v2))
	}
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
