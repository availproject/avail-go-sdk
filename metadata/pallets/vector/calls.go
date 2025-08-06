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

func (cfc CallFulfillCall) PalletIndex() uint8 {
	return PalletIndex
}

func (cfc CallFulfillCall) PalletName() string {
	return PalletName
}

func (cfc CallFulfillCall) CallIndex() uint8 {
	return 0
}

func (cfc CallFulfillCall) CallName() string {
	return "fulfill_call"
}

type CallExecute struct {
	Slot         uint64 `scale:"compact"`
	AddrMessage  metadata.VectorMessage
	AccountProof []byte
	StorageProof []byte
}

func (ce CallExecute) PalletIndex() uint8 {
	return PalletIndex
}

func (ce CallExecute) PalletName() string {
	return PalletName
}

func (ce CallExecute) CallIndex() uint8 {
	return 1
}

func (ce CallExecute) CallName() string {
	return "execute"
}

type CallSourceChainFroze struct {
	SourceChainId uint32 `scale:"compact"`
	Frozen        bool
}

func (cscf CallSourceChainFroze) PalletIndex() uint8 {
	return PalletIndex
}

func (cscf CallSourceChainFroze) PalletName() string {
	return PalletName
}

func (cscf CallSourceChainFroze) CallIndex() uint8 {
	return 2
}

func (cscf CallSourceChainFroze) CallName() string {
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

func (csm CallSendMessage) PalletIndex() uint8 {
	return PalletIndex
}

func (csm CallSendMessage) PalletName() string {
	return PalletName
}

func (csm CallSendMessage) CallIndex() uint8 {
	return 3
}

func (csm CallSendMessage) CallName() string {
	return "send_message"
}

type CallSetPoseidonHash struct {
	Period       uint64 `scale:"compact"`
	PoseidonHash []byte
}

func (csph CallSetPoseidonHash) PalletIndex() uint8 {
	return PalletIndex
}

func (csph CallSetPoseidonHash) PalletName() string {
	return PalletName
}

func (csph CallSetPoseidonHash) CallIndex() uint8 {
	return 4
}

func (csph CallSetPoseidonHash) CallName() string {
	return "set_poseidon_hash"
}

type CallSetBroadcaster struct {
	BroadcasterDomain uint32 `scale:"compact"`
	Broadcaster       prim.H256
}

func (csb CallSetBroadcaster) PalletIndex() uint8 {
	return PalletIndex
}

func (csb CallSetBroadcaster) PalletName() string {
	return PalletName
}

func (csb CallSetBroadcaster) CallIndex() uint8 {
	return 5
}

func (csb CallSetBroadcaster) CallName() string {
	return "set_broadcaster"
}

type CallSetWhitelistedDomains struct {
	Value []uint32
}

func (cswd CallSetWhitelistedDomains) PalletIndex() uint8 {
	return PalletIndex
}

func (cswd CallSetWhitelistedDomains) PalletName() string {
	return PalletName
}

func (cswd CallSetWhitelistedDomains) CallIndex() uint8 {
	return 6
}

func (cswd CallSetWhitelistedDomains) CallName() string {
	return "set_whitelisted_domains"
}

type CallSetConfiguration struct {
	Value metadata.VectorConfiguration
}

func (csc CallSetConfiguration) PalletIndex() uint8 {
	return PalletIndex
}

func (csc CallSetConfiguration) PalletName() string {
	return PalletName
}

func (csc CallSetConfiguration) CallIndex() uint8 {
	return 7
}

func (csc CallSetConfiguration) CallName() string {
	return "set_configuration"
}

type CallSetFunctionsIds struct {
	Value prim.Option[metadata.Tuple2[prim.H256, prim.H256]]
}

func (csfi CallSetFunctionsIds) PalletIndex() uint8 {
	return PalletIndex
}

func (csfi CallSetFunctionsIds) PalletName() string {
	return PalletName
}

func (csfi CallSetFunctionsIds) CallIndex() uint8 {
	return 8
}

func (csfi CallSetFunctionsIds) CallName() string {
	return "set_function_ids"
}

type CallSetStepVerificationKey struct {
	Value prim.Option[[]byte]
}

func (cssvk CallSetStepVerificationKey) PalletIndex() uint8 {
	return PalletIndex
}

func (cssvk CallSetStepVerificationKey) PalletName() string {
	return PalletName
}

func (cssvk CallSetStepVerificationKey) CallIndex() uint8 {
	return 9
}

func (cssvk CallSetStepVerificationKey) CallName() string {
	return "set_step_verification_key"
}

type CallSetRotateVerificationKey struct {
	Value prim.Option[[]byte]
}

func (csrvk CallSetRotateVerificationKey) PalletIndex() uint8 {
	return PalletIndex
}

func (csrvk CallSetRotateVerificationKey) PalletName() string {
	return PalletName
}

func (csrvk CallSetRotateVerificationKey) CallIndex() uint8 {
	return 10
}

func (csrvk CallSetRotateVerificationKey) CallName() string {
	return "set_rotate_verification_key"
}

type CallFailedSendMessageTxs struct {
	FailedTxs []uint32 `scale:"compact"`
}

func (cfsmt CallFailedSendMessageTxs) PalletIndex() uint8 {
	return PalletIndex
}

func (cfsmt CallFailedSendMessageTxs) PalletName() string {
	return PalletName
}

func (cfsmt CallFailedSendMessageTxs) CallIndex() uint8 {
	return 11
}

func (cfsmt CallFailedSendMessageTxs) CallName() string {
	return "failed_send_message_txs"
}

type CallSetUpdater struct {
	Updater prim.H256
}

func (csu CallSetUpdater) PalletIndex() uint8 {
	return PalletIndex
}

func (csu CallSetUpdater) PalletName() string {
	return PalletName
}

func (csu CallSetUpdater) CallIndex() uint8 {
	return 12
}

func (csu CallSetUpdater) CallName() string {
	return "set_updater"
}

type CallFulfill struct {
	Proof        []byte
	PublicValues []byte
}

func (cf CallFulfill) PalletIndex() uint8 {
	return PalletIndex
}

func (cf CallFulfill) PalletName() string {
	return PalletName
}

func (cf CallFulfill) CallIndex() uint8 {
	return 13
}

func (cf CallFulfill) CallName() string {
	return "fulfill"
}

type CallSetSp1VerificationKey struct {
	Sp1Vk prim.H256
}

func (c CallSetSp1VerificationKey) PalletIndex() uint8 {
	return PalletIndex
}

func (c CallSetSp1VerificationKey) PalletName() string {
	return PalletName
}

func (c CallSetSp1VerificationKey) CallIndex() uint8 {
	return 14
}

func (c CallSetSp1VerificationKey) CallName() string {
	return "set_sp1_verification_key"
}

type CallSetSyncCommitteeHash struct {
	Period uint64
	Hash   prim.H256
}

func (csch CallSetSyncCommitteeHash) PalletIndex() uint8 {
	return PalletIndex
}

func (csch CallSetSyncCommitteeHash) PalletName() string {
	return PalletName
}

func (csch CallSetSyncCommitteeHash) CallIndex() uint8 {
	return 15
}

func (csch CallSetSyncCommitteeHash) CallName() string {
	return "set_sync_committee_hash"
}

type CallEnableMock struct {
	Value bool
}

func (cem CallEnableMock) PalletIndex() uint8 {
	return PalletIndex
}

func (cem CallEnableMock) PalletName() string {
	return PalletName
}

func (cem CallEnableMock) CallIndex() uint8 {
	return 16
}

func (cem CallEnableMock) CallName() string {
	return "enable_mock"
}

type CallMockFulfill struct {
	PublicValues []byte
}

func (cmf CallMockFulfill) PalletIndex() uint8 {
	return PalletIndex
}

func (cmf CallMockFulfill) PalletName() string {
	return PalletName
}

func (cmf CallMockFulfill) CallIndex() uint8 {
	return 17
}

func (cmf CallMockFulfill) CallName() string {
	return "mock_fulfill"
}
