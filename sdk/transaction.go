package sdk

import (
	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/vedhavyas/go-subkey/v2"

	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

const Finalization = uint8(0)
const Inclusion = uint8(1)

type Transaction struct {
	client  *Client
	Payload metadata.Payload
}

func NewTransaction(client *Client, payload metadata.Payload) Transaction {
	return Transaction{
		client:  client,
		Payload: payload,
	}
}

func (t *Transaction) CallToHex() string {
	return "0x" + prim.Encoder.Encode(t.Payload.Call)
}

func (t *Transaction) ToHex(account subkey.KeyPair, options TransactionOptions) (string, error) {
	extra, additional, _, err := options.ToPrimitive(t.client, account.SS58Address(42))
	if err != nil {
		return "", err
	}
	tx, err := prim.CreateSigned(t.Payload.Call, extra, additional, account)
	if err != nil {
		return "", err
	}

	return tx.ToHexWith0x(), nil
}

// Transaction will be signed and sent.
//
// There is no guarantee that the transaction was executed at all. It might have been
// dropped or discarded for various reasons. The caller is responsible for querying future
// blocks in order to determine the execution status of that transaction.
func (t *Transaction) Execute(account subkey.KeyPair, options TransactionOptions) (prim.H256, error) {
	return TransactionSignAndSend(t.client, account, t.Payload, options)
}

// Transaction will be signed, sent, and watched
// If the transaction was dropped or never executed, the system will retry it
// for 2 more times using the same nonce and app id.
//
// Waits for finalization to finalize the transaction.
func (t *Transaction) ExecuteAndWatchFinalization(account subkey.KeyPair, options TransactionOptions) (TransactionDetails, error) {
	return TransactionSignSendWatch(t.client, account, t.Payload, Finalization, options)
}

// Transaction will be signed, sent, and watched
// If the transaction was dropped or never executed, the system will retry it
// for 2 more times using the same nonce and app id.
//
// Waits for transaction inclusion. Most of the time you would want to call `ExecuteAndWatchFinalization` as
// inclusion doesn't mean that the transaction will be in the canonical chain.
func (t *Transaction) ExecuteAndWatchInclusion(account subkey.KeyPair, options TransactionOptions) (TransactionDetails, error) {
	return TransactionSignSendWatch(t.client, account, t.Payload, Inclusion, options)
}

func (t *Transaction) PaymentQueryInfo(account subkey.KeyPair, options TransactionOptions) (metadata.RuntimeDispatchInfo, error) {
	val, err := t.ToHex(account, options)
	if err != nil {
		return metadata.RuntimeDispatchInfo{}, err
	}

	return t.client.Call.TransactionPaymentApi_queryInfo(val, prim.None[prim.H256]())
}

func (t *Transaction) PaymentQueryFeeDetails(account subkey.KeyPair, options TransactionOptions) (metadata.FeeDetails, error) {
	val, err := t.ToHex(account, options)
	if err != nil {
		return metadata.FeeDetails{}, err
	}

	return t.client.Call.TransactionPaymentApi_queryFeeDetails(val, prim.None[prim.H256]())
}

func (t *Transaction) PaymentQueryCallInfo() (metadata.RuntimeDispatchInfo, error) {
	return t.client.Call.TransactionPaymentCallApi_queryCallInfo(t.CallToHex(), prim.None[prim.H256]())
}

func (t *Transaction) PaymentQueryCallFeeDetails() (metadata.FeeDetails, error) {
	return t.client.Call.TransactionPaymentCallApi_queryCallFeeDetails(t.CallToHex(), prim.None[prim.H256]())
}

type TransactionDetails struct {
	client      *Client
	TxHash      prim.H256
	TxIndex     uint32
	BlockHash   prim.H256
	BlockNumber uint32
	Events      prim.Option[EventRecords]
}

// Returns None if there was no way to determine the
// success status of a transaction. Otherwise it returns
// true or false.
func (t *TransactionDetails) IsSuccessful() prim.Option[bool] {
	events := t.Events.Unwrap()

	extFailedEvent := syPallet.EventExtrinsicFailed{}
	extSuccessEvent := syPallet.EventExtrinsicSuccess{}

	for i := range events {
		if events[i].PalletIndex == extFailedEvent.PalletIndex() && events[i].EventIndex == extFailedEvent.EventIndex() {
			return prim.Some(false)
		}
		if events[i].PalletIndex == extSuccessEvent.PalletIndex() && events[i].EventIndex == extSuccessEvent.EventIndex() {
			return prim.Some(true)
		}
	}

	return prim.None[bool]()
}
