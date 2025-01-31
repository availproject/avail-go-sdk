package data_availability

import ()

// Do not add, remove or change any of the field members.
type CallCreateApplicationKey struct {
	Key []uint8
}

func (this CallCreateApplicationKey) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallCreateApplicationKey) PalletName() string {
	return PalletName
}

func (this CallCreateApplicationKey) CallIndex() uint8 {
	return 0
}

func (this CallCreateApplicationKey) CallName() string {
	return "create_application_key"
}

// Do not add, remove or change any of the field members.
type CallSubmitData struct {
	Data []uint8
}

func (this CallSubmitData) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSubmitData) PalletName() string {
	return PalletName
}

func (this CallSubmitData) CallIndex() uint8 {
	return 1
}

func (this CallSubmitData) CallName() string {
	return "submit_data"
}
