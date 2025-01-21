package sdk

import (
	meta "go-sdk/metadata"
	daPallet "go-sdk/metadata/pallets/data_availability"
	prim "go-sdk/primitives"
)

type Block struct {
	client *Client
	Block  RPCBlock
	events prim.Option[EventRecords]
}

func NewBlock(client *Client, blockHash prim.H256) (Block, error) {
	block, err := client.BlockAt(prim.NewSome(blockHash))
	if err != nil {
		return Block{}, nil
	}

	events, err := client.EventsAt(prim.NewSome(blockHash))
	blockEvents := prim.NewNone[EventRecords]()
	if err != nil {
		println(err.Error())
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
	var result = []BlockTransaction{}
	for _, tx := range this.Block.Extrinsics {
		result = append(result, NewBlockTransaction(this.client, &tx))
	}

	return result
}

func (this *Block) TransactionBySigner(accountId meta.AccountId) []BlockTransaction {
	var result = []BlockTransaction{}
	for _, tx := range this.Block.Extrinsics {
		if !sameSignature(&tx, accountId) {
			continue
		}
		result = append(result, NewBlockTransaction(this.client, &tx))
	}

	return result
}

func (this *Block) TransactionByIndex(txIndex uint32) prim.Option[BlockTransaction] {
	for _, tx := range this.Block.Extrinsics {
		if tx.TxIndex == txIndex {
			return prim.NewSome(NewBlockTransaction(this.client, &tx))
		}
	}

	return prim.NewNone[BlockTransaction]()
}

func (this *Block) TransactionByHash(txHash prim.H256) prim.Option[BlockTransaction] {
	for _, tx := range this.Block.Extrinsics {
		if tx.TxHash == txHash {
			return prim.NewSome(NewBlockTransaction(this.client, &tx))
		}
	}

	return prim.NewNone[BlockTransaction]()
}

func (this *Block) TransactionByAppId(appId uint32) []BlockTransaction {
	var result = []BlockTransaction{}
	for _, tx := range this.Block.Extrinsics {
		if tx.Signed.IsNone() {
			continue
		}
		var signed = tx.Signed.Unwrap()
		if signed.AppId != appId {
			continue
		}
		result = append(result, NewBlockTransaction(this.client, &tx))
	}

	return result
}

func (this *Block) DataSubmissionAll() []DataSubmission {
	var result = []DataSubmission{}
	for _, tx := range this.Block.Extrinsics {
		if res, ok := NewDataSubmission(&tx); ok {
			result = append(result, res)
		}
	}

	return result
}

func (this *Block) DataSubmissionBySigner(accountId meta.AccountId) []DataSubmission {
	var result = []DataSubmission{}
	for _, tx := range this.Block.Extrinsics {
		if !sameSignature(&tx, accountId) {
			continue
		}

		if res, ok := NewDataSubmission(&tx); ok {
			result = append(result, res)
		}

	}

	return result
}

func (this *Block) DataSubmissionByIndex(txIndex uint32) prim.Option[DataSubmission] {
	for _, tx := range this.Block.Extrinsics {
		if tx.TxIndex != txIndex {
			continue
		}
		if res, ok := NewDataSubmission(&tx); ok {
			return prim.NewSome(res)
		}
	}

	return prim.NewNone[DataSubmission]()
}

func (this *Block) DataSubmissionByHash(txHash prim.H256) prim.Option[DataSubmission] {
	for _, tx := range this.Block.Extrinsics {
		if tx.TxHash != txHash {
			continue
		}
		if res, ok := NewDataSubmission(&tx); ok {
			return prim.NewSome(res)
		}
	}

	return prim.NewNone[DataSubmission]()
}

func (this *Block) DataSubmissionByAppId(appId uint32) []DataSubmission {
	var result = []DataSubmission{}
	for _, tx := range this.Block.Extrinsics {
		if tx.Signed.IsNone() {
			continue
		}
		var signed = tx.Signed.Unwrap()
		if signed.AppId != appId {
			continue
		}
		if res, ok := NewDataSubmission(&tx); ok {
			result = append(result, res)
		}
	}

	return result
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
