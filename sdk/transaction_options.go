package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type TransactionOptions struct {
	AppId     prim.Option[uint32]
	Nonce     prim.Option[uint32]
	Mortality prim.Option[uint32]
	Tip       prim.Option[metadata.Balance]
}

func NewTransactionOptions() TransactionOptions {
	return TransactionOptions{
		AppId:     prim.None[uint32](),
		Nonce:     prim.None[uint32](),
		Mortality: prim.None[uint32](),
		Tip:       prim.None[metadata.Balance](),
	}
}

func (this TransactionOptions) WithAppId(value uint32) TransactionOptions {
	this.AppId = prim.Some(value)
	return this
}

func (this TransactionOptions) WithNonce(value uint32) TransactionOptions {
	this.Nonce = prim.Some(value)
	return this
}

func (this TransactionOptions) WithMortality(value uint32) TransactionOptions {
	this.Mortality = prim.Some(value)
	return this
}

func (this TransactionOptions) WithTip(value metadata.Balance) TransactionOptions {
	this.Tip = prim.Some(value)
	return this
}

func (this *TransactionOptions) ToPrimitive(client *Client, accountAddress string) (prim.Extra, prim.Additional, error) {
	forkHash, err := client.Rpc.Chain.GetFinalizedHead()
	if err != nil {
		return prim.Extra{}, prim.Additional{}, err
	}
	header, err := client.Rpc.Chain.GetHeader(prim.Some(forkHash))
	if err != nil {
		return prim.Extra{}, prim.Additional{}, err
	}
	forkBlockNumber := header.Number

	additional := prim.Additional{
		SpecVersion: client.RuntimeVersion.SpecVersion,
		TxVersion:   client.RuntimeVersion.TxVersion,
		GenesisHash: *client.GenesisHash,
		ForkHash:    forkHash,
	}

	extra := prim.Extra{}
	extra.AppId = this.AppId.UnwrapOr(uint32(0))
	extra.Tip = this.Tip.UnwrapOr(metadata.Balance{Value: uint128.Zero}).Value
	if this.Nonce.IsNone() {
		extra.Nonce, err = client.Rpc.System.AccountNextIndex(accountAddress)
		if err != nil {
			return prim.Extra{}, prim.Additional{}, err
		}
	} else {
		extra.Nonce = this.Nonce.Unwrap()
	}
	extra.Era = prim.NewEra(uint64(this.Mortality.UnwrapOr(32)), uint64(forkBlockNumber))

	return extra, additional, nil
}

func RegenerateEra(client *Client, extra *prim.Extra, additional *prim.Additional) error {
	forkHash, err := client.Rpc.Chain.GetBlockHash(prim.None[uint32]())
	if err != nil {
		return err
	}
	header, err := client.Rpc.Chain.GetHeader(prim.Some(forkHash))
	if err != nil {
		return err
	}

	additional.ForkHash = forkHash
	extra.Era = prim.NewEra(extra.Era.Period, uint64(header.Number))

	return nil
}
