package metadata

import (
	primitives "github.com/availproject/avail-go-sdk/primitives"
)

type Payload struct {
	Call       primitives.Call
	palletName string
	callName   string
}

func NewPayload(call primitives.Call, palletName string, callName string) Payload {
	return Payload{
		Call:       call,
		palletName: palletName,
		callName:   callName,
	}
}

func (this *Payload) PalletName() string {
	return this.palletName
}

func (this *Payload) CallName() string {
	return this.callName
}
