package sdk

import (
	"fmt"
	"log"
	"time"

	"github.com/availproject/avail-go-sdk/src/extrinsic"
	"github.com/availproject/avail-go-sdk/src/rpc"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func CreateExtrinsic(api *SubstrateAPI, ext_call string, keyring signature.KeyringPair, AppID int, arg ...interface{}) (extrinsic.Extrinsic, error) {
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return extrinsic.Extrinsic{}, err
	}
	call, err := types.NewCall(meta, ext_call, arg...)
	if err != nil {
		return extrinsic.Extrinsic{}, err
	}
	ext := extrinsic.NewExtrinsic(call)
	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return extrinsic.Extrinsic{}, err
	}
	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return extrinsic.Extrinsic{}, err
	}
	key, err := types.CreateStorageKey(meta, "System", "Account", keyring.PublicKey)
	if err != nil {
		return extrinsic.Extrinsic{}, err
	}

	var accountInfo types.AccountInfo
	ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		return extrinsic.Extrinsic{}, err
	}
	nonce := uint32(accountInfo.Nonce)
	options := extrinsic.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                extrinsic.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(100),
		AppID:              types.NewUCompactFromUInt(uint64(AppID)),
		TransactionVersion: rv.TransactionVersion,
	}

	err = ext.Sign(keyring, options)
	if err != nil {
		return extrinsic.Extrinsic{}, fmt.Errorf("cannot sign:%v", err)
	}

	return ext, nil
}

func SubmitExtrinsic(api *SubstrateAPI, ext extrinsic.Extrinsic) (types.Hash, error) {
	hash, err := rpc.SubmitExtrinsic(ext, api.Client)
	if err != nil {
		return types.Hash{}, fmt.Errorf("cannot submit extrinsic:%v", err)
	}

	fmt.Printf("Data submitted using APPID: %v \n", ext.Signature.AppID.Int64())
	return hash, nil
}

func SubmitExtrinsicWatch(api *SubstrateAPI, ext extrinsic.Extrinsic, final chan types.Hash, txHash chan types.Hash, WaitForInclusion WaitFor) error {
	go func() {
		hash, err := ext.TxHash()
		if err != nil {
			log.Fatal(err)
		}
		txHash <- hash
	}()

	sub, err := rpc.SubmitAndWatchExtrinsic(ext, api.Client)
	if err != nil {
		return fmt.Errorf("cannot submit extrinsic:%v", err)
	}

	fmt.Printf("Transaction being submitted .... ⏳Waiting for block inclusion..")

	defer sub.Unsubscribe()
	timeout := time.After(200 * time.Second)
	for {
		select {
		case status := <-sub.Chan():
			switch WaitForInclusion {
			case BlockInclusion:
				if status.IsInBlock {
					final <- status.AsInBlock
					return err
				}
			case BlockFinalization:
				if status.IsFinalized {
					final <- status.AsFinalized
					return err
				}
			}
		case <-timeout:
			fmt.Printf("timeout of 200 seconds reached without getting finalized status for extrinsic")
			return err
		}
	}
}

func NewExtrinsic(api *SubstrateAPI, ext_call string, keyring signature.KeyringPair, AppID int, arg ...interface{}) (types.Hash, error) {
	ext, err := CreateExtrinsic(api, ext_call, keyring, AppID, arg...)
	if err != nil {
		return types.Hash{}, err
	}

	return SubmitExtrinsic(api, ext)
}

func NewExtrinsicWatch(api *SubstrateAPI, ext_call string, keyring signature.KeyringPair, final chan types.Hash, txHash chan types.Hash, AppID int, WaitForInclusion WaitFor, arg ...interface{}) error {
	ext, err := CreateExtrinsic(api, ext_call, keyring, AppID, arg...)
	if err != nil {
		return err
	}

	return SubmitExtrinsicWatch(api, ext, final, txHash, WaitForInclusion)
}
