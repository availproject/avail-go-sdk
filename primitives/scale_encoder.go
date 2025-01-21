package primitives

import (
	"fmt"
	"math/big"
	"math/bits"
	"reflect"

	SType "github.com/itering/scale.go/types"
	"github.com/itering/scale.go/utiles/uint128"
)

type encoderT struct{}

var Encoder encoderT

// Fixed arrays
func (encoderT) array(value reflect.Value, dest *string) bool {
	len := uint32(value.Len())
	if len == 0 {
		return true
	}

	// Lets see if we can encode the element, if not we bail out
	tmp := ""
	firstElement := value.Index(0)
	if !Encoder.EncodeTo(firstElement.Interface(), &tmp) {
		return false
	}

	for i := 0; i < int(len); i++ {
		if !Encoder.EncodeTo(value.Index(i).Interface(), dest) {
			return false
		}
	}

	return true
}

// Dynamic arrays
func (encoderT) slice(value reflect.Value, dest *string) bool {
	len := uint32(value.Len())
	if len == 0 {
		Encoder.EncodeTo(CompactU32{Value: 0}, dest)
		return true
	}

	// Lets see if we can encode the element, if not we bail out
	tmp := ""
	firstElement := value.Index(0)
	if !Encoder.EncodeTo(firstElement.Interface(), &tmp) {
		return false
	}

	Encoder.EncodeTo(CompactU32{Value: len}, dest)
	for i := 0; i < int(len); i++ {
		if !Encoder.EncodeTo(value.Index(i).Interface(), dest) {
			return false
		}
	}

	return true
}

func (encoderT) callMethod(value reflect.Value, dest *string) bool {
	methodName := "EncodeTo"
	method := value.MethodByName(methodName)

	if !method.IsValid() {
		return false
	}

	args := []reflect.Value{reflect.ValueOf(dest)}
	var _ = method.Call(args)

	return true

}

func (encoderT) structureFields(value reflect.Value, dest *string) bool {
	valueType := value.Type()

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		scaleTag := field.Tag.Get("scale")
		if scaleTag == "ignore" {
			continue
		}

		isCompact := scaleTag == "compact"
		fieldValue := value.Field(i)
		if !Encoder.encodeToInner(fieldValue, dest, isCompact) {
			return false
		}
	}

	return true
}

func (encoderT) structure(value reflect.Value, dest *string) bool {
	// See if there is a method that doesn't need a pointer
	if Encoder.callMethod(value, dest) {
		return true
	}

	ptrValue := reflect.New(value.Type())
	ptrValue.Elem().Set(value) // Copy the original value into the pointer

	// See if there is a method that needs a pointer
	if Encoder.callMethod(ptrValue, dest) {
		return true
	}

	// See if we can encode all the fields
	return Encoder.structureFields(value, dest)
}

func (encoderT) pointer(value reflect.Value, dest *string) bool {
	if value.Kind() != reflect.Ptr {
		return false
	}

	if Encoder.callMethod(value, dest) {
		return true
	}

	return false
}

func (encoderT) primitives(value reflect.Value, dest *string, isCompact bool) bool {
	kind := value.Kind()
	name := value.Type().Name()
	if !isCompact {
		if kind == reflect.Bool {
			genericEncodeTo("bool", value.Interface(), dest)
			return true
		}
		if kind == reflect.Uint8 {
			genericEncodeTo("u8", value.Interface(), dest)
			return true
		}
		if kind == reflect.Uint16 {
			genericEncodeTo("u16", value.Interface(), dest)
			return true
		}
		if kind == reflect.Uint32 {
			genericEncodeTo("u32", value.Interface(), dest)
			return true
		}
		if kind == reflect.Uint64 {
			genericEncodeTo("u64", value.Interface(), dest)
			return true
		}
		if kind == reflect.String {
			genericEncodeTo("string", value.Interface(), dest)
			return true
		}
		if name == "Uint128" {
			genericEncodeTo("u128", value.Interface().(uint128.Uint128).String(), dest)
			return true
		}
	} else {
		if kind == reflect.Uint32 {
			val := value.Interface().(uint32)
			if val <= 1073741823 {
				genericEncodeTo("Compact<u32>", val, dest)
				return true
			}
			Encoder.EncodeTo(uint8(0b11), dest)
			Encoder.EncodeTo(uint32(val), dest)
			return true
		}
		if kind == reflect.Uint64 {
			val := value.Interface().(uint64)
			if val <= 1073741823 {
				genericEncodeTo("Compact", uint128.From64(val).String(), dest)
				return true
			}

			leadingZeros := bits.LeadingZeros64(val)
			bytesNeeded := 8 - leadingZeros/8
			midRes := uint8(0b11) + uint8(((bytesNeeded - 4) << 2))
			Encoder.EncodeTo(uint8(midRes), dest)

			v := val
			for i := 0; i < bytesNeeded; i++ {
				Encoder.EncodeTo(uint8(v), dest)
				v = v >> 8
			}
			return true
		}
		if name == "Uint128" {
			val := value.Interface().(uint128.Uint128)
			maxVal := uint128.From64(1073741823)
			if val.Cmp(maxVal) != 1 {
				genericEncodeTo("Compact", val.String(), dest)
				return true
			}

			leadingZeros := leadingZeros(val.Big(), 128)
			bytesNeeded := 16 - leadingZeros/8
			midRes := uint8(0b11) + uint8(((bytesNeeded - 4) << 2))
			Encoder.EncodeTo(uint8(midRes), dest)

			v := val
			for i := 0; i < bytesNeeded; i++ {
				b := v.Big().Uint64() & 0xFF
				Encoder.EncodeTo(uint8(b), dest)
				v = v.Rsh(8)
			}
			return true
		}
	}

	return false
}

func (encoderT) Encode(value interface{}) string {
	encoded := ""
	Encoder.EncodeTo(value, &encoded)
	return encoded
}

func (encoderT) EncodeTo(value interface{}, dest *string) bool {
	valueOf := reflect.ValueOf(value)
	return Encoder.encodeToInner(valueOf, dest, false)
}

func (encoderT) encodeToInner(value reflect.Value, dest *string, isCompact bool) bool {
	kind := value.Kind()
	name := value.Type().Name()

	if Encoder.primitives(value, dest, isCompact) {
		return true
	}

	switch kind {
	case reflect.Array:
		return Encoder.array(value, dest)
	case reflect.Slice:
		return Encoder.slice(value, dest)
	case reflect.Struct:
		return Encoder.structure(value, dest)
	case reflect.Pointer:
		return Encoder.encodeToInner(value.Elem(), dest, isCompact)
	default:
		panic(fmt.Sprintf(`Unknown Type. Name: %v, Kind: %v`, name, kind))
	}
}

func leadingZeros(x *big.Int, totalBits int) int {
	if x.Sign() == 0 { // Handle zero case
		return totalBits
	}
	return totalBits - x.BitLen()
}

func (encoderT) FixedArrayTo(value []byte, dest *string) {
	*dest = *dest + Hex.ToHex(value[:])
}

func genericEncodeTo(typeString string, value interface{}, dest *string) {
	*dest += SType.Encode(typeString, value)
}
