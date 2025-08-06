package examples

import (
	"errors"
	"fmt"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

// ANCHOR: simpleenum
type SimpleEnum struct {
	VariantIndex uint8
}

func (se SimpleEnum) ToString() string {
	switch se.VariantIndex {
	case 0:
		return "Nothing"
	case 1:
		return "Day"
	case 2:
		return "Month"
	case 3:
		return "Year"
	default:
		panic("Unknown SimpleEnum Variant Index")
	}
}

// ANCHOR_END: simpleenum

// ANCHOR: complexenum
type ComplexEnum struct {
	VariantIndex uint8
	Day          prim.Option[uint16]
	Month        prim.Option[uint8]
	Year         prim.Option[uint32]
}

func (se ComplexEnum) ToString() string {
	switch se.VariantIndex {
	case 0:
		return "Nothing"
	case 1:
		return fmt.Sprintf("Set: %v", se.Day.Unwrap())
	case 2:
		return fmt.Sprintf("Set: %v", se.Month.Unwrap())
	case 3:
		return fmt.Sprintf("Set: %v", se.Year.Unwrap())
	default:
		panic("Unknown ComplexEnum Variant Index")
	}
}

func (se *ComplexEnum) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(se.VariantIndex, dest)

	if se.Day.IsSome() {
		prim.Encoder.EncodeTo(se.Day.Unwrap(), dest)
	}

	if se.Month.IsSome() {
		prim.Encoder.EncodeTo(se.Month.Unwrap(), dest)
	}

	if se.Year.IsSome() {
		prim.Encoder.EncodeTo(se.Year.Unwrap(), dest)
	}
}

func (se *ComplexEnum) Decode(decoder *prim.Decoder) error {
	*se = ComplexEnum{}

	if err := decoder.Decode(&se.VariantIndex); err != nil {
		return err
	}

	switch se.VariantIndex {
	case 0:
	case 1:
		var t uint16
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		se.Day.Set(t)
	case 2:
		var t uint8
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		se.Month.Set(t)
	case 3:
		var t uint32
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		se.Year.Set(t)
	default:
		return errors.New("Unknown ComplexEnum Variant Index while Decoding")
	}

	return nil
}

// ANCHOR_END: complexenum

func tst() {
	// VariantIndex out of range
	enum := ComplexEnum{}
	enum.VariantIndex = 125

	// VariantIndex and data not matching
	enum.VariantIndex = 0
	enum.Year.Set(1990)

	// Too many data fields are set
	enum.VariantIndex = 1
	enum.Day.Set(24)
	enum.Year.Set(1990)
}
