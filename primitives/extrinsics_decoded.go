package primitives

import (
	"errors"

	"github.com/itering/scale.go/utiles/uint128"
	"golang.org/x/crypto/blake2b"
)

type DecodedExtrinsic struct {
	Call    Call
	TxHash  H256
	TxIndex uint32
	Signed  Option[DecodedExtrinsicSigned]
}

type DecodedExtrinsicSigned struct {
	Address   MultiAddress
	Signature MultiSignature
	Nonce     uint32
	AppId     uint32
	Tip       uint128.Uint128
	Era       Era
}

func NewDecodedExtrinsic(extrinsic EncodedExtrinsic, txIndex uint32) (DecodedExtrinsic, error) {
	decodedData := extrinsic.HexToBytes()
	txHashArray := blake2b.Sum256(decodedData)
	txHash, err := NewH256FromByteSlice(txHashArray[:])
	if err != nil {
		return DecodedExtrinsic{}, err
	}

	totalLength := len(decodedData)
	signedPart := None[DecodedExtrinsicSigned]()

	// Reading Transaction Length
	decoder := NewDecoder(decodedData, 0)
	txLength := CompactU32{}
	if err := decoder.Decode(&txLength); err != nil {
		return DecodedExtrinsic{}, err
	}

	if totalLength != int(txLength.Value)+decoder.Offset() {
		return DecodedExtrinsic{}, errors.New("remainingLength is not equal to txLength + scaleDecoder.Data.Offset")
	}

	// Checking if the message is signed
	signed := uint8(0)
	if err := decoder.Decode(&signed); err != nil {
		return DecodedExtrinsic{}, err
	}
	// 132 is 0x84
	if signed == 132 {
		multiAddress := MultiAddress{}
		if err := decoder.Decode(&multiAddress); err != nil {
			return DecodedExtrinsic{}, err
		}
		multiSignature := MultiSignature{}
		if err := decoder.Decode(&multiSignature); err != nil {
			return DecodedExtrinsic{}, err
		}
		extra := Extra{}
		if err := decoder.Decode(&extra); err != nil {
			return DecodedExtrinsic{}, err
		}

		signedData := DecodedExtrinsicSigned{
			Address:   multiAddress,
			Signature: multiSignature,
			Nonce:     extra.Nonce,
			AppId:     extra.AppId,
			Tip:       extra.Tip,
			Era:       extra.Era,
		}

		signedPart.Set(signedData)
	}

	call := Call{}
	if err := decoder.Decode(&call); err != nil {
		return DecodedExtrinsic{}, err
	}

	return DecodedExtrinsic{
		Call:    call,
		TxHash:  txHash,
		TxIndex: txIndex,
		Signed:  signedPart,
	}, nil
}
