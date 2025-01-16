package primitives

import (
	"github.com/vedhavyas/go-subkey/v2"
)

func CreateSigned(call Call, extra Extra, additional Additional, kp subkey.KeyPair) EncodedExtrinsic {
	unsignedPayload := UnsignedPayload{
		Call:       call,
		Extra:      extra,
		Additional: additional,
	}
	unsignedEncodedPayload := unsignedPayload.Encode()

	rawSignature, err := unsignedEncodedPayload.Sign(kp)
	if err != nil {
		panic(err)
	}

	accountId := NewH256FromByteSlice(kp.AccountID())
	signature := NewH512FromByteSlice(rawSignature)
	multiAddress := NewMultiAddressId(accountId)
	multiSignature := NewMultiSignatureSr(signature)
	encodedTransaction := NewEncodedExtrinsic(&unsignedEncodedPayload.Extra, &unsignedEncodedPayload.Call, multiAddress, multiSignature)
	return encodedTransaction

}
