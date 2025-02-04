package primitives

import (
	"github.com/vedhavyas/go-subkey/v2"
)

func CreateSigned(call Call, extra Extra, additional Additional, kp subkey.KeyPair) (EncodedExtrinsic, error) {
	unsignedPayload := UnsignedPayload{
		Call:       call,
		Extra:      extra,
		Additional: additional,
	}
	unsignedEncodedPayload := unsignedPayload.Encode()

	rawSignature, err := unsignedEncodedPayload.Sign(kp)
	if err != nil {
		return EncodedExtrinsic{}, err
	}

	accountId, err := NewH256FromByteSlice(kp.AccountID())
	if err != nil {
		return EncodedExtrinsic{}, err
	}

	signature, err := NewH512FromByteSlice(rawSignature)
	if err != nil {
		return EncodedExtrinsic{}, err
	}
	multiAddress := NewMultiAddressId(AccountId{Value: accountId})
	multiSignature := NewMultiSignatureSr(signature)
	encodedTransaction := NewEncodedExtrinsic(&unsignedEncodedPayload.Extra, &unsignedEncodedPayload.Call, multiAddress, multiSignature)
	return encodedTransaction, nil

}
