package metadata

import (
	"github.com/availproject/avail-go-sdk/primitives"
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

func (p *Payload) PalletName() string {
	return p.palletName
}

func (p *Payload) CallName() string {
	return p.callName
}
