package sdk

import (
	"github.com/vedhavyas/go-subkey/v2"

	"errors"

	"github.com/availproject/avail-go-sdk/metadata"
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

func (this *Transaction) CallToHex() string {
	return "0x" + prim.Encoder.Encode(this.Payload.Call)
}

func (this *Transaction) ToHex(account subkey.KeyPair, options TransactionOptions) (string, error) {
	extra, additional, err := options.ToPrimitive(this.client, account.SS58Address(42))
	if err != nil {
		return "", err
	}
	tx, err := prim.CreateSigned(this.Payload.Call, extra, additional, account)
	if err != nil {
		return "", err
	}

	return tx.ToHexWith0x(), nil
}

// Transaction will be signed and send.
//
// There is no guarantee that the transaction was executed at all. It might have been
// dropped or discarded for various reasons. The caller is responsible for querying future
// blocks in order to determine the execution status of that transaction.
func (this *Transaction) Execute(account subkey.KeyPair, options TransactionOptions) (prim.H256, error) {
	return TransactionSignAndSend(this.client, account, this.Payload, options)
}

// Transaction will be signed, send and watched
//
// Same as manually calling `Execute` plus running the Watcher
func (this *Transaction) ExecuteAndWatch(account subkey.KeyPair, waitFor uint8, options TransactionOptions, blockTimeout uint32) (TransactionDetails, error) {
	return TransactionSignSendWatch(this.client, account, this.Payload, waitFor, options, blockTimeout, 3)
}

// Transaction will be signed, send and watched
//
// Same as manually calling `Execute` plus running the Watcher
func (this *Transaction) ExecuteAndWatchFinalization(account subkey.KeyPair, options TransactionOptions) (TransactionDetails, error) {
	return TransactionSignSendWatch(this.client, account, this.Payload, Finalization, options, 5, 3)
}

// Transaction will be signed, send and watched
//
// Same as manually calling `Execute` plus running the Watcher
func (this *Transaction) ExecuteAndWatchInclusion(account subkey.KeyPair, options TransactionOptions) (TransactionDetails, error) {
	return TransactionSignSendWatch(this.client, account, this.Payload, Inclusion, options, 3, 3)
}

func (this *Transaction) PaymentQueryFeeDetails(account subkey.KeyPair, options TransactionOptions) (metadata.FeeDetails, error) {
	val, err := this.ToHex(account, options)
	if err != nil {
		return metadata.FeeDetails{}, err
	}

	return this.client.Call.TransactionPaymentApi_queryFeeDetails(val, prim.NewNone[prim.H256]())
}

func (this *Transaction) PaymentQueryFeeInfo(account subkey.KeyPair, options TransactionOptions) (metadata.RuntimeDispatchInfo, error) {
	val, err := this.ToHex(account, options)
	if err != nil {
		return metadata.RuntimeDispatchInfo{}, err
	}

	return this.client.Call.TransactionPaymentApi_queryInfo(val, prim.NewNone[prim.H256]())
}

func (this *Transaction) PaymentQueryCallFeeDetails() (metadata.FeeDetails, error) {
	return this.client.Call.TransactionPaymentCallApi_queryCallFeeDetails(this.CallToHex(), prim.NewNone[prim.H256]())
}

func (this *Transaction) PaymentQueryCallFeeInfo() (metadata.RuntimeDispatchInfo, error) {
	return this.client.Call.TransactionPaymentCallApi_queryCallInfo(this.CallToHex(), prim.NewNone[prim.H256]())
}

type TransactionDetails struct {
	client      *Client
	TxHash      prim.H256
	TxIndex     uint32
	BlockHash   prim.H256
	BlockNumber uint32
	Events      prim.Option[EventRecords]
}

func (this *TransactionDetails) IsSuccessful() (bool, error) {
	if this.Events.IsNone() {
		return false, errors.New("No events were decoded.")
	}
	events := this.Events.Unwrap()

	maybeFound, err := EventFindFirstChecked(events, syPallet.EventExtrinsicFailed{})
	if err != nil {
		return false, err
	}

	if maybeFound.IsNone() {
		return true, nil
	}

	return false, nil
}
