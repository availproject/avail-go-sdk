package nomination_pools

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"

	"github.com/itering/scale.go/utiles/uint128"
)

type StorageBondedPoolsKey = uint32
type StorageBondedPoolsEntry = StorageEntry[StorageBondedPoolsKey, StorageBondedPools]

type StorageBondedPools struct {
	Commission    PoolCommission
	MemberCounter uint32
	Points        uint128.Uint128
	Roles         PoolRoles
	State         PoolState
}

func (sbp *StorageBondedPools) PalletName() string {
	return PalletName
}

func (sbp *StorageBondedPools) StorageName() string {
	return "BondedPools"
}

func (sbp *StorageBondedPools) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sbp *StorageBondedPools) Fetch(blockStorage interfaces.BlockStorageT, key StorageBondedPoolsKey) (prim.Option[StorageBondedPoolsEntry], error) {
	return GenericMapFetch[StorageBondedPools](blockStorage, key, sbp)
}

func (sbp *StorageBondedPools) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageBondedPoolsEntry, error) {
	return GenericMapKeysFetch[StorageBondedPools, StorageBondedPoolsKey](blockStorage, sbp)
}

//
//
//

type StorageClaimPermissionsKey = prim.AccountId
type StorageClaimPermissionsValue = PoolClaimPermission
type StorageClaimPermissionsEntry = StorageEntry[StorageClaimPermissionsKey, StorageClaimPermissionsValue]

type StorageClaimPermissions struct {
	Commission    PoolCommission
	MemberCounter uint32
	Points        uint128.Uint128
	Roles         PoolRoles
	State         PoolState
}

func (scp *StorageClaimPermissions) PalletName() string {
	return PalletName
}

func (scp *StorageClaimPermissions) StorageName() string {
	return "ClaimPermissions"
}

func (scp *StorageClaimPermissions) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (scp *StorageClaimPermissions) Fetch(blockStorage interfaces.BlockStorageT, key StorageClaimPermissionsKey) (StorageClaimPermissionsEntry, error) {
	val, err := GenericMapFetch[StorageClaimPermissionsValue](blockStorage, key, scp)
	return val.Unwrap(), err
}

func (scp *StorageClaimPermissions) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageClaimPermissionsEntry, error) {
	return GenericMapKeysFetch[StorageClaimPermissionsValue, StorageClaimPermissionsKey](blockStorage, scp)
}

//
//
//

type StorageCounterForBondedPoolsValue = uint32
type StorageCounterForBondedPools struct{}

func (scfbp *StorageCounterForBondedPools) PalletName() string {
	return PalletName
}

func (scfbp *StorageCounterForBondedPools) StorageName() string {
	return "CounterForBondedPools"
}

func (scfbp *StorageCounterForBondedPools) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForBondedPoolsValue, error) {
	return GenericFetchDefault[StorageCounterForBondedPoolsValue](blockStorage, scfbp)
}

//
//
//

type StorageCounterForMetadataValue = uint32
type StorageCounterForMetadata struct{}

func (scfm *StorageCounterForMetadata) PalletName() string {
	return PalletName
}

func (scfm *StorageCounterForMetadata) StorageName() string {
	return "CounterForMetadata"
}

func (scfm *StorageCounterForMetadata) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForMetadataValue, error) {
	return GenericFetchDefault[StorageCounterForMetadataValue](blockStorage, scfm)
}

//
//
//

type StorageCounterForPoolMembersValue = uint32
type StorageCounterForPoolMembers struct{}

func (scfpm *StorageCounterForPoolMembers) PalletName() string {
	return PalletName
}

func (scfpm *StorageCounterForPoolMembers) StorageName() string {
	return "CounterForPoolMembers"
}

func (scfpm *StorageCounterForPoolMembers) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForPoolMembersValue, error) {
	return GenericFetchDefault[StorageCounterForPoolMembersValue](blockStorage, scfpm)
}

//
//
//

type StorageCounterForReversePoolIdLookupValue = uint32
type StorageCounterForReversePoolIdLookup struct{}

func (scfrpil *StorageCounterForReversePoolIdLookup) PalletName() string {
	return PalletName
}

func (scfrpil *StorageCounterForReversePoolIdLookup) StorageName() string {
	return "CounterForReversePoolIdLookup"
}

func (scfrpil *StorageCounterForReversePoolIdLookup) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForReversePoolIdLookupValue, error) {
	return GenericFetchDefault[StorageCounterForReversePoolIdLookupValue](blockStorage, scfrpil)
}

//
//
//

type StorageCounterForRewardPoolsValue = uint32
type StorageCounterForRewardPools struct{}

func (scfrp *StorageCounterForRewardPools) PalletName() string {
	return PalletName
}

func (scfrp *StorageCounterForRewardPools) StorageName() string {
	return "CounterForRewardPools"
}

func (scfrp *StorageCounterForRewardPools) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForRewardPoolsValue, error) {
	return GenericFetchDefault[StorageCounterForRewardPoolsValue](blockStorage, scfrp)
}

//
//
//

type StorageCounterForSubPoolsStorageValue = uint32
type StorageCounterForSubPoolsStorage struct{}

func (scfsps *StorageCounterForSubPoolsStorage) PalletName() string {
	return PalletName
}

func (scfsps *StorageCounterForSubPoolsStorage) StorageName() string {
	return "CounterForSubPoolsStorage"
}

func (scfsps *StorageCounterForSubPoolsStorage) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForSubPoolsStorageValue, error) {
	return GenericFetchDefault[StorageCounterForSubPoolsStorageValue](blockStorage, scfsps)
}

//
//
//

type StorageGlobalMaxCommissionValue = Perbill
type StorageGlobalMaxCommission struct{}

func (sgmc *StorageGlobalMaxCommission) PalletName() string {
	return PalletName
}

func (sgmc *StorageGlobalMaxCommission) StorageName() string {
	return "GlobalMaxCommission"
}

func (sgmc *StorageGlobalMaxCommission) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageGlobalMaxCommissionValue], error) {
	return GenericFetch[StorageGlobalMaxCommissionValue](blockStorage, sgmc)
}

//
//
//

type StorageLastPoolIdValue = uint32
type StorageLastPoolId struct{}

func (slp *StorageLastPoolId) PalletName() string {
	return PalletName
}

func (slp *StorageLastPoolId) StorageName() string {
	return "LastPoolId"
}

func (slp *StorageLastPoolId) Fetch(blockStorage interfaces.BlockStorageT) (StorageLastPoolIdValue, error) {
	return GenericFetchDefault[StorageLastPoolIdValue](blockStorage, slp)
}

//
//
//

type StorageMaxPoolMembersValue = uint32
type StorageMaxPoolMembers struct{}

func (smpm *StorageMaxPoolMembers) PalletName() string {
	return PalletName
}

func (smpm *StorageMaxPoolMembers) StorageName() string {
	return "MaxPoolMembers"
}

func (smpm *StorageMaxPoolMembers) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxPoolMembersValue], error) {
	return GenericFetch[StorageMaxPoolMembersValue](blockStorage, smpm)
}

//
//
//

type StorageMaxPoolMembersPerPoolValue = uint32
type StorageMaxPoolMembersPerPool struct{}

func (smpmpp *StorageMaxPoolMembersPerPool) PalletName() string {
	return PalletName
}

func (smpmpp *StorageMaxPoolMembersPerPool) StorageName() string {
	return "MaxPoolMembersPerPool"
}

func (smpmpp *StorageMaxPoolMembersPerPool) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxPoolMembersPerPoolValue], error) {
	return GenericFetch[StorageMaxPoolMembersPerPoolValue](blockStorage, smpmpp)
}

//
//
//

type StorageMaxPoolsValue = uint32
type StorageMaxPools struct{}

func (smp *StorageMaxPools) PalletName() string {
	return PalletName
}

func (smp *StorageMaxPools) StorageName() string {
	return "MaxPools"
}

func (smp *StorageMaxPools) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxPoolsValue], error) {
	return GenericFetch[StorageMaxPoolsValue](blockStorage, smp)
}

//
//
//

type StorageMetadataKey = uint32
type StorageMetadataValue = []byte
type StorageMetadataEntry = StorageEntry[StorageMetadataKey, StorageMetadataValue]

type StorageMetadata struct{}

func (sm *StorageMetadata) PalletName() string {
	return PalletName
}

func (sm *StorageMetadata) StorageName() string {
	return "Metadata"
}

func (sm *StorageMetadata) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sm *StorageMetadata) Fetch(blockStorage interfaces.BlockStorageT, key StorageMetadataKey) (StorageMetadataEntry, error) {
	val, err := GenericMapFetch[StorageMetadataValue](blockStorage, key, sm)
	return val.Unwrap(), err
}

func (sm *StorageMetadata) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageMetadataEntry, error) {
	return GenericMapKeysFetch[StorageMetadataValue, StorageMetadataKey](blockStorage, sm)
}

//
//
//

type StorageMinCreateBondValue = Balance
type StorageMinCreateBond struct{}

func (smcb *StorageMinCreateBond) PalletName() string {
	return PalletName
}

func (smcb *StorageMinCreateBond) StorageName() string {
	return "MinCreateBond"
}

func (smcb *StorageMinCreateBond) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinCreateBondValue, error) {
	return GenericFetchDefault[StorageMinCreateBondValue](blockStorage, smcb)
}

//
//
//

type StorageMinJoinBondValue = Balance
type StorageMinJoinBond struct{}

func (smjb *StorageMinJoinBond) PalletName() string {
	return PalletName
}

func (smjb *StorageMinJoinBond) StorageName() string {
	return "MinJoinBond"
}

func (smjb *StorageMinJoinBond) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinJoinBondValue, error) {
	return GenericFetchDefault[StorageMinJoinBondValue](blockStorage, smjb)
}

//
//
//

type StoragePoolMembersKey = prim.AccountId
type StoragePoolMembersEntry = StorageEntry[StoragePoolMembersKey, StoragePoolMembers]

type StoragePoolMembers struct {
	PoolId                   uint32
	Points                   uint128.Uint128
	LasRecordedRewardCounter uint128.Uint128
	UnbondingEras            []Tuple2[uint32, uint128.Uint128]
}

func (spm *StoragePoolMembers) PalletName() string {
	return PalletName
}

func (spm *StoragePoolMembers) StorageName() string {
	return "PoolMembers"
}

func (spm *StoragePoolMembers) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (spm *StoragePoolMembers) Fetch(blockStorage interfaces.BlockStorageT, key StoragePoolMembersKey) (prim.Option[StoragePoolMembersEntry], error) {
	return GenericMapFetch[StoragePoolMembers](blockStorage, key, spm)
}

func (spm *StoragePoolMembers) FetchAll(blockStorage interfaces.BlockStorageT) ([]StoragePoolMembersEntry, error) {
	return GenericMapKeysFetch[StoragePoolMembers, StoragePoolMembersKey](blockStorage, spm)
}

//
//
//

type StorageReversePoolIdLookupKey = prim.AccountId
type StorageReversePoolIdLookupValue = uint32
type StorageReversePoolIdLookupEntry = StorageEntry[StorageReversePoolIdLookupKey, StorageReversePoolIdLookupValue]

type StorageReversePoolIdLookup struct{}

func (srpil *StorageReversePoolIdLookup) PalletName() string {
	return PalletName
}

func (srpil *StorageReversePoolIdLookup) StorageName() string {
	return "ReversePoolIdLookup"
}

func (srpil *StorageReversePoolIdLookup) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (srpil *StorageReversePoolIdLookup) Fetch(blockStorage interfaces.BlockStorageT, key StorageReversePoolIdLookupKey) (prim.Option[StorageReversePoolIdLookupEntry], error) {
	return GenericMapFetch[StorageReversePoolIdLookupValue](blockStorage, key, srpil)
}

func (srpil *StorageReversePoolIdLookup) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageReversePoolIdLookupEntry, error) {
	return GenericMapKeysFetch[StorageReversePoolIdLookupValue, StorageReversePoolIdLookupKey](blockStorage, srpil)
}

//
//
//

type StorageRewardPoolsKey = uint32
type StorageRewardPoolsEntry = StorageEntry[StorageRewardPoolsKey, StorageRewardPools]

type StorageRewardPools struct {
	LastRecordedRewardCounter uint128.Uint128
	LastRecordedTotalPayouts  Balance
	TotalRewardClaimed        Balance
	TotalCommissionPending    Balance
	TotalCommissionClaimed    Balance
}

func (srp *StorageRewardPools) PalletName() string {
	return PalletName
}

func (srp *StorageRewardPools) StorageName() string {
	return "RewardPools"
}

func (srp *StorageRewardPools) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (srp *StorageRewardPools) Fetch(blockStorage interfaces.BlockStorageT, key StorageRewardPoolsKey) (prim.Option[StorageRewardPoolsEntry], error) {
	return GenericMapFetch[StorageRewardPools](blockStorage, key, srp)
}

func (srp *StorageRewardPools) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageRewardPoolsEntry, error) {
	return GenericMapKeysFetch[StorageRewardPools, StorageRewardPoolsKey](blockStorage, srp)
}

//
//
//

type StorageSubPoolsStorageKey = uint32
type StorageSubPoolsStorageEntry = StorageEntry[StorageSubPoolsStorageKey, StorageSubPoolsStorage]

type StorageSubPoolsStorage struct {
	NoEra   UnbondPool
	WithEra []Tuple2[uint32, UnbondPool]
}

type UnbondPool struct {
	Points  uint128.Uint128
	Balance Balance
}

func (ssps *StorageSubPoolsStorage) PalletName() string {
	return PalletName
}

func (ssps *StorageSubPoolsStorage) StorageName() string {
	return "SubPoolsStorage"
}

func (ssps *StorageSubPoolsStorage) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (ssps *StorageSubPoolsStorage) Fetch(blockStorage interfaces.BlockStorageT, key StorageSubPoolsStorageKey) (prim.Option[StorageSubPoolsStorageEntry], error) {
	return GenericMapFetch[StorageSubPoolsStorage](blockStorage, key, ssps)
}

func (ssps *StorageSubPoolsStorage) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageSubPoolsStorageEntry, error) {
	return GenericMapKeysFetch[StorageSubPoolsStorage, StorageSubPoolsStorageKey](blockStorage, ssps)
}

//
//
//

type StorageTotalValueLockedValue = uint128.Uint128
type StorageTotalValueLocked struct{}

func (stvl *StorageTotalValueLocked) PalletName() string {
	return PalletName
}

func (stvl *StorageTotalValueLocked) StorageName() string {
	return "TotalValueLocked"
}

func (stvl *StorageTotalValueLocked) Fetch(blockStorage interfaces.BlockStorageT) (StorageTotalValueLockedValue, error) {
	return GenericFetchDefault[StorageTotalValueLockedValue](blockStorage, stvl)
}
