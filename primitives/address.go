package primitives

import (
	"errors"
	"fmt"

	"github.com/vedhavyas/go-subkey/v2"
)

// Do not change the order of field members.
type MultiAddress struct {
	VariantIndex uint8
	Id           Option[AccountId]
	Index        Option[uint32]
	Raw          Option[[]byte]
	Address32    Option[[32]byte]
	Address20    Option[[20]byte]
}

func NewMultiAddressId(value AccountId) MultiAddress {
	address := MultiAddress{}
	address.Id = Some(value)
	address.VariantIndex = 0
	return address
}

func (this MultiAddress) ToAccountId() Option[AccountId] {
	if this.Id.IsSome() {
		return Some(this.Id.Unwrap())
	}
	return None[AccountId]()
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
		value := AccountId{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Id = Some(value)
	case 1:
		value := uint32(0)
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Index = Some(value)
	case 2:
		value := []byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Raw = Some(value)
	case 3:
		value := [32]byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Address32 = Some(value)
	case 4:
		value := [20]byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		this.Address20 = Some(value)
	default:
		return errors.New(fmt.Sprintf(`MultiAddress Decode failure. Unknown Variant index: %v`, this.VariantIndex))
	}

	return nil
}

type AccountId struct {
	Value H256
}

func (this AccountId) ToSS58() string {
	return subkey.SS58Encode(this.Value.Value[:], 42)
}

func (this AccountId) ToAddress() string {
	return this.ToSS58()
}

func (this AccountId) ToHuman() string {
	return this.ToSS58()
}

func (this AccountId) ToString() string {
	return this.Value.ToHex()
}

func (this AccountId) ToMultiAddress() MultiAddress {
	return NewMultiAddressId(this)
}

func NewAccountIdFromKeyPair(keyPair subkey.KeyPair) AccountId {
	h256, err := NewH256FromByteSlice(keyPair.AccountID())
	if err != nil {
		// This should never happen
		panic(err)
	}

	return AccountId{Value: h256}
}

func NewAccountIdFromAddress(address string) (AccountId, error) {
	var _, accountBytes, err = subkey.SS58Decode(address)
	if err != nil {
		return AccountId{}, err
	}
	h256, err := NewH256FromByteSlice(accountBytes)
	if err != nil {
		return AccountId{}, err
	}
	res := AccountId{Value: h256}
	return res, nil
}

func NewAccountIdFromMultiAddress(address MultiAddress) (AccountId, error) {
	if address.Id.IsNone() {
		return AccountId{}, errors.New("Cannot decode multiaddress")
	}

	return address.Id.Unwrap(), nil
}
