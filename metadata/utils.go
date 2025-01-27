package metadata

import (
	"strings"

	"github.com/availproject/avail-go-sdk/interfaces"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

const Blake2_128ConcatHasher = uint8(0)
const Twox64ConcatHasher = uint8(1)

func removePalletStoragePrefix(key string) string {
	key = strings.TrimPrefix(key, "0x")
	return "0x" + key[64:]
}

func storageKeyEncode[S StorageT](storage S) string {
	return "0x" + prim.Hex.ToHex(prim.Crypto.TwoX128(storage.PalletName())) + prim.Hex.ToHex(prim.Crypto.TwoX128(storage.StorageName()))
}

func storageMapKeyEncode[S StorageMapT, K any](storage S, mapKey prim.Option[K]) string {
	encoded := "0x" + prim.Hex.ToHex(prim.Crypto.TwoX128(storage.PalletName())) + prim.Hex.ToHex(prim.Crypto.TwoX128(storage.StorageName()))
	if mapKey.IsSome() {
		keyEncoded := prim.Hex.FromHex(prim.Encoder.Encode(mapKey.Unwrap()))
		switch storage.MapKeyHasher() {
		case Blake2_128ConcatHasher:
			encoded += prim.Hex.ToHex(prim.Crypto.Blake2_128_Concat(keyEncoded))
		case Twox64ConcatHasher:
			encoded += prim.Hex.ToHex(prim.Crypto.Twox64Concat(keyEncoded))
		default:
			panic("Unknown hasher")
		}
	}

	return encoded
}

func storageDoubleMapKeyEncode[S StorageDoubleMapT, K1 any, K2 any](storage S, key1 prim.Option[K1], key2 prim.Option[K2]) string {
	encoded := "0x" + prim.Hex.ToHex(prim.Crypto.TwoX128(storage.PalletName())) + prim.Hex.ToHex(prim.Crypto.TwoX128(storage.StorageName()))
	if key1.IsSome() {
		keyEncoded := prim.Hex.FromHex(prim.Encoder.Encode(key1.Unwrap()))
		switch storage.MapKey1Hasher() {
		case Blake2_128ConcatHasher:
			encoded += prim.Hex.ToHex(prim.Crypto.Blake2_128_Concat(keyEncoded))
		case Twox64ConcatHasher:
			encoded += prim.Hex.ToHex(prim.Crypto.Twox64Concat(keyEncoded))
		default:
			panic("Unknown hasher")
		}
	}

	if key2.IsSome() {
		keyEncoded := prim.Hex.FromHex(prim.Encoder.Encode(key2.Unwrap()))
		switch storage.MapKey1Hasher() {
		case Blake2_128ConcatHasher:
			encoded += prim.Hex.ToHex(prim.Crypto.Blake2_128_Concat(keyEncoded))
		case Twox64ConcatHasher:
			encoded += prim.Hex.ToHex(prim.Crypto.Twox64Concat(keyEncoded))
		default:
			panic("Unknown hasher")
		}
	}

	return encoded
}

func storageMapKeyDecode[K any, S StorageMapT](storageKey string, storage S) (K, error) {
	withoutPalletPrefix := removePalletStoragePrefix(storageKey)
	input := prim.Hex.FromHex(withoutPalletPrefix)

	keyEncoded := []byte{}
	switch storage.MapKeyHasher() {
	case Blake2_128ConcatHasher:
		keyEncoded = prim.Crypto.DecodeBlake2_128Concat(input)
	case Twox64ConcatHasher:
		keyEncoded = prim.Crypto.DecodeTwox64Concat(input)
	default:
		panic("Unknown Hasher.")
	}

	var t K
	decoder := prim.NewDecoder(keyEncoded, 0)
	if err := decoder.Decode(&t); err != nil {
		return t, err
	}

	return t, nil
}

func storageDoubleMapKeyDecode[K1 any, K2 any, S StorageDoubleMapT](storageKey string, storage S) (K1, K2, error) {
	withoutPalletPrefix := removePalletStoragePrefix(storageKey)
	input := prim.Hex.FromHex(withoutPalletPrefix)

	decoder := prim.NewDecoder(input, 0)

	var t1 K1
	var t2 K2
	switch storage.MapKey1Hasher() {
	case Blake2_128ConcatHasher:
		// The first 16 bytes can be ignore
		if len(input) <= 16 {
			panic("Invalid Blake2_128Concat key format")
		}

		decoder.ScaleBytes.Offset += 16
		if err := decoder.Decode(&t1); err != nil {
			return t1, t2, err
		}
	case Twox64ConcatHasher:
		// The first 8 bytes can be ignore
		if len(input) <= 8 {
			panic("Invalid Twox64ConcatHasher key format")
		}

		decoder.ScaleBytes.Offset += 8
		if err := decoder.Decode(&t1); err != nil {
			return t1, t2, err
		}
	default:
		panic("Unknown Hasher.")
	}

	switch storage.MapKey2Hasher() {
	case Blake2_128ConcatHasher:
		// The first 16 bytes can be ignore
		if len(input) <= 16 {
			panic("Invalid Blake2_128Concat key format")
		}

		decoder.ScaleBytes.Offset += 16
		if err := decoder.Decode(&t2); err != nil {
			return t1, t2, err
		}
	case Twox64ConcatHasher:
		// The first 8 bytes can be ignore
		if len(input) <= 8 {
			panic("Invalid Twox64ConcatHasher key format")
		}

		decoder.ScaleBytes.Offset += 8
		if err := decoder.Decode(&t2); err != nil {
			return t1, t2, err
		}
	default:
		panic("Unknown Hasher.")
	}

	return t1, t2, nil
}

func GenericFetchDefault[V any, S StorageT](blockStorage interfaces.BlockStorageT, storage S) (V, error) {
	val, err := GenericFetch[V, S](blockStorage, storage)
	return val.Unwrap(), err
}

func GenericFetch[V any, S StorageT](blockStorage interfaces.BlockStorageT, storage S) (prim.Option[V], error) {
	storageKey := storageKeyEncode(storage)
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
	storageKey := storageMapKeyEncode(storage, prim.NewSome(mapKey))
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

func GenericMapKeysFetch[V any, K any, S StorageMapT](blockStorage interfaces.BlockStorageT, storage S) ([]StorageEntry[K, V], error) {
	storageKey := storageMapKeyEncode(storage, prim.NewNone[K]())
	storageKeys, err := blockStorage.FetchKeys(storageKey)
	if err != nil {
		return nil, err
	}

	if len(storageKeys) == 1 && storageKeys[0] == storageKey {
		return nil, nil
	}

	storageEntries := []StorageEntry[K, V]{}
	for i := range storageKeys {
		mapKey, err := storageMapKeyDecode[K](storageKeys[i], storage)
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

func GenericDoubleMapFetch[V any, K1 any, K2 any, S StorageDoubleMapT](blockStorage interfaces.BlockStorageT, mapKey1 K1, mapKey2 K2, storage S) (prim.Option[StorageEntryDoubleMap[K1, K2, V]], error) {
	storageKey := storageDoubleMapKeyEncode(storage, prim.NewSome(mapKey1), prim.NewSome(mapKey2))
	encoded, err := blockStorage.Fetch(storageKey)
	if err != nil {
		return prim.NewNone[StorageEntryDoubleMap[K1, K2, V]](), err
	}

	if encoded == "" {
		return prim.NewNone[StorageEntryDoubleMap[K1, K2, V]](), nil
	}

	var t V
	decoder := prim.NewDecoder(prim.Hex.FromHex(encoded), 0)
	if err := decoder.Decode(&t); err != nil {
		return prim.NewNone[StorageEntryDoubleMap[K1, K2, V]](), err
	}

	res := prim.NewSome(StorageEntryDoubleMap[K1, K2, V]{Key1: mapKey1, Key2: mapKey2, Value: t})
	return res, nil
}

func GenericDoubleMapKeysFetch[V any, K1 any, K2 any, S StorageDoubleMapT](blockStorage interfaces.BlockStorageT, mapKey1 K1, storage S) ([]StorageEntryDoubleMap[K1, K2, V], error) {
	storageKey := storageDoubleMapKeyEncode(storage, prim.NewSome(mapKey1), prim.NewNone[K2]())
	storageKeys, err := blockStorage.FetchKeys(storageKey)
	if err != nil {
		return nil, err
	}

	if len(storageKeys) == 1 && storageKeys[0] == storageKey {
		return nil, nil
	}

	storageEntries := []StorageEntryDoubleMap[K1, K2, V]{}
	for i := range storageKeys {
		mapKey1, mapKey2, err := storageDoubleMapKeyDecode[K1, K2](storageKeys[i], storage)
		if err != nil {
			return nil, err
		}

		value, err := GenericDoubleMapFetch[V](blockStorage, mapKey1, mapKey2, storage)
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

type StorageDoubleMapT interface {
	PalletName() string
	StorageName() string
	MapKey1Hasher() uint8
	MapKey2Hasher() uint8
}

type StorageEntry[K any, V any] struct {
	Key   K
	Value V
}

type StorageEntryDoubleMap[K1 any, K2 any, V any] struct {
	Key1  K1
	Key2  K2
	Value V
}

type Tuple2[T0 any, T1 any] struct {
	T0 T0
	T1 T1
}

func NewTuple2[T0 any, T1 any](v1 T0, v2 T1) Tuple2[T0, T1] {
	return Tuple2[T0, T1]{T0: v1, T1: v2}
}
