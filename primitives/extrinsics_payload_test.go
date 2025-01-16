package primitives

import (
	"reflect"
	"testing"

	"github.com/itering/scale.go/utiles/uint128"
)

func TestEraDecode(t *testing.T) {
	var expected = NewEra(5, 20)
	var encode = Encoder.Encode(expected)
	var input = FromHex(encode)
	var decoder = NewDecoder(input, 0)

	var actual = Era{}
	decoder.Decode(&actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(`Decoder Era Failure. Input %v, Output %v, Expected Output %v`, input, actual, expected)
	}
}

func TestExtraDecode(t *testing.T) {
	var era = NewEra(5, 20)
	var expected = Extra{
		Era:   era,
		Nonce: 5,
		Tip:   uint128.From64(uint64(123)),
		AppId: 3,
	}
	var encode = Encoder.Encode(expected)
	var input = FromHex(encode)
	var decoder = NewDecoder(input, 0)

	var actual = Extra{}
	decoder.Decode(&actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(`Decoder Extra Failure. Input %v, Output %v, Expected Output %v`, input, actual, expected)
	}
}
