package primitives

import (
	"errors"
	"fmt"

	SUtiles "github.com/itering/scale.go/utiles"
)

type H256 struct {
	Value [32]byte
}

func (h H256) ToHex() string {
	return SUtiles.BytesToHex(h.Value[:])
}

func (h H256) ToHexWith0x() string {
	return "0x" + h.ToHex()
}

func (h H256) ToHuman() string {
	return h.ToHexWith0x()
}

func (h H256) ToString() string {
	return h.ToHexWith0x()
}

func (h H256) String() string {
	return h.ToHexWith0x()
}

func (h H256) ToRpcParam() string {
	return "\"" + h.ToHexWith0x() + "\""
}

func NewH256FromHexString(hexString string) (H256, error) {
	value := SUtiles.HexToBytes(hexString)
	if len(value) != 32 {
		return H256{}, errors.New(fmt.Sprintf(`H256 expected length: %v, actual length: %v.`, 32, len(value)))
	}

	return H256{Value: [32]byte(value)}, nil
}

func NewH256FromByteSlice(array []byte) (H256, error) {
	if len(array) != 32 {
		return H256{}, errors.New(fmt.Sprintf(`H256 expected length: %v, actual length: %v.`, 32, len(array)))
	}

	return H256{Value: [32]byte(array)}, nil
}

func NewBlockHashFromHexString(hexString string) (H256, error) {
	return NewH256FromHexString(hexString)
}

func NewBlockHashFromByteSlice(array []byte) (H256, error) {
	return NewH256FromByteSlice(array)
}
