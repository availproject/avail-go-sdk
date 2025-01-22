package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"go-sdk/metadata"
	baPallet "go-sdk/metadata/pallets/balances"
	daPallet "go-sdk/metadata/pallets/data_availability"
	stPallet "go-sdk/metadata/pallets/staking"
	utPallet "go-sdk/metadata/pallets/utility"
	prim "go-sdk/primitives"
)

type DataAvailabilityTx struct {
	client *Client
}

func (this *Transactions) NewTransaction(payload metadata.Payload) Transaction {
	return NewTransaction(this.client, payload)
}

func (this *DataAvailabilityTx) SubmitData(data []byte) Transaction {
	call := daPallet.CallSubmitData{Data: data}
	return NewTransaction(this.client, call.ToPayload())
}

func (this *DataAvailabilityTx) CreateApplicationKey(key []byte) Transaction {
	call := daPallet.CallCreateApplicationKey{Key: key}
	return NewTransaction(this.client, call.ToPayload())
}

type UtilityTx struct {
	client *Client
}

// Send a batch of dispatch calls.
//
// May be called from any origin except `None`.
func (this *UtilityTx) Batch(calls []prim.Call) Transaction {
	call := utPallet.CallBatch{Calls: calls}
	return NewTransaction(this.client, call.ToPayload())
}

// Send a batch of dispatch calls and atomically execute them.
// The whole transaction will rollback and fail if any of the calls failed.
//
// May be called from any origin except `None`.
func (this *UtilityTx) BatchAll(calls []prim.Call) Transaction {
	call := utPallet.CallBatchAll{Calls: calls}
	return NewTransaction(this.client, call.ToPayload())
}

// Send a batch of dispatch calls.
// Unlike `batch`, it allows errors and won't interrupt.
//
// May be called from any origin except `None`.
func (this *UtilityTx) ForceBatch(calls []prim.Call) Transaction {
	call := utPallet.CallForceBatch{Calls: calls}
	return NewTransaction(this.client, call.ToPayload())
}

// Send a call through an indexed pseudonym of the sender.
//
// Filter from origin are passed along. The call will be dispatched with an origin which
// use the same filter as the origin of this call.
//
// NOTE: If you need to ensure that any account-based filtering is not honored (i.e.
// because you expect `proxy` to have been used prior in the call stack and you do not want
// the call restrictions to apply to any sub-accounts), then use `as_multi_threshold_1`
// in the Multisig pallet instead.
func (this *UtilityTx) AsDerivate(index uint16, call prim.Call) Transaction {
	c := utPallet.CallAsDerivate{Index: index, Call: call}
	return NewTransaction(this.client, c.ToPayload())
}

type BalancesTx struct {
	client *Client
}

// Transfer some liquid free balance to another account.
//
// `transfer_allow_death` will set the `FreeBalance` of the sender and receiver.
// If the sender's account is below the existential deposit as a result
// of the transfer, the account will be reaped.
//
// The dispatch origin for this call must be `Signed` by the transactor.
func (this *BalancesTx) TransferAllowDeath(dest prim.MultiAddress, amount uint128.Uint128) Transaction {
	call := baPallet.CallTransferAlowDeath{Dest: dest, Value: amount}
	return NewTransaction(this.client, call.ToPayload())
}

// Exactly as `TransferAlowDeath`, except the origin must be root and the source account
// may be specified
func (this *BalancesTx) ForceTransfer(dest prim.MultiAddress, amount uint128.Uint128) Transaction {
	call := baPallet.CallForceTransfer{Dest: dest, Value: amount}
	return NewTransaction(this.client, call.ToPayload())
}

// Same as the `TransferAlowDeath` call, but with a check that the transfer will not
// kill the origin account.
func (this *BalancesTx) TransferKeepAlive(dest prim.MultiAddress, amount uint128.Uint128) Transaction {
	call := baPallet.CallTransferKeepAlive{Dest: dest, Value: amount}
	return NewTransaction(this.client, call.ToPayload())
}

type StakingTx struct {
	client *Client
}

// Take the origin account as a stash and lock up `value` of its balance. `controller` will
// be the account that controls it.
func (this *StakingTx) Bond(value uint128.Uint128, payee metadata.RewardDestination) Transaction {
	call := stPallet.CallBond{Value: value, Payee: payee}
	return NewTransaction(this.client, call.ToPayload())
}

// Add some extra amount that have appeared in the stash `free_balance` into the balance up
// for staking.
func (this *StakingTx) BondExtra(maxAdditional uint128.Uint128) Transaction {
	call := stPallet.CallBondExtra{MaxAdditional: maxAdditional}
	return NewTransaction(this.client, call.ToPayload())
}

// Schedule a portion of the stash to be unlocked ready for transfer out after the bond
// period ends. If this leaves an amount actively bonded less than
// T::Currency::minimum_balance(), then it is increased to the full amount.
func (this *StakingTx) Unbond(value uint128.Uint128) Transaction {
	call := stPallet.CallUnbond{Value: value}
	return NewTransaction(this.client, call.ToPayload())
}

// Remove any unlocked chunks from the `unlocking` queue from our management.
//
// This essentially frees up that balance to be used by the stash account to do whatever
// it wants.
func (this *StakingTx) WithdrawUnbonded(numSlashingSpans uint32) Transaction {
	call := stPallet.CallWithdrawUnbonded{NumSlashingSpans: numSlashingSpans}
	return NewTransaction(this.client, call.ToPayload())
}

// Declare the desire to validate for the origin controller.
//
// Effects will be felt at the beginning of the next era.
func (this *StakingTx) Validate(prefs metadata.ValidatorPrefs) Transaction {
	call := stPallet.CallValidate{Prefs: prefs}
	return NewTransaction(this.client, call.ToPayload())
}

// Declare the desire to nominate `targets` for the origin controller.
//
// Effects will be felt at the beginning of the next era.
func (this *StakingTx) Nominate(targets []prim.MultiAddress) Transaction {
	call := stPallet.CallNominate{Targets: targets}
	return NewTransaction(this.client, call.ToPayload())
}

// Declare no desire to either validate or nominate.
//
// Effects will be felt at the beginning of the next era.
func (this *StakingTx) Chill() Transaction {
	call := stPallet.CallChill{}
	return NewTransaction(this.client, call.ToPayload())
}

// (Re-)set the payment target for a controller.
//
// Effects will be felt instantly (as soon as this function is completed successfully).
func (this *StakingTx) SetPayee(payee metadata.RewardDestination) Transaction {
	call := stPallet.CallSetPayee{Payee: payee}
	return NewTransaction(this.client, call.ToPayload())
}

// (Re-)sets the controller of a stash to the stash itself. This function previously
// accepted a `controller` argument to set the controller to an account other than the
// stash itself. This functionality has now been removed, now only setting the controller
// to the stash, if it is not already.
//
// Effects will be felt instantly (as soon as this function is completed successfully).
func (this *StakingTx) SetController() Transaction {
	call := stPallet.CallSetController{}
	return NewTransaction(this.client, call.ToPayload())
}

// Pay out next page of the stakers behind a validator for the given era.
//
// - `validator_stash` is the stash account of the validator.
// - `era` may be any era between `[current_era - history_depth; current_era]`.
func (this *StakingTx) PayoutStakers(validatorStash metadata.AccountId, era uint32) Transaction {
	call := stPallet.CallPayoutStakers{ValidatorStash: validatorStash, Era: era}
	return NewTransaction(this.client, call.ToPayload())
}

// Rebond a portion of the stash scheduled to be unlocked.
func (this *StakingTx) Rebond(value uint128.Uint128) Transaction {
	call := stPallet.CallRebond{Value: value}
	return NewTransaction(this.client, call.ToPayload())
}

// Remove all data structures concerning a staker/stash once it is at a state where it can
// be considered `dust` in the staking system. The requirements are:
//
// 1. the `total_balance` of the stash is below existential deposit.
// 2. or, the `ledger.total` of the stash is below existential deposit.
//
// The former can happen in cases like a slash; the latter when a fully unbonded account
// is still receiving staking rewards in `RewardDestination::Staked`.
//
// It can be called by anyone, as long as `stash` meets the above requirements.
//
// Refunds the transaction fees upon successful execution.
func (this *StakingTx) ReapStash(stash metadata.AccountId, numSlashingSpans uint32) Transaction {
	call := stPallet.CallReapStash{Stash: stash, NumSlashingSpans: numSlashingSpans}
	return NewTransaction(this.client, call.ToPayload())
}

// Remove the given nominations from the calling validator.
//
// Effects will be felt at the beginning of the next era.
func (this *StakingTx) Kick(who []prim.MultiAddress) Transaction {
	call := stPallet.CallKick{Who: who}
	return NewTransaction(this.client, call.ToPayload())
}

// Declare a `controller` to stop participating as either a validator or nominator.
//
// Effects will be felt at the beginning of the next era.
//
// The dispatch origin for this call must be _Signed_, but can be called by anyone.
//
// If the caller is the same as the controller being targeted, then no further checks are
// enforced, and this function behaves just like `chill`.
//
// If the caller is different than the controller being targeted, the following conditions
// must be met:
//
// * `controller` must belong to a nominator who has become non-decodable,
//
// Or:
//
//   - A `ChillThreshold` must be set and checked which defines how close to the max
//     nominators or validators we must reach before users can start chilling one-another.
//   - A `MaxNominatorCount` and `MaxValidatorCount` must be set which is used to determine
//     how close we are to the threshold.
//   - A `MinNominatorBond` and `MinValidatorBond` must be set and checked, which determines
//     if this is a person that should be chilled because they have not met the threshold
//     bond required.
//
// This can be helpful if bond requirements are updated, and we need to remove old users
// who do not satisfy these requirements.
func (this *StakingTx) ChillOther(stash metadata.AccountId) Transaction {
	call := stPallet.CallChillOther{Stash: stash}
	return NewTransaction(this.client, call.ToPayload())
}

// Force a validator to have at least the minimum commission. This will not affect a
// validator who already has a commission greater than or equal to the minimum. Any account
// can call this.
func (this *StakingTx) ForceApplyMinCommission(validatorStash metadata.AccountId) Transaction {
	call := stPallet.CallForceApplyMinCommission{ValidatorStash: validatorStash}
	return NewTransaction(this.client, call.ToPayload())
}

// Pay out a page of the stakers behind a validator for the given era and page.
//
//   - `validator_stash` is the stash account of the validator.
//   - `era` may be any era between `[current_era - history_depth; current_era]`.
//   - `page` is the page index of nominators to pay out with value between 0 and
//     `num_nominators / T::MaxExposurePageSize`.
//
// The origin of this call must be _Signed_. Any account can call this function, even if
// it is not one of the stakers.
//
// If a validator has more than [`Config::MaxExposurePageSize`] nominators backing
// them, then the list of nominators is paged, with each page being capped at
// [`Config::MaxExposurePageSize`.] If a validator has more than one page of nominators,
// the call needs to be made for each page separately in order for all the nominators
// backing a validator to receive the reward. The nominators are not sorted across pages
// and so it should not be assumed the highest staker would be on the topmost page and vice
// versa. If rewards are not claimed in [`Config::HistoryDepth`] eras, they are lost.
func (this *StakingTx) PayoutStakersByPage(validatorStash metadata.AccountId, era uint32, page uint32) Transaction {
	call := stPallet.CallPayoutStakersByPage{ValidatorStash: validatorStash, Era: era, Page: page}
	return NewTransaction(this.client, call.ToPayload())
}

// Migrates an account's `RewardDestination::Controller` to
// `RewardDestination::Account(controller)`.
//
// Effects will be felt instantly (as soon as this function is completed successfully).
//
// This will waive the transaction fee if the `payee` is successfully migrated.
func (this *StakingTx) UpdatePayee(controller metadata.AccountId) Transaction {
	call := stPallet.CallUpdatePayee{Controller: controller}
	return NewTransaction(this.client, call.ToPayload())
}
