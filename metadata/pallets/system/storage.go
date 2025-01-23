package system

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageAccountKey = AccountId
type StorageAccountEntry = StorageEntry[StorageAccountKey, StorageAccount]

type StorageAccount struct {
	Nonce       uint32
	Consumers   uint32
	Providers   uint32
	Sufficients uint32
	AccountData AccountData
}

func (this *StorageAccount) PalletName() string {
	return PalletName
}

func (this *StorageAccount) StorageName() string {
	return "Account"
}

func (this *StorageAccount) MapKeyHasher() uint8 {
	return Blake2_128ConcatHasher
}

func (this *StorageAccount) Fetch(blockStorage interfaces.BlockStorageT, key StorageAccountKey) (StorageAccountEntry, error) {
	val, err := GenericMapFetch[StorageAccount](blockStorage, key, this)
	return val.Unwrap(), err
}

func (this *StorageAccount) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageAccountEntry, error) {
	return GenericMapKeysFetch[StorageAccount, StorageAccountKey](blockStorage, this)
}

//
//
//

type StorageBlockHashKey = uint32
type StorageBlockHashEntry = StorageEntry[StorageBlockHashKey, StorageBlockHash]

type StorageBlockHash struct {
	Value prim.H256
}

func (this *StorageBlockHash) PalletName() string {
	return PalletName
}

func (this *StorageBlockHash) StorageName() string {
	return "BlockHash"
}

func (this *StorageBlockHash) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageBlockHash) Fetch(blockStorage interfaces.BlockStorageT, key StorageBlockHashKey) (prim.Option[StorageBlockHashEntry], error) {
	return GenericMapFetch[StorageBlockHash](blockStorage, key, this)
}

func (this *StorageBlockHash) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageBlockHashEntry, error) {
	return GenericMapKeysFetch[StorageBlockHash, StorageBlockHashKey](blockStorage, this)
}

//
//
//

type StorageBlockWeight struct {
	Normal      Weight
	Operational Weight
	Mandatory   Weight
}

func (this *StorageBlockWeight) PalletName() string {
	return PalletName
}

func (this *StorageBlockWeight) StorageName() string {
	return "BlockWeight"
}

func (this *StorageBlockWeight) Fetch(blockStorage interfaces.BlockStorageT) (StorageBlockWeight, error) {
	return GenericFetchDefault[StorageBlockWeight](blockStorage, this)
}

//
//
//

type StorageDynamicBlockLength struct {
	Max       PerDispatchClassU32
	Cols      uint32 `scale:"compact"`
	Rows      uint32 `scale:"compact"`
	ChunkSize uint32 `scale:"compact"`
}

func (this *StorageDynamicBlockLength) PalletName() string {
	return PalletName
}

func (this *StorageDynamicBlockLength) StorageName() string {
	return "DynamicBlockLength"
}

func (this *StorageDynamicBlockLength) Fetch(blockStorage interfaces.BlockStorageT) (StorageDynamicBlockLength, error) {
	val, err := GenericFetch[StorageDynamicBlockLength](blockStorage, this)

	// TODO Fallback might not be correct.
	// Fallback: 0x00003c0000005000000050000104010480
	return val.Unwrap(), err
}

//
//
//

type StorageEventCountValue = uint32
type StorageEventCount struct{}

func (this *StorageEventCount) PalletName() string {
	return PalletName
}

func (this *StorageEventCount) StorageName() string {
	return "EventCount"
}

func (this *StorageEventCount) Fetch(blockStorage interfaces.BlockStorageT) (StorageEventCountValue, error) {
	return GenericFetchDefault[StorageEventCountValue](blockStorage, this)
}

//
//
//

type StorageExtrinsicCountValue = uint32
type StorageExtrinsicCount struct{}

func (this *StorageExtrinsicCount) PalletName() string {
	return PalletName
}

func (this *StorageExtrinsicCount) StorageName() string {
	return "ExtrinsicCount"
}

func (this *StorageExtrinsicCount) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageExtrinsicCountValue], error) {
	return GenericFetch[StorageExtrinsicCountValue](blockStorage, this)
}

//
//
//

type StorageNumberValue = uint32
type StorageNumber struct{}

func (this *StorageNumber) PalletName() string {
	return PalletName
}

func (this *StorageNumber) StorageName() string {
	return "Number"
}

func (this *StorageNumber) Fetch(blockStorage interfaces.BlockStorageT) (StorageNumberValue, error) {
	return GenericFetchDefault[StorageNumberValue](blockStorage, this)
}

//
//
//

type StorageParentHashValue = prim.H256
type StorageParentHash struct{}

func (this *StorageParentHash) PalletName() string {
	return PalletName
}

func (this *StorageParentHash) StorageName() string {
	return "ParentHash"
}

func (this *StorageParentHash) Fetch(blockStorage interfaces.BlockStorageT) (StorageParentHashValue, error) {
	return GenericFetchDefault[StorageParentHashValue](blockStorage, this)
}
