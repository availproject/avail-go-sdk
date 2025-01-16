package primitives

import (
	"errors"
	"fmt"

	SUtiles "github.com/itering/scale.go/utiles"
)

type H512 struct {
	Value [64]byte
}

func (this *H512) ToHex() string {
	return SUtiles.BytesToHex(this.Value[:])
}

func (this *H512) ToHexWith0x() string {
	return "0x" + this.ToHex()
}

func (this *H512) ToRpcParam() string {
	return "\"" + this.ToHexWith0x() + "\""
}

func NewH512FromHexString(hexString string) (H512, error) {
	value := SUtiles.HexToBytes(hexString)
	if len(value) != 64 {
		return H512{}, errors.New(fmt.Sprintf(`H512 expected length: %v, actual length: %v.`, 64, len(value)))
	}

	return H512{Value: [64]byte(value)}, nil
}

func NewH512FromByteSlice(array []byte) (H512, error) {
	if len(array) != 64 {
		return H512{}, errors.New(fmt.Sprintf(`H512 expected length: %v, actual length: %v.`, 64, len(array)))
	}

	return H512{Value: [64]byte(array)}, nil
}
