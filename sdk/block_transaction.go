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

func (b *BlockTransaction) PalletName() string {
	return b.palletName
}

func (b *BlockTransaction) CallName() string {
	return b.callName
}

func (b *BlockTransaction) PalletIndex() uint8 {
	return b.Extrinsic.Call.PalletIndex
}

func (b *BlockTransaction) CallIndex() uint8 {
	return b.Extrinsic.Call.CallIndex
}

func (b *BlockTransaction) TxHash() prim.H256 {
	return b.Extrinsic.TxHash
}

func (b *BlockTransaction) TxIndex() uint32 {
	return b.Extrinsic.TxIndex
}

func (b *BlockTransaction) Signed() prim.Option[prim.DecodedExtrinsicSigned] {
	return b.Extrinsic.Signed
}

func (b *BlockTransaction) Fields() prim.AlreadyEncoded {
	return b.Extrinsic.Call.Fields
}

func (b *BlockTransaction) Events() prim.Option[EventRecords] {
	return b.events
}

func (b *BlockTransaction) MultiAddress() prim.Option[prim.MultiAddress] {
	signed := b.Signed()
	if signed.IsNone() {
		return prim.None[prim.MultiAddress]()
	}

	address := signed.Unwrap().Address

	return prim.Some(address)
}

func (b *BlockTransaction) AccountId() prim.Option[prim.AccountId] {
	multiMyb := b.MultiAddress()
	if multiMyb.IsNone() {
		return prim.None[prim.AccountId]()
	}

	multi := multiMyb.Unwrap()

	if multi.Id.IsNone() {
		return prim.None[prim.AccountId]()
	}

	return prim.Some(multi.Id.Unwrap())
}

func (b *BlockTransaction) SS58Address() prim.Option[string] {
	accountId := b.AccountId()
	if accountId.IsNone() {
		return prim.None[string]()
	}

	return prim.Some(accountId.Unwrap().ToHuman())
}

func (b *BlockTransaction) AppId() prim.Option[uint32] {
	signed := b.Signed()
	if signed.IsNone() {
		return prim.None[uint32]()
	}

	return prim.Some(signed.Unwrap().AppId)
}

func (b *BlockTransaction) Tip() prim.Option[metadata.Balance] {
	signed := b.Signed()
	if signed.IsNone() {
		return prim.None[metadata.Balance]()
	}

	return prim.Some(metadata.Balance{Value: signed.Unwrap().Tip})
}

func (b *BlockTransaction) Mortality() prim.Option[prim.Era] {
	signed := b.Signed()
	if signed.IsNone() {
		return prim.None[prim.Era]()
	}

	return prim.Some(signed.Unwrap().Era)
}

func (b *BlockTransaction) Nonce() prim.Option[uint32] {
	signed := b.Signed()
	if signed.IsNone() {
		return prim.None[uint32]()
	}

	return prim.Some(signed.Unwrap().Nonce)
}
