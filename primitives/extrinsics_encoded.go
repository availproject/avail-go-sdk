package primitives

import (
	SType "github.com/itering/scale.go/types"
	"github.com/itering/scale.go/utiles"
)

// Hex and Scale encoded extrinsic that stats with "0x"
type EncodedExtrinsic struct {
	Value string
}

func (this *EncodedExtrinsic) Encode() string {
	return this.Value
}

func (this *EncodedExtrinsic) ToHex() string {
	return this.Value
}

func (this *EncodedExtrinsic) ToHexWith0x() string {
	return "0x" + this.ToHex()
}

func (this *EncodedExtrinsic) HexToBytes() []byte {
	return utiles.HexToBytes(this.Value)
}

func (this *EncodedExtrinsic) Decode(txIndex uint32) (DecodedExtrinsic, error) {
	return NewDecodedExtrinsic(*this, txIndex)
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
	if len(hexString) < 2 {
		panic("Not Possible")
	}
	if hexString[0] == '0' && hexString[1] == 'x' {
		hexString = hexString[2:]
	}

	return EncodedExtrinsic{
		Value: hexString,
	}
}
