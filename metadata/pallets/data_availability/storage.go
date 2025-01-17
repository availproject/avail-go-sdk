package data_availability

import (
	"go-sdk/interfaces"
	meta "go-sdk/metadata"
	prim "go-sdk/primitives"
	"strings"
)

type StorageNextAppId struct {
	Value uint32
}

type StorageAppKeysEntry = meta.StorageEntry[[]byte, StorageAppKeys]

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

func (this *StorageAppKeys) Encode(key prim.Option[[]byte]) string {
	encoded := "0x"
	encoded += prim.Hex.ToHex(prim.TwoX128(this.PalletName()))
	encoded += prim.Hex.ToHex(prim.TwoX128(this.StorageName()))
	if key.IsSome() {
		keyEncoded := prim.Hex.FromHex(prim.Encoder.Encode(key.Unwrap()))
		encoded += prim.Hex.ToHex(prim.Blake2_128_Concat(keyEncoded))
	}

	return encoded
}

func (this *StorageAppKeys) Fetch(blockStorage interfaces.BlockStorageT, key []byte) (StorageAppKeysEntry, error) {
	storageEntryKey := this.Encode(prim.NewSome(key))
	println("storageEntryKey: ", storageEntryKey)

	val, err := genericFetch[StorageAppKeys](blockStorage, storageEntryKey)
	if err != nil {
		return StorageAppKeysEntry{}, err
	}

	res := StorageAppKeysEntry{Key: key, Value: val}
	return res, nil
}

func (this *StorageAppKeys) FetchKeys(blockStorage interfaces.BlockStorageT) ([]StorageAppKeysEntry, error) {
	storageEntryKey := this.Encode(prim.NewNone[[]byte]())
	println("storageEntryKey: ", storageEntryKey)

	keys, err := blockStorage.FetchKeys(storageEntryKey)
	if err != nil {
		return nil, err
	}

	storageEntries := []StorageAppKeysEntry{}
	for _, key := range keys {
		input := removePalletStoragePrefix(key)
		input2 := prim.Hex.FromHex(input)
		input3 := prim.DecodeBlake2_128Concat(input2)

		t := []byte{}
		decoder := prim.NewDecoder(input3, 0)
		if err := decoder.Decode(&t); err != nil {
			panic(err)
		}
		println("Key: ", string(t))

		value, err := this.Fetch(blockStorage, t)
		if err != nil {
			return nil, err
		}
		storageEntries = append(storageEntries, value)
	}

	return storageEntries, nil
}

func genericFetch[T any](blockStorage interfaces.BlockStorageT, key string) (T, error) {
	encoded, err := blockStorage.Fetch(key)
	if err != nil {
		var t T
		return t, err
	}

	println("Encoded: ", encoded)

	var t T
	decoder := prim.NewDecoder(prim.Hex.FromHex(encoded), 0)
	if err := decoder.Decode(&t); err != nil {
		return t, err
	}

	return t, nil
}

func removePalletStoragePrefix(key string) string {
	key = strings.TrimPrefix(key, "0x")
	return "0x" + key[64:]
}
