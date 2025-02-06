package primitives

import (
	"errors"
	"fmt"
)

// Do not change the order of field members.
type MultiSignature struct {
	Ed25519 Option[H512]
	Sr25519 Option[H512]
	Ecdsa   Option[H520]
}

func emptyMultiSignature() MultiSignature {
	return MultiSignature{
		Ed25519: None[H512](),
		Sr25519: None[H512](),
		Ecdsa:   None[H520](),
	}
}

func NewMultiSignatureEd(value H512) MultiSignature {
	signature := emptyMultiSignature()
	signature.Ed25519 = Some(value)
	return signature
}

func NewMultiSignatureSr(value H512) MultiSignature {
	signature := emptyMultiSignature()
	signature.Sr25519 = Some(value)
	return signature
}

func NewMultiSignatureEcdsa(value H520) MultiSignature {
	signature := emptyMultiSignature()
	signature.Ecdsa = Some(value)
	return signature
}

func (this *MultiSignature) EncodeTo(dest *string) {
	if this.Ed25519.IsSome() {
		Encoder.EncodeTo(uint8(0), dest)
		Encoder.EncodeTo(this.Ed25519.Unwrap(), dest)
	} else if this.Sr25519.IsSome() {
		Encoder.EncodeTo(uint8(1), dest)
		Encoder.EncodeTo(this.Sr25519.Unwrap(), dest)
	} else if this.Ecdsa.IsSome() {
		Encoder.EncodeTo(uint8(2), dest)
		Encoder.EncodeTo(this.Ecdsa.Unwrap(), dest)
	} else {
		panic("Something Went Wrong with MultiSignature EncodeTo")
	}

}

func (this *MultiSignature) Decode(decoder *Decoder) error {
	result := emptyMultiSignature()
	variantIndex := uint8(0)
	decoder.Decode(&variantIndex)
	if variantIndex == 0 {
		value := H512{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		result.Ed25519 = Some(value)
	} else if variantIndex == 1 {
		value := H512{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		result.Sr25519 = Some(value)
	} else if variantIndex == 2 {
		value := H520{}
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		result.Ecdsa = Some(value)
	} else {
		return errors.New(fmt.Sprintf(`MultiSignature Decode failure. Unknown Variant index: %v`, variantIndex))
	}

	*this = result
	return nil
}
