package staking

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageActiveEra struct {
	Index uint32
	Start prim.Option[uint64]
}

func (this *StorageActiveEra) PalletName() string {
	return PalletName
}

func (this *StorageActiveEra) StorageName() string {
	return "ActiveEra"
}

func (this *StorageActiveEra) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageActiveEra], error) {
	return GenericFetch[StorageActiveEra](blockStorage, this)
}

//
//
//

type StorageBondedKey = AccountId
type StorageBondedValue = AccountId
type StorageBondedEntry = StorageEntry[StorageBondedKey, StorageBondedValue]

type StorageBonded struct{}

func (this *StorageBonded) PalletName() string {
	return PalletName
}

func (this *StorageBonded) StorageName() string {
	return "Bonded"
}

func (this *StorageBonded) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageBonded) Fetch(blockStorage interfaces.BlockStorageT, key StorageBondedKey) (prim.Option[StorageBondedEntry], error) {
	return GenericMapFetch[StorageBondedValue](blockStorage, key, this)
}

func (this *StorageBonded) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageBondedEntry, error) {
	return GenericMapKeysFetch[StorageBondedValue, StorageBondedKey](blockStorage, this)
}

//
//
//

type BondedErasValue = []Tuple2[uint32, uint32]
type StorageBondedEras struct{}

func (this *StorageBondedEras) PalletName() string {
	return PalletName
}

func (this *StorageBondedEras) StorageName() string {
	return "BondedEras"
}

func (this *StorageBondedEras) Fetch(blockStorage interfaces.BlockStorageT) (BondedErasValue, error) {
	val, err := GenericFetch[BondedErasValue](blockStorage, this)
	return val.UnwrapOr(BondedErasValue{}), err
}

//
//
//

type StorageCanceledSlashPayoutValue = Balance
type StorageCanceledSlashPayout struct{}

func (this *StorageCanceledSlashPayout) PalletName() string {
	return PalletName
}

func (this *StorageCanceledSlashPayout) StorageName() string {
	return "CanceledSlashPayout"
}

func (this *StorageCanceledSlashPayout) Fetch(blockStorage interfaces.BlockStorageT) (StorageCanceledSlashPayoutValue, error) {
	return GenericFetchDefault[StorageCanceledSlashPayoutValue](blockStorage, this)
}

//
//
//

type StorageChillThresholdValue = uint8
type StorageChillThreshold struct{}

func (this *StorageChillThreshold) PalletName() string {
	return PalletName
}

func (this *StorageChillThreshold) StorageName() string {
	return "CanceledSlashPayout"
}

func (this *StorageChillThreshold) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageChillThresholdValue], error) {
	return GenericFetch[StorageChillThresholdValue](blockStorage, this)
}

//
//
//

type StorageClaimedRewardsKey1 = uint32
type StorageClaimedRewardsKey2 = AccountId
type StorageClaimedRewardsValue = []uint32
type StorageClaimedRewardsEntry = StorageEntryDoubleMap[StorageClaimedRewardsKey1, StorageClaimedRewardsKey2, StorageClaimedRewardsValue]
type StorageClaimedRewards struct{}

func (this *StorageClaimedRewards) PalletName() string {
	return PalletName
}

func (this *StorageClaimedRewards) StorageName() string {
	return "ClaimedRewards"
}

func (this *StorageClaimedRewards) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageClaimedRewards) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageClaimedRewards) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageClaimedRewardsKey1, key2 StorageClaimedRewardsKey2) (StorageClaimedRewardsEntry, error) {
	val, err := GenericDoubleMapFetch[StorageClaimedRewardsValue](blockStorage, key1, key2, this)
	return val.Unwrap(), err
}

func (this *StorageClaimedRewards) FetchAll(blockStorage interfaces.BlockStorageT, key StorageClaimedRewardsKey1) ([]StorageClaimedRewardsEntry, error) {
	return GenericDoubleMapKeysFetch[StorageClaimedRewardsValue, StorageClaimedRewardsKey1, StorageClaimedRewardsKey2](blockStorage, key, this)
}

//
//
//

type StorageCounterForNominatorsValue = uint32
type StorageCounterForNominators struct{}

func (this *StorageCounterForNominators) PalletName() string {
	return PalletName
}

func (this *StorageCounterForNominators) StorageName() string {
	return "CounterForNominators"
}

func (this *StorageCounterForNominators) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForNominatorsValue, error) {
	return GenericFetchDefault[StorageCounterForNominatorsValue](blockStorage, this)
}

//
//
//

type StorageCounterForValidatorsValue = uint32
type StorageCounterForValidators struct{}

func (this *StorageCounterForValidators) PalletName() string {
	return PalletName
}

func (this *StorageCounterForValidators) StorageName() string {
	return "CounterForValidators"
}

func (this *StorageCounterForValidators) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForValidatorsValue, error) {
	return GenericFetchDefault[StorageCounterForValidatorsValue](blockStorage, this)
}

//
//
//

type StorageCurrentEraValue = uint32
type StorageCurrentEra struct{}

func (this *StorageCurrentEra) PalletName() string {
	return PalletName
}

func (this *StorageCurrentEra) StorageName() string {
	return "CurrentEra"
}

func (this *StorageCurrentEra) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageCurrentEraValue], error) {
	return GenericFetch[StorageCurrentEraValue](blockStorage, this)
}

//
//
//

type StorageCurrentPlannedSessionValue = uint32
type StorageCurrentPlannedSession struct{}

func (this *StorageCurrentPlannedSession) PalletName() string {
	return PalletName
}

func (this *StorageCurrentPlannedSession) StorageName() string {
	return "CurrentPlannedSession"
}

func (this *StorageCurrentPlannedSession) Fetch(blockStorage interfaces.BlockStorageT) (StorageCurrentPlannedSessionValue, error) {
	return GenericFetchDefault[StorageCurrentPlannedSessionValue](blockStorage, this)
}

//
//
//

type StorageErasRewardPointsKey = uint32
type StorageErasRewardPointsEntry = StorageEntry[StorageErasRewardPointsKey, StorageErasRewardPoints]

type StorageErasRewardPoints struct {
	Total      uint32
	Individual []Tuple2[AccountId, uint32]
}

func (this *StorageErasRewardPoints) PalletName() string {
	return PalletName
}

func (this *StorageErasRewardPoints) StorageName() string {
	return "ErasRewardPoints"
}

func (this *StorageErasRewardPoints) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasRewardPoints) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasRewardPointsKey) (StorageErasRewardPointsEntry, error) {
	val, err := GenericMapFetch[StorageErasRewardPoints](blockStorage, key, this)
	return val.Unwrap(), err

}

func (this *StorageErasRewardPoints) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageErasRewardPointsEntry, error) {
	return GenericMapKeysFetch[StorageErasRewardPoints, StorageErasRewardPointsKey](blockStorage, this)
}

//
//
//

type StorageErasStakersKey1 = uint32
type StorageErasStakersKey2 = AccountId
type StorageErasStakersEntry = StorageEntryDoubleMap[StorageErasStakersKey1, StorageErasStakersKey2, StorageErasStakers]
type StorageErasStakers struct {
	AccountId AccountId
	Balance   Balance
}

func (this *StorageErasStakers) PalletName() string {
	return PalletName
}

func (this *StorageErasStakers) StorageName() string {
	return "ErasStakers"
}

func (this *StorageErasStakers) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasStakers) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasStakers) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageErasStakersKey1, key2 StorageErasStakersKey2) (StorageErasStakersEntry, error) {
	val, err := GenericDoubleMapFetch[StorageErasStakers](blockStorage, key1, key2, this)
	return val.Unwrap(), err
}

func (this *StorageErasStakers) FetchAll(blockStorage interfaces.BlockStorageT, key StorageErasStakersKey1) ([]StorageErasStakersEntry, error) {
	return GenericDoubleMapKeysFetch[StorageErasStakers, StorageErasStakersKey1, StorageErasStakersKey2](blockStorage, key, this)
}

//
//
//

type StorageErasStartSessionIndexKey = uint32
type StorageErasStartSessionIndexValue = uint32
type StorageErasStartSessionIndexEntry = StorageEntry[StorageErasStartSessionIndexKey, StorageErasStartSessionIndexValue]

type StorageErasStartSessionIndex struct{}

func (this *StorageErasStartSessionIndex) PalletName() string {
	return PalletName
}

func (this *StorageErasStartSessionIndex) StorageName() string {
	return "ErasStartSessionIndex"
}

func (this *StorageErasStartSessionIndex) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasStartSessionIndex) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasStartSessionIndexKey) (prim.Option[StorageErasStartSessionIndexEntry], error) {
	return GenericMapFetch[StorageErasStartSessionIndexValue](blockStorage, key, this)

}

func (this *StorageErasStartSessionIndex) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageErasStartSessionIndexEntry, error) {
	return GenericMapKeysFetch[StorageErasStartSessionIndexValue, StorageErasStartSessionIndexKey](blockStorage, this)
}

//
//
//

type StorageErasTotalStakeKey = uint32
type StorageErasTotalStakeValue = Balance
type StorageErasTotalStakeEntry = StorageEntry[StorageErasTotalStakeKey, StorageErasTotalStakeValue]

type StorageErasTotalStake struct{}

func (this *StorageErasTotalStake) PalletName() string {
	return PalletName
}

func (this *StorageErasTotalStake) StorageName() string {
	return "ErasTotalStake"
}

func (this *StorageErasTotalStake) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasTotalStake) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasTotalStakeKey) (StorageErasTotalStakeEntry, error) {
	val, err := GenericMapFetch[StorageErasTotalStakeValue](blockStorage, key, this)
	return val.Unwrap(), err

}

func (this *StorageErasTotalStake) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageErasTotalStakeEntry, error) {
	return GenericMapKeysFetch[StorageErasTotalStakeValue, StorageErasTotalStakeKey](blockStorage, this)
}

//
//
//

type StorageErasValidatorPrefsKey1 = uint32
type StorageErasValidatorPrefsKey2 = AccountId
type StorageErasValidatorPrefsValue = ValidatorPrefs
type StorageErasValidatorPrefsEntry = StorageEntryDoubleMap[StorageErasValidatorPrefsKey1, StorageErasValidatorPrefsKey2, StorageErasValidatorPrefsValue]
type StorageErasValidatorPrefs struct{}

func (this *StorageErasValidatorPrefs) PalletName() string {
	return PalletName
}

func (this *StorageErasValidatorPrefs) StorageName() string {
	return "ErasValidatorPrefs"
}

func (this *StorageErasValidatorPrefs) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasValidatorPrefs) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasValidatorPrefs) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageErasValidatorPrefsKey1, key2 StorageErasValidatorPrefsKey2) (StorageErasValidatorPrefsEntry, error) {
	val, err := GenericDoubleMapFetch[StorageErasValidatorPrefsValue](blockStorage, key1, key2, this)
	if err != nil {
		return StorageErasValidatorPrefsEntry{}, err
	}

	return val.Unwrap(), nil
}

func (this *StorageErasValidatorPrefs) FetchAll(blockStorage interfaces.BlockStorageT, key StorageErasValidatorPrefsKey1) ([]StorageErasValidatorPrefsEntry, error) {
	return GenericDoubleMapKeysFetch[StorageErasValidatorPrefsValue, StorageErasValidatorPrefsKey1, StorageErasValidatorPrefsKey2](blockStorage, key, this)
}

//
//
//

type StorageErasValidatorRewardKey = uint32
type StorageErasValidatorRewardValue = Balance
type StorageErasValidatorRewardEntry = StorageEntry[StorageErasValidatorRewardKey, StorageErasValidatorRewardValue]

type StorageErasValidatorReward struct{}

func (this *StorageErasValidatorReward) PalletName() string {
	return PalletName
}

func (this *StorageErasValidatorReward) StorageName() string {
	return "ErasValidatorReward"
}

func (this *StorageErasValidatorReward) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasValidatorReward) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasValidatorRewardKey) (prim.Option[StorageErasValidatorRewardEntry], error) {
	return GenericMapFetch[StorageErasValidatorRewardValue](blockStorage, key, this)

}

func (this *StorageErasValidatorReward) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageErasValidatorRewardEntry, error) {
	return GenericMapKeysFetch[StorageErasValidatorRewardValue, StorageErasValidatorRewardKey](blockStorage, this)
}

//
//
//

type StorageInvulnerablesValue = []AccountId
type StorageInvulnerables struct{}

func (this *StorageInvulnerables) PalletName() string {
	return PalletName
}

func (this *StorageInvulnerables) StorageName() string {
	return "Invulnerables"
}

func (this *StorageInvulnerables) Fetch(blockStorage interfaces.BlockStorageT) (StorageInvulnerablesValue, error) {
	val, err := GenericFetch[StorageInvulnerablesValue](blockStorage, this)
	return val.UnwrapOr(StorageInvulnerablesValue{}), err
}

//
//
//

type StorageLedgerKey = AccountId
type StorageLedgerEntry = StorageEntry[StorageLedgerKey, StorageLedger]

type StorageLedger struct {
	Stash                AccountId
	Total                Balance `scale:"compact"`
	Active               Balance `scale:"compact"`
	Unlocking            []UnlockChunk
	LegacyClaimedRewards []uint32
}

type UnlockChunk struct {
	Value Balance `scale:"compact"`
	Era   uint32  `scale:"compact"`
}

func (this *StorageLedger) PalletName() string {
	return PalletName
}

func (this *StorageLedger) StorageName() string {
	return "Ledger"
}

func (this *StorageLedger) MapKeyHasher() uint8 {
	return Blake2_128ConcatHasher
}

func (this *StorageLedger) Fetch(blockStorage interfaces.BlockStorageT, key StorageLedgerKey) (prim.Option[StorageLedgerEntry], error) {
	return GenericMapFetch[StorageLedger](blockStorage, key, this)

}

func (this *StorageLedger) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageLedgerEntry, error) {
	return GenericMapKeysFetch[StorageLedger, StorageLedgerKey](blockStorage, this)
}

//
//
//

type StorageMaxNominatorsCountValue = uint32
type StorageMaxNominatorsCount struct{}

func (this *StorageMaxNominatorsCount) PalletName() string {
	return PalletName
}

func (this *StorageMaxNominatorsCount) StorageName() string {
	return "MaxNominatorsCount"
}

func (this *StorageMaxNominatorsCount) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxNominatorsCountValue], error) {
	return GenericFetch[StorageMaxNominatorsCountValue](blockStorage, this)
}

//
//
//

type StorageMaxValidatorsCountValue = uint32
type StorageMaxValidatorsCount struct{}

func (this *StorageMaxValidatorsCount) PalletName() string {
	return PalletName
}

func (this *StorageMaxValidatorsCount) StorageName() string {
	return "MaxValidatorsCount"
}

func (this *StorageMaxValidatorsCount) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxValidatorsCountValue], error) {
	return GenericFetch[StorageMaxValidatorsCountValue](blockStorage, this)
}

//
//
//

type StorageMinCommissionValue = Perbill
type StorageMinCommission struct{}

func (this *StorageMinCommission) PalletName() string {
	return PalletName
}

func (this *StorageMinCommission) StorageName() string {
	return "MinCommission"
}

func (this *StorageMinCommission) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinCommissionValue, error) {
	return GenericFetchDefault[StorageMinCommissionValue](blockStorage, this)
}

//
//
//

type StorageMinNominatorBondValue = Balance
type StorageMinNominatorBond struct{}

func (this *StorageMinNominatorBond) PalletName() string {
	return PalletName
}

func (this *StorageMinNominatorBond) StorageName() string {
	return "MinNominatorBond"
}

func (this *StorageMinNominatorBond) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinNominatorBondValue, error) {
	return GenericFetchDefault[StorageMinNominatorBondValue](blockStorage, this)
}

//
//
//

type StorageMinValidatorBondValue = Balance
type StorageMinValidatorBond struct{}

func (this *StorageMinValidatorBond) PalletName() string {
	return PalletName
}

func (this *StorageMinValidatorBond) StorageName() string {
	return "MinValidatorBond"
}

func (this *StorageMinValidatorBond) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinValidatorBondValue, error) {
	return GenericFetchDefault[StorageMinValidatorBondValue](blockStorage, this)
}

//
//
//

type StorageMinimumActiveStakeValue = Balance
type StorageMinimumActiveStake struct{}

func (this *StorageMinimumActiveStake) PalletName() string {
	return PalletName
}

func (this *StorageMinimumActiveStake) StorageName() string {
	return "MinimumActiveStake"
}

func (this *StorageMinimumActiveStake) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinimumActiveStakeValue, error) {
	return GenericFetchDefault[StorageMinimumActiveStakeValue](blockStorage, this)
}

//
//
//

type StorageMinimumValidatorCountValue = uint32
type StorageMinimumValidatorCount struct{}

func (this *StorageMinimumValidatorCount) PalletName() string {
	return PalletName
}

func (this *StorageMinimumValidatorCount) StorageName() string {
	return "MinimumValidatorCount"
}

func (this *StorageMinimumValidatorCount) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinimumValidatorCountValue, error) {
	return GenericFetchDefault[StorageMinimumValidatorCountValue](blockStorage, this)
}

//
//
//

type StorageNominatorSlashInEraKey1 = uint32
type StorageNominatorSlashInEraKey2 = AccountId
type StorageNominatorSlashInEraValue = Balance
type StorageNominatorSlashInEraEntry = StorageEntryDoubleMap[StorageNominatorSlashInEraKey1, StorageNominatorSlashInEraKey2, StorageNominatorSlashInEraValue]
type StorageNominatorSlashInEra struct{}

func (this *StorageNominatorSlashInEra) PalletName() string {
	return PalletName
}

func (this *StorageNominatorSlashInEra) StorageName() string {
	return "NominatorSlashInEra"
}

func (this *StorageNominatorSlashInEra) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageNominatorSlashInEra) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageNominatorSlashInEra) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageNominatorSlashInEraKey1, key2 StorageNominatorSlashInEraKey2) (prim.Option[StorageNominatorSlashInEraEntry], error) {
	return GenericDoubleMapFetch[StorageNominatorSlashInEraValue](blockStorage, key1, key2, this)
}

func (this *StorageNominatorSlashInEra) FetchAll(blockStorage interfaces.BlockStorageT, key StorageNominatorSlashInEraKey1) ([]StorageNominatorSlashInEraEntry, error) {
	return GenericDoubleMapKeysFetch[StorageNominatorSlashInEraValue, StorageNominatorSlashInEraKey1, StorageNominatorSlashInEraKey2](blockStorage, key, this)
}

//
//
//

type StorageNominatorsKey = AccountId
type StorageNominatorsEntry = StorageEntry[StorageNominatorsKey, StorageNominators]

type StorageNominators struct {
	Targets     []AccountId
	SubmittedIn uint32
	Suppressed  bool
}

func (this *StorageNominators) PalletName() string {
	return PalletName
}

func (this *StorageNominators) StorageName() string {
	return "Nominators"
}

func (this *StorageNominators) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageNominators) Fetch(blockStorage interfaces.BlockStorageT, key StorageNominatorsKey) (prim.Option[StorageNominatorsEntry], error) {
	return GenericMapFetch[StorageNominators](blockStorage, key, this)

}

func (this *StorageNominators) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageNominatorsEntry, error) {
	return GenericMapKeysFetch[StorageNominators, StorageNominatorsKey](blockStorage, this)
}

//
//
//

type StorageOffendingValidatorsValue = []Tuple2[uint32, bool]
type StorageOffendingValidators struct{}

func (this *StorageOffendingValidators) PalletName() string {
	return PalletName
}

func (this *StorageOffendingValidators) StorageName() string {
	return "OffendingValidators"
}

func (this *StorageOffendingValidators) Fetch(blockStorage interfaces.BlockStorageT) (StorageOffendingValidatorsValue, error) {
	val, err := GenericFetch[StorageOffendingValidatorsValue](blockStorage, this)
	return val.UnwrapOr(StorageOffendingValidatorsValue{}), err
}

//
//
//

type StoragePayeeKey = AccountId
type StoragePayeeValue = RewardDestination
type StoragePayeeEntry = StorageEntry[StoragePayeeKey, StoragePayeeValue]

type StoragePayee struct{}

func (this *StoragePayee) PalletName() string {
	return PalletName
}

func (this *StoragePayee) StorageName() string {
	return "Payee"
}

func (this *StoragePayee) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StoragePayee) Fetch(blockStorage interfaces.BlockStorageT, key StoragePayeeKey) (prim.Option[StoragePayeeEntry], error) {
	return GenericMapFetch[StoragePayeeValue](blockStorage, key, this)

}

func (this *StoragePayee) FetchAll(blockStorage interfaces.BlockStorageT) ([]StoragePayeeEntry, error) {
	return GenericMapKeysFetch[StoragePayeeValue, StoragePayeeKey](blockStorage, this)
}

//
//
//

type StorageSlashRewardFractionValue = Perbill
type StorageSlashRewardFraction struct{}

func (this *StorageSlashRewardFraction) PalletName() string {
	return PalletName
}

func (this *StorageSlashRewardFraction) StorageName() string {
	return "SlashRewardFraction "
}

func (this *StorageSlashRewardFraction) Fetch(blockStorage interfaces.BlockStorageT) (StorageSlashRewardFractionValue, error) {
	return GenericFetchDefault[StorageSlashRewardFractionValue](blockStorage, this)
}

//
//
//

type StorageSlashingSpansKey = AccountId
type StorageSlashingSpansEntry = StorageEntry[StoragePayeeKey, StorageSlashingSpans]

type StorageSlashingSpans struct {
	SpanIndex        uint32
	LastStart        uint32
	LastNonZeroSlash uint32
	Prior            []uint32
}

func (this *StorageSlashingSpans) PalletName() string {
	return PalletName
}

func (this *StorageSlashingSpans) StorageName() string {
	return "SlashingSpans"
}

func (this *StorageSlashingSpans) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageSlashingSpans) Fetch(blockStorage interfaces.BlockStorageT, key StorageSlashingSpansKey) (prim.Option[StorageSlashingSpansEntry], error) {
	return GenericMapFetch[StorageSlashingSpans](blockStorage, key, this)

}

func (this *StorageSlashingSpans) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageSlashingSpansEntry, error) {
	return GenericMapKeysFetch[StorageSlashingSpans, StorageSlashingSpansKey](blockStorage, this)
}

//
//
//

type StorageSpanSlashKey = Tuple2[AccountId, uint32]
type StorageSpanSlashEntry = StorageEntry[StorageSpanSlashKey, StorageSpanSlash]

type StorageSpanSlash struct {
	Slashed Balance
	PaidOut Balance
}

func (this *StorageSpanSlash) PalletName() string {
	return PalletName
}

func (this *StorageSpanSlash) StorageName() string {
	return "SpanSlash"
}

func (this *StorageSpanSlash) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageSpanSlash) Fetch(blockStorage interfaces.BlockStorageT, key StorageSpanSlashKey) (prim.Option[StorageSpanSlashEntry], error) {
	return GenericMapFetch[StorageSpanSlash](blockStorage, key, this)

}

func (this *StorageSpanSlash) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageSpanSlashEntry, error) {
	return GenericMapKeysFetch[StorageSpanSlash, StorageSpanSlashKey](blockStorage, this)
}

//
//
//

type StorageUnappliedSlashesKey = uint32
type StorageUnappliedSlashesValue = []StorageUnappliedSlashes
type StorageUnappliedSlashesEntry = StorageEntry[StorageUnappliedSlashesKey, StorageUnappliedSlashesValue]

type StorageUnappliedSlashes struct {
	Validator AccountId
	Own       Balance
	Others    []Tuple2[AccountId, Balance]
	Reporters []AccountId
	Payout    Balance
}

func (this *StorageUnappliedSlashes) PalletName() string {
	return PalletName
}

func (this *StorageUnappliedSlashes) StorageName() string {
	return "UnappliedSlashes"
}

func (this *StorageUnappliedSlashes) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageUnappliedSlashes) Fetch(blockStorage interfaces.BlockStorageT, key StorageUnappliedSlashesKey) (StorageUnappliedSlashesEntry, error) {
	val, err := GenericMapFetch[StorageUnappliedSlashesValue](blockStorage, key, this)
	return val.UnwrapOr(StorageUnappliedSlashesEntry{}), err

}

func (this *StorageUnappliedSlashes) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageUnappliedSlashesEntry, error) {
	return GenericMapKeysFetch[StorageUnappliedSlashesValue, StorageUnappliedSlashesKey](blockStorage, this)
}

//
//
//

type StorageValidatorCountValue = uint32
type StorageValidatorCount struct{}

func (this *StorageValidatorCount) PalletName() string {
	return PalletName
}

func (this *StorageValidatorCount) StorageName() string {
	return "ValidatorCount"
}

func (this *StorageValidatorCount) Fetch(blockStorage interfaces.BlockStorageT) (StorageValidatorCountValue, error) {
	return GenericFetchDefault[StorageValidatorCountValue](blockStorage, this)
}

//
//
//

type StorageValidatorSlashInEraKey1 = uint32
type StorageValidatorSlashInEraKey2 = AccountId
type StorageValidatorSlashInEraValue = Tuple2[Perbill, Balance]
type StorageValidatorSlashInEraEntry = StorageEntryDoubleMap[StorageValidatorSlashInEraKey1, StorageValidatorSlashInEraKey2, StorageValidatorSlashInEraValue]
type StorageValidatorSlashInEra struct{}

func (this *StorageValidatorSlashInEra) PalletName() string {
	return PalletName
}

func (this *StorageValidatorSlashInEra) StorageName() string {
	return "ValidatorSlashInEra"
}

func (this *StorageValidatorSlashInEra) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageValidatorSlashInEra) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageValidatorSlashInEra) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageValidatorSlashInEraKey1, key2 StorageValidatorSlashInEraKey2) (prim.Option[StorageValidatorSlashInEraEntry], error) {
	return GenericDoubleMapFetch[StorageValidatorSlashInEraValue](blockStorage, key1, key2, this)
}

func (this *StorageValidatorSlashInEra) FetchAll(blockStorage interfaces.BlockStorageT, key StorageValidatorSlashInEraKey1) ([]StorageValidatorSlashInEraEntry, error) {
	return GenericDoubleMapKeysFetch[StorageValidatorSlashInEraValue, StorageValidatorSlashInEraKey1, StorageValidatorSlashInEraKey2](blockStorage, key, this)
}

//
//
//

type StorageValidatorsKey = AccountId
type StorageValidatorsValue = ValidatorPrefs
type StorageValidatorsEntry = StorageEntry[StorageValidatorsKey, StorageValidatorsValue]

type StorageValidators struct {
	Slashed Balance
	PaidOut Balance
}

func (this *StorageValidators) PalletName() string {
	return PalletName
}

func (this *StorageValidators) StorageName() string {
	return "Validators"
}

func (this *StorageValidators) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageValidators) Fetch(blockStorage interfaces.BlockStorageT, key StorageValidatorsKey) (StorageValidatorsEntry, error) {
	val, err := GenericMapFetch[StorageValidatorsValue](blockStorage, key, this)
	return val.Unwrap(), err

}

func (this *StorageValidators) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageValidatorsEntry, error) {
	return GenericMapKeysFetch[StorageValidatorsValue, StorageValidatorsKey](blockStorage, this)
}
