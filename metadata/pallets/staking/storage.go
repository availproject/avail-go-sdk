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

func (sae *StorageActiveEra) PalletName() string {
	return PalletName
}

func (sae *StorageActiveEra) StorageName() string {
	return "ActiveEra"
}

func (sae *StorageActiveEra) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageActiveEra], error) {
	return GenericFetch[StorageActiveEra](blockStorage, sae)
}

//
//
//

type StorageBondedKey = prim.AccountId
type StorageBondedValue = prim.AccountId
type StorageBondedEntry = StorageEntry[StorageBondedKey, StorageBondedValue]

type StorageBonded struct{}

func (sb *StorageBonded) PalletName() string {
	return PalletName
}

func (sb *StorageBonded) StorageName() string {
	return "Bonded"
}

func (sb *StorageBonded) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sb *StorageBonded) Fetch(blockStorage interfaces.BlockStorageT, key StorageBondedKey) (prim.Option[StorageBondedEntry], error) {
	return GenericMapFetch[StorageBondedValue](blockStorage, key, sb)
}

func (sb *StorageBonded) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageBondedEntry, error) {
	return GenericMapKeysFetch[StorageBondedValue, StorageBondedKey](blockStorage, sb)
}

//
//
//

type BondedErasValue = []Tuple2[uint32, uint32]
type StorageBondedEras struct{}

func (sbe *StorageBondedEras) PalletName() string {
	return PalletName
}

func (sbe *StorageBondedEras) StorageName() string {
	return "BondedEras"
}

func (sbe *StorageBondedEras) Fetch(blockStorage interfaces.BlockStorageT) (BondedErasValue, error) {
	val, err := GenericFetch[BondedErasValue](blockStorage, sbe)
	return val.UnwrapOr(BondedErasValue{}), err
}

//
//
//

type StorageCanceledSlashPayoutValue = Balance
type StorageCanceledSlashPayout struct{}

func (scsp *StorageCanceledSlashPayout) PalletName() string {
	return PalletName
}

func (scsp *StorageCanceledSlashPayout) StorageName() string {
	return "CanceledSlashPayout"
}

func (scsp *StorageCanceledSlashPayout) Fetch(blockStorage interfaces.BlockStorageT) (StorageCanceledSlashPayoutValue, error) {
	return GenericFetchDefault[StorageCanceledSlashPayoutValue](blockStorage, scsp)
}

//
//
//

type StorageChillThresholdValue = uint8
type StorageChillThreshold struct{}

func (sct *StorageChillThreshold) PalletName() string {
	return PalletName
}

func (sct *StorageChillThreshold) StorageName() string {
	return "CanceledSlashPayout"
}

func (sct *StorageChillThreshold) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageChillThresholdValue], error) {
	return GenericFetch[StorageChillThresholdValue](blockStorage, sct)
}

//
//
//

type StorageClaimedRewardsKey1 = uint32
type StorageClaimedRewardsKey2 = prim.AccountId
type StorageClaimedRewardsValue = []uint32
type StorageClaimedRewardsEntry = StorageEntryDoubleMap[StorageClaimedRewardsKey1, StorageClaimedRewardsKey2, StorageClaimedRewardsValue]
type StorageClaimedRewards struct{}

func (scr *StorageClaimedRewards) PalletName() string {
	return PalletName
}

func (scr *StorageClaimedRewards) StorageName() string {
	return "ClaimedRewards"
}

func (scr *StorageClaimedRewards) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (scr *StorageClaimedRewards) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (scr *StorageClaimedRewards) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageClaimedRewardsKey1, key2 StorageClaimedRewardsKey2) (StorageClaimedRewardsEntry, error) {
	val, err := GenericDoubleMapFetch[StorageClaimedRewardsValue](blockStorage, key1, key2, scr)
	return val.Unwrap(), err
}

func (scr *StorageClaimedRewards) FetchAll(blockStorage interfaces.BlockStorageT, key StorageClaimedRewardsKey1) ([]StorageClaimedRewardsEntry, error) {
	return GenericDoubleMapKeysFetch[StorageClaimedRewardsValue, StorageClaimedRewardsKey1, StorageClaimedRewardsKey2](blockStorage, key, scr)
}

//
//
//

type StorageCounterForNominatorsValue = uint32
type StorageCounterForNominators struct{}

func (scfn *StorageCounterForNominators) PalletName() string {
	return PalletName
}

func (scfn *StorageCounterForNominators) StorageName() string {
	return "CounterForNominators"
}

func (scfn *StorageCounterForNominators) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForNominatorsValue, error) {
	return GenericFetchDefault[StorageCounterForNominatorsValue](blockStorage, scfn)
}

//
//
//

type StorageCounterForValidatorsValue = uint32
type StorageCounterForValidators struct{}

func (scfv *StorageCounterForValidators) PalletName() string {
	return PalletName
}

func (scfv *StorageCounterForValidators) StorageName() string {
	return "CounterForValidators"
}

func (scfv *StorageCounterForValidators) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForValidatorsValue, error) {
	return GenericFetchDefault[StorageCounterForValidatorsValue](blockStorage, scfv)
}

//
//
//

type StorageCurrentEraValue = uint32
type StorageCurrentEra struct{}

func (sce *StorageCurrentEra) PalletName() string {
	return PalletName
}

func (sce *StorageCurrentEra) StorageName() string {
	return "CurrentEra"
}

func (sce *StorageCurrentEra) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageCurrentEraValue], error) {
	return GenericFetch[StorageCurrentEraValue](blockStorage, sce)
}

//
//
//

type StorageCurrentPlannedSessionValue = uint32
type StorageCurrentPlannedSession struct{}

func (scps *StorageCurrentPlannedSession) PalletName() string {
	return PalletName
}

func (scps *StorageCurrentPlannedSession) StorageName() string {
	return "CurrentPlannedSession"
}

func (scps *StorageCurrentPlannedSession) Fetch(blockStorage interfaces.BlockStorageT) (StorageCurrentPlannedSessionValue, error) {
	return GenericFetchDefault[StorageCurrentPlannedSessionValue](blockStorage, scps)
}

//
//
//

type StorageErasRewardPointsKey = uint32
type StorageErasRewardPointsEntry = StorageEntry[StorageErasRewardPointsKey, StorageErasRewardPoints]

type StorageErasRewardPoints struct {
	Total      uint32
	Individual []Tuple2[prim.AccountId, uint32]
}

func (serp *StorageErasRewardPoints) PalletName() string {
	return PalletName
}

func (serp *StorageErasRewardPoints) StorageName() string {
	return "ErasRewardPoints"
}

func (serp *StorageErasRewardPoints) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (serp *StorageErasRewardPoints) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasRewardPointsKey) (StorageErasRewardPointsEntry, error) {
	val, err := GenericMapFetch[StorageErasRewardPoints](blockStorage, key, serp)
	return val.Unwrap(), err

}

func (serp *StorageErasRewardPoints) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageErasRewardPointsEntry, error) {
	return GenericMapKeysFetch[StorageErasRewardPoints, StorageErasRewardPointsKey](blockStorage, serp)
}

//
//
//

type StorageErasStakersKey1 = uint32
type StorageErasStakersKey2 = prim.AccountId
type StorageErasStakersEntry = StorageEntryDoubleMap[StorageErasStakersKey1, StorageErasStakersKey2, StorageErasStakers]
type StorageErasStakers struct {
	AccountId prim.AccountId
	Balance   Balance
}

func (ses *StorageErasStakers) PalletName() string {
	return PalletName
}

func (ses *StorageErasStakers) StorageName() string {
	return "ErasStakers"
}

func (ses *StorageErasStakers) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (ses *StorageErasStakers) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (ses *StorageErasStakers) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageErasStakersKey1, key2 StorageErasStakersKey2) (StorageErasStakersEntry, error) {
	val, err := GenericDoubleMapFetch[StorageErasStakers](blockStorage, key1, key2, ses)
	return val.Unwrap(), err
}

func (ses *StorageErasStakers) FetchAll(blockStorage interfaces.BlockStorageT, key StorageErasStakersKey1) ([]StorageErasStakersEntry, error) {
	return GenericDoubleMapKeysFetch[StorageErasStakers, StorageErasStakersKey1, StorageErasStakersKey2](blockStorage, key, ses)
}

//
//
//

type StorageErasStartSessionIndexKey = uint32
type StorageErasStartSessionIndexValue = uint32
type StorageErasStartSessionIndexEntry = StorageEntry[StorageErasStartSessionIndexKey, StorageErasStartSessionIndexValue]

type StorageErasStartSessionIndex struct{}

func (sessi *StorageErasStartSessionIndex) PalletName() string {
	return PalletName
}

func (sessi *StorageErasStartSessionIndex) StorageName() string {
	return "ErasStartSessionIndex"
}

func (sessi *StorageErasStartSessionIndex) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sessi *StorageErasStartSessionIndex) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasStartSessionIndexKey) (prim.Option[StorageErasStartSessionIndexEntry], error) {
	return GenericMapFetch[StorageErasStartSessionIndexValue](blockStorage, key, sessi)

}

func (sessi *StorageErasStartSessionIndex) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageErasStartSessionIndexEntry, error) {
	return GenericMapKeysFetch[StorageErasStartSessionIndexValue, StorageErasStartSessionIndexKey](blockStorage, sessi)
}

//
//
//

type StorageErasTotalStakeKey = uint32
type StorageErasTotalStakeValue = Balance
type StorageErasTotalStakeEntry = StorageEntry[StorageErasTotalStakeKey, StorageErasTotalStakeValue]

type StorageErasTotalStake struct{}

func (sets *StorageErasTotalStake) PalletName() string {
	return PalletName
}

func (sets *StorageErasTotalStake) StorageName() string {
	return "ErasTotalStake"
}

func (sets *StorageErasTotalStake) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sets *StorageErasTotalStake) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasTotalStakeKey) (StorageErasTotalStakeEntry, error) {
	val, err := GenericMapFetch[StorageErasTotalStakeValue](blockStorage, key, sets)
	return val.Unwrap(), err

}

func (sets *StorageErasTotalStake) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageErasTotalStakeEntry, error) {
	return GenericMapKeysFetch[StorageErasTotalStakeValue, StorageErasTotalStakeKey](blockStorage, sets)
}

//
//
//

type StorageErasValidatorPrefsKey1 = uint32
type StorageErasValidatorPrefsKey2 = prim.AccountId
type StorageErasValidatorPrefsValue = ValidatorPrefs
type StorageErasValidatorPrefsEntry = StorageEntryDoubleMap[StorageErasValidatorPrefsKey1, StorageErasValidatorPrefsKey2, StorageErasValidatorPrefsValue]
type StorageErasValidatorPrefs struct{}

func (sevp *StorageErasValidatorPrefs) PalletName() string {
	return PalletName
}

func (sevp *StorageErasValidatorPrefs) StorageName() string {
	return "ErasValidatorPrefs"
}

func (sevp *StorageErasValidatorPrefs) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (sevp *StorageErasValidatorPrefs) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (sevp *StorageErasValidatorPrefs) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageErasValidatorPrefsKey1, key2 StorageErasValidatorPrefsKey2) (StorageErasValidatorPrefsEntry, error) {
	val, err := GenericDoubleMapFetch[StorageErasValidatorPrefsValue](blockStorage, key1, key2, sevp)
	if err != nil {
		return StorageErasValidatorPrefsEntry{}, err
	}

	return val.Unwrap(), nil
}

func (sevp *StorageErasValidatorPrefs) FetchAll(blockStorage interfaces.BlockStorageT, key StorageErasValidatorPrefsKey1) ([]StorageErasValidatorPrefsEntry, error) {
	return GenericDoubleMapKeysFetch[StorageErasValidatorPrefsValue, StorageErasValidatorPrefsKey1, StorageErasValidatorPrefsKey2](blockStorage, key, sevp)
}

//
//
//

type StorageErasValidatorRewardKey = uint32
type StorageErasValidatorRewardValue = Balance
type StorageErasValidatorRewardEntry = StorageEntry[StorageErasValidatorRewardKey, StorageErasValidatorRewardValue]

type StorageErasValidatorReward struct{}

func (sevr *StorageErasValidatorReward) PalletName() string {
	return PalletName
}

func (sevr *StorageErasValidatorReward) StorageName() string {
	return "ErasValidatorReward"
}

func (sevr *StorageErasValidatorReward) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sevr *StorageErasValidatorReward) Fetch(blockStorage interfaces.BlockStorageT, key StorageErasValidatorRewardKey) (prim.Option[StorageErasValidatorRewardEntry], error) {
	return GenericMapFetch[StorageErasValidatorRewardValue](blockStorage, key, sevr)

}

func (sevr *StorageErasValidatorReward) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageErasValidatorRewardEntry, error) {
	return GenericMapKeysFetch[StorageErasValidatorRewardValue, StorageErasValidatorRewardKey](blockStorage, sevr)
}

//
//
//

type StorageInvulnerablesValue = []prim.AccountId
type StorageInvulnerables struct{}

func (si *StorageInvulnerables) PalletName() string {
	return PalletName
}

func (si *StorageInvulnerables) StorageName() string {
	return "Invulnerables"
}

func (si *StorageInvulnerables) Fetch(blockStorage interfaces.BlockStorageT) (StorageInvulnerablesValue, error) {
	val, err := GenericFetch[StorageInvulnerablesValue](blockStorage, si)
	return val.UnwrapOr(StorageInvulnerablesValue{}), err
}

//
//
//

type StorageLedgerKey = prim.AccountId
type StorageLedgerEntry = StorageEntry[StorageLedgerKey, StorageLedger]

type StorageLedger struct {
	Stash                prim.AccountId
	Total                Balance `scale:"compact"`
	Active               Balance `scale:"compact"`
	Unlocking            []UnlockChunk
	LegacyClaimedRewards []uint32
}

type UnlockChunk struct {
	Value Balance `scale:"compact"`
	Era   uint32  `scale:"compact"`
}

func (sl *StorageLedger) PalletName() string {
	return PalletName
}

func (sl *StorageLedger) StorageName() string {
	return "Ledger"
}

func (sl *StorageLedger) MapKeyHasher() uint8 {
	return Blake2_128ConcatHasher
}

func (sl *StorageLedger) Fetch(blockStorage interfaces.BlockStorageT, key StorageLedgerKey) (prim.Option[StorageLedgerEntry], error) {
	return GenericMapFetch[StorageLedger](blockStorage, key, sl)

}

func (sl *StorageLedger) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageLedgerEntry, error) {
	return GenericMapKeysFetch[StorageLedger, StorageLedgerKey](blockStorage, sl)
}

//
//
//

type StorageMaxNominatorsCountValue = uint32
type StorageMaxNominatorsCount struct{}

func (smnc *StorageMaxNominatorsCount) PalletName() string {
	return PalletName
}

func (smnc *StorageMaxNominatorsCount) StorageName() string {
	return "MaxNominatorsCount"
}

func (smnc *StorageMaxNominatorsCount) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxNominatorsCountValue], error) {
	return GenericFetch[StorageMaxNominatorsCountValue](blockStorage, smnc)
}

//
//
//

type StorageMaxValidatorsCountValue = uint32
type StorageMaxValidatorsCount struct{}

func (smvc *StorageMaxValidatorsCount) PalletName() string {
	return PalletName
}

func (smvc *StorageMaxValidatorsCount) StorageName() string {
	return "MaxValidatorsCount"
}

func (smvc *StorageMaxValidatorsCount) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxValidatorsCountValue], error) {
	return GenericFetch[StorageMaxValidatorsCountValue](blockStorage, smvc)
}

//
//
//

type StorageMinCommissionValue = Perbill
type StorageMinCommission struct{}

func (smc *StorageMinCommission) PalletName() string {
	return PalletName
}

func (smc *StorageMinCommission) StorageName() string {
	return "MinCommission"
}

func (smc *StorageMinCommission) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinCommissionValue, error) {
	return GenericFetchDefault[StorageMinCommissionValue](blockStorage, smc)
}

//
//
//

type StorageMinNominatorBondValue = Balance
type StorageMinNominatorBond struct{}

func (smnb *StorageMinNominatorBond) PalletName() string {
	return PalletName
}

func (smnb *StorageMinNominatorBond) StorageName() string {
	return "MinNominatorBond"
}

func (smnb *StorageMinNominatorBond) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinNominatorBondValue, error) {
	return GenericFetchDefault[StorageMinNominatorBondValue](blockStorage, smnb)
}

//
//
//

type StorageMinValidatorBondValue = Balance
type StorageMinValidatorBond struct{}

func (smvb *StorageMinValidatorBond) PalletName() string {
	return PalletName
}

func (smvb *StorageMinValidatorBond) StorageName() string {
	return "MinValidatorBond"
}

func (smvb *StorageMinValidatorBond) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinValidatorBondValue, error) {
	return GenericFetchDefault[StorageMinValidatorBondValue](blockStorage, smvb)
}

//
//
//

type StorageMinimumActiveStakeValue = Balance
type StorageMinimumActiveStake struct{}

func (smas *StorageMinimumActiveStake) PalletName() string {
	return PalletName
}

func (smas *StorageMinimumActiveStake) StorageName() string {
	return "MinimumActiveStake"
}

func (smas *StorageMinimumActiveStake) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinimumActiveStakeValue, error) {
	return GenericFetchDefault[StorageMinimumActiveStakeValue](blockStorage, smas)
}

//
//
//

type StorageMinimumValidatorCountValue = uint32
type StorageMinimumValidatorCount struct{}

func (smvc *StorageMinimumValidatorCount) PalletName() string {
	return PalletName
}

func (smvc *StorageMinimumValidatorCount) StorageName() string {
	return "MinimumValidatorCount"
}

func (smvc *StorageMinimumValidatorCount) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinimumValidatorCountValue, error) {
	return GenericFetchDefault[StorageMinimumValidatorCountValue](blockStorage, smvc)
}

//
//
//

type StorageNominatorSlashInEraKey1 = uint32
type StorageNominatorSlashInEraKey2 = prim.AccountId
type StorageNominatorSlashInEraValue = Balance
type StorageNominatorSlashInEraEntry = StorageEntryDoubleMap[StorageNominatorSlashInEraKey1, StorageNominatorSlashInEraKey2, StorageNominatorSlashInEraValue]
type StorageNominatorSlashInEra struct{}

func (snsie *StorageNominatorSlashInEra) PalletName() string {
	return PalletName
}

func (snsie *StorageNominatorSlashInEra) StorageName() string {
	return "NominatorSlashInEra"
}

func (snsie *StorageNominatorSlashInEra) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (snsie *StorageNominatorSlashInEra) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (snsie *StorageNominatorSlashInEra) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageNominatorSlashInEraKey1, key2 StorageNominatorSlashInEraKey2) (prim.Option[StorageNominatorSlashInEraEntry], error) {
	return GenericDoubleMapFetch[StorageNominatorSlashInEraValue](blockStorage, key1, key2, snsie)
}

func (snsie *StorageNominatorSlashInEra) FetchAll(blockStorage interfaces.BlockStorageT, key StorageNominatorSlashInEraKey1) ([]StorageNominatorSlashInEraEntry, error) {
	return GenericDoubleMapKeysFetch[StorageNominatorSlashInEraValue, StorageNominatorSlashInEraKey1, StorageNominatorSlashInEraKey2](blockStorage, key, snsie)
}

//
//
//

type StorageNominatorsKey = prim.AccountId
type StorageNominatorsEntry = StorageEntry[StorageNominatorsKey, StorageNominators]

type StorageNominators struct {
	Targets     []prim.AccountId
	SubmittedIn uint32
	Suppressed  bool
}

func (sn *StorageNominators) PalletName() string {
	return PalletName
}

func (sn *StorageNominators) StorageName() string {
	return "Nominators"
}

func (sn *StorageNominators) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sn *StorageNominators) Fetch(blockStorage interfaces.BlockStorageT, key StorageNominatorsKey) (prim.Option[StorageNominatorsEntry], error) {
	return GenericMapFetch[StorageNominators](blockStorage, key, sn)

}

func (sn *StorageNominators) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageNominatorsEntry, error) {
	return GenericMapKeysFetch[StorageNominators, StorageNominatorsKey](blockStorage, sn)
}

//
//
//

type StorageOffendingValidatorsValue = []Tuple2[uint32, bool]
type StorageOffendingValidators struct{}

func (sov *StorageOffendingValidators) PalletName() string {
	return PalletName
}

func (sov *StorageOffendingValidators) StorageName() string {
	return "OffendingValidators"
}

func (sov *StorageOffendingValidators) Fetch(blockStorage interfaces.BlockStorageT) (StorageOffendingValidatorsValue, error) {
	val, err := GenericFetch[StorageOffendingValidatorsValue](blockStorage, sov)
	return val.UnwrapOr(StorageOffendingValidatorsValue{}), err
}

//
//
//

type StoragePayeeKey = prim.AccountId
type StoragePayeeValue = RewardDestination
type StoragePayeeEntry = StorageEntry[StoragePayeeKey, StoragePayeeValue]

type StoragePayee struct{}

func (sp *StoragePayee) PalletName() string {
	return PalletName
}

func (sp *StoragePayee) StorageName() string {
	return "Payee"
}

func (sp *StoragePayee) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sp *StoragePayee) Fetch(blockStorage interfaces.BlockStorageT, key StoragePayeeKey) (prim.Option[StoragePayeeEntry], error) {
	return GenericMapFetch[StoragePayeeValue](blockStorage, key, sp)

}

func (sp *StoragePayee) FetchAll(blockStorage interfaces.BlockStorageT) ([]StoragePayeeEntry, error) {
	return GenericMapKeysFetch[StoragePayeeValue, StoragePayeeKey](blockStorage, sp)
}

//
//
//

type StorageSlashRewardFractionValue = Perbill
type StorageSlashRewardFraction struct{}

func (ssrf *StorageSlashRewardFraction) PalletName() string {
	return PalletName
}

func (ssrf *StorageSlashRewardFraction) StorageName() string {
	return "SlashRewardFraction "
}

func (ssrf *StorageSlashRewardFraction) Fetch(blockStorage interfaces.BlockStorageT) (StorageSlashRewardFractionValue, error) {
	return GenericFetchDefault[StorageSlashRewardFractionValue](blockStorage, ssrf)
}

//
//
//

type StorageSlashingSpansKey = prim.AccountId
type StorageSlashingSpansEntry = StorageEntry[StoragePayeeKey, StorageSlashingSpans]

type StorageSlashingSpans struct {
	SpanIndex        uint32
	LastStart        uint32
	LastNonZeroSlash uint32
	Prior            []uint32
}

func (sss *StorageSlashingSpans) PalletName() string {
	return PalletName
}

func (sss *StorageSlashingSpans) StorageName() string {
	return "SlashingSpans"
}

func (sss *StorageSlashingSpans) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sss *StorageSlashingSpans) Fetch(blockStorage interfaces.BlockStorageT, key StorageSlashingSpansKey) (prim.Option[StorageSlashingSpansEntry], error) {
	return GenericMapFetch[StorageSlashingSpans](blockStorage, key, sss)

}

func (sss *StorageSlashingSpans) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageSlashingSpansEntry, error) {
	return GenericMapKeysFetch[StorageSlashingSpans, StorageSlashingSpansKey](blockStorage, sss)
}

//
//
//

type StorageSpanSlashKey = Tuple2[prim.AccountId, uint32]
type StorageSpanSlashEntry = StorageEntry[StorageSpanSlashKey, StorageSpanSlash]

type StorageSpanSlash struct {
	Slashed Balance
	PaidOut Balance
}

func (sss *StorageSpanSlash) PalletName() string {
	return PalletName
}

func (sss *StorageSpanSlash) StorageName() string {
	return "SpanSlash"
}

func (sss *StorageSpanSlash) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sss *StorageSpanSlash) Fetch(blockStorage interfaces.BlockStorageT, key StorageSpanSlashKey) (prim.Option[StorageSpanSlashEntry], error) {
	return GenericMapFetch[StorageSpanSlash](blockStorage, key, sss)

}

func (sss *StorageSpanSlash) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageSpanSlashEntry, error) {
	return GenericMapKeysFetch[StorageSpanSlash, StorageSpanSlashKey](blockStorage, sss)
}

//
//
//

type StorageUnappliedSlashesKey = uint32
type StorageUnappliedSlashesValue = []StorageUnappliedSlashes
type StorageUnappliedSlashesEntry = StorageEntry[StorageUnappliedSlashesKey, StorageUnappliedSlashesValue]

type StorageUnappliedSlashes struct {
	Validator prim.AccountId
	Own       Balance
	Others    []Tuple2[prim.AccountId, Balance]
	Reporters []prim.AccountId
	Payout    Balance
}

func (sus *StorageUnappliedSlashes) PalletName() string {
	return PalletName
}

func (sus *StorageUnappliedSlashes) StorageName() string {
	return "UnappliedSlashes"
}

func (sus *StorageUnappliedSlashes) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sus *StorageUnappliedSlashes) Fetch(blockStorage interfaces.BlockStorageT, key StorageUnappliedSlashesKey) (StorageUnappliedSlashesEntry, error) {
	val, err := GenericMapFetch[StorageUnappliedSlashesValue](blockStorage, key, sus)
	return val.UnwrapOr(StorageUnappliedSlashesEntry{}), err

}

func (sus *StorageUnappliedSlashes) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageUnappliedSlashesEntry, error) {
	return GenericMapKeysFetch[StorageUnappliedSlashesValue, StorageUnappliedSlashesKey](blockStorage, sus)
}

//
//
//

type StorageValidatorCountValue = uint32
type StorageValidatorCount struct{}

func (svc *StorageValidatorCount) PalletName() string {
	return PalletName
}

func (svc *StorageValidatorCount) StorageName() string {
	return "ValidatorCount"
}

func (svc *StorageValidatorCount) Fetch(blockStorage interfaces.BlockStorageT) (StorageValidatorCountValue, error) {
	return GenericFetchDefault[StorageValidatorCountValue](blockStorage, svc)
}

//
//
//

type StorageValidatorSlashInEraKey1 = uint32
type StorageValidatorSlashInEraKey2 = prim.AccountId
type StorageValidatorSlashInEraValue = Tuple2[Perbill, Balance]
type StorageValidatorSlashInEraEntry = StorageEntryDoubleMap[StorageValidatorSlashInEraKey1, StorageValidatorSlashInEraKey2, StorageValidatorSlashInEraValue]
type StorageValidatorSlashInEra struct{}

func (svsie *StorageValidatorSlashInEra) PalletName() string {
	return PalletName
}

func (svsie *StorageValidatorSlashInEra) StorageName() string {
	return "ValidatorSlashInEra"
}

func (svsie *StorageValidatorSlashInEra) MapKey1Hasher() uint8 {
	return Twox64ConcatHasher
}

func (svsie *StorageValidatorSlashInEra) MapKey2Hasher() uint8 {
	return Twox64ConcatHasher
}

func (svsie *StorageValidatorSlashInEra) Fetch(blockStorage interfaces.BlockStorageT, key1 StorageValidatorSlashInEraKey1, key2 StorageValidatorSlashInEraKey2) (prim.Option[StorageValidatorSlashInEraEntry], error) {
	return GenericDoubleMapFetch[StorageValidatorSlashInEraValue](blockStorage, key1, key2, svsie)
}

func (svsie *StorageValidatorSlashInEra) FetchAll(blockStorage interfaces.BlockStorageT, key StorageValidatorSlashInEraKey1) ([]StorageValidatorSlashInEraEntry, error) {
	return GenericDoubleMapKeysFetch[StorageValidatorSlashInEraValue, StorageValidatorSlashInEraKey1, StorageValidatorSlashInEraKey2](blockStorage, key, svsie)
}

//
//
//

type StorageValidatorsKey = prim.AccountId
type StorageValidatorsValue = ValidatorPrefs
type StorageValidatorsEntry = StorageEntry[StorageValidatorsKey, StorageValidatorsValue]

type StorageValidators struct {
	Slashed Balance
	PaidOut Balance
}

func (sv *StorageValidators) PalletName() string {
	return PalletName
}

func (sv *StorageValidators) StorageName() string {
	return "Validators"
}

func (sv *StorageValidators) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sv *StorageValidators) Fetch(blockStorage interfaces.BlockStorageT, key StorageValidatorsKey) (StorageValidatorsEntry, error) {
	val, err := GenericMapFetch[StorageValidatorsValue](blockStorage, key, sv)
	return val.Unwrap(), err

}

func (sv *StorageValidators) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageValidatorsEntry, error) {
	return GenericMapKeysFetch[StorageValidatorsValue, StorageValidatorsKey](blockStorage, sv)
}
