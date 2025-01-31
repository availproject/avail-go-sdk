package metadata

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/itering/scale.go/utiles/uint128"
	"github.com/vedhavyas/go-subkey/v2"

	"errors"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type Balance struct {
	Value uint128.Uint128
}

func NewBalanceFromString(value string) (Balance, error) {
	var res, ok = new(big.Int).SetString(value, 10)
	if !ok {
		return Balance{}, errors.New("Failed to convert string to Balance")
	}
	return NewBalanceFromBigInt(res), nil
}

func NewBalanceFromBigInt(value *big.Int) Balance {
	return Balance{Value: uint128.FromBig(value)}
}

func (this Balance) String() string {
	return this.ToHuman()
}

func (this Balance) ToString() string {
	return this.ToHuman()
}

func (this Balance) ToHuman() string {
	var stringValue = this.Value.String()

	if len(stringValue) <= 18 {
		var result = "0."
		var trailing = removeTrailingZeros(stringValue)
		if trailing == "" {
			result += "0"
		} else {
			missingPlaces := 18 - len(stringValue)

			for i := 0; i < missingPlaces; i++ {
				result += "0"
			}
		}

		return result + trailing + " Avail"
	}

	var result = ""
	for i := 0; i < len(stringValue); i++ {
		result = string(stringValue[len(stringValue)-i-1]) + result
		if i == 17 {
			result = "." + result
		}
	}

	result = removeTrailingZeros(result)
	if strings.HasSuffix(result, ".") {
		result += "0"
	}

	return result + " Avail"
}

// Add returns this+v.
func (this Balance) Add(v Balance) Balance {
	return Balance{Value: this.Value.Add(v.Value)}
}

// Add64 returns this+v.
func (this Balance) Add64(v uint64) Balance {
	return Balance{Value: this.Value.Add64(v)}
}

// Add128 returns this+v.
func (this Balance) Add128(v uint128.Uint128) Balance {
	return Balance{Value: this.Value.Add(v)}
}

// Sub returns this-v.
func (this Balance) Sub(v Balance) Balance {
	return Balance{Value: this.Value.Sub(v.Value)}
}

// Sub64 returns this-v.
func (this Balance) Sub64(v uint64) Balance {
	return Balance{Value: this.Value.Sub64(v)}
}

// Sub128 returns this-v.
func (this Balance) Sub128(v uint128.Uint128) Balance {
	return Balance{Value: this.Value.Sub(v)}
}

// Mul returns this*v.
func (this Balance) Mul(v Balance) Balance {
	return Balance{Value: this.Value.Mul(v.Value)}
}

// Mul64 returns this*v.
func (this Balance) Mul64(v uint64) Balance {
	return Balance{Value: this.Value.Mul64(v)}
}

// Mul128 returns this*v.
func (this Balance) Mul128(v uint128.Uint128) Balance {
	return Balance{Value: this.Value.Mul(v)}
}

// Div returns this/v.
func (this Balance) Div(v Balance) Balance {
	return Balance{Value: this.Value.Div(v.Value)}
}

// Div64 returns this/v.
func (this Balance) Div64(v uint64) Balance {
	return Balance{Value: this.Value.Div64(v)}
}

// Div128 returns this/v.
func (this Balance) Div128(v uint128.Uint128) Balance {
	return Balance{Value: this.Value.Div(v)}
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

func (this AccountId) ToString() string {
	return this.Value.ToHex()
}

func (this AccountId) ToMultiAddress() prim.MultiAddress {
	return prim.NewMultiAddressId(this.Value)
}

func NewAccountIdFromKeyPair(keyPair subkey.KeyPair) AccountId {
	h256, err := prim.NewH256FromByteSlice(keyPair.AccountID())
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
	h256, err := prim.NewH256FromByteSlice(accountBytes)
	if err != nil {
		return AccountId{}, err
	}
	res := AccountId{Value: h256}
	return res, nil
}

func NewAccountIdFromMultiAddress(address prim.MultiAddress) (AccountId, error) {
	if address.Id.IsNone() {
		return AccountId{}, errors.New("Cannot decode multiaddress")
	}

	return AccountId{Value: address.Id.Unwrap()}, nil
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
	RefTime   uint64 `scale:"compact"`
	ProofSize uint64 `scale:"compact"`
}

// Do not add, remove or change any of the field members.
type DispatchClass struct {
	VariantIndex uint8
}

func (this DispatchClass) ToHuman() string {
	return this.ToString()
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

func (this DispatchClass) String() string {
	return this.ToString()
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

func (this DispatchError) ToHuman() string {
	return this.ToString()
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
		return fmt.Sprintf("Module. Index %v", this.Module.Unwrap().Index)
	case 4:
		return "ConsumerRemaining"
	case 5:
		return "NoProviders"
	case 6:
		return "TooManyConsumers"
	case 7:
		return fmt.Sprintf("Token. %v", this.Token.Unwrap().ToHuman())
	case 8:
		return fmt.Sprintf("Arithmetic. %v", this.Arithmetic.Unwrap().ToHuman())
	case 9:
		return fmt.Sprintf("Transactional. %v", this.Transactional.Unwrap().ToString())
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

func (this *DispatchError) EncodeTo(dest *string) {
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
		if err := decoder.Decode(&t); err != nil {
			return err
		}
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

func (this TokenError) ToHuman() string {
	return this.ToString()
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

func (this ArithmeticError) ToHuman() string {
	return this.ToString()
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

func (this TransactionalError) ToHuman() string {
	return this.ToString()
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

// Do not add, remove or change any of the field members.
type AccountData struct {
	Free     Balance
	Reserved Balance
	Frozen   Balance
	Flags    uint128.Uint128
}

// Do not add, remove or change any of the field members.
type PerDispatchClassU32 struct {
	Normal      uint32
	Operational uint32
	Mandatory   uint32
}

type Perbill struct {
	Value uint32
}

func NewPerbillFromU8(percent uint8) Perbill {
	if percent > 100 {
		panic("Percent cannot be more than 100")
	}

	value := 10_000_000 * uint32(percent)

	return Perbill{Value: value}
}

func (this Perbill) ToString() string {
	return this.ToHuman()
}

func (this Perbill) ToHuman() string {
	stringValue := strconv.FormatUint(uint64(this.Value), 10)

	if len(stringValue) <= 7 {
		addZeros := 7 - len(stringValue)

		var result = "0."
		for i := 0; i < addZeros; i++ {
			result += "0"
		}

		var trailing = removeTrailingZeros(stringValue)
		if trailing == "" {
			result = "0.0"
			trailing = ""
		}
		return result + trailing + "%"
	}

	var result = ""
	for i := 0; i < len(stringValue); i++ {
		result = string(stringValue[len(stringValue)-i-1]) + result
		if i == 6 {
			result = "." + result
		}
	}

	result = removeTrailingZeros(result)
	if strings.HasSuffix(result, ".") {
		result += "0"
	}

	return result + "%"
}

// Variant 0: Staked
// Variant 1: Stash
// Variant 2: Controller
// Variant 3: Account - Account field needs to be set up
// Variant 4 Nonce
type RewardDestination struct {
	VariantIndex uint8
	Account      prim.Option[AccountId]
}

func (this RewardDestination) ToHuman() string {
	return this.ToString()
}

func (this RewardDestination) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Staked"
	case 1:
		return "Stash"
	case 2:
		return "Controller"
	case 3:
		return fmt.Sprintf("Account: %v", this.Account.Unwrap().ToHuman())
	case 4:
		return "None"
	default:
		panic("Unknown RewardDestination Variant Index")
	}
}

func (this *RewardDestination) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)

	if this.Account.IsSome() {
		prim.Encoder.EncodeTo(this.Account.Unwrap(), dest)
	}
}

func (this *RewardDestination) Decode(decoder *prim.Decoder) error {
	*this = RewardDestination{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
	case 2:
	case 3:
		var t AccountId
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Account.Set(t)
	case 4:
	default:
		return errors.New("Unknown RewardDestination Variant Index while Decoding")
	}

	return nil
}

type ValidatorPrefs struct {
	Commission Perbill `scale:"compact"`
	Blocked    bool
}

type SessionKeys struct {
	Babe               prim.H256
	Grandpa            prim.H256
	ImOnline           prim.H256
	AuthorityDiscovery prim.H256
}

type CommissionClaimPermission struct {
	VariantIndex uint8
	Account      prim.Option[AccountId]
}

func (this CommissionClaimPermission) ToHuman() string {
	return this.ToString()
}

func (this CommissionClaimPermission) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Permissionless"
	case 1:
		return fmt.Sprintf("Account: %v", this.Account.Unwrap().ToHuman())
	default:
		panic("Unknown CommissionClaimPermission Variant Index")
	}
}

func (this *CommissionClaimPermission) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)

	if this.Account.IsSome() {
		prim.Encoder.EncodeTo(this.Account.Unwrap(), dest)
	}
}

func (this *CommissionClaimPermission) Decode(decoder *prim.Decoder) error {
	*this = CommissionClaimPermission{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
		var t AccountId
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Account.Set(t)
	default:
		return errors.New("Unknown RewardDestination Variant Index while Decoding")
	}

	return nil
}

type PoolRoles struct {
	Depositor AccountId
	Root      prim.Option[AccountId]
	Nominator prim.Option[AccountId]
	Bouncer   prim.Option[AccountId]
}

type PoolState struct {
	VariantIndex uint8
}

func (this PoolState) ToHuman() string {
	return this.ToString()
}

func (this PoolState) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Open"
	case 1:
		return "Blocked"
	case 2:
		return "Destroying"
	default:
		panic("Unknown PoolState Variant Index")
	}
}

func (this *PoolState) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)
}

func (this *PoolState) Decode(decoder *prim.Decoder) error {
	*this = PoolState{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
	case 2:
	default:
		return errors.New("Unknown PoolState Variant Index while Decoding")
	}

	return nil
}

type PoolCommission struct {
	Current         prim.Option[Tuple2[Perbill, AccountId]]
	Max             prim.Option[Perbill]
	ChangeRate      prim.Option[PoolCommissionChangeRate]
	ThrottleFrom    prim.Option[uint32]
	ClaimPermission prim.Option[CommissionClaimPermission]
}

type PoolCommissionChangeRate struct {
	MaxIncrease Perbill
	MinDelay    uint32
}

type PoolClaimPermission struct {
	VariantIndex uint8
}

func (this PoolClaimPermission) ToHuman() string {
	return this.ToString()
}

func (this PoolClaimPermission) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Permissioned"
	case 1:
		return "PermissionlessCompound"
	case 2:
		return "PermissionlessWithdraw"
	case 3:
		return "PermissionlessAll"
	default:
		panic("Unknown PoolState Variant Index")
	}
}

func (this *PoolClaimPermission) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)
}

func (this *PoolClaimPermission) Decode(decoder *prim.Decoder) error {
	*this = PoolClaimPermission{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
	case 2:
	case 3:
	default:
		return errors.New("Unknown PoolClaimPermission Variant Index while Decoding")
	}

	return nil
}

type Registration struct {
	Judgements []Tuple2[uint32, Judgement]
	Deposit    Balance
	Info       IdentityInfo
}

type IdentityInfo struct {
	Additional     []Tuple2[IdentityData, IdentityData]
	Display        IdentityData
	Legal          IdentityData
	Web            IdentityData
	Riot           IdentityData
	Email          IdentityData
	PgpFingerprint prim.Option[[20]byte]
	Image          IdentityData
	Twitter        IdentityData
}

type Judgement struct {
	VariantIndex uint8
	FeePaid      prim.Option[Balance]
}

func (this Judgement) ToHuman() string {
	return this.ToString()
}

func (this Judgement) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Unknown"
	case 1:
		return "FeePaid"
	case 2:
		return "Reasonable"
	case 3:
		return "KnownGood"
	case 4:
		return "OutOfDate"
	case 5:
		return "LowQuality"
	case 6:
		return "Erroneous"
	default:
		panic("Unknown Judgement Variant Index")
	}
}

func (this *Judgement) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)

	if this.FeePaid.IsSome() {
		prim.Encoder.EncodeTo(this.FeePaid.Unwrap(), dest)
	}
}

func (this *Judgement) Decode(decoder *prim.Decoder) error {
	*this = Judgement{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
		var t Balance
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.FeePaid.Set(t)
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	default:
		return errors.New("Unknown Judgement Variant Index while Decoding")
	}

	return nil
}

type IdentityData struct {
	VariantIndex uint8
	Raw0         prim.Option[[0]byte]
	Raw1         prim.Option[[1]byte]
	Raw2         prim.Option[[2]byte]
	Raw3         prim.Option[[3]byte]
	Raw4         prim.Option[[4]byte]
	Raw5         prim.Option[[5]byte]
	Raw6         prim.Option[[6]byte]
	Raw7         prim.Option[[7]byte]
	Raw8         prim.Option[[8]byte]
	Raw9         prim.Option[[9]byte]
	Raw10        prim.Option[[10]byte]
	Raw11        prim.Option[[11]byte]
	Raw12        prim.Option[[12]byte]
	Raw13        prim.Option[[13]byte]
	Raw14        prim.Option[[14]byte]
	Raw15        prim.Option[[15]byte]
	Raw16        prim.Option[[16]byte]
	Raw17        prim.Option[[17]byte]
	Raw18        prim.Option[[18]byte]
	Raw19        prim.Option[[19]byte]
	Raw20        prim.Option[[20]byte]
	Raw21        prim.Option[[21]byte]
	Raw22        prim.Option[[22]byte]
	Raw23        prim.Option[[23]byte]
	Raw24        prim.Option[[24]byte]
	Raw25        prim.Option[[25]byte]
	Raw26        prim.Option[[26]byte]
	Raw27        prim.Option[[27]byte]
	Raw28        prim.Option[[28]byte]
	Raw29        prim.Option[[29]byte]
	Raw30        prim.Option[[30]byte]
	Raw31        prim.Option[[31]byte]
	Raw32        prim.Option[[32]byte]
	BlakeTwo256  prim.Option[prim.H256]
	Sha256       prim.Option[prim.H256]
	Keccak256    prim.Option[prim.H256]
	ShaThree256  prim.Option[prim.H256]
}

func (this IdentityData) ToHuman() string {
	return this.ToString()
}

func (this IdentityData) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "None"
	case 1:
		return "Raw0"
	case 2:
		v := this.Raw1.Unwrap()
		return string(v[:])
	case 3:
		v := this.Raw2.Unwrap()
		return string(v[:])
	case 4:
		v := this.Raw3.Unwrap()
		return string(v[:])
	case 5:
		v := this.Raw4.Unwrap()
		return string(v[:])
	case 6:
		v := this.Raw5.Unwrap()
		return string(v[:])
	case 7:
		v := this.Raw6.Unwrap()
		return string(v[:])
	case 8:
		v := this.Raw7.Unwrap()
		return string(v[:])
	case 9:
		v := this.Raw8.Unwrap()
		return string(v[:])
	case 10:
		v := this.Raw9.Unwrap()
		return string(v[:])
	case 11:
		v := this.Raw10.Unwrap()
		return string(v[:])
	case 12:
		v := this.Raw11.Unwrap()
		return string(v[:])
	case 13:
		v := this.Raw12.Unwrap()
		return string(v[:])
	case 14:
		v := this.Raw13.Unwrap()
		return string(v[:])
	case 15:
		v := this.Raw14.Unwrap()
		return string(v[:])
	case 16:
		v := this.Raw15.Unwrap()
		return string(v[:])
	case 17:
		v := this.Raw16.Unwrap()
		return string(v[:])
	case 18:
		v := this.Raw17.Unwrap()
		return string(v[:])
	case 19:
		v := this.Raw18.Unwrap()
		return string(v[:])
	case 20:
		v := this.Raw19.Unwrap()
		return string(v[:])
	case 21:
		v := this.Raw20.Unwrap()
		return string(v[:])
	case 22:
		v := this.Raw21.Unwrap()
		return string(v[:])
	case 23:
		v := this.Raw22.Unwrap()
		return string(v[:])
	case 24:
		v := this.Raw23.Unwrap()
		return string(v[:])
	case 25:
		v := this.Raw24.Unwrap()
		return string(v[:])
	case 26:
		v := this.Raw25.Unwrap()
		return string(v[:])
	case 27:
		v := this.Raw26.Unwrap()
		return string(v[:])
	case 28:
		v := this.Raw27.Unwrap()
		return string(v[:])
	case 29:
		v := this.Raw28.Unwrap()
		return string(v[:])
	case 30:
		v := this.Raw29.Unwrap()
		return string(v[:])
	case 31:
		v := this.Raw30.Unwrap()
		return string(v[:])
	case 32:
		v := this.Raw31.Unwrap()
		return string(v[:])
	case 33:
		v := this.Raw32.Unwrap()
		return string(v[:])
	case 34:
		v := this.BlakeTwo256.Unwrap()
		return v.ToHexWith0x()
	case 35:
		v := this.Sha256.Unwrap()
		return v.ToHexWith0x()
	case 36:
		v := this.Keccak256.Unwrap()
		return v.ToHexWith0x()
	case 37:
		v := this.ShaThree256.Unwrap()
		return v.ToHexWith0x()
	default:
		panic("Unknown IdentityData Variant Index")
	}
}

func (this *IdentityData) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)
	if this.Raw1.IsSome() {
		prim.Encoder.EncodeTo(this.Raw1.Unwrap(), dest)
	}
	if this.Raw2.IsSome() {
		prim.Encoder.EncodeTo(this.Raw2.Unwrap(), dest)
	}
	if this.Raw3.IsSome() {
		prim.Encoder.EncodeTo(this.Raw3.Unwrap(), dest)
	}
	if this.Raw4.IsSome() {
		prim.Encoder.EncodeTo(this.Raw4.Unwrap(), dest)
	}
	if this.Raw5.IsSome() {
		prim.Encoder.EncodeTo(this.Raw5.Unwrap(), dest)
	}
	if this.Raw6.IsSome() {
		prim.Encoder.EncodeTo(this.Raw6.Unwrap(), dest)
	}
	if this.Raw7.IsSome() {
		prim.Encoder.EncodeTo(this.Raw7.Unwrap(), dest)
	}
	if this.Raw8.IsSome() {
		prim.Encoder.EncodeTo(this.Raw8.Unwrap(), dest)
	}
	if this.Raw9.IsSome() {
		prim.Encoder.EncodeTo(this.Raw9.Unwrap(), dest)
	}
	if this.Raw10.IsSome() {
		prim.Encoder.EncodeTo(this.Raw10.Unwrap(), dest)
	}
	if this.Raw11.IsSome() {
		prim.Encoder.EncodeTo(this.Raw11.Unwrap(), dest)
	}
	if this.Raw12.IsSome() {
		prim.Encoder.EncodeTo(this.Raw12.Unwrap(), dest)
	}
	if this.Raw13.IsSome() {
		prim.Encoder.EncodeTo(this.Raw13.Unwrap(), dest)
	}
	if this.Raw14.IsSome() {
		prim.Encoder.EncodeTo(this.Raw14.Unwrap(), dest)
	}
	if this.Raw15.IsSome() {
		prim.Encoder.EncodeTo(this.Raw15.Unwrap(), dest)
	}
	if this.Raw16.IsSome() {
		prim.Encoder.EncodeTo(this.Raw16.Unwrap(), dest)
	}
	if this.Raw17.IsSome() {
		prim.Encoder.EncodeTo(this.Raw17.Unwrap(), dest)
	}
	if this.Raw18.IsSome() {
		prim.Encoder.EncodeTo(this.Raw18.Unwrap(), dest)
	}
	if this.Raw19.IsSome() {
		prim.Encoder.EncodeTo(this.Raw19.Unwrap(), dest)
	}
	if this.Raw20.IsSome() {
		prim.Encoder.EncodeTo(this.Raw20.Unwrap(), dest)
	}
	if this.Raw21.IsSome() {
		prim.Encoder.EncodeTo(this.Raw21.Unwrap(), dest)
	}
	if this.Raw22.IsSome() {
		prim.Encoder.EncodeTo(this.Raw22.Unwrap(), dest)
	}
	if this.Raw23.IsSome() {
		prim.Encoder.EncodeTo(this.Raw23.Unwrap(), dest)
	}
	if this.Raw24.IsSome() {
		prim.Encoder.EncodeTo(this.Raw24.Unwrap(), dest)
	}
	if this.Raw25.IsSome() {
		prim.Encoder.EncodeTo(this.Raw25.Unwrap(), dest)
	}
	if this.Raw26.IsSome() {
		prim.Encoder.EncodeTo(this.Raw26.Unwrap(), dest)
	}
	if this.Raw27.IsSome() {
		prim.Encoder.EncodeTo(this.Raw27.Unwrap(), dest)
	}
	if this.Raw28.IsSome() {
		prim.Encoder.EncodeTo(this.Raw28.Unwrap(), dest)
	}
	if this.Raw29.IsSome() {
		prim.Encoder.EncodeTo(this.Raw29.Unwrap(), dest)
	}
	if this.Raw30.IsSome() {
		prim.Encoder.EncodeTo(this.Raw30.Unwrap(), dest)
	}
	if this.Raw31.IsSome() {
		prim.Encoder.EncodeTo(this.Raw31.Unwrap(), dest)
	}
	if this.Raw32.IsSome() {
		prim.Encoder.EncodeTo(this.Raw32.Unwrap(), dest)
	}
	if this.BlakeTwo256.IsSome() {
		prim.Encoder.EncodeTo(this.BlakeTwo256.Unwrap(), dest)
	}
	if this.Sha256.IsSome() {
		prim.Encoder.EncodeTo(this.Sha256.Unwrap(), dest)
	}
	if this.Keccak256.IsSome() {
		prim.Encoder.EncodeTo(this.Keccak256.Unwrap(), dest)
	}
	if this.ShaThree256.IsSome() {
		prim.Encoder.EncodeTo(this.ShaThree256.Unwrap(), dest)
	}
}

func (this *IdentityData) Decode(decoder *prim.Decoder) error {
	*this = IdentityData{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
		var t [0]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw0.Set(t)
	case 2:
		var t [1]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw1.Set(t)
	case 3:
		var t [2]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw2.Set(t)
	case 4:
		var t [3]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw3.Set(t)
	case 5:
		var t [4]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw4.Set(t)
	case 6:
		var t [5]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw5.Set(t)
	case 7:
		var t [6]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw6.Set(t)
	case 8:
		var t [7]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw7.Set(t)
	case 9:
		var t [8]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw8.Set(t)
	case 10:
		var t [9]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw9.Set(t)
	case 11:
		var t [10]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw10.Set(t)
	case 12:
		var t [11]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw11.Set(t)
	case 13:
		var t [12]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw12.Set(t)
	case 14:
		var t [13]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw13.Set(t)
	case 15:
		var t [14]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw14.Set(t)
	case 16:
		var t [15]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw15.Set(t)
	case 17:
		var t [16]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw16.Set(t)
	case 18:
		var t [17]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw17.Set(t)
	case 19:
		var t [18]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw18.Set(t)
	case 20:
		var t [19]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw19.Set(t)
	case 21:
		var t [20]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw20.Set(t)
	case 22:
		var t [21]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw21.Set(t)
	case 23:
		var t [22]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw22.Set(t)
	case 24:
		var t [23]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw23.Set(t)
	case 25:
		var t [24]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw24.Set(t)
	case 26:
		var t [25]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw25.Set(t)
	case 27:
		var t [26]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw26.Set(t)
	case 28:
		var t [27]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw27.Set(t)
	case 29:
		var t [28]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw28.Set(t)
	case 30:
		var t [29]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw29.Set(t)
	case 31:
		var t [30]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw30.Set(t)
	case 32:
		var t [31]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw31.Set(t)
	case 33:
		var t [32]byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Raw32.Set(t)
	case 34:
		var t prim.H256
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.BlakeTwo256.Set(t)
	case 35:
		var t prim.H256
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Sha256.Set(t)
	case 36:
		var t prim.H256
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Keccak256.Set(t)
	case 37:
		var t prim.H256
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.ShaThree256.Set(t)
	default:
		return errors.New("Unknown IdentityData Variant Index while Decoding")
	}

	return nil
}

type DispatchResult struct {
	VariantIndex uint8
	Err          prim.Option[DispatchError]
}

func (this DispatchResult) ToHuman() string {
	return this.ToString()
}

func (this DispatchResult) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Ok"
	case 1:
		return "Err"
	default:
		panic("Unknown DispatchResult Variant Index")
	}
}

func (this *DispatchResult) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)

	if this.Err.IsSome() {
		prim.Encoder.EncodeTo(this.Err.Unwrap(), dest)
	}
}

func (this *DispatchResult) Decode(decoder *prim.Decoder) error {
	*this = DispatchResult{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
		var t DispatchError
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Err.Set(t)
	default:
		return errors.New("Unknown DispatchResult Variant Index while Decoding")
	}

	return nil
}

type PoolBondExtra struct {
	VariantIndex uint8
	FreeBalance  prim.Option[Balance]
}

func (this PoolBondExtra) ToHuman() string {
	return this.ToString()
}

func (this PoolBondExtra) ToString() string {
	switch this.VariantIndex {
	case 0:
		return fmt.Sprintf("Free Balance: %v", this.FreeBalance.Unwrap().ToHuman())
	case 1:
		return "Rewards"
	default:
		panic("Unknown PoolBondExtra Variant Index")
	}
}

func (this *PoolBondExtra) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)

	if this.FreeBalance.IsSome() {
		prim.Encoder.EncodeTo(this.FreeBalance.Unwrap(), dest)
	}
}

func (this *PoolBondExtra) Decode(decoder *prim.Decoder) error {
	*this = PoolBondExtra{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
		var t Balance
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.FreeBalance.Set(t)
	case 1:
	default:
		return errors.New("Unknown PoolBondExtra Variant Index while Decoding")
	}

	return nil
}

type PoolRoleConfig struct {
	VariantIndex uint8
	Set          prim.Option[AccountId]
}

func (this PoolRoleConfig) ToHuman() string {
	return this.ToString()
}

func (this PoolRoleConfig) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Noop"
	case 1:
		return fmt.Sprintf("Set: %v", this.Set.Unwrap().ToHuman())
	case 2:
		return "Remove"
	default:
		panic("Unknown PoolRoleConfig Variant Index")
	}
}

func (this *PoolRoleConfig) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)

	if this.Set.IsSome() {
		prim.Encoder.EncodeTo(this.Set.Unwrap(), dest)
	}
}

func (this *PoolRoleConfig) Decode(decoder *prim.Decoder) error {
	*this = PoolRoleConfig{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
		var t AccountId
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.Set.Set(t)
	case 2:
	default:
		return errors.New("Unknown PoolRoleConfig Variant Index while Decoding")
	}

	return nil
}

type BlockLength struct {
	Max       PerDispatchClassU32
	Cols      uint32 `scale:"compact"`
	Rows      uint32 `scale:"compact"`
	ChunkSize uint32 `scale:"compact"`
}

type VectorMessageKind struct {
	VariantIndex     uint8
	ArbitraryMessage prim.Option[[]byte]
	FungibleToken    prim.Option[MessageFungibleToken]
}

type MessageFungibleToken struct {
	AssetId prim.H256
	Amount  Balance `scale:"compact"`
}

func (this VectorMessageKind) ToHuman() string {
	return this.ToString()
}

func (this VectorMessageKind) ToString() string {
	switch this.VariantIndex {
	case 0:
		return fmt.Sprintf("ArbitraryMessage: %v", string(this.ArbitraryMessage.Unwrap()))
	case 1:
		v := this.FungibleToken.Unwrap()
		return fmt.Sprintf("FungibleToken: Asset Id: %v,  Amount: %v", v.AssetId, v.Amount.String())
	default:
		panic("Unknown VectorMessageKind Variant Index")
	}
}

func (this *VectorMessageKind) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)

	if this.ArbitraryMessage.IsSome() {
		prim.Encoder.EncodeTo(this.ArbitraryMessage.Unwrap(), dest)
	}

	if this.FungibleToken.IsSome() {
		prim.Encoder.EncodeTo(this.FungibleToken.Unwrap(), dest)
	}
}

func (this *VectorMessageKind) Decode(decoder *prim.Decoder) error {
	*this = VectorMessageKind{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
		var t []byte
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.ArbitraryMessage.Set(t)
	case 1:
		var t MessageFungibleToken
		if err := decoder.Decode(&t); err != nil {
			return err
		}
		this.FungibleToken.Set(t)
	default:
		return errors.New("Unknown VectorMessageKind Variant Index while Decoding")
	}

	return nil
}

type ProofResponse struct {
	DataProof DataProof
	Message   prim.Option[AddressedMessage]
}

type DataProof struct {
	Roots          TxDataRoots
	Proof          []prim.H256
	NumberOfLeaves uint32 `scale:"compact"`
	LeafIndex      uint32 `scale:"compact"`
	Leaf           prim.H256
}

type TxDataRoots struct {
	DataRoot   prim.H256
	BlobRoot   prim.H256
	BridgeRoot prim.H256
}

type AddressedMessage struct {
	Message           VectorMessageKind
	From              prim.H256
	To                prim.H256
	OriginDomain      uint32 `scale:"compact"`
	DestinationDomain uint32 `scale:"compact"`
	Id                uint64 `scale:"compact"`
}

type InclusionFee struct {
	BaseFee           Balance
	LenFee            Balance
	AdjustedWeightFee Balance
}
type FeeDetails struct {
	InclusionFee prim.Option[InclusionFee]
}

type RuntimeDispatchInfo struct {
	Weight     Weight
	Class      DispatchClass
	PartialFee Balance
}

type ProxyType struct {
	VariantIndex uint8
}

func (this ProxyType) ToHuman() string {
	return this.ToString()
}

func (this ProxyType) ToString() string {
	switch this.VariantIndex {
	case 0:
		return "Any"
	case 1:
		return "NonTransfer"
	case 2:
		return "Governance"
	case 3:
		return "Staking"
	case 4:
		return "IdentityJudgement"
	case 5:
		return "NominationPools"
	default:
		panic("Unknown ProxyType Variant Index")
	}
}

func (this *ProxyType) EncodeTo(dest *string) {
	prim.Encoder.EncodeTo(this.VariantIndex, dest)
}

func (this *ProxyType) Decode(decoder *prim.Decoder) error {
	*this = ProxyType{}

	if err := decoder.Decode(&this.VariantIndex); err != nil {
		return err
	}

	switch this.VariantIndex {
	case 0:
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
	default:
		return errors.New("Unknown ProxyType Variant Index while Decoding")
	}

	return nil
}
