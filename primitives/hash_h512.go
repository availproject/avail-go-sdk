package primitives

import (
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

func NewH512FromHexString(hexString string) H512 {
	value := SUtiles.HexToBytes(hexString)
	if len(value) != 64 {
		panic("Uppps it's not 64")
	}

	return H512{Value: [64]byte(value)}
}

func NewH512FromByteSlice(array []byte) H512 {
	if len(array) != 64 {
		panic("Byte Slice for H512 needs to be 64 elements long.")
	}

	return H512{Value: [64]byte(array)}
}
