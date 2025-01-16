package primitives

import (
	"errors"
	"fmt"

	SUtiles "github.com/itering/scale.go/utiles"
)

type H520 struct {
	Value [65]byte
}

func (this *H520) ToHex() string {
	return SUtiles.BytesToHex(this.Value[:])
}

func (this *H520) ToHexWith0x() string {
	return "0x" + this.ToHex()
}

func (this *H520) ToRpcParam() string {
	return "\"" + this.ToHexWith0x() + "\""
}

func NewH520FromHexString(hexString string) (H520, error) {
	value := SUtiles.HexToBytes(hexString)
	if len(value) != 65 {
		return H520{}, errors.New(fmt.Sprintf(`H520 expected length: %v, actual length: %v.`, 65, len(value)))
	}

	return H520{Value: [65]byte(value)}, nil
}

func NewH520FromByteSlice(array []byte) (H520, error) {
	if len(array) != 65 {
		return H520{}, errors.New(fmt.Sprintf(`H520 expected length: %v, actual length: %v.`, 65, len(array)))
	}

	return H520{Value: [65]byte(array)}, nil
}
