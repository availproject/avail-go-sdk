package sdk

import (
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type BlockTransaction struct {
	client     *Client
	Extrinsic  *prim.DecodedExtrinsic
	palletName string
	callName   string
	events     prim.Option[EventRecords]
}

func NewBlockTransaction(client *Client, extrinsic *prim.DecodedExtrinsic, events prim.Option[EventRecords]) BlockTransaction {
	palletName, callName, err := client.Metadata().PalletCallName(extrinsic.Call.PalletIndex, extrinsic.Call.CallIndex)
	if err != nil {
		println(err.Error())
	}

	return BlockTransaction{
		client:     client,
		Extrinsic:  extrinsic,
		palletName: palletName,
		callName:   callName,
		events:     events,
	}
}

func (this *BlockTransaction) PalletName() string {
	return this.palletName
}

func (this *BlockTransaction) CallName() string {
	return this.callName
}

func (this *BlockTransaction) PalletIndex() uint8 {
	return this.Extrinsic.Call.PalletIndex
}

func (this *BlockTransaction) CallIndex() uint8 {
	return this.Extrinsic.Call.CallIndex
}

func (this *BlockTransaction) TxHash() prim.H256 {
	return this.Extrinsic.TxHash
}

func (this *BlockTransaction) TxIndex() uint32 {
	return this.Extrinsic.TxIndex
}

func (this *BlockTransaction) Signed() prim.Option[prim.DecodedExtrinsicSigned] {
	return this.Extrinsic.Signed
}

func (this *BlockTransaction) Fields() prim.AlreadyEncoded {
	return this.Extrinsic.Call.Fields
}

func (this *BlockTransaction) Events() prim.Option[EventRecords] {
	return this.events
}
