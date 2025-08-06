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

func (f Filter) WAppId(value uint32) Filter {
	f.AppId = prim.Some(value)
	return f
}

func (f Filter) WTxHash(value prim.H256) Filter {
	f.TxHash = prim.Some(value)
	return f
}

func (f Filter) WTxIndex(value uint32) Filter {
	f.TxIndex = prim.Some(value)
	return f
}

func (f Filter) WTxSigner(value prim.AccountId) Filter {
	f.TxSigner = prim.Some(value)
	return f
}

type Block struct {
	client *Client
	Block  RPCBlock
	events prim.Option[EventRecords]
}

func NewBlock(client *Client, blockHash prim.H256) (Block, error) {
	block, err := client.RPCBlockAt(prim.Some(blockHash))
	if err != nil {
		return Block{}, err
	}

	events, err := client.EventsAt(prim.Some(blockHash))
	blockEvents := prim.None[EventRecords]()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		blockEvents = prim.Some(events)
	}

	return Block{
		client: client,
		Block:  block,
		events: blockEvents,
	}, nil
}

func NewBestBlock(client *Client) (Block, error) {
	hash, err := client.Rpc.Chain.GetBlockHash(prim.None[uint32]())
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

func (b *Block) Transactions(filter Filter) []BlockTransaction {
	extrinsics := b.Block.Extrinsics
	result := []BlockTransaction{}
	for i := range b.Block.Extrinsics {
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

		txEvents := b.EventsForTransaction(extrinsics[i].TxIndex)
		result = append(result, NewBlockTransaction(b.client, &extrinsics[i], txEvents))
	}

	return result
}

func (b *Block) DataSubmissions(filter Filter) []DataSubmission {
	extrinsics := b.Block.Extrinsics
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

func (b *Block) EventsForTransaction(txIndex uint32) prim.Option[EventRecords] {
	if b.events.IsNone() {
		return prim.None[EventRecords]()
	}

	if txIndex >= uint32(len(b.Block.Extrinsics)) {
		return prim.None[EventRecords]()
	}

	allEvents := b.events.Unwrap()
	txEvents := EventFilterByTxIndex(allEvents, txIndex)
	return prim.Some(txEvents)
}

func (b *Block) Events() prim.Option[EventRecords] {
	return b.events
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
