package data_availability

import (
	"go-sdk/interfaces"
	meta "go-sdk/metadata"
	prim "go-sdk/primitives"
	"strings"

	"github.com/itering/scale.go/utiles/uint128"
)

type StorageNextAppId struct {
	Value uint32
}

type StorageAppKeysEntry = meta.StorageEntry[[]byte, StorageAppKeys]
type StorageAccountEntry = meta.StorageEntry[meta.AccountId, StorageAccount]

type StorageAccount struct {
	Nonce       uint32
	Consumers   uint32
	Providers   uint32
	Sufficients uint32
	AccountData AccountData
}

type AccountData struct {
	Free     meta.Balance
	Reserved meta.Balance
	Frozen   meta.Balance
	Flags    uint128.Uint128
}

func (this *StorageAccount) PalletName() string {
	return "System"
}

func (this *StorageAccount) StorageName() string {
	return "Account"
}

func (this *StorageAccount) EncodeKey(key prim.Option[meta.AccountId]) string {
	return storageKeyEncode(this, key)
}

func (this *StorageAccount) DecodeKey(key string) (meta.AccountId, error) {
	return storageKeyDecode[meta.AccountId](key)
}

func (this *StorageAccount) Fetch(blockStorage interfaces.BlockStorageT, key meta.AccountId) (prim.Option[StorageAccountEntry], error) {
	val, err := GenericFetch[StorageAccount](blockStorage, this.EncodeKey(prim.NewSome(key)), key)
	if err != nil {
		return prim.NewNone[StorageAccountEntry](), err
	}

	return val, nil
}

func (this *StorageAccount) FetchKeys(blockStorage interfaces.BlockStorageT) ([]StorageAccountEntry, error) {
	storageKeys, err := blockStorage.FetchKeys(this.EncodeKey(prim.NewNone[meta.AccountId]()))
	if err != nil {
		return nil, err
	}

	storageEntries := []StorageAccountEntry{}
	for _, storageKey := range storageKeys {
		mapKey, err := this.DecodeKey(storageKey)
		if err != nil {
			panic(err)
		}

		value, err := this.Fetch(blockStorage, mapKey)
		if err != nil {
			panic(err)
		}

		if value.IsNone() {
			continue
		}

		storageEntries = append(storageEntries, value.Unwrap())
	}

	return storageEntries, nil
}

// /
// /
// /
type StorageAppKeys struct {
	Owner meta.AccountId
	AppId uint32 `scale:"compact"`
}

func (this *StorageAppKeys) PalletName() string {
	return PalletName
}

func (this *StorageAppKeys) StorageName() string {
	return "AppKeys"
}

func (this *StorageAppKeys) EncodeKey(key prim.Option[[]byte]) string {
	return storageKeyEncode(this, key)
}

func (this *StorageAppKeys) DecodeKey(key string) ([]byte, error) {
	return storageKeyDecode[[]byte](key)
}

func (this *StorageAppKeys) Fetch(blockStorage interfaces.BlockStorageT, key []byte) (prim.Option[StorageAppKeysEntry], error) {
	val, err := GenericFetch[StorageAppKeys](blockStorage, this.EncodeKey(prim.NewSome(key)), key)
	if err != nil {
		return prim.NewNone[StorageAppKeysEntry](), err
	}

	return val, nil
}

func (this *StorageAppKeys) FetchKeys(blockStorage interfaces.BlockStorageT) ([]StorageAppKeysEntry, error) {
	storageKeys, err := blockStorage.FetchKeys(this.EncodeKey(prim.NewNone[[]byte]()))
	if err != nil {
		return nil, err
	}

	storageEntries := []StorageAppKeysEntry{}
	for _, storageKey := range storageKeys {
		mapKey, err := this.DecodeKey(storageKey)
		if err != nil {
			return nil, err
		}

		value, err := this.Fetch(blockStorage, mapKey)
		if err != nil {
			return nil, err
		}

		if value.IsNone() {
			continue
		}

		storageEntries = append(storageEntries, value.Unwrap())
	}

	return storageEntries, nil
}

func GenericFetch[T any, V any](blockStorage interfaces.BlockStorageT, storageKey string, mapKey V) (prim.Option[meta.StorageEntry[V, T]], error) {
	encoded, err := blockStorage.Fetch(storageKey)
	if err != nil {
		return prim.NewNone[meta.StorageEntry[V, T]](), err
	}

	if encoded == "" {
		return prim.NewNone[meta.StorageEntry[V, T]](), nil
	}

	var t T
	decoder := prim.NewDecoder(prim.Hex.FromHex(encoded), 0)
	if err := decoder.Decode(&t); err != nil {
		return prim.NewNone[meta.StorageEntry[V, T]](), err
	}

	res := prim.NewSome(meta.StorageEntry[V, T]{Key: mapKey, Value: t})
	return res, nil
}

func removePalletStoragePrefix(key string) string {
	key = strings.TrimPrefix(key, "0x")
	return "0x" + key[64:]
}

func storageKeyEncode[T StorageT, V any](pallet T, mapKey prim.Option[V]) string {
	encoded := "0x" + prim.Hex.ToHex(prim.TwoX128(pallet.PalletName())) + prim.Hex.ToHex(prim.TwoX128(pallet.StorageName()))
	if mapKey.IsSome() {
		keyEncoded := prim.Hex.FromHex(prim.Encoder.Encode(mapKey.Unwrap()))
		encoded += prim.Hex.ToHex(prim.Blake2_128_Concat(keyEncoded))
	}

	return encoded
}

func storageKeyDecode[T any](storageKey string) (T, error) {
	input := prim.Hex.FromHex(removePalletStoragePrefix(storageKey))
	input2 := prim.DecodeBlake2_128Concat(input)

	var t T
	decoder := prim.NewDecoder(input2, 0)
	if err := decoder.Decode(&t); err != nil {
		return t, err
	}

	return t, nil
}

type StorageT interface {
	PalletName() string
	StorageName() string
}
