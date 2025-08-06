package system

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageAccountKey = prim.AccountId
type StorageAccountEntry = StorageEntry[StorageAccountKey, StorageAccount]

type StorageAccount struct {
	Nonce       uint32
	Consumers   uint32
	Providers   uint32
	Sufficients uint32
	AccountData AccountData
}

func (sa *StorageAccount) PalletName() string {
	return PalletName
}

func (sa *StorageAccount) StorageName() string {
	return "Account"
}

func (sa *StorageAccount) MapKeyHasher() uint8 {
	return Blake2_128ConcatHasher
}

func (sa *StorageAccount) Fetch(blockStorage interfaces.BlockStorageT, key StorageAccountKey) (StorageAccountEntry, error) {
	val, err := GenericMapFetch[StorageAccount](blockStorage, key, sa)
	return val.Unwrap(), err
}

func (sa *StorageAccount) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageAccountEntry, error) {
	return GenericMapKeysFetch[StorageAccount, StorageAccountKey](blockStorage, sa)
}

//
//
//

type StorageBlockHashKey = uint32
type StorageBlockHashEntry = StorageEntry[StorageBlockHashKey, StorageBlockHash]

type StorageBlockHash struct {
	Value prim.H256
}

func (sbh *StorageBlockHash) PalletName() string {
	return PalletName
}

func (sbh *StorageBlockHash) StorageName() string {
	return "BlockHash"
}

func (sbh *StorageBlockHash) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sbh *StorageBlockHash) Fetch(blockStorage interfaces.BlockStorageT, key StorageBlockHashKey) (prim.Option[StorageBlockHashEntry], error) {
	return GenericMapFetch[StorageBlockHash](blockStorage, key, sbh)
}

func (sbh *StorageBlockHash) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageBlockHashEntry, error) {
	return GenericMapKeysFetch[StorageBlockHash, StorageBlockHashKey](blockStorage, sbh)
}

//
//
//

type StorageBlockWeight struct {
	Normal      Weight
	Operational Weight
	Mandatory   Weight
}

func (sbw *StorageBlockWeight) PalletName() string {
	return PalletName
}

func (sbw *StorageBlockWeight) StorageName() string {
	return "BlockWeight"
}

func (sbw *StorageBlockWeight) Fetch(blockStorage interfaces.BlockStorageT) (StorageBlockWeight, error) {
	return GenericFetchDefault[StorageBlockWeight](blockStorage, sbw)
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

func (sdbl *StorageDynamicBlockLength) PalletName() string {
	return PalletName
}

func (sdbl *StorageDynamicBlockLength) StorageName() string {
	return "DynamicBlockLength"
}

func (sdbl *StorageDynamicBlockLength) Fetch(blockStorage interfaces.BlockStorageT) (StorageDynamicBlockLength, error) {
	val, err := GenericFetch[StorageDynamicBlockLength](blockStorage, sdbl)

	// TODO Fallback might not be correct.
	// Fallback: 0x00003c0000005000000050000104010480
	return val.Unwrap(), err
}

//
//
//

type StorageEventCountValue = uint32
type StorageEventCount struct{}

func (sec *StorageEventCount) PalletName() string {
	return PalletName
}

func (sec *StorageEventCount) StorageName() string {
	return "EventCount"
}

func (sec *StorageEventCount) Fetch(blockStorage interfaces.BlockStorageT) (StorageEventCountValue, error) {
	return GenericFetchDefault[StorageEventCountValue](blockStorage, sec)
}

//
//
//

type StorageExtrinsicCountValue = uint32
type StorageExtrinsicCount struct{}

func (sec *StorageExtrinsicCount) PalletName() string {
	return PalletName
}

func (sec *StorageExtrinsicCount) StorageName() string {
	return "ExtrinsicCount"
}

func (sec *StorageExtrinsicCount) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageExtrinsicCountValue], error) {
	return GenericFetch[StorageExtrinsicCountValue](blockStorage, sec)
}

//
//
//

type StorageNumberValue = uint32
type StorageNumber struct{}

func (sn *StorageNumber) PalletName() string {
	return PalletName
}

func (sn *StorageNumber) StorageName() string {
	return "Number"
}

func (sn *StorageNumber) Fetch(blockStorage interfaces.BlockStorageT) (StorageNumberValue, error) {
	return GenericFetchDefault[StorageNumberValue](blockStorage, sn)
}

//
//
//

type StorageParentHashValue = prim.H256
type StorageParentHash struct{}

func (sph *StorageParentHash) PalletName() string {
	return PalletName
}

func (sph *StorageParentHash) StorageName() string {
	return "ParentHash"
}

func (sph *StorageParentHash) Fetch(blockStorage interfaces.BlockStorageT) (StorageParentHashValue, error) {
	return GenericFetchDefault[StorageParentHashValue](blockStorage, sph)
}
