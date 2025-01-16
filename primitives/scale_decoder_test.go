package primitives

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/itering/scale.go/utiles/uint128"
)

func TestDecoderBool(t *testing.T) {
	var testParameters = map[string]bool{}
	// Min
	testParameters["0x00"] = false
	// Max
	testParameters["0x01"] = true

	for key, expected := range testParameters {
		var decoder = NewDecoder(Hex.FromHex(key), 0)
		var actual = false
		decoder.Decode(&actual)
		if actual != expected {
			t.Fatalf(`Decoder Bool Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
		}
	}
}

func TestDecoderUint8(t *testing.T) {
	var testParameters = map[string]uint8{}
	// Min
	testParameters["0x00"] = 0
	// Min + 1
	testParameters["0x01"] = 1
	// Mid
	testParameters["0x80"] = 128
	// Max - 1
	testParameters["0xfe"] = 254
	// Max
	testParameters["0xff"] = 255

	for key, expected := range testParameters {
		var decoder = NewDecoder(Hex.FromHex(key), 0)
		var actual = uint8(0)
		decoder.Decode(&actual)
		if actual != expected {
			t.Fatalf(`Decoder Uint8 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
		}
	}
}

func TestDecoderUint16(t *testing.T) {
	var testParameters = map[string]uint16{}
	// Min
	testParameters["0x0000"] = 0
	// Min + 1
	testParameters["0x0100"] = 1
	// Mid
	testParameters["0x0080"] = 32768
	// Max - 1
	testParameters["0xfeff"] = 65534
	// Max
	testParameters["0xffff"] = 65535

	for key, expected := range testParameters {
		var decoder = NewDecoder(Hex.FromHex(key), 0)
		var actual = uint16(0)
		decoder.Decode(&actual)
		if actual != expected {
			t.Fatalf(`Decoder Uint16 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
		}
	}
}

func TestDecoderUint32(t *testing.T) {
	var testParameters = map[string]uint32{}
	// Min
	testParameters["0x00000000"] = 0
	// Min + 1
	testParameters["0x01000000"] = 1
	// Mid
	testParameters["0x00000080"] = 2147483648
	// Max - 1
	testParameters["0xfeffffff"] = 4294967294
	// Max
	testParameters["0xffffffff"] = 4294967295

	for key, expected := range testParameters {
		var decoder = NewDecoder(Hex.FromHex(key), 0)
		var actual = uint32(0)
		decoder.Decode(&actual)
		if actual != expected {
			t.Fatalf(`Decoder Uint32 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
		}
	}
}

func TestDecoderUint64(t *testing.T) {
	var testParameters = map[string]uint64{}
	// Min
	testParameters["0x0000000000000000"] = 0
	// Min + 1
	testParameters["0x0100000000000000"] = 1
	// Mid
	testParameters["0xffffffffffffff7f"] = 9223372036854775807
	// Max - 1
	testParameters["0xfeffffffffffffff"] = 18446744073709551614
	// Max
	testParameters["0xffffffffffffffff"] = 18446744073709551615

	for key, expected := range testParameters {
		var decoder = NewDecoder(Hex.FromHex(key), 0)
		var actual = uint64(0)
		decoder.Decode(&actual)
		if actual != expected {
			t.Fatalf(`Decoder Uint64 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
		}
	}
}

func TestDecoderUint128(t *testing.T) {
	var testParameters = map[string]uint128.Uint128{}
	// Min
	var res0, _ = new(big.Int).SetString("0", 10)
	testParameters["0x00000000000000000000000000000000"] = uint128.FromBig(res0)
	// Min + 1
	var res1, _ = new(big.Int).SetString("1", 10)
	testParameters["0x01000000000000000000000000000000"] = uint128.FromBig(res1)
	var res2, _ = new(big.Int).SetString("9223372036854775807", 10)
	testParameters["0xffffffffffffff7f0000000000000000"] = uint128.FromBig(res2)
	var res7, _ = new(big.Int).SetString("18446744073709551615", 10)
	testParameters["0xffffffffffffffff0000000000000000"] = uint128.FromBig(res7)
	var res8, _ = new(big.Int).SetString("18446744073709551616", 10)
	testParameters["0x00000000000000000100000000000000"] = uint128.FromBig(res8)

	for key, expected := range testParameters {
		var decoder = NewDecoder(Hex.FromHex(key), 0)
		var actual = uint128.Uint128{}
		decoder.Decode(&actual)
		if actual != expected {
			t.Fatalf(`Decoder Uint128 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
		}
	}
}

func TestDecoderCompactU32(t *testing.T) {
	var testParameters = map[string]uint32{}
	// Min
	testParameters["0x00"] = 0
	// Min + 1
	testParameters["0x04"] = 1
	// Mid
	testParameters["0x0300000080"] = 2147483648
	// Max - 1
	testParameters["0x03feffffff"] = 4294967294
	// Max
	testParameters["0x03ffffffff"] = 4294967295

	for key, expected := range testParameters {
		var decoder = NewDecoder(Hex.FromHex(key), 0)
		var actual = CompactU32{}
		decoder.Decode(&actual)
		if actual.Value != expected {
			t.Fatalf(`Decoder Compact Uint32 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
		}
	}
}

func TestDecoderCompactU64(t *testing.T) {
	var testParameters = map[string]uint64{}
	// Min
	testParameters["0x00"] = 0
	// Min + 1
	testParameters["0x04"] = 1
	// Mid
	testParameters["0x13ffffffffffffff7f"] = 9223372036854775807
	// Max - 1
	testParameters["0x13feffffffffffffff"] = 18446744073709551614
	// Max
	testParameters["0x13ffffffffffffffff"] = 18446744073709551615

	for key, expected := range testParameters {
		var decoder = NewDecoder(Hex.FromHex(key), 0)
		var actual = CompactU64{}
		decoder.Decode(&actual)
		if actual.Value != expected {
			t.Fatalf(`Decoder Compact Uint64 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
		}
	}
}

func TestDecoderCompactUint128(t *testing.T) {
	var testParameters = map[string]uint128.Uint128{}
	// Min
	var res0, _ = new(big.Int).SetString("0", 10)
	testParameters["0x00"] = uint128.FromBig(res0)
	// Min + 1
	var res1, _ = new(big.Int).SetString("1", 10)
	testParameters["0x04"] = uint128.FromBig(res1)
	var res2, _ = new(big.Int).SetString("9223372036854775807", 10)
	testParameters["0x13ffffffffffffff7f"] = uint128.FromBig(res2)
	var res7, _ = new(big.Int).SetString("18446744073709551615", 10)
	testParameters["0x13ffffffffffffffff"] = uint128.FromBig(res7)
	var res8, _ = new(big.Int).SetString("18446744073709551616", 10)
	testParameters["0x17000000000000000001"] = uint128.FromBig(res8)

	for key, expected := range testParameters {
		var decoder = NewDecoder(Hex.FromHex(key), 0)
		var actual = CompactU128{}
		decoder.Decode(&actual)
		if actual.Value != expected {
			t.Fatalf(`Decoder Compact Uint128 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
		}
	}
}

func TestDecoderArray(t *testing.T) {
	// Primitives
	{
		var expected = [5]byte{0, 1, 2, 3, 4}
		var input = "0x0001020304"
		var actual = [5]byte{}
		var decoder = NewDecoder(Hex.FromHex(input), 0)
		decoder.Decode(&actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf(`Decoder Array Failure. Input %v, Output %v, Expected Output %v`, input, actual, expected)
		}
	}

	// Structures
	{
		var input = "0x000102030414000102030480000000000000003c000102030414000102030480000000000000003c"
		var el = DummyStruct{
			Array:     [5]byte{0, 1, 2, 3, 4},
			Slice:     []byte{0, 1, 2, 3, 4},
			Primitive: 128,
			Compact:   15,
		}
		var el2 = el
		var expected = [2]DummyStruct{el, el2}
		var actual = [2]DummyStruct{}
		var decoder = NewDecoder(Hex.FromHex(input), 0)
		decoder.Decode(&actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf(`Decoder Array Failure. Input %v, Output %v, Expected Output %v`, input, actual, expected)
		}
	}
}

func TestDecoderSlice(t *testing.T) {
	// Primitives
	{
		var input = "0x140001020304"
		var expected = []byte{0, 1, 2, 3, 4}
		var actual = []byte{}
		var decoder = NewDecoder(Hex.FromHex(input), 0)
		decoder.Decode(&actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf(`Decoder Slice Failure. Input %v, Output %v, Expected Output %v`, input, actual, expected)
		}
	}

	// Structures
	{
		var input = "0x08000102030414000102030480000000000000003c000102030414000102030480000000000000003c"
		var el = DummyStruct{
			Array:     [5]byte{0, 1, 2, 3, 4},
			Slice:     []byte{0, 1, 2, 3, 4},
			Primitive: 128,
			Compact:   15,
		}
		var el2 = el
		var expected = []DummyStruct{el, el2}
		var actual = []DummyStruct{}
		var decoder = NewDecoder(Hex.FromHex(input), 0)
		decoder.Decode(&actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf(`Decoder Slice Failure. Input %v, Output %v, Expected Output %v`, input, actual, expected)
		}
	}
}

func TestDecoderStructures(t *testing.T) {
	{
		var input = "0x000102030414000102030480000000000000003c"
		var expected = DummyStruct{
			Array:     [5]byte{0, 1, 2, 3, 4},
			Slice:     []byte{0, 1, 2, 3, 4},
			Primitive: 128,
			Compact:   15,
		}
		var actual = DummyStruct{}
		var decoder = NewDecoder(Hex.FromHex(input), 0)
		decoder.Decode(&actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf(`Decoder Structure. Output %v, Expected Output %v`, actual, expected)
		}
	}

	{
		// Method ref/pointer struct + pointer method
		var expected = DummyStruct2{
			value: uint32(0xbeef),
		}
		var input = Encoder.Encode(&expected)
		var actual = DummyStruct2{}
		var decoder = NewDecoder(Hex.FromHex(input), 0)
		decoder.Decode(&actual)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf(`Decoder Structure. Output %v, Expected Output %v`, actual, expected)
		}
	}
}
