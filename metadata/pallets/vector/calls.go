package vector

import (
	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type CallFulfillCall struct {
	FunctionId prim.H256
	Input      []byte
	Output     []byte
	Proof      []byte
	Slot       uint64 `scale:"compact"`
}

func (this CallFulfillCall) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallFulfillCall) PalletName() string {
	return PalletName
}

func (this CallFulfillCall) CallIndex() uint8 {
	return 0
}

func (this CallFulfillCall) CallName() string {
	return "fulfill_call"
}

type CallExecute struct {
	Slot         uint64 `scale:"compact"`
	AddrMessage  metadata.VectorMessage
	AccountProof []byte
	StorageProof []byte
}

func (this CallExecute) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallExecute) PalletName() string {
	return PalletName
}

func (this CallExecute) CallIndex() uint8 {
	return 1
}

func (this CallExecute) CallName() string {
	return "execute"
}

type CallSourceChainFroze struct {
	SourceChainId uint32 `scale:"compact"`
	Frozen        bool
}

func (this CallSourceChainFroze) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSourceChainFroze) PalletName() string {
	return PalletName
}

func (this CallSourceChainFroze) CallIndex() uint8 {
	return 2
}

func (this CallSourceChainFroze) CallName() string {
	return "source_chain_froze"
}

// Send a batch of dispatch calls.
//
// May be called from any origin except `None`.
type CallSendMessage struct {
	Message metadata.VectorMessage
	To      prim.H256
	Domain  uint32 `scale:"compact"`
}

func (this CallSendMessage) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSendMessage) PalletName() string {
	return PalletName
}

func (this CallSendMessage) CallIndex() uint8 {
	return 3
}

func (this CallSendMessage) CallName() string {
	return "send_message"
}

type CallSetPoseidonHash struct {
	Period       uint64 `scale:"compact"`
	PoseidonHash []byte
}

func (this CallSetPoseidonHash) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetPoseidonHash) PalletName() string {
	return PalletName
}

func (this CallSetPoseidonHash) CallIndex() uint8 {
	return 4
}

func (this CallSetPoseidonHash) CallName() string {
	return "set_poseidon_hash"
}

type CallSetBroadcaster struct {
	BroadcasterDomain uint32 `scale:"compact"`
	Broadcaster       prim.H256
}

func (this CallSetBroadcaster) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetBroadcaster) PalletName() string {
	return PalletName
}

func (this CallSetBroadcaster) CallIndex() uint8 {
	return 5
}

func (this CallSetBroadcaster) CallName() string {
	return "set_broadcaster"
}

type CallSetWhitelistedDomains struct {
	Value []uint32
}

func (this CallSetWhitelistedDomains) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetWhitelistedDomains) PalletName() string {
	return PalletName
}

func (this CallSetWhitelistedDomains) CallIndex() uint8 {
	return 6
}

func (this CallSetWhitelistedDomains) CallName() string {
	return "set_whitelisted_domains"
}

type CallSetConfiguration struct {
	Value metadata.VectorConfiguration
}

func (this CallSetConfiguration) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetConfiguration) PalletName() string {
	return PalletName
}

func (this CallSetConfiguration) CallIndex() uint8 {
	return 7
}

func (this CallSetConfiguration) CallName() string {
	return "set_configuration"
}

type CallSetFunctionsIds struct {
	Value prim.Option[metadata.Tuple2[prim.H256, prim.H256]]
}

func (this CallSetFunctionsIds) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetFunctionsIds) PalletName() string {
	return PalletName
}

func (this CallSetFunctionsIds) CallIndex() uint8 {
	return 8
}

func (this CallSetFunctionsIds) CallName() string {
	return "set_function_ids"
}

type CallSetStepVerificationKey struct {
	Value prim.Option[[]byte]
}

func (this CallSetStepVerificationKey) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetStepVerificationKey) PalletName() string {
	return PalletName
}

func (this CallSetStepVerificationKey) CallIndex() uint8 {
	return 9
}

func (this CallSetStepVerificationKey) CallName() string {
	return "set_step_verification_key"
}

type CallSetRotateVerificationKey struct {
	Value prim.Option[[]byte]
}

func (this CallSetRotateVerificationKey) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetRotateVerificationKey) PalletName() string {
	return PalletName
}

func (this CallSetRotateVerificationKey) CallIndex() uint8 {
	return 10
}

func (this CallSetRotateVerificationKey) CallName() string {
	return "set_rotate_verification_key"
}

type CallFailedSendMessageTxs struct {
	FailedTxs []uint32 `scale:"compact"`
}

func (this CallFailedSendMessageTxs) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallFailedSendMessageTxs) PalletName() string {
	return PalletName
}

func (this CallFailedSendMessageTxs) CallIndex() uint8 {
	return 11
}

func (this CallFailedSendMessageTxs) CallName() string {
	return "failed_send_message_txs"
}

type CallSetUpdater struct {
	Updater prim.H256
}

func (this CallSetUpdater) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetUpdater) PalletName() string {
	return PalletName
}

func (this CallSetUpdater) CallIndex() uint8 {
	return 12
}

func (this CallSetUpdater) CallName() string {
	return "set_updater"
}

type CallFulfill struct {
	Proof        []byte
	PublicValues []byte
}

func (this CallFulfill) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallFulfill) PalletName() string {
	return PalletName
}

func (this CallFulfill) CallIndex() uint8 {
	return 13
}

func (this CallFulfill) CallName() string {
	return "fulfill"
}

type CallSetSp1VerificationKey struct {
	Sp1Vk prim.H256
}

func (this CallSetSp1VerificationKey) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetSp1VerificationKey) PalletName() string {
	return PalletName
}

func (this CallSetSp1VerificationKey) CallIndex() uint8 {
	return 14
}

func (this CallSetSp1VerificationKey) CallName() string {
	return "set_sp1_verification_key"
}

type CallSetSyncCommitteeHash struct {
	Period uint64
	Hash   prim.H256
}

func (this CallSetSyncCommitteeHash) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetSyncCommitteeHash) PalletName() string {
	return PalletName
}

func (this CallSetSyncCommitteeHash) CallIndex() uint8 {
	return 15
}

func (this CallSetSyncCommitteeHash) CallName() string {
	return "set_sync_committee_hash"
}

type CallEnableMock struct {
	Value bool
}

func (this CallEnableMock) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallEnableMock) PalletName() string {
	return PalletName
}

func (this CallEnableMock) CallIndex() uint8 {
	return 16
}

func (this CallEnableMock) CallName() string {
	return "enable_mock"
}

type CallMockFulfill struct {
	PublicValues []byte
}

func (this CallMockFulfill) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallMockFulfill) PalletName() string {
	return PalletName
}

func (this CallMockFulfill) CallIndex() uint8 {
	return 17
}

func (this CallMockFulfill) CallName() string {
	return "mock_fulfill"
}
