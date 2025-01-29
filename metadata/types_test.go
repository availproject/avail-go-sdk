package metadata

import (
	"fmt"
	"testing"
)

func TestBalanceDisplay(t *testing.T) {
	// 1.0
	balance, err := NewBalanceFromString("1000000000000000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "1.0 Avail", "Balance is not correctly displayed")

	// 1.01
	balance, err = NewBalanceFromString("1010000000000000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "1.01 Avail", "Balance is not correctly displayed")

	// 1.01005
	balance, err = NewBalanceFromString("1010050000000000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "1.01005 Avail", "Balance is not correctly displayed")

	// 1.21
	balance, err = NewBalanceFromString("1210000000000000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "1.21 Avail", "Balance is not correctly displayed")

	// 100.01
	balance, err = NewBalanceFromString("100010000000000000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "100.01 Avail", "Balance is not correctly displayed")

	// 102.01
	balance, err = NewBalanceFromString("102010000000000000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "102.01 Avail", "Balance is not correctly displayed")

	// 0.1
	balance, err = NewBalanceFromString("100000000000000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "0.1 Avail", "Balance is not correctly displayed")

	// 0.01
	balance, err = NewBalanceFromString("10000000000000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "0.01 Avail", "Balance is not correctly displayed")

	// 0.00102
	balance, err = NewBalanceFromString("1020000000000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "0.00102 Avail", "Balance is not correctly displayed")

	// 0.000000000001
	balance, err = NewBalanceFromString("1000000")
	panicOnError(err)
	assertEq(balance.ToHuman(), "0.000000000001 Avail", "Balance is not correctly displayed")

	// 0.000000000000000001
	balance, err = NewBalanceFromString("1")
	panicOnError(err)
	assertEq(balance.ToHuman(), "0.000000000000000001 Avail", "Balance is not correctly displayed")
}

func TestPerbilDisplay(t *testing.T) {
	// 100.0%
	perbill := Perbill{Value: 1_000_000_000}
	assertEq(perbill.ToHuman(), "100.0%", "Perbill is not correctly displayed")

	// 90.5%
	perbill = Perbill{Value: 905_000_000}
	assertEq(perbill.ToHuman(), "90.5%", "Perbill is not correctly displayed")

	// 90.01%
	perbill = Perbill{Value: 900_100_000}
	assertEq(perbill.ToHuman(), "90.01%", "Perbill is not correctly displayed")

	// 1.1%
	perbill = Perbill{Value: 11_000_000}
	assertEq(perbill.ToHuman(), "1.1%", "Perbill is not correctly displayed")

	// 1.0105%
	perbill = Perbill{Value: 10_105_000}
	assertEq(perbill.ToHuman(), "1.0105%", "Perbill is not correctly displayed")

	// 1.01%
	perbill = Perbill{Value: 10_100_000}
	assertEq(perbill.ToHuman(), "1.01%", "Perbill is not correctly displayed")

	// 1.00001%
	perbill = Perbill{Value: 10_000_100}
	assertEq(perbill.ToHuman(), "1.00001%", "Perbill is not correctly displayed")

	// 1.0%
	perbill = Perbill{Value: 10_000_000}
	assertEq(perbill.ToHuman(), "1.0%", "Perbill is not correctly displayed")

	// 0.9909%
	perbill = Perbill{Value: 9_909_000}
	assertEq(perbill.ToHuman(), "0.9909%", "Perbill is not correctly displayed")

	// 0.0005%
	perbill = Perbill{Value: 5_000}
	assertEq(perbill.ToHuman(), "0.0005%", "Perbill is not correctly displayed")

	// 0.0000001%
	perbill = Perbill{Value: 1}
	assertEq(perbill.ToHuman(), "0.0000001%", "Perbill is not correctly displayed")

}

func assertTrue(v bool, message string) {
	if !v {
		panic(fmt.Sprintf("Failure. Message: %v", message))
	}
}

// v1 is Actual value, v2 is Expected value
func assertEq[T comparable](v1 T, v2 T, message string) {
	if v1 != v2 {
		panic(fmt.Sprintf("Failure. Message: %v, Actual: %v, Expected: %v", message, v1, v2))
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
