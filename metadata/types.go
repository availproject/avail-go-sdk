package metadata

import (
	"github.com/itering/scale.go/utiles/uint128"
	"github.com/vedhavyas/go-subkey/v2"

	"errors"

	prim "go-sdk/primitives"
)

type Balance struct {
	Value uint128.Uint128
}

func (this Balance) ToHuman() string {
	var stringValue = this.Value.String()

	if len(stringValue) <= 18 {
		var result = "0."
		var trailing = removeTrailingZeros(stringValue)
		if trailing == "" {
			result += "0"
		}
		return result + trailing + " Avail"
	}

	var result = ""
	for i := 0; i < len(stringValue); i++ {
		result = string(stringValue[len(stringValue)-i-1]) + result
		if i == 18 {
			result = "." + result
		}
	}

	return removeTrailingZeros(result) + " Avail"
}

func removeTrailingZeros(s string) string {
	for {
		if len(s) == 0 {
			break
		}
		if s[len(s)-1] != '0' {
			break
		}
		s = s[:len(s)-1]
	}
	return s
}

// Do not add, remove or change any of the field members.
type AccountId struct {
	Value prim.H256
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

func NewAccountIdFromAddress(address string) (AccountId, error) {
	var _, accountBytes, err = subkey.SS58Decode(address)
	if err != nil {
		return AccountId{}, err
	}
	h256, err := prim.NewH256FromByteSlice(accountBytes)
	if err != nil {
		return AccountId{}, err
	}
	res := AccountId{Value: h256}
	return res, nil
}

// Do not add, remove or change any of the field members.
type DispatchInfo struct {
	Weight      Weight
	Class       DispatchClass
	PaysFee     Pays
	FeeModifier DispatchFeeModifier
}

// Do not add, remove or change any of the field members.
type Weight struct {
	RefTime   uint64
	ProofSize uint64
}

// Do not add, remove or change any of the field members.
type DispatchClass struct {
	VariantIndex uint8
}

func (this DispatchClass) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Normal"
	case 1:
		return "Operational"
	case 2:
		return "Mandatory"
	default:
		panic("Unknown DispatchCall Variant Index")
	}
}

// Do not add, remove or change any of the field members.
type Pays struct {
	VariantIndex uint8
}

func (this Pays) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Yes"
	case 1:
		return "No"
	default:
		panic("Unknown Pays Variant Index")
	}
}

// Do not add, remove or change any of the field members.
type DispatchFeeModifier struct {
	WeightMaximumFee    prim.Option[uint128.Uint128]
	WeightFeeDivider    prim.Option[uint32]
	WeightFeeMultiplier prim.Option[uint32]
}

// Do not add, remove or change any of the field members.
type DispatchError struct {
	VariantIndex  uint8
	Module        prim.Option[ModuleError]
	Token         prim.Option[TokenError]
	Arithmetic    prim.Option[ArithmeticError]
	Transactional prim.Option[TransactionalError]
}

func (this DispatchError) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Other"
	case 1:
		return "CannotLookup"
	case 2:
		return "BadOrigin"
	case 3:
		return "Module"
	case 4:
		return "ConsumerRemaining"
	case 5:
		return "NoProviders"
	case 6:
		return "TooManyConsumers"
	case 7:
		return "Token"
	case 8:
		return "Arithmetic"
	case 9:
		return "Transactional"
	case 10:
		return "Exhausted"
	case 11:
		return "Corruption"
	case 12:
		return "Unavailable"
	case 13:
		return "RootNotAllowed"
	default:
		panic("Unknown DispatchError Variant Index")
	}
}

func (this DispatchError) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)

	if this.Module.IsSome() {
		prim.Encoder.EncodeTo(this.Module.Unwrap(), dest)
	}

	if this.Token.IsSome() {
		prim.Encoder.EncodeTo(this.Token.Unwrap(), dest)
	}

	if this.Arithmetic.IsSome() {
		prim.Encoder.EncodeTo(this.Arithmetic.Unwrap(), dest)
	}

	if this.Transactional.IsSome() {
		prim.Encoder.EncodeTo(this.Transactional.Unwrap(), dest)
	}
}

func (this *DispatchError) Decode(decoder *prim.Decoder) error {
	*this = DispatchError{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
	case 2:
	case 3:
		var t ModuleError
		decoder.Decode(&t)
		this.Module.Set(t)
	case 4:
	case 5:
	case 6:
	case 7:
		var t TokenError
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Token.Set(t)
	case 8:
		var t ArithmeticError
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Arithmetic.Set(t)
	case 9:
		var t TransactionalError
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Transactional.Set(t)
	case 10:
	case 11:
	case 12:
	case 13:
	default:
		return errors.New("Unknown DispatchError Variant Index while Decoding")
	}

	return nil
}

// Do not add, remove or change any of the field members.
type ModuleError struct {
	Index uint8
	Error [4]byte
}

// Do not add, remove or change any of the field members.
type TokenError struct {
	VariantIndex uint8
}

func (this TokenError) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "FundsUnavailable"
	case 1:
		return "OnlyProvider"
	case 2:
		return "BelowMinimum"
	case 3:
		return "CannotCreate"
	case 4:
		return "UnknownAsset"
	case 5:
		return "Frozen"
	case 6:
		return "Unsupported"
	case 7:
		return "CannotCreateHold"
	case 8:
		return "NotExpendable"
	case 9:
		return "Blocked"
	default:
		panic("Unknown TokenError Variant Index")
	}
}

// Do not add, remove or change any of the field members.
type ArithmeticError struct {
	VariantIndex uint8
}

func (this ArithmeticError) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Underflow"
	case 1:
		return "Overflow"
	case 2:
		return "DivisionByZero"
	default:
		panic("Unknown ArithmeticError Variant Index")
	}
}

// Do not add, remove or change any of the field members.
type TransactionalError struct {
	VariantIndex uint8
}

func (this TransactionalError) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "LimitReached"
	case 1:
		return "NoLayer"
	default:
		panic("Unknown TransactionalError Variant Index")
	}
}
