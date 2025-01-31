package system

import ()

// Make some on-chain remark.
type CallRemark struct {
	Remark []byte
}

func (this CallRemark) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallRemark) PalletName() string {
	return PalletName
}

func (this CallRemark) CallIndex() uint8 {
	return 0
}

func (this CallRemark) CallName() string {
	return "remark"
}

// Make some on-chain remark and emit event
type CallRemarkWithEvent struct {
	Remark []byte
}

func (this CallRemarkWithEvent) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallRemarkWithEvent) PalletName() string {
	return PalletName
}

func (this CallRemarkWithEvent) CallIndex() uint8 {
	return 7
}

func (this CallRemarkWithEvent) CallName() string {
	return "remark_with_event"
}
