package examples

func Run() {
	RunAccountNonce()
	RunAccountBalance()
	RunBatch()
	RunBlock()
	RunDataSubmission()
	RunEvents()
	RunStorage()
	RunRpc()
	RunTransactionOptions()
	RunTransactionPayment()
	RunCustomTransaction()
	RunBlockTransactions()
}

func AssertEq[T comparable](v1 T, v2 T, message string) {
	if v1 != v2 {
		panic(message)
	}
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
