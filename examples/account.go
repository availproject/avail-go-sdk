package examples

import "fmt"

func RunAccount() {
	RunAccountCreation()
	RunAccountBalance()
	RunAccountNonce()

	fmt.Println("RunAccount finished correctly.")
}
