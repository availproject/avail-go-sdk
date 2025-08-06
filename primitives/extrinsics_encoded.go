package primitives

import (
	SType "github.com/itering/scale.go/types"
	"github.com/itering/scale.go/utiles"
)

// Hex and Scale encoded extrinsic that stats with "0x"
type EncodedExtrinsic struct {
	Value string
}

func (e *EncodedExtrinsic) Encode() string {
	return e.Value
}

func (e *EncodedExtrinsic) ToHex() string {
	return e.Value
}

func (e *EncodedExtrinsic) ToHexWith0x() string {
	return "0x" + e.ToHex()
}

func (e *EncodedExtrinsic) HexToBytes() []byte {
	return utiles.HexToBytes(e.Value)
}

func (e *EncodedExtrinsic) Decode(txIndex uint32) (DecodedExtrinsic, error) {
	return NewDecodedExtrinsic(*e, txIndex)
}

func NewEncodedExtrinsic(payloadExtra *AlreadyEncoded, payloadCall *AlreadyEncoded, address MultiAddress, signature MultiSignature) EncodedExtrinsic {
	encoded_inner := ""

	// "is signed" + transaction protocol version (4)
	encoded_inner += "84"

	// Attach Address from Signer
	address.EncodeTo(&encoded_inner)

	// Signature
	signature.EncodeTo(&encoded_inner)

	// Payload + Call
	payloadExtra.EncodeTo(&encoded_inner)
	payloadCall.EncodeTo(&encoded_inner)

	innerLength := SType.Encode("Compact<u32>", len(encoded_inner)/2)
	encoded := innerLength + encoded_inner
	return EncodedExtrinsic{
		Value: encoded,
	}
}

func NewEncodedExtrinsicFromHex(hexString string) EncodedExtrinsic {
	if len(hexString) >= 2 {
		if hexString[0] == '0' && hexString[1] == 'x' {
			hexString = hexString[2:]
		}
	}

	return EncodedExtrinsic{
		Value: hexString,
	}
}
