package primitives

import (
	SUtiles "github.com/itering/scale.go/utiles"
)

type H256 struct {
	Value [32]byte
}

func (this H256) ToHex() string {
	return SUtiles.BytesToHex(this.Value[:])
}

func (this H256) ToHexWith0x() string {
	return "0x" + this.ToHex()
}

func (this H256) ToRpcParam() string {
	return "\"" + this.ToHexWith0x() + "\""
}

func NewH256FromHexString(hexString string) H256 {
	value := SUtiles.HexToBytes(hexString)
	if len(value) != 32 {
		panic("Uppps it's not 32")
	}

	return H256{Value: [32]byte(value)}
}

func NewH256FromByteSlice(array []byte) H256 {
	if len(array) != 32 {
		panic("Byte Slice for H256 needs to be 32 elements long.")
	}

	return H256{Value: [32]byte(array)}
}
