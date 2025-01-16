package primitives

import (
	SUtiles "github.com/itering/scale.go/utiles"
)

func ToHex(arr []byte) string {
	return SUtiles.BytesToHex(arr)
}

func FromHex(arr string) []byte {
	return SUtiles.HexToBytes(arr)
}

func ToHexWith0x(arr []byte) string {
	return "0x" + ToHex(arr)
}
