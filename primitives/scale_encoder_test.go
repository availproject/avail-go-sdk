package primitives

import (
	"math/big"
	"testing"

	"github.com/itering/scale.go/utiles/uint128"
)

func TestEncoderBool(t *testing.T) {
	var testParameters = map[bool]string{}
	// Min
	testParameters[false] = "0x00"
	// Max
	testParameters[true] = "0x01"

	for key, expected := range testParameters {
		{
			var actual = "0x" + Encoder.Encode(key)
			if actual != expected {
				t.Fatalf(`Encoder Bool Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
		{
			var actual = "0x"
			Encoder.EncodeTo(key, &actual)
			if actual != expected {
				t.Fatalf(`Encoder Bool Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
	}
}

func TestEncoderUint8(t *testing.T) {
	var testParameters = map[uint8]string{}
	// Min
	testParameters[0] = "0x00"
	// Min + 1
	testParameters[1] = "0x01"
	// Mid
	testParameters[128] = "0x80"
	// Max - 1
	testParameters[254] = "0xfe"
	// Max
	testParameters[255] = "0xff"

	for key, expected := range testParameters {
		{
			var actual = "0x" + Encoder.Encode(key)
			if actual != expected {
				t.Fatalf(`Encoder U8 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
		{
			var actual = "0x"
			Encoder.EncodeTo(key, &actual)
			if actual != expected {
				t.Fatalf(`Encoder U8 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
	}
}

func TestEncoderUint16(t *testing.T) {
	var testParameters = map[uint16]string{}
	// Min
	testParameters[0] = "0x0000"
	// Min + 1
	testParameters[1] = "0x0100"
	// Mid
	testParameters[32768] = "0x0080"
	// Max - 1
	testParameters[65534] = "0xfeff"
	// Max
	testParameters[65535] = "0xffff"

	for key, expected := range testParameters {
		{
			var actual = "0x" + Encoder.Encode(key)
			if actual != expected {
				t.Fatalf(`Encoder U16 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
		{
			var actual = "0x"
			Encoder.EncodeTo(key, &actual)
			if actual != expected {
				t.Fatalf(`Encoder U16 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
	}
}

func TestEncoderUint32(t *testing.T) {
	var testParameters = map[uint32]string{}
	// Min
	testParameters[0] = "0x00000000"
	// Min + 1
	testParameters[1] = "0x01000000"
	// Mid
	testParameters[2147483648] = "0x00000080"
	// Max - 1
	testParameters[4294967294] = "0xfeffffff"
	// Max
	testParameters[4294967295] = "0xffffffff"

	for key, expected := range testParameters {
		{
			var actual = "0x" + Encoder.Encode(key)
			if actual != expected {
				t.Fatalf(`Encoder U32 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
		{
			var actual = "0x"
			Encoder.EncodeTo(key, &actual)
			if actual != expected {
				t.Fatalf(`Encoder U32 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
	}
}

func TestEncoderUint64(t *testing.T) {
	var testParameters = map[uint64]string{}
	// Min
	testParameters[0] = "0x0000000000000000"
	// Min + 1
	testParameters[1] = "0x0100000000000000"
	// Mid
	testParameters[9223372036854775807] = "0xffffffffffffff7f"
	// Max - 1
	testParameters[18446744073709551614] = "0xfeffffffffffffff"
	// Max
	testParameters[18446744073709551615] = "0xffffffffffffffff"

	for key, expected := range testParameters {
		{
			var actual = "0x" + Encoder.Encode(key)
			if actual != expected {
				t.Fatalf(`Encoder U64 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
		{
			var actual = "0x"
			Encoder.EncodeTo(key, &actual)
			if actual != expected {
				t.Fatalf(`Encoder U64 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
	}
}

func TestEncoderUint128(t *testing.T) {
	var testParameters = map[uint128.Uint128]string{}
	// Min
	var res0, _ = new(big.Int).SetString("0", 10)
	testParameters[uint128.FromBig(res0)] = "0x00000000000000000000000000000000"
	// Min + 1
	var res1, _ = new(big.Int).SetString("1", 10)
	testParameters[uint128.FromBig(res1)] = "0x01000000000000000000000000000000"
	var res2, _ = new(big.Int).SetString("9223372036854775807", 10)
	testParameters[uint128.FromBig(res2)] = "0xffffffffffffff7f0000000000000000"
	var res7, _ = new(big.Int).SetString("18446744073709551615", 10)
	testParameters[uint128.FromBig(res7)] = "0xffffffffffffffff0000000000000000"
	var res8, _ = new(big.Int).SetString("18446744073709551616", 10)
	testParameters[uint128.FromBig(res8)] = "0x00000000000000000100000000000000"

	for key, expected := range testParameters {
		{
			var actual = "0x" + Encoder.Encode(key)
			if actual != expected {
				t.Fatalf(`Encoder U128 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
		{
			var actual = "0x"
			Encoder.EncodeTo(key, &actual)
			if actual != expected {
				t.Fatalf(`Encoder U128 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
	}
}

func TestEncoderCompactUint32(t *testing.T) {
	var testParameters = map[uint32]string{}
	// Min
	testParameters[0] = "0x00"
	// Min + 1
	testParameters[1] = "0x04"
	// Mid
	testParameters[2147483648] = "0x0300000080"
	// Max - 1
	testParameters[4294967294] = "0x03feffffff"
	// Max
	testParameters[4294967295] = "0x03ffffffff"

	for key, expected := range testParameters {
		{
			var actual = "0x" + Encoder.Encode(CompactU32{Value: key})
			if actual != expected {
				t.Fatalf(`Encoder U32 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
		{
			var actual = "0x"
			Encoder.EncodeTo(CompactU32{Value: key}, &actual)
			if actual != expected {
				t.Fatalf(`Encoder U32 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
	}
}

func TestEncoderCompactUint64(t *testing.T) {
	var testParameters = map[uint64]string{}
	// Min
	testParameters[0] = "0x00"
	// Min + 1
	testParameters[1] = "0x04"
	// Mid
	testParameters[9223372036854775807] = "0x13ffffffffffffff7f"
	// Max - 1
	testParameters[18446744073709551614] = "0x13feffffffffffffff"
	// Max
	testParameters[18446744073709551615] = "0x13ffffffffffffffff"

	for key, expected := range testParameters {
		{
			var actual = "0x" + Encoder.Encode(CompactU64{Value: key})
			if actual != expected {
				t.Fatalf(`Encoder U64 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
		{
			var actual = "0x"
			Encoder.EncodeTo(CompactU64{Value: key}, &actual)
			if actual != expected {
				t.Fatalf(`Encoder U64 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
	}
}

func TestEncoderCompactUint128(t *testing.T) {
	var testParameters = map[uint128.Uint128]string{}
	// Min
	var res0, _ = new(big.Int).SetString("0", 10)
	testParameters[uint128.FromBig(res0)] = "0x00"
	// Min + 1
	var res1, _ = new(big.Int).SetString("1", 10)
	testParameters[uint128.FromBig(res1)] = "0x04"
	var res2, _ = new(big.Int).SetString("9223372036854775807", 10)
	testParameters[uint128.FromBig(res2)] = "0x13ffffffffffffff7f"
	var res7, _ = new(big.Int).SetString("18446744073709551615", 10)
	testParameters[uint128.FromBig(res7)] = "0x13ffffffffffffffff"
	var res8, _ = new(big.Int).SetString("18446744073709551616", 10)
	testParameters[uint128.FromBig(res8)] = "0x17000000000000000001"

	for key, expected := range testParameters {
		{
			var actual = "0x" + Encoder.Encode(CompactU128{Value: key})
			if actual != expected {
				t.Fatalf(`Encoder U128 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
		{
			var actual = "0x"
			Encoder.EncodeTo(CompactU128{Value: key}, &actual)
			if actual != expected {
				t.Fatalf(`Encoder U128 Failure. Input %v, Output %v, Expected Output %v`, key, actual, expected)
			}
		}
	}
}

type StructWithUint32 struct {
	Value uint32
}

type DummyStruct struct {
	Array         [5]byte
	Slice         []byte
	Primitive     uint64
	Compact       uint32           `scale:"compact"`
	CompactStruct StructWithUint32 `scale:"compact"`
}

type DummyStruct2 struct {
	value uint32
}

func (this DummyStruct2) EncodeTo(dest *string) {
	*dest = "Success"
}

func (this *DummyStruct2) Decode(decoder *Decoder) error {
	this.value = uint32(0xbeef)
	return nil
}

type DummyStruct3 struct {
	value uint32
}

func (this *DummyStruct3) EncodeTo(dest *string) {
	*dest = "Success"
}

func TestEncoderArray(t *testing.T) {
	var expected = "0x0001020304"
	// Primitives
	{
		var array = [5]byte{0, 1, 2, 3, 4}
		var actual = "0x" + Encoder.Encode(array)
		if actual != expected {
			t.Fatalf(`Encoder Fixed Array Primitives. Input %v, Output %v, Expected Output %v`, array, actual, expected)
		}

		var actual2 = "0x"
		Encoder.EncodeTo(array, &actual2)
		if actual != expected {
			t.Fatalf(`Encoder Fixed Array Primitives. Input %v, Output %v, Expected Output %v`, array, actual, expected)
		}
	}

	// Structures
	{
		var expected = "0x000102030414000102030480000000000000003c04000102030414000102030480000000000000003c04"
		var el = DummyStruct{
			Array:         [5]byte{0, 1, 2, 3, 4},
			Slice:         []byte{0, 1, 2, 3, 4},
			Primitive:     128,
			Compact:       15,
			CompactStruct: StructWithUint32{Value: 1},
		}
		var el2 = el
		var array = [2]DummyStruct{el, el2}
		var actual = "0x" + Encoder.Encode(array)
		if actual != expected {
			t.Fatalf(`Encoder Fixed Array Struct. Output %v, Expected Output %v`, actual, expected)
		}

		var actual2 = "0x"
		Encoder.EncodeTo(array, &actual2)
		if actual != expected {
			t.Fatalf(`Encoder Fixed Array Struct. Output %v, Expected Output %v`, actual, expected)
		}
	}
}

func TestEncoderSlice(t *testing.T) {
	var expected = "0x140001020304"
	// Primitives
	{
		var array = []byte{0, 1, 2, 3, 4}
		var actual = "0x" + Encoder.Encode(array)
		if actual != expected {
			t.Fatalf(`Encoder Fixed Array Primitives. Input %v, Output %v, Expected Output %v`, array, actual, expected)
		}

		var actual2 = "0x"
		Encoder.EncodeTo(array, &actual2)
		if actual != expected {
			t.Fatalf(`Encoder Fixed Array Primitives. Input %v, Output %v, Expected Output %v`, array, actual, expected)
		}
	}

	// Structures
	{
		var expected = "0x08000102030414000102030480000000000000003c04000102030414000102030480000000000000003c04"
		var el = DummyStruct{
			Array:         [5]byte{0, 1, 2, 3, 4},
			Slice:         []byte{0, 1, 2, 3, 4},
			Primitive:     128,
			Compact:       15,
			CompactStruct: StructWithUint32{Value: 1},
		}
		var el2 = el
		var array = []DummyStruct{el, el2}
		var actual = "0x" + Encoder.Encode(array)
		if actual != expected {
			t.Fatalf(`Encoder Slice Struct. Output %v, Expected Output %v`, actual, expected)
		}

		var actual2 = "0x"
		Encoder.EncodeTo(array, &actual2)
		if actual != expected {
			t.Fatalf(`Encoder Slice Struct. Output %v, Expected Output %v`, actual, expected)
		}
	}
}

func TestEncoderStructures(t *testing.T) {
	{
		var expected = "0x000102030414000102030480000000000000003c04"
		var el = DummyStruct{
			Array:         [5]byte{0, 1, 2, 3, 4},
			Slice:         []byte{0, 1, 2, 3, 4},
			Primitive:     128,
			Compact:       15,
			CompactStruct: StructWithUint32{Value: 1},
		}
		var actual = "0x" + Encoder.Encode(el)
		if actual != expected {
			t.Fatalf(`Encoder Structure. Output %v, Expected Output %v`, actual, expected)
		}

		var actual2 = "0x"
		Encoder.EncodeTo(el, &actual2)
		if actual != expected {
			t.Fatalf(`Encoder Structure. Output %v, Expected Output %v`, actual, expected)
		}
	}

	{
		// Method ref/pointer struct + ref method
		var expected = "Success"
		var actual1 = Encoder.Encode(DummyStruct2{})
		if actual1 != expected {
			t.Fatalf(`Encoder Structure Method. Output %v, Expected Output %v`, actual1, expected)
		}

		var actual2 = Encoder.Encode(&DummyStruct2{})
		if actual2 != expected {
			t.Fatalf(`Encoder Structure Method. Output %v, Expected Output %v`, actual2, expected)
		}
	}

	{
		// Method ref/pointer struct + pointer method
		var expected = "Success"
		var actual1 = Encoder.Encode(DummyStruct3{})
		if actual1 != expected {
			t.Fatalf(`Encoder Structure Method. Output %v, Expected Output %v`, actual1, expected)
		}

		var actual2 = Encoder.Encode(&DummyStruct3{})
		if actual2 != expected {
			t.Fatalf(`Encoder Structure Method. Output %v, Expected Output %v`, actual2, expected)
		}
	}
}

func TestEncoderString(t *testing.T) {
	{
		var expected = "0x4054686973204973204120537472696e67"
		var el = "This Is A String"
		var actual = "0x" + Encoder.Encode(el)
		if actual != expected {
			t.Fatalf(`Encoder String. Output %v, Expected Output %v`, actual, expected)
		}

		var actual2 = "0x"
		Encoder.EncodeTo(el, &actual2)
		if actual != expected {
			t.Fatalf(`Encoder String. Output %v, Expected Output %v`, actual, expected)
		}
	}

	{
		var expected = "0x00"
		var el = ""
		var actual = "0x" + Encoder.Encode(el)
		if actual != expected {
			t.Fatalf(`Encoder String. Output %v, Expected Output %v`, actual, expected)
		}

		var actual2 = "0x"
		Encoder.EncodeTo(el, &actual2)
		if actual != expected {
			t.Fatalf(`Encoder String. Output %v, Expected Output %v`, actual, expected)
		}
	}
}

func TestEncoderOption(t *testing.T) {
	{
		var expected = "0x00"
		var el = NewNone[uint16]()
		var actual = "0x" + Encoder.Encode(el)
		if actual != expected {
			t.Fatalf(`Encoder String. Output %v, Expected Output %v`, actual, expected)
		}

		var actual2 = "0x"
		Encoder.EncodeTo(el, &actual2)
		if actual != expected {
			t.Fatalf(`Encoder String. Output %v, Expected Output %v`, actual, expected)
		}
	}

	{
		var expected = "0x016400"
		var el = NewSome(uint16(100))
		var actual = "0x" + Encoder.Encode(el)
		if actual != expected {
			t.Fatalf(`Encoder String. Output %v, Expected Output %v`, actual, expected)
		}

		var actual2 = "0x"
		Encoder.EncodeTo(el, &actual2)
		if actual != expected {
			t.Fatalf(`Encoder String. Output %v, Expected Output %v`, actual, expected)
		}
	}
}
