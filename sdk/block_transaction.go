package sdk

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
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
		fmt.Println(err.Error())
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

func (this *BlockTransaction) MultiAddress() prim.Option[prim.MultiAddress] {
	signed := this.Signed()
	if signed.IsNone() {
		return prim.None[prim.MultiAddress]()
	}

	address := signed.Unwrap().Address

	return prim.Some(address)
}

func (this *BlockTransaction) AccountId() prim.Option[prim.AccountId] {
	multiMyb := this.MultiAddress()
	if multiMyb.IsNone() {
		return prim.None[prim.AccountId]()
	}

	multi := multiMyb.Unwrap()

	if multi.Id.IsNone() {
		return prim.None[prim.AccountId]()
	}

	return prim.Some(multi.Id.Unwrap())
}

func (this *BlockTransaction) SS58Address() prim.Option[string] {
	accountId := this.AccountId()
	if accountId.IsNone() {
		return prim.None[string]()
	}

	return prim.Some(accountId.Unwrap().ToHuman())
}

func (this *BlockTransaction) AppId() prim.Option[uint32] {
	signed := this.Signed()
	if signed.IsNone() {
		return prim.None[uint32]()
	}

	return prim.Some(signed.Unwrap().AppId)
}

func (this *BlockTransaction) Tip() prim.Option[metadata.Balance] {
	signed := this.Signed()
	if signed.IsNone() {
		return prim.None[metadata.Balance]()
	}

	return prim.Some(metadata.Balance{Value: signed.Unwrap().Tip})
}

func (this *BlockTransaction) Mortality() prim.Option[prim.Era] {
	signed := this.Signed()
	if signed.IsNone() {
		return prim.None[prim.Era]()
	}

	return prim.Some(signed.Unwrap().Era)
}

func (this *BlockTransaction) Nonce() prim.Option[uint32] {
	signed := this.Signed()
	if signed.IsNone() {
		return prim.None[uint32]()
	}

	return prim.Some(signed.Unwrap().Nonce)
}
