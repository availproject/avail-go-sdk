package sdk

import (
	"fmt"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type Filter struct {
	AppId    prim.Option[uint32]
	TxHash   prim.Option[prim.H256]
	TxIndex  prim.Option[uint32]
	TxSigner prim.Option[prim.AccountId]
}

func (this Filter) WAppId(value uint32) Filter {
	this.AppId = prim.NewSome(value)
	return this
}

func (this Filter) WTxHash(value prim.H256) Filter {
	this.TxHash = prim.NewSome(value)
	return this
}

func (this Filter) WTxIndex(value uint32) Filter {
	this.TxIndex = prim.NewSome(value)
	return this
}

func (this Filter) WTxSigner(value prim.AccountId) Filter {
	this.TxSigner = prim.NewSome(value)
	return this
}

type Block struct {
	client *Client
	Block  RPCBlock
	events prim.Option[EventRecords]
}

func NewBlock(client *Client, blockHash prim.H256) (Block, error) {
	block, err := client.RPCBlockAt(prim.NewSome(blockHash))
	if err != nil {
		return Block{}, err
	}

	events, err := client.EventsAt(prim.NewSome(blockHash))
	blockEvents := prim.NewNone[EventRecords]()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		blockEvents = prim.NewSome(events)
	}

	return Block{
		client: client,
		Block:  block,
		events: blockEvents,
	}, nil
}

func NewBestBlock(client *Client) (Block, error) {
	hash, err := client.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
	if err != nil {
		return Block{}, err
	}
	return NewBlock(client, hash)
}

func NewFinalizedBlock(client *Client) (Block, error) {
	hash, err := client.Rpc.Chain.GetFinalizedHead()
	if err != nil {
		return Block{}, err
	}
	return NewBlock(client, hash)
}

func (this *Block) Transactions(filter Filter) []BlockTransaction {
	extrinsics := this.Block.Extrinsics
	result := []BlockTransaction{}
	for i := range this.Block.Extrinsics {
		ext := &extrinsics[i]

		if filter.AppId.IsSome() {
			if ext.Signed.IsNone() {
				continue
			}
			var signed = ext.Signed.Unwrap()
			if signed.AppId != filter.AppId.Unwrap() {
				continue
			}
		}

		if filter.TxHash.IsSome() {
			if ext.TxHash != filter.TxHash.Unwrap() {
				continue
			}
		}

		if filter.TxIndex.IsSome() {
			if ext.TxIndex != filter.TxIndex.Unwrap() {
				continue
			}
		}

		if filter.TxSigner.IsSome() {
			if !sameSignature(ext, filter.TxSigner.Unwrap()) {
				continue
			}
		}

		txEvents := this.EventsForTransaction(extrinsics[i].TxIndex)
		result = append(result, NewBlockTransaction(this.client, &extrinsics[i], txEvents))
	}

	return result
}

func (this *Block) DataSubmissions(filter Filter) []DataSubmission {
	extrinsics := this.Block.Extrinsics
	result := []DataSubmission{}
	for i := range extrinsics {
		ext := &extrinsics[i]

		if filter.AppId.IsSome() {
			if ext.Signed.IsNone() {
				continue
			}
			var signed = ext.Signed.Unwrap()
			if signed.AppId != filter.AppId.Unwrap() {
				continue
			}
		}

		if filter.TxHash.IsSome() {
			if ext.TxHash != filter.TxHash.Unwrap() {
				continue
			}
		}

		if filter.TxIndex.IsSome() {
			if ext.TxIndex != filter.TxIndex.Unwrap() {
				continue
			}
		}

		if filter.TxSigner.IsSome() {
			if !sameSignature(ext, filter.TxSigner.Unwrap()) {
				continue
			}
		}

		if res, ok := NewDataSubmission(ext); ok {
			result = append(result, res)
		}
	}

	return result
}

func (this *Block) EventsForTransaction(txIndex uint32) prim.Option[EventRecords] {
	extrinsics := this.Block.Extrinsics

	if txIndex >= uint32(len(extrinsics)) {
		return prim.NewNone[EventRecords]()
	}

	if this.events.IsNone() {
		return prim.NewNone[EventRecords]()
	}

	for i := range extrinsics {
		if extrinsics[i].TxIndex != txIndex {
			continue
		}
		allEvents := this.events.Unwrap()
		txEvents := EventFilterByTxIndex(allEvents, txIndex)
		return prim.NewSome(txEvents)
	}

	return prim.NewNone[EventRecords]()
}

func (this *Block) Events() prim.Option[EventRecords] {
	return this.events
}

func sameSignature(tx *prim.DecodedExtrinsic, accountId prim.AccountId) bool {
	txAccountId := tx.Signed.Unwrap().Address.Id.Unwrap()
	if accountId != txAccountId {
		return false
	}

	return true
}

type DataSubmission struct {
	TxHash   prim.H256
	TxIndex  uint32
	Data     []byte
	TxSigner prim.MultiAddress
	AppId    uint32
}

func NewDataSubmission(tx *prim.DecodedExtrinsic) (DataSubmission, bool) {
	callSubmitData := daPallet.CallSubmitData{}
	if tx.Call.PalletIndex != callSubmitData.PalletIndex() {
		return DataSubmission{}, false
	}

	if tx.Call.CallIndex != callSubmitData.CallIndex() {
		return DataSubmission{}, false
	}

	// Data submission cannot be done without signed being set.
	signed := tx.Signed.Unwrap()

	decoder := prim.NewDecoder(prim.Hex.FromHex(tx.Call.Fields.Value), 0)
	if err := decoder.Decode(&callSubmitData); err != nil {
		return DataSubmission{}, false
	}

	res := DataSubmission{
		TxHash:   tx.TxHash,
		TxIndex:  tx.TxIndex,
		Data:     callSubmitData.Data,
		TxSigner: signed.Address,
		AppId:    signed.AppId,
	}

	return res, true
}
