package pallets

import (
	meta "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type CallT interface {
	PalletIndex() uint8
	PalletName() string
	CallIndex() uint8
	CallName() string
}

func ToCall[T CallT](call T) prim.Call {
	return prim.Call{
		PalletIndex: call.PalletIndex(),
		CallIndex:   call.CallIndex(),
		Fields:      prim.AlreadyEncoded{Value: prim.Encoder.Encode(call)},
	}
}

func ToPayload[T CallT](call T) meta.Payload {
	return meta.NewPayload(ToCall(call), call.PalletName(), call.CallName())
}

func Decode[T CallT](this T, tx *prim.DecodedExtrinsic) bool {
	if this.PalletIndex() != tx.Call.PalletIndex {
		return false
	}

	if this.CallIndex() != tx.Call.CallIndex {
		return false
	}

	var decoder = prim.NewDecoder(tx.Call.Fields.ToBytes(), 0)
	if err := decoder.Decode(this); err != nil {
		return false
	}
	return true
}
