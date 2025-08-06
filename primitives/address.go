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

func (m MultiAddress) ToAccountId() Option[AccountId] {
	if m.Id.IsSome() {
		return Some(m.Id.Unwrap())
	}
	return None[AccountId]()
}

func (m *MultiAddress) EncodeTo(dest *string) {
	Encoder.EncodeTo(m.VariantIndex, dest)

	if m.Id.IsSome() {
		Encoder.EncodeTo(m.Id.Unwrap(), dest)
	} else if m.Index.IsSome() {
		Encoder.EncodeTo(m.Index.Unwrap(), dest)
	} else if m.Raw.IsSome() {
		Encoder.EncodeTo(m.Raw.Unwrap(), dest)
	} else if m.Address32.IsSome() {
		Encoder.EncodeTo(m.Address32.Unwrap(), dest)
	} else if m.Address20.IsSome() {
		Encoder.EncodeTo(m.Address20.Unwrap(), dest)
	} else {
		panic("Something Went Wrong with MultiAddress EncodeTo")
	}
}

func (m *MultiAddress) Decode(decoder *Decoder) error {
	if err := decoder.Decode(&m.VariantIndex); err != nil {
		return err
	}

	switch m.VariantIndex {
	case 0:
		value := AccountId{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		m.Id = Some(value)
	case 1:
		value := uint32(0)
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		m.Index = Some(value)
	case 2:
		value := []byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		m.Raw = Some(value)
	case 3:
		value := [32]byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		m.Address32 = Some(value)
	case 4:
		value := [20]byte{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		m.Address20 = Some(value)
	default:
		return errors.New(fmt.Sprintf(`MultiAddress Decode failure. Unknown Variant index: %v`, m.VariantIndex))
	}

	return nil
}

type AccountId struct {
	Value H256
}

func (a AccountId) ToSS58() string {
	return subkey.SS58Encode(a.Value.Value[:], 42)
}

func (a AccountId) ToAddress() string {
	return a.ToSS58()
}

func (a AccountId) ToHuman() string {
	return a.ToSS58()
}

func (a AccountId) ToString() string {
	return a.Value.ToHex()
}

func (a AccountId) ToMultiAddress() MultiAddress {
	return NewMultiAddressId(a)
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
