package complex

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

func NewBlock(client *Client, blockHash prim.H256) Block {
	block := client.GetBlock(prim.NewSome(blockHash))
	events, err := client.GetEvents(prim.NewSome(blockHash))
	if err != nil {
		panic(err)
	}
	return Block{
		client: client,
		Block:  block,
		events: prim.NewSome(events),
	}
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

func (this *Block) Events(appId uint32) prim.Option[EventRecords] {
	return this.events
}

func sameSignature(tx *prim.DecodedExtrinsic, accountId meta.AccountId) bool {
	txAccountId := tx.Signed.UnwrapOrDefault().Address.Id.UnwrapOrDefault()
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
	ok := false
	res := DataSubmission{}

	callSubmitData := daPallet.CallSubmitData{}
	if tx.Call.PalletIndex != callSubmitData.PalletIndex() {
		return res, ok
	}

	if tx.Call.CallIndex != callSubmitData.CallIndex() {
		return res, ok
	}

	// Data submission cannot be done without signed being set.
	signed := tx.Signed.UnwrapOrDefault()

	decoder := prim.NewDecoder(prim.FromHex(tx.Call.Fields.Value), 0)
	if err := decoder.Decode(&callSubmitData); err != nil {
		return res, false
	}

	res = DataSubmission{
		TxHash:   tx.TxHash,
		TxIndex:  tx.TxIndex,
		Data:     callSubmitData.Data,
		TxSigner: signed.Address,
		AppId:    signed.AppId,
	}

	return res, true
}
