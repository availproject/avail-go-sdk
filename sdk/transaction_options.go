package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"math/big"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type TransactionOptions struct {
	AppId prim.Option[uint32]
	Nonce prim.Option[uint32]
}

func NewTransactionOptions() TransactionOptions {
	return TransactionOptions{
		AppId: prim.NewNone[uint32](),
		Nonce: prim.NewNone[uint32](),
	}
}

func (this TransactionOptions) WithAppId(value uint32) TransactionOptions {
	this.AppId = prim.NewSome(value)
	return this
}

func (this TransactionOptions) WithNonce(value uint32) TransactionOptions {
	this.Nonce = prim.NewSome(value)
	return this
}

func (this *TransactionOptions) ToPrimitive(client *Client, accountAddress string) (prim.Extra, prim.Additional, error) {
	genesisHash, err := client.Rpc.ChainSpec.V1GenesisHash()
	if err != nil {
		return prim.Extra{}, prim.Additional{}, err
	}
	forkHash, err := client.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
	if err != nil {
		return prim.Extra{}, prim.Additional{}, err
	}
	header, err := client.Rpc.Chain.GetHeader(prim.NewSome(forkHash))
	if err != nil {
		return prim.Extra{}, prim.Additional{}, err
	}
	forBlockNumber := header.Number

	runtimeVersion, err := client.Rpc.State.GetRuntimeVersion(prim.NewNone[prim.H256]())
	if err != nil {
		return prim.Extra{}, prim.Additional{}, err
	}

	additional := prim.Additional{
		SpecVersion: runtimeVersion.SpecVersion,
		TxVersion:   runtimeVersion.TxVersion,
		GenesisHash: genesisHash,
		ForkHash:    forkHash,
	}

	extra := prim.Extra{}
	extra.AppId = this.AppId.UnwrapOr(uint32(0))
	extra.Tip = uint128.FromBig(big.NewInt(0))
	if this.Nonce.IsNone() {
		extra.Nonce, err = client.Rpc.System.AccountNextIndex(accountAddress)
		if err != nil {
			return prim.Extra{}, prim.Additional{}, err
		}
	} else {
		extra.Nonce = this.Nonce.Unwrap()
	}
	extra.Era = prim.NewEra(32, uint64(forBlockNumber))

	return extra, additional, nil
}

func RegenerateEra(client *Client, extra *prim.Extra, additional *prim.Additional) error {
	forkHash, err := client.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
	if err != nil {
		return err
	}
	header, err := client.Rpc.Chain.GetHeader(prim.NewSome(forkHash))
	if err != nil {
		return err
	}

	additional.ForkHash = forkHash
	extra.Era = prim.NewEra(32, uint64(header.Number))

	return nil
}
