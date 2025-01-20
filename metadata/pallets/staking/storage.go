package staking

import (
	"go-sdk/interfaces"
	. "go-sdk/metadata"
	prim "go-sdk/primitives"
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
	return "Staking"
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

type BondedEra struct {
	Tup0 uint32
	Tup1 uint32
}

type BondedErasValue = []BondedEra
type StorageBondedEras struct{}

func (this *StorageBondedEras) PalletName() string {
	return PalletName
}

func (this *StorageBondedEras) StorageName() string {
	return "BondedEras"
}

func (this *StorageBondedEras) Fetch(blockStorage interfaces.BlockStorageT) (BondedErasValue, error) {
	val, err := GenericFetch[BondedErasValue](blockStorage, this)
	if err != nil {
		return nil, err
	}

	if val.IsNone() {
		return BondedErasValue{}, nil
	}

	return val.Unwrap(), nil
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
	if err != nil {
		return StorageClaimedRewardsEntry{}, err
	}

	return val.Unwrap(), nil
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
	Individual []EraRewardPointsIndividual
}

type EraRewardPointsIndividual struct {
	Tup0 AccountId
	Tup1 uint32
}

func (this *StorageErasRewardPoints) PalletName() string {
	return "Staking"
}

func (this *StorageErasRewardPoints) StorageName() string {
	return "ErasRewardPoints"
}

func (this *StorageErasRewardPoints) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasRewardPoints) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasRewardPointsKey) (StorageErasRewardPointsEntry, error) {
	val, err := GenericMapFetch[StorageErasRewardPoints](blockStorage, key, this)
	if err != nil {
		return StorageErasRewardPointsEntry{}, nil
	}
	return val.Unwrap(), nil

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
	if err != nil {
		return StorageErasStakersEntry{}, err
	}

	return val.Unwrap(), nil
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
	return "Staking"
}

func (this *StorageErasStartSessionIndex) StorageName() string {
	return "ErasStartSessionIndex"
}

func (this *StorageErasStartSessionIndex) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasStartSessionIndex) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasRewardPointsKey) (prim.Option[StorageErasStartSessionIndexEntry], error) {
	val, err := GenericMapFetch[StorageErasStartSessionIndexValue](blockStorage, key, this)
	if err != nil {
		return prim.NewNone[StorageErasStartSessionIndexEntry](), err
	}
	if val.IsNone() {
		return prim.NewNone[StorageErasStartSessionIndexEntry](), nil
	}

	return prim.NewSome(val.Unwrap()), nil

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
	return "Staking"
}

func (this *StorageErasTotalStake) StorageName() string {
	return "ErasTotalStake"
}

func (this *StorageErasTotalStake) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasTotalStake) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasRewardPointsKey) (StorageErasTotalStakeEntry, error) {
	val, err := GenericMapFetch[StorageErasTotalStakeValue](blockStorage, key, this)
	if err != nil {
		return StorageErasTotalStakeEntry{}, err
	}

	return val.Unwrap(), nil

}

func (this *StorageErasTotalStake) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageErasTotalStakeEntry, error) {
	return GenericMapKeysFetch[StorageErasTotalStakeValue, StorageErasTotalStakeKey](blockStorage, this)
}

//
//
//

type StorageErasValidatorPrefsKey1 = uint32
type StorageErasValidatorPrefsKey2 = AccountId
type StorageErasValidatorPrefsEntry = StorageEntryDoubleMap[StorageErasValidatorPrefsKey1, StorageErasValidatorPrefsKey2, StorageErasValidatorPrefs]
type StorageErasValidatorPrefs struct {
	Commission Perbill
	Blocked    bool
}

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

func (this *StorageErasValidatorPrefs) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageErasStakersKey1, key2 StorageErasStakersKey2) (StorageErasValidatorPrefsEntry, error) {
	val, err := GenericDoubleMapFetch[StorageErasValidatorPrefs](blockStorage, key1, key2, this)
	if err != nil {
		return StorageErasValidatorPrefsEntry{}, err
	}

	return val.Unwrap(), nil
}

func (this *StorageErasValidatorPrefs) FetchAll(blockStorage interfaces.BlockStorageT, key StorageErasStakersKey1) ([]StorageErasValidatorPrefsEntry, error) {
	return GenericDoubleMapKeysFetch[StorageErasValidatorPrefs, StorageErasValidatorPrefsKey1, StorageErasValidatorPrefsKey2](blockStorage, key, this)
}

//
//
//

type StorageErasValidatorRewardKey = uint32
type StorageErasValidatorRewardValue = Balance
type StorageErasValidatorRewardEntry = StorageEntry[StorageErasValidatorRewardKey, StorageErasValidatorRewardValue]

type StorageErasValidatorReward struct{}

func (this *StorageErasValidatorReward) PalletName() string {
	return "Staking"
}

func (this *StorageErasValidatorReward) StorageName() string {
	return "ErasValidatorReward"
}

func (this *StorageErasValidatorReward) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageErasValidatorReward) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasRewardPointsKey) (prim.Option[StorageErasValidatorRewardEntry], error) {
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
	if err != nil {
		return StorageInvulnerablesValue{}, err
	}

	return val.Unwrap(), nil
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
	return "Staking"
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
