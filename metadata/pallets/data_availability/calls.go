package data_availability

// Do not add, remove or change any of the field members.
type CallCreateApplicationKey struct {
	Key []uint8
}

func (ccak CallCreateApplicationKey) PalletIndex() uint8 {
	return PalletIndex
}

func (ccak CallCreateApplicationKey) PalletName() string {
	return PalletName
}

func (ccak CallCreateApplicationKey) CallIndex() uint8 {
	return 0
}

func (ccak CallCreateApplicationKey) CallName() string {
	return "create_application_key"
}

// Do not add, remove or change any of the field members.
type CallSubmitData struct {
	Data []uint8
}

func (csd CallSubmitData) PalletIndex() uint8 {
	return PalletIndex
}

func (csd CallSubmitData) PalletName() string {
	return PalletName
}

func (csd CallSubmitData) CallIndex() uint8 {
	return 1
}

func (csd CallSubmitData) CallName() string {
	return "submit_data"
}
