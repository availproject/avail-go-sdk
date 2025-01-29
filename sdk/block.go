package sdk

import (
	"fmt"

	meta "github.com/availproject/avail-go-sdk/metadata"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

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

func (this *Block) TransactionAll() []BlockTransaction {
	extrinsics := this.Block.Extrinsics
	result := []BlockTransaction{}
	for i := range this.Block.Extrinsics {
		txEvents := this.EventsForTransaction(extrinsics[i].TxIndex)
		result = append(result, NewBlockTransaction(this.client, &extrinsics[i], txEvents))
	}

	return result
}

func (this *Block) TransactionBySigner(accountId meta.AccountId) []BlockTransaction {
	extrinsics := this.Block.Extrinsics
	result := []BlockTransaction{}
	for i := range extrinsics {
		if !sameSignature(&extrinsics[i], accountId) {
			continue
		}

		txEvents := this.EventsForTransaction(extrinsics[i].TxIndex)
		result = append(result, NewBlockTransaction(this.client, &extrinsics[i], txEvents))
	}

	return result
}

func (this *Block) TransactionByIndex(txIndex uint32) prim.Option[BlockTransaction] {
	extrinsics := this.Block.Extrinsics
	for i := range extrinsics {
		if extrinsics[i].TxIndex != txIndex {
			continue
		}

		txEvents := this.EventsForTransaction(extrinsics[i].TxIndex)
		return prim.NewSome(NewBlockTransaction(this.client, &extrinsics[i], txEvents))
	}

	return prim.NewNone[BlockTransaction]()
}

func (this *Block) TransactionByHash(txHash prim.H256) prim.Option[BlockTransaction] {
	extrinsics := this.Block.Extrinsics
	for i := range extrinsics {
		if extrinsics[i].TxHash != txHash {
			continue
		}
		txEvents := this.EventsForTransaction(extrinsics[i].TxIndex)
		return prim.NewSome(NewBlockTransaction(this.client, &extrinsics[i], txEvents))
	}

	return prim.NewNone[BlockTransaction]()
}

func (this *Block) TransactionByAppId(appId uint32) []BlockTransaction {
	extrinsics := this.Block.Extrinsics
	result := []BlockTransaction{}
	for i := range extrinsics {
		if extrinsics[i].Signed.IsNone() {
			continue
		}
		var signed = extrinsics[i].Signed.Unwrap()
		if signed.AppId != appId {
			continue
		}

		txEvents := this.EventsForTransaction(extrinsics[i].TxIndex)
		result = append(result, NewBlockTransaction(this.client, &extrinsics[i], txEvents))
	}

	return result
}

func (this *Block) DataSubmissionAll() []DataSubmission {
	extrinsics := this.Block.Extrinsics
	result := []DataSubmission{}
	for i := range extrinsics {
		if res, ok := NewDataSubmission(&extrinsics[i]); ok {
			result = append(result, res)
		}
	}

	return result
}

func (this *Block) DataSubmissionBySigner(accountId meta.AccountId) []DataSubmission {
	extrinsics := this.Block.Extrinsics
	result := []DataSubmission{}
	for i := range extrinsics {
		if !sameSignature(&extrinsics[i], accountId) {
			continue
		}

		if res, ok := NewDataSubmission(&extrinsics[i]); ok {
			result = append(result, res)
		}
	}

	return result
}

func (this *Block) DataSubmissionByIndex(txIndex uint32) prim.Option[DataSubmission] {
	extrinsics := this.Block.Extrinsics
	for i := range extrinsics {
		if extrinsics[i].TxIndex != txIndex {
			continue
		}
		if res, ok := NewDataSubmission(&extrinsics[i]); ok {
			return prim.NewSome(res)
		}
	}

	return prim.NewNone[DataSubmission]()
}

func (this *Block) DataSubmissionByHash(txHash prim.H256) prim.Option[DataSubmission] {
	extrinsics := this.Block.Extrinsics
	for i := range extrinsics {
		if extrinsics[i].TxHash != txHash {
			continue
		}
		if res, ok := NewDataSubmission(&extrinsics[i]); ok {
			return prim.NewSome(res)
		}
	}

	return prim.NewNone[DataSubmission]()
}

func (this *Block) DataSubmissionByAppId(appId uint32) []DataSubmission {
	extrinsics := this.Block.Extrinsics
	result := []DataSubmission{}
	for i := range extrinsics {
		if extrinsics[i].Signed.IsNone() {
			continue
		}
		var signed = extrinsics[i].Signed.Unwrap()
		if signed.AppId != appId {
			continue
		}
		if res, ok := NewDataSubmission(&extrinsics[i]); ok {
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

func sameSignature(tx *prim.DecodedExtrinsic, accountId meta.AccountId) bool {
	txAccountId := tx.Signed.Unwrap().Address.Id.Unwrap()
	if accountId.Value != txAccountId {
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
