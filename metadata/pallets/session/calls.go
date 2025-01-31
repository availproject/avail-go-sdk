package session

import (
	"github.com/availproject/avail-go-sdk/metadata"
	. "github.com/availproject/avail-go-sdk/metadata/pallets"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

// Sets the session key(s) of the function caller to `keys`.
// Allows an account to set its session key prior to becoming a validator.
// This doesn't take effect until the next session.
type CallSetKeys struct {
	Keys  metadata.SessionKeys
	Proof []byte
}

func (this CallSetKeys) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetKeys) PalletName() string {
	return PalletName
}

func (this CallSetKeys) CallIndex() uint8 {
	return 0
}

func (this CallSetKeys) CallName() string {
	return "set_keys"
}

func (this *CallSetKeys) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallSetKeys) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallSetKeys) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}

// Removes any session key(s) of the function caller.
//
// This doesn't take effect until the next session.
//
// The dispatch origin of this function must be Signed and the account must be either be
// convertible to a validator ID using the chain's typical addressing system (this usually
// means being a controller account) or directly convertible into a validator ID (which
// usually means being a stash account).
type CallPurgeKeys struct{}

func (this CallPurgeKeys) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallPurgeKeys) PalletName() string {
	return PalletName
}

func (this CallPurgeKeys) CallIndex() uint8 {
	return 1
}

func (this CallPurgeKeys) CallName() string {
	return "purge_keys"
}

func (this *CallPurgeKeys) ToCall() prim.Call {
	return ToCall(this)
}

func (this *CallPurgeKeys) ToPayload() metadata.Payload {
	return ToPayload(this)
}

func (this *CallPurgeKeys) DecodeExtrinsic(tx *prim.DecodedExtrinsic) bool {
	return Decode(this, tx)
}
