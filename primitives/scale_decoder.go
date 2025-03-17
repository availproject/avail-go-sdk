package primitives

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"

	SType "github.com/itering/scale.go/types"
	"github.com/itering/scale.go/types/scaleBytes"
	"github.com/itering/scale.go/utiles/uint128"
	"github.com/shopspring/decimal"
)

type Decoder struct {
	ScaleBytes scaleBytes.ScaleBytes
}

// Fixed arrays
func (this *Decoder) array(value reflect.Value) error {
	elemType := value.Type().Elem()
	arrayType := reflect.ArrayOf(value.Len(), elemType)
	arrayPointer := reflect.New(arrayType)
	newArray := arrayPointer.Elem()

	for i := 0; i < int(value.Len()); i++ {
		elem := reflect.New(elemType).Elem()

		if err := this.Decode(elem.Addr().Interface()); err != nil {
			return err
		}

		newArray.Index(i).Set(elem)
	}

	value.Set(newArray)
	return nil
}

// Dynamic arrays
func (this *Decoder) slice(value reflect.Value) error {
	len := CompactU32{}
	if err := this.Decode(&len); err != nil {
		return err
	}

	elemType := value.Type().Elem()
	newSlice := reflect.MakeSlice(value.Type(), int(len.Value), int(len.Value))

	for i := 0; i < int(len.Value); i++ {
		elem := reflect.New(elemType).Elem()
		if err := this.Decode(elem.Addr().Interface()); err != nil {
			return err
		}

		newSlice.Index(i).Set(elem)
	}

	value.Set(newSlice)
	return nil
}

func (this *Decoder) callMethod(value reflect.Value) error {
	methodName := "Decode"
	method := value.MethodByName(methodName)

	args := []reflect.Value{reflect.ValueOf(this)}
	results := method.Call(args)

	if len(results) == 0 {
		return errors.New(`Decoder failed. Method Decode was called but the result of the call has no return values`)
	}

	res := results[0].Interface()
	if res == nil {
		return nil
	}

	return res.(error)
}

func (this *Decoder) structureFields(value reflect.Value, isCompact bool) error {
	valueType := value.Type()

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		if !field.IsExported() {
			return errors.New(fmt.Sprintf(`Decoder failed. Struct field is not exported. Field Name: %v`, field.Name))
		}

		scaleTag := field.Tag.Get("scale")
		if scaleTag == "ignore" {
			continue
		}

		filedIsCompact := isCompact || scaleTag == "compact"

		fieldValue := value.Field(i)

		if err := this.decodeInner(fieldValue.Addr(), filedIsCompact); err != nil {
			return err
		}
	}

	return nil
}

func (this *Decoder) primitives(value reflect.Value, isCompact bool) (error, bool) {
	kind := value.Kind().String()
	name := value.Type().Name()

	if !isCompact {
		if kind == "bool" {
			if !this.HasAtLeastRemainingBytes(1) {
				return errors.New(`Decoder failed. Out of Bytes`), true
			}

			decoder := SType.Bool{}
			decoder.Init(this.ScaleBytes, nil)
			decoder.Process()
			this.ScaleBytes.Offset = decoder.Data.Offset

			res := decoder.Value.(bool)
			value.Set(reflect.ValueOf(res))
			return nil, true
		}
		if kind == "uint8" {
			if !this.HasAtLeastRemainingBytes(1) {
				return errors.New(`Decoder failed. Out of Bytes`), true
			}

			decoder := SType.U8{}
			decoder.Init(this.ScaleBytes, nil)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset

			res := uint8(decoder.Value.(int))
			value.Set(reflect.ValueOf(res))
			return nil, true
		}
		if kind == "uint16" {
			if !this.HasAtLeastRemainingBytes(2) {
				return errors.New(`Decoder failed. Out of Bytes`), true
			}

			decoder := SType.U16{}
			decoder.Init(this.ScaleBytes, nil)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset

			res := decoder.Value.(uint16)
			value.Set(reflect.ValueOf(res))
			return nil, true
		}
		if kind == "uint32" {
			if !this.HasAtLeastRemainingBytes(4) {
				return errors.New(`Decoder failed. Out of Bytes`), true
			}

			decoder := SType.U32{}
			decoder.Init(this.ScaleBytes, nil)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset

			res := decoder.Value.(uint32)
			value.Set(reflect.ValueOf(res))
			return nil, true
		}
		if kind == "uint64" {
			if !this.HasAtLeastRemainingBytes(8) {
				return errors.New(`Decoder failed. Out of Bytes`), true
			}

			decoder := SType.U64{}
			decoder.Init(this.ScaleBytes, nil)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset

			res := decoder.Value.(uint64)
			value.Set(reflect.ValueOf(res))
			return nil, true
		}
		if kind == "string" {
			decoder := SType.String{}
			decoder.Init(this.ScaleBytes, nil)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset

			res := decoder.Value.(string)
			value.Set(reflect.ValueOf(res))
			return nil, true
		}
		if name == "Uint128" {
			if !this.HasAtLeastRemainingBytes(16) {
				return errors.New(`Decoder failed. Out of Bytes`), true
			}

			decoder := SType.U128{}
			decoder.Init(this.ScaleBytes, nil)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset

			valueDecoded, _ := new(big.Int).SetString(decoder.Value.(string), 10)
			res := uint128.FromBig(valueDecoded)

			value.Set(reflect.ValueOf(res))
			return nil, true
		}
	} else {
		if kind == "uint16" {
			decoder := SType.Compact{}
			options := SType.ScaleDecoderOption{}
			options.SubType = "u16"
			decoder.Init(this.ScaleBytes, &options)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset
			res := decoder.Value.(uint16)
			value.Set(reflect.ValueOf(res))
			return nil, true
		}
		if kind == "uint32" {
			decoder := SType.CompactU32{}
			decoder.Init(this.ScaleBytes, nil)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset
			res := uint32(decoder.Value.(int))
			value.Set(reflect.ValueOf(res))
			return nil, true

		}
		if kind == "uint64" {
			decoder := SType.Compact{}
			options := SType.ScaleDecoderOption{}
			options.SubType = "u64"
			decoder.Init(this.ScaleBytes, &options)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset
			res := decoder.Value.(uint64)
			value.Set(reflect.ValueOf(res))
			return nil, true
		}
		if name == "Uint128" {
			decoder := SType.Compact{}
			options := SType.ScaleDecoderOption{}
			options.SubType = "u128"
			decoder.Init(this.ScaleBytes, &options)
			decoder.Process()

			this.ScaleBytes.Offset = decoder.Data.Offset

			valueBytes := decoder.Value.(decimal.Decimal)
			valueDecoded, _ := new(big.Int).SetString(valueBytes.String(), 10)

			res := uint128.FromBig(valueDecoded)
			value.Set(reflect.ValueOf(res))
			return nil, true
		}
	}

	return nil, false
}

func (this *Decoder) Decode(value interface{}) error {
	valueOf := reflect.ValueOf(value)

	if err := this.decodeInner(valueOf, false); err != nil {
		// If failed reset the value to zero
		if valueOf.Kind() == reflect.Ptr {
			elem := valueOf.Elem()
			elem.Set(reflect.Zero(elem.Type()))
		}

		return err
	}

	return nil
}

func (this *Decoder) decodeInner(value reflect.Value, isCompact bool) error {
	if value.Kind() != reflect.Ptr {
		return errors.New("Decoder failed. The passed value is not of pointer type")
	}

	if !value.Elem().CanSet() {
		return errors.New("Decoder failed. The passed value cannot be changed. CanSet is set to false")
	}

	if this.hasCallMethod(value) {
		return this.callMethod(value)
	}

	pointee := value.Elem()
	if res, ok := this.primitives(pointee, isCompact); ok {
		return res
	}

	switch pointee.Kind() {
	case reflect.Slice:
		return this.slice(pointee)
	case reflect.Array:
		return this.array(pointee)
	case reflect.Struct:
		return this.structureFields(pointee, isCompact)
	default:
		elemKind := pointee.Kind()
		elemName := pointee.Type().Name()
		return errors.New(fmt.Sprintf(`Decoder failed. Unknown Value. Name: %v Type: %v`, elemName, elemKind))
	}
}

func (this *Decoder) hasCallMethod(value reflect.Value) bool {
	methodName := "Decode"
	method := value.MethodByName(methodName)

	return method.IsValid()
}

func NewDecoder(data []byte, offset int) Decoder {
	return Decoder{
		ScaleBytes: scaleBytes.ScaleBytes{Data: data, Offset: offset},
	}
}

func (this *Decoder) Offset() int {
	return this.ScaleBytes.Offset
}

func (this *Decoder) RemainingLength() int {
	return this.ScaleBytes.GetRemainingLength()
}

func (this *Decoder) HasAtLeastRemainingBytes(atLeast int) bool {
	return this.ScaleBytes.GetRemainingLength() >= atLeast
}

func (this *Decoder) NextBytes(length int) []byte {
	return this.ScaleBytes.GetNextBytes(length)
}

// This is just a different name for GetNextBytes
func (this *Decoder) StaticArray(byteCount int) []byte {
	return this.ScaleBytes.GetNextBytes(byteCount)
}
