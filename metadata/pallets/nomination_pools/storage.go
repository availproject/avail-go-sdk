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

func (this *StorageBondedPools) PalletName() string {
	return PalletName
}

func (this *StorageBondedPools) StorageName() string {
	return "BondedPools"
}

func (this *StorageBondedPools) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageBondedPools) Fetch(blockStorage interfaces.BlockStorageT, key StorageBondedPoolsKey) (prim.Option[StorageBondedPoolsEntry], error) {
	return GenericMapFetch[StorageBondedPools](blockStorage, key, this)
}

func (this *StorageBondedPools) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageBondedPoolsEntry, error) {
	return GenericMapKeysFetch[StorageBondedPools, StorageBondedPoolsKey](blockStorage, this)
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

func (this *StorageClaimPermissions) PalletName() string {
	return PalletName
}

func (this *StorageClaimPermissions) StorageName() string {
	return "ClaimPermissions"
}

func (this *StorageClaimPermissions) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageClaimPermissions) Fetch(blockStorage interfaces.BlockStorageT, key StorageClaimPermissionsKey) (StorageClaimPermissionsEntry, error) {
	val, err := GenericMapFetch[StorageClaimPermissionsValue](blockStorage, key, this)
	return val.Unwrap(), err
}

func (this *StorageClaimPermissions) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageClaimPermissionsEntry, error) {
	return GenericMapKeysFetch[StorageClaimPermissionsValue, StorageClaimPermissionsKey](blockStorage, this)
}

//
//
//

type StorageCounterForBondedPoolsValue = uint32
type StorageCounterForBondedPools struct{}

func (this *StorageCounterForBondedPools) PalletName() string {
	return PalletName
}

func (this *StorageCounterForBondedPools) StorageName() string {
	return "CounterForBondedPools"
}

func (this *StorageCounterForBondedPools) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForBondedPoolsValue, error) {
	return GenericFetchDefault[StorageCounterForBondedPoolsValue](blockStorage, this)
}

//
//
//

type StorageCounterForMetadataValue = uint32
type StorageCounterForMetadata struct{}

func (this *StorageCounterForMetadata) PalletName() string {
	return PalletName
}

func (this *StorageCounterForMetadata) StorageName() string {
	return "CounterForMetadata"
}

func (this *StorageCounterForMetadata) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForMetadataValue, error) {
	return GenericFetchDefault[StorageCounterForMetadataValue](blockStorage, this)
}

//
//
//

type StorageCounterForPoolMembersValue = uint32
type StorageCounterForPoolMembers struct{}

func (this *StorageCounterForPoolMembers) PalletName() string {
	return PalletName
}

func (this *StorageCounterForPoolMembers) StorageName() string {
	return "CounterForPoolMembers"
}

func (this *StorageCounterForPoolMembers) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForPoolMembersValue, error) {
	return GenericFetchDefault[StorageCounterForPoolMembersValue](blockStorage, this)
}

//
//
//

type StorageCounterForReversePoolIdLookupValue = uint32
type StorageCounterForReversePoolIdLookup struct{}

func (this *StorageCounterForReversePoolIdLookup) PalletName() string {
	return PalletName
}

func (this *StorageCounterForReversePoolIdLookup) StorageName() string {
	return "CounterForReversePoolIdLookup"
}

func (this *StorageCounterForReversePoolIdLookup) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForReversePoolIdLookupValue, error) {
	return GenericFetchDefault[StorageCounterForReversePoolIdLookupValue](blockStorage, this)
}

//
//
//

type StorageCounterForRewardPoolsValue = uint32
type StorageCounterForRewardPools struct{}

func (this *StorageCounterForRewardPools) PalletName() string {
	return PalletName
}

func (this *StorageCounterForRewardPools) StorageName() string {
	return "CounterForRewardPools"
}

func (this *StorageCounterForRewardPools) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForRewardPoolsValue, error) {
	return GenericFetchDefault[StorageCounterForRewardPoolsValue](blockStorage, this)
}

//
//
//

type StorageCounterForSubPoolsStorageValue = uint32
type StorageCounterForSubPoolsStorage struct{}

func (this *StorageCounterForSubPoolsStorage) PalletName() string {
	return PalletName
}

func (this *StorageCounterForSubPoolsStorage) StorageName() string {
	return "CounterForSubPoolsStorage"
}

func (this *StorageCounterForSubPoolsStorage) Fetch(blockStorage interfaces.BlockStorageT) (StorageCounterForSubPoolsStorageValue, error) {
	return GenericFetchDefault[StorageCounterForSubPoolsStorageValue](blockStorage, this)
}

//
//
//

type StorageGlobalMaxCommissionValue = Perbill
type StorageGlobalMaxCommission struct{}

func (this *StorageGlobalMaxCommission) PalletName() string {
	return PalletName
}

func (this *StorageGlobalMaxCommission) StorageName() string {
	return "GlobalMaxCommission"
}

func (this *StorageGlobalMaxCommission) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageGlobalMaxCommissionValue], error) {
	return GenericFetch[StorageGlobalMaxCommissionValue](blockStorage, this)
}

//
//
//

type StorageLastPoolIdValue = uint32
type StorageLastPoolId struct{}

func (this *StorageLastPoolId) PalletName() string {
	return PalletName
}

func (this *StorageLastPoolId) StorageName() string {
	return "LastPoolId"
}

func (this *StorageLastPoolId) Fetch(blockStorage interfaces.BlockStorageT) (StorageLastPoolIdValue, error) {
	return GenericFetchDefault[StorageLastPoolIdValue](blockStorage, this)
}

//
//
//

type StorageMaxPoolMembersValue = uint32
type StorageMaxPoolMembers struct{}

func (this *StorageMaxPoolMembers) PalletName() string {
	return PalletName
}

func (this *StorageMaxPoolMembers) StorageName() string {
	return "MaxPoolMembers"
}

func (this *StorageMaxPoolMembers) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxPoolMembersValue], error) {
	return GenericFetch[StorageMaxPoolMembersValue](blockStorage, this)
}

//
//
//

type StorageMaxPoolMembersPerPoolValue = uint32
type StorageMaxPoolMembersPerPool struct{}

func (this *StorageMaxPoolMembersPerPool) PalletName() string {
	return PalletName
}

func (this *StorageMaxPoolMembersPerPool) StorageName() string {
	return "MaxPoolMembersPerPool"
}

func (this *StorageMaxPoolMembersPerPool) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxPoolMembersPerPoolValue], error) {
	return GenericFetch[StorageMaxPoolMembersPerPoolValue](blockStorage, this)
}

//
//
//

type StorageMaxPoolsValue = uint32
type StorageMaxPools struct{}

func (this *StorageMaxPools) PalletName() string {
	return PalletName
}

func (this *StorageMaxPools) StorageName() string {
	return "MaxPools"
}

func (this *StorageMaxPools) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageMaxPoolsValue], error) {
	return GenericFetch[StorageMaxPoolsValue](blockStorage, this)
}

//
//
//

type StorageMetadataKey = uint32
type StorageMetadataValue = []byte
type StorageMetadataEntry = StorageEntry[StorageMetadataKey, StorageMetadataValue]

type StorageMetadata struct{}

func (this *StorageMetadata) PalletName() string {
	return PalletName
}

func (this *StorageMetadata) StorageName() string {
	return "Metadata"
}

func (this *StorageMetadata) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageMetadata) Fetch(blockStorage interfaces.BlockStorageT, key StorageMetadataKey) (StorageMetadataEntry, error) {
	val, err := GenericMapFetch[StorageMetadataValue](blockStorage, key, this)
	return val.Unwrap(), err
}

func (this *StorageMetadata) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageMetadataEntry, error) {
	return GenericMapKeysFetch[StorageMetadataValue, StorageMetadataKey](blockStorage, this)
}

//
//
//

type StorageMinCreateBondValue = Balance
type StorageMinCreateBond struct{}

func (this *StorageMinCreateBond) PalletName() string {
	return PalletName
}

func (this *StorageMinCreateBond) StorageName() string {
	return "MinCreateBond"
}

func (this *StorageMinCreateBond) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinCreateBondValue, error) {
	return GenericFetchDefault[StorageMinCreateBondValue](blockStorage, this)
}

//
//
//

type StorageMinJoinBondValue = Balance
type StorageMinJoinBond struct{}

func (this *StorageMinJoinBond) PalletName() string {
	return PalletName
}

func (this *StorageMinJoinBond) StorageName() string {
	return "MinJoinBond"
}

func (this *StorageMinJoinBond) Fetch(blockStorage interfaces.BlockStorageT) (StorageMinJoinBondValue, error) {
	return GenericFetchDefault[StorageMinJoinBondValue](blockStorage, this)
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

func (this *StoragePoolMembers) PalletName() string {
	return PalletName
}

func (this *StoragePoolMembers) StorageName() string {
	return "PoolMembers"
}

func (this *StoragePoolMembers) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StoragePoolMembers) Fetch(blockStorage interfaces.BlockStorageT, key StoragePoolMembersKey) (prim.Option[StoragePoolMembersEntry], error) {
	return GenericMapFetch[StoragePoolMembers](blockStorage, key, this)
}

func (this *StoragePoolMembers) FetchAll(blockStorage interfaces.BlockStorageT) ([]StoragePoolMembersEntry, error) {
	return GenericMapKeysFetch[StoragePoolMembers, StoragePoolMembersKey](blockStorage, this)
}

//
//
//

type StorageReversePoolIdLookupKey = prim.AccountId
type StorageReversePoolIdLookupValue = uint32
type StorageReversePoolIdLookupEntry = StorageEntry[StorageReversePoolIdLookupKey, StorageReversePoolIdLookupValue]

type StorageReversePoolIdLookup struct{}

func (this *StorageReversePoolIdLookup) PalletName() string {
	return PalletName
}

func (this *StorageReversePoolIdLookup) StorageName() string {
	return "ReversePoolIdLookup"
}

func (this *StorageReversePoolIdLookup) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageReversePoolIdLookup) Fetch(blockStorage interfaces.BlockStorageT, key StorageReversePoolIdLookupKey) (prim.Option[StorageReversePoolIdLookupEntry], error) {
	return GenericMapFetch[StorageReversePoolIdLookupValue](blockStorage, key, this)
}

func (this *StorageReversePoolIdLookup) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageReversePoolIdLookupEntry, error) {
	return GenericMapKeysFetch[StorageReversePoolIdLookupValue, StorageReversePoolIdLookupKey](blockStorage, this)
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

func (this *StorageRewardPools) PalletName() string {
	return PalletName
}

func (this *StorageRewardPools) StorageName() string {
	return "RewardPools"
}

func (this *StorageRewardPools) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageRewardPools) Fetch(blockStorage interfaces.BlockStorageT, key StorageRewardPoolsKey) (prim.Option[StorageRewardPoolsEntry], error) {
	return GenericMapFetch[StorageRewardPools](blockStorage, key, this)
}

func (this *StorageRewardPools) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageRewardPoolsEntry, error) {
	return GenericMapKeysFetch[StorageRewardPools, StorageRewardPoolsKey](blockStorage, this)
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

func (this *StorageSubPoolsStorage) PalletName() string {
	return PalletName
}

func (this *StorageSubPoolsStorage) StorageName() string {
	return "SubPoolsStorage"
}

func (this *StorageSubPoolsStorage) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageSubPoolsStorage) Fetch(blockStorage interfaces.BlockStorageT, key StorageSubPoolsStorageKey) (prim.Option[StorageSubPoolsStorageEntry], error) {
	return GenericMapFetch[StorageSubPoolsStorage](blockStorage, key, this)
}

func (this *StorageSubPoolsStorage) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageSubPoolsStorageEntry, error) {
	return GenericMapKeysFetch[StorageSubPoolsStorage, StorageSubPoolsStorageKey](blockStorage, this)
}

//
//
//

type StorageTotalValueLockedValue = uint128.Uint128
type StorageTotalValueLocked struct{}

func (this *StorageTotalValueLocked) PalletName() string {
	return PalletName
}

func (this *StorageTotalValueLocked) StorageName() string {
	return "TotalValueLocked"
}

func (this *StorageTotalValueLocked) Fetch(blockStorage interfaces.BlockStorageT) (StorageTotalValueLockedValue, error) {
	return GenericFetchDefault[StorageTotalValueLockedValue](blockStorage, this)
}
