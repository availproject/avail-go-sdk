package primitives

import (
	"errors"
	"fmt"
)

// Do not change the order of field members.
type MultiAddress struct {
	VariantIndex uint8
	Id           Option[H256]
	Index        Option[uint32]
	Raw          Option[[]byte]
	Address32    Option[[32]byte]
	Address20    Option[[20]byte]
}

func NewMultiAddressId(value H256) MultiAddress {
	address := MultiAddress{}
	address.Id = NewSome(value)
	address.VariantIndex = 0
	return address
}

func (this *MultiAddress) EncodeTo(dest *string) {
	Encoder.EncodeTo(this.VariantIndex, dest)

	if this.Id.IsSome() {
		Encoder.EncodeTo(this.Id.Unwrap(), dest)
	} else if this.Index.IsSome() {
		Encoder.EncodeTo(this.Index.Unwrap(), dest)
	} else if this.Raw.IsSome() {
		Encoder.EncodeTo(this.Raw.Unwrap(), dest)
	} else if this.Address32.IsSome() {
		Encoder.EncodeTo(this.Address32.Unwrap(), dest)
	} else if this.Address20.IsSome() {
		Encoder.EncodeTo(this.Address20.Unwrap(), dest)
	} else {
		panic("Something Went Wrong with MultiAddress EncodeTo")
	}
}

func (this *MultiAddress) Decode(decoder *Decoder) error {
	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
		value := H256{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Id = NewSome(value)
	case 1:
		value := uint32(0)
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Index = NewSome(value)
	case 2:
		value := []byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Raw = NewSome(value)
	case 3:
		value := [32]byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Address32 = NewSome(value)
	case 4:
		value := [20]byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Address20 = NewSome(value)
	default:
		return errors.New(fmt.Sprintf(`MultiAddress Decode failure. Unknown Variant index: %v`, this.VariantIndex))
	}

	return nil
}
