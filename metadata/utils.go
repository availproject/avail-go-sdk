package metadata

import (
	"strings"

	"go-sdk/interfaces"
	prim "go-sdk/primitives"
)

const Blake2_128ConcatHasher = uint8(0)
const Twox64ConcatHasher = uint8(1)

func removePalletStoragePrefix(key string) string {
	key = strings.TrimPrefix(key, "0x")
	return "0x" + key[64:]
}

func storageKeyEncode[S StorageT, K any](storage S, mapKey prim.Option[K], hasher uint8) string {
	encoded := "0x" + prim.Hex.ToHex(prim.TwoX128(storage.PalletName())) + prim.Hex.ToHex(prim.TwoX128(storage.StorageName()))
	if mapKey.IsSome() {
		keyEncoded := prim.Hex.FromHex(prim.Encoder.Encode(mapKey.Unwrap()))
		if hasher == Blake2_128ConcatHasher {
			encoded += prim.Hex.ToHex(prim.Blake2_128_Concat(keyEncoded))
		} else if hasher == Twox64ConcatHasher {
			encoded += prim.Hex.ToHex(prim.Twox64Concat(keyEncoded))
		} else {
			panic("Unknown hasher")
		}
	}

	return encoded
}

func storageKeyDecode[K any](storageKey string, hasher uint8) (K, error) {
	withoutPalletPrefix := removePalletStoragePrefix(storageKey)
	input := prim.Hex.FromHex(withoutPalletPrefix)

	input2 := []byte{}
	if hasher == Blake2_128ConcatHasher {
		input2 = prim.DecodeBlake2_128Concat(input)
	} else if hasher == Twox64ConcatHasher {
		input2 = prim.DecodeTwox64Concat(input)
	} else {
		panic("Unknown Hasher.")
	}

	var t K
	decoder := prim.NewDecoder(input2, 0)
	if err := decoder.Decode(&t); err != nil {
		return t, err
	}

	return t, nil
}

func GenericFetch[V any, S StorageT](blockStorage interfaces.BlockStorageT, storage S) (prim.Option[V], error) {
	storageKey := storageKeyEncode(storage, prim.NewNone[uint32](), 0)
	encoded, err := blockStorage.Fetch(storageKey)
	if err != nil {
		return prim.NewNone[V](), err
	}

	if encoded == "" {
		return prim.NewNone[V](), nil
	}

	var t V
	decoder := prim.NewDecoder(prim.Hex.FromHex(encoded), 0)
	if err := decoder.Decode(&t); err != nil {
		return prim.NewNone[V](), err
	}

	return prim.NewSome(t), nil
}

func GenericMapFetch[V any, K any, S StorageMapT](blockStorage interfaces.BlockStorageT, mapKey K, storage S) (prim.Option[StorageEntry[K, V]], error) {
	storageKey := storageKeyEncode(storage, prim.NewSome(mapKey), storage.MapKeyHasher())
	encoded, err := blockStorage.Fetch(storageKey)
	if err != nil {
		return prim.NewNone[StorageEntry[K, V]](), err
	}

	if encoded == "" {
		return prim.NewNone[StorageEntry[K, V]](), nil
	}

	var t V
	decoder := prim.NewDecoder(prim.Hex.FromHex(encoded), 0)
	if err := decoder.Decode(&t); err != nil {
		return prim.NewNone[StorageEntry[K, V]](), err
	}

	res := prim.NewSome(StorageEntry[K, V]{Key: mapKey, Value: t})
	return res, nil
}

func GenericFetchKeys[V any, K any, S StorageMapT](blockStorage interfaces.BlockStorageT, storage S) ([]StorageEntry[K, V], error) {
	storageKey := storageKeyEncode(storage, prim.NewNone[K](), storage.MapKeyHasher())
	storageKeys, err := blockStorage.FetchKeys(storageKey)
	if err != nil {
		return nil, err
	}

	if len(storageKeys) == 1 && storageKeys[0] == storageKey {
		return nil, nil
	}

	storageEntries := []StorageEntry[K, V]{}
	for _, storageKey := range storageKeys {
		mapKey, err := storageKeyDecode[K](storageKey, storage.MapKeyHasher())
		if err != nil {
			return nil, err
		}

		value, err := GenericMapFetch[V](blockStorage, mapKey, storage)
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

type StorageT interface {
	PalletName() string
	StorageName() string
}

type StorageMapT interface {
	PalletName() string
	StorageName() string
	MapKeyHasher() uint8
}

type StorageEntry[K any, V any] struct {
	Key   K
	Value V
}
