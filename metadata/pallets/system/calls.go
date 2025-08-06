package system

// Make some on-chain remark.
type CallRemark struct {
	Remark []byte
}

func (cr CallRemark) PalletIndex() uint8 {
	return PalletIndex
}

func (cr CallRemark) PalletName() string {
	return PalletName
}

func (cr CallRemark) CallIndex() uint8 {
	return 0
}

func (cr CallRemark) CallName() string {
	return "remark"
}

// Make some on-chain remark and emit event
type CallRemarkWithEvent struct {
	Remark []byte
}

func (crwe CallRemarkWithEvent) PalletIndex() uint8 {
	return PalletIndex
}

func (crwe CallRemarkWithEvent) PalletName() string {
	return PalletName
}

func (crwe CallRemarkWithEvent) CallIndex() uint8 {
	return 7
}

func (crwe CallRemarkWithEvent) CallName() string {
	return "remark_with_event"
}
