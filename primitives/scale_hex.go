package primitives

import (
	SUtiles "github.com/itering/scale.go/utiles"
)

type hexT struct{}

var Hex hexT

func (hexT) ToHex(arr []byte) string {
	return SUtiles.BytesToHex(arr)
}

func (hexT) FromHex(arr string) []byte {
	return SUtiles.HexToBytes(arr)
}

func (hexT) ToHexWith0x(arr []byte) string {
	return "0x" + Hex.ToHex(arr)
}
