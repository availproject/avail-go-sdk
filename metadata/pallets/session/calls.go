package session

import (
	"github.com/availproject/avail-go-sdk/metadata"
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
