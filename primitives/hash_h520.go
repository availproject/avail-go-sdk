package primitives

import (
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

func NewH520FromHexString(hexString string) H520 {
	value := SUtiles.HexToBytes(hexString)
	if len(value) != 65 {
		panic("Uppps it's not 65")
	}

	return H520{Value: [65]byte(value)}
}

func NewH520FromByteSlice(array []byte) H520 {
	if len(array) != 65 {
		panic("Byte Slice for H520 needs to be 65 elements long.")
	}

	return H520{Value: [65]byte(array)}
}
