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

func (this SimpleEnum) ToString() string {
	switch this.VariantIndex {
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

func (this ComplexEnum) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Nothing"
	case 1:
		return fmt.Sprintf("Set: %v", this.Day.Unwrap())
	case 2:
		return fmt.Sprintf("Set: %v", this.Month.Unwrap())
	case 3:
		return fmt.Sprintf("Set: %v", this.Year.Unwrap())
	default:
		panic("Unknown ComplexEnum Variant Index")
	}
}

func (this *ComplexEnum) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)

	if this.Day.IsSome() {
		prim.Encoder.EncodeTo(this.Day.Unwrap(), dest)
	}

	if this.Month.IsSome() {
		prim.Encoder.EncodeTo(this.Month.Unwrap(), dest)
	}

	if this.Year.IsSome() {
		prim.Encoder.EncodeTo(this.Year.Unwrap(), dest)
	}
}

func (this *ComplexEnum) Decode(decoder *prim.Decoder) error {
	*this = ComplexEnum{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
		var t uint16
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Day.Set(t)
	case 2:
		var t uint8
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Month.Set(t)
	case 3:
		var t uint32
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Year.Set(t)
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
