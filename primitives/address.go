package primitives

import (
	"errors"
	"fmt"
)

// Do not change the order of field members.
type MultiAddress struct {
	Id        Option[H256]
	Index     Option[uint32]
	Raw       Option[[]byte]
	Address32 Option[[32]byte]
	Address20 Option[[20]byte]
}

func emptyMultiAddress() MultiAddress {
	return MultiAddress{
		Id:        NewNone[H256](),
		Index:     NewNone[uint32](),
		Raw:       NewNone[[]byte](),
		Address32: NewNone[[32]byte](),
		Address20: NewNone[[20]byte](),
	}
}

func NewMultiAddressId(value H256) MultiAddress {
	address := emptyMultiAddress()
	address.Id = NewSome(value)
	return address
}

func (this *MultiAddress) EncodeTo(dest *string) {
	if this.Id.IsSome() {
		Encoder.EncodeTo(uint8(0), dest)
		Encoder.EncodeTo(this.Id.Unwrap(), dest)
	} else if this.Index.IsSome() {
		Encoder.EncodeTo(uint8(1), dest)
		Encoder.EncodeTo(this.Index.Unwrap(), dest)
	} else if this.Raw.IsSome() {
		Encoder.EncodeTo(uint8(2), dest)
		Encoder.EncodeTo(this.Raw.Unwrap(), dest)
	} else if this.Address32.IsSome() {
		Encoder.EncodeTo(uint8(3), dest)
		Encoder.EncodeTo(this.Address32.Unwrap(), dest)
	} else if this.Address20.IsSome() {
		Encoder.EncodeTo(uint8(4), dest)
		Encoder.EncodeTo(this.Address20.Unwrap(), dest)
	} else {
		panic("Something Went Wrong with MultiAddress EncodeTo")
	}
}

func (this *MultiAddress) Decode(decoder *Decoder) error {
	result := emptyMultiAddress()
	variantIndex := uint8(0)
	decoder.Decode(&variantIndex)
	if variantIndex == 0 {
		value := H256{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		result.Id = NewSome(value)
	} else if variantIndex == 1 {
		value := uint32(0)
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		result.Index = NewSome(value)
	} else if variantIndex == 2 {
		value := []byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		result.Raw = NewSome(value)
	} else if variantIndex == 3 {
		value := [32]byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		result.Address32 = NewSome(value)
	} else if variantIndex == 4 {
		value := [20]byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		result.Address20 = NewSome(value)
	} else {
		return errors.New(fmt.Sprintf(`MultiAddress Decode failure. Unknown Variant index: %v`, variantIndex))
	}

	*this = result
	return nil
}
