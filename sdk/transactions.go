package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"github.com/availproject/avail-go-sdk/metadata"
	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	npPallet "github.com/availproject/avail-go-sdk/metadata/pallets/nomination_pools"
	stPallet "github.com/availproject/avail-go-sdk/metadata/pallets/staking"
	sdPallet "github.com/availproject/avail-go-sdk/metadata/pallets/sudo"
	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	utPallet "github.com/availproject/avail-go-sdk/metadata/pallets/utility"
	vcPallet "github.com/availproject/avail-go-sdk/metadata/pallets/vector"
	prim "github.com/availproject/avail-go-sdk/primitives"
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
func (this *BalancesTx) TransferAllowDeath(dest prim.MultiAddress, amount metadata.Balance) Transaction {
	call := baPallet.CallTransferAlowDeath{Dest: dest, Value: amount.Value}
	return NewTransaction(this.client, call.ToPayload())
}

// Exactly as `TransferAlowDeath`, except the origin must be root and the source account
// may be specified
func (this *BalancesTx) ForceTransfer(dest prim.MultiAddress, amount metadata.Balance) Transaction {
	call := baPallet.CallForceTransfer{Dest: dest, Value: amount}
	return NewTransaction(this.client, call.ToPayload())
}

// Same as the `TransferAlowDeath` call, but with a check that the transfer will not
// kill the origin account.
func (this *BalancesTx) TransferKeepAlive(dest prim.MultiAddress, amount metadata.Balance) Transaction {
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

type NominationPoolsTx struct {
	client *Client
}

// Stake funds with a pool. The amount to bond is transferred from the member to the
// pools account and immediately increases the pools bond.
//
// # Note
//
//   - An account can only be a member of a single pool.
//   - An account cannot join the same pool multiple times.
//   - This call will *not* dust the member account, so the member must have at least
//     `existential deposit + amount` in their account.
//   - Only a pool with [`PoolState::Open`] can be joined
func (this *NominationPoolsTx) Join(amount metadata.Balance, poolId uint32) Transaction {
	call := npPallet.CallJoin{Amount: amount, PoolId: poolId}
	return NewTransaction(this.client, call.ToPayload())
}

// Bond `extra` more funds from `origin` into the pool to which they already belong.
//
// Additional funds can come from either the free balance of the account, of from the
// accumulated rewards, see [`BondExtra`].
//
// Bonding extra funds implies an automatic payout of all pending rewards as well.
// See `bond_extra_other` to bond pending rewards of `other` members.
func (this *NominationPoolsTx) BondExtra(extra metadata.PoolBondExtra) Transaction {
	call := npPallet.CallBondExtra{Extra: extra}
	return NewTransaction(this.client, call.ToPayload())
}

// A bonded member can use this to claim their payout based on the rewards that the pool
// has accumulated since their last claimed payout (OR since joining if this is their first
// time claiming rewards). The payout will be transferred to the member's account.
//
// The member will earn rewards pro rata based on the members stake vs the sum of the
// members in the pools stake. Rewards do not "expire".
//
// See `claim_payout_other` to caim rewards on bahalf of some `other` pool member.
func (this *NominationPoolsTx) ClaimPayout() Transaction {
	call := npPallet.CallClaimPayout{}
	return NewTransaction(this.client, call.ToPayload())
}

// Unbond up to `unbonding_points` of the `member_account`'s funds from the pool. It
// implicitly collects the rewards one last time, since not doing so would mean some
// rewards would be forfeited.
//
// Under certain conditions, this call can be dispatched permissionlessly (i.e. by any
// account).
//
// # Conditions for a permissionless dispatch.
//
//   - The pool is blocked and the caller is either the root or bouncer. This is refereed to
//     as a kick.
//   - The pool is destroying and the member is not the depositor.
//   - The pool is destroying, the member is the depositor and no other members are in the
//     pool.
//
// ## Conditions for permissioned dispatch (i.e. the caller is also the
// `member_account`):
//
//   - The caller is not the depositor.
//   - The caller is the depositor, the pool is destroying and no other members are in the
//     pool.
//
// # Note
//
// If there are too many unlocking chunks to unbond with the pool account,
// [`Call::pool_withdraw_unbonded`] can be called to try and minimize unlocking chunks.
// The [`StakingInterface::unbond`] will implicitly call [`Call::pool_withdraw_unbonded`]
// to try to free chunks if necessary (ie. if unbound was called and no unlocking chunks
// are available). However, it may not be possible to release the current unlocking chunks,
// in which case, the result of this call will likely be the `NoMoreChunks` error from the
// staking system.
func (this *NominationPoolsTx) Unbond(memberAccount prim.MultiAddress, unbondingPoints uint128.Uint128) Transaction {
	call := npPallet.CallUnbond{MemberAccount: memberAccount, UnbondingPoints: unbondingPoints}
	return NewTransaction(this.client, call.ToPayload())
}

// Call `withdraw_unbonded` for the pools account. This call can be made by any account.
//
// This is useful if there are too many unlocking chunks to call `unbond`, and some
// can be cleared by withdrawing. In the case there are too many unlocking chunks, the user
// would probably see an error like `NoMoreChunks` emitted from the staking system when
// they attempt to unbond.
func (this *NominationPoolsTx) PoolWithdrawUnbonded(poolId uint32, numSlashingSpans uint32) Transaction {
	call := npPallet.CallPoolWithdrawUnbonded{PoolId: poolId, NumSlashingSpans: numSlashingSpans}
	return NewTransaction(this.client, call.ToPayload())
}

// Withdraw unbonded funds from `member_account`. If no bonded funds can be unbonded, an
// error is returned.
//
// Under certain conditions, this call can be dispatched permissionlessly (i.e. by any
// account).
//
// # Conditions for a permissionless dispatch
//
// * The pool is in destroy mode and the target is not the depositor.
// * The target is the depositor and they are the only member in the sub pools.
// * The pool is blocked and the caller is either the root or bouncer.
//
// # Conditions for permissioned dispatch
//
// * The caller is the target and they are not the depositor.
//
// # Note
//
// If the target is the depositor, the pool will be destroyed.
func (this *NominationPoolsTx) WithdrawUnbonded(memberAccount prim.MultiAddress, numSlashingSpans uint32) Transaction {
	call := npPallet.CallWithdrawUnbonded{MemberAccount: memberAccount, NumSlashingSpans: numSlashingSpans}
	return NewTransaction(this.client, call.ToPayload())
}

// Create a new delegation pool.
//
// # Arguments
//
//   - `amount` - The amount of funds to delegate to the pool. This also acts of a sort of
//     deposit since the pools creator cannot fully unbond funds until the pool is being
//     destroyed.
//   - `index` - A disambiguation index for creating the account. Likely only useful when
//     creating multiple pools in the same extrinsic.
//   - `root` - The account to set as [`PoolRoles::root`].
//   - `nominator` - The account to set as the [`PoolRoles::nominator`].
//   - `bouncer` - The account to set as the [`PoolRoles::bouncer`].
//
// # Note
//
// In addition to `amount`, the caller will transfer the existential deposit; so the caller
// needs at have at least `amount + existential_deposit` transferable.
func (this *NominationPoolsTx) Create(amount metadata.Balance, root prim.MultiAddress, nominator prim.MultiAddress, bouncer prim.MultiAddress) Transaction {
	call := npPallet.CallCreate{Amount: amount, Root: root, Nominator: nominator, Bouncer: bouncer}
	return NewTransaction(this.client, call.ToPayload())
}

// Create a new delegation pool with a previously used pool id
//
// # Arguments
//
// same as `create` with the inclusion of
// * `pool_id` - `A valid PoolId.
func (this *NominationPoolsTx) CreateWithPoolId(amount metadata.Balance, root prim.MultiAddress, nominator prim.MultiAddress, bouncer prim.MultiAddress, poolId uint32) Transaction {
	call := npPallet.CallCreateWithPoolId{Amount: amount, Root: root, Nominator: nominator, Bouncer: bouncer, PoolId: poolId}
	return NewTransaction(this.client, call.ToPayload())
}

// Nominate on behalf of the pool.
//
// The dispatch origin of this call must be signed by the pool nominator or the pool
// root role.
//
// This directly forward the call to the staking pallet, on behalf of the pool bonded
// account.
func (this *NominationPoolsTx) Nominate(poolId uint32, validators []metadata.AccountId) Transaction {
	call := npPallet.CallNominate{PoolId: poolId, Validators: validators}
	return NewTransaction(this.client, call.ToPayload())
}

// Set a new state for the pool.
//
// If a pool is already in the `Destroying` state, then under no condition can its state
// change again.
//
// The dispatch origin of this call must be either:
//
//  1. signed by the bouncer, or the root role of the pool,
//  2. if the pool conditions to be open are NOT met (as described by `ok_to_be_open`), and
//     then the state of the pool can be permissionlessly changed to `Destroying`.
func (this *NominationPoolsTx) SetState(poolId uint32, state metadata.PoolState) Transaction {
	call := npPallet.CallSetState{PoolId: poolId, State: state}
	return NewTransaction(this.client, call.ToPayload())
}

// Set a new metadata for the pool.
//
// The dispatch origin of this call must be signed by the bouncer, or the root role of the
// pool.
func (this *NominationPoolsTx) SetMetadata(poolId uint32, metadata []byte) Transaction {
	call := npPallet.CallSetMetadata{PoolId: poolId, Metadata: metadata}
	return NewTransaction(this.client, call.ToPayload())
}

// Update the roles of the pool.
//
// The root is the only entity that can change any of the roles, including itself,
// excluding the depositor, who can never change.
//
// It emits an event, notifying UIs of the role change. This event is quite relevant to
// most pool members and they should be informed of changes to pool roles.
func (this *NominationPoolsTx) UpdateRoles(poolId uint32, newRoot metadata.PoolRoleConfig, newNominator metadata.PoolRoleConfig, newBouncer metadata.PoolRoleConfig) Transaction {
	call := npPallet.CallUpdateRoles{PoolId: poolId, NewRoot: newRoot, NewNominator: newNominator, NewBouncer: newBouncer}
	return NewTransaction(this.client, call.ToPayload())
}

// Chill on behalf of the pool.
//
// The dispatch origin of this call must be signed by the pool nominator or the pool
// root role, same as [`Pallet::nominate`].
//
// This directly forward the call to the staking pallet, on behalf of the pool bonded
// account.
func (this *NominationPoolsTx) Chill(poolId uint32) Transaction {
	call := npPallet.CallChill{PoolId: poolId}
	return NewTransaction(this.client, call.ToPayload())
}

// `origin` bonds funds from `extra` for some pool member `member` into their respective
// pools.
//
// `origin` can bond extra funds from free balance or pending rewards when `origin ==
// other`.
//
// In the case of `origin != other`, `origin` can only bond extra pending rewards of
// `other` members assuming set_claim_permission for the given member is
// `PermissionlessAll` or `PermissionlessCompound`.
func (this *NominationPoolsTx) BondExtraOther(member prim.MultiAddress, extra metadata.PoolBondExtra) Transaction {
	call := npPallet.CallBondExtraOther{Member: member, Extra: extra}
	return NewTransaction(this.client, call.ToPayload())
}

// Allows a pool member to set a claim permission to allow or disallow permissionless
// bonding and withdrawing.
//
// By default, this is `Permissioned`, which implies only the pool member themselves can
// claim their pending rewards. If a pool member wishes so, they can set this to
// `PermissionlessAll` to allow any account to claim their rewards and bond extra to the
// pool.
//
// # Arguments
//
// * `origin` - Member of a pool.
// * `actor` - Account to claim reward. // improve this
func (this *NominationPoolsTx) SetClaimPermission(permission metadata.PoolClaimPermission) Transaction {
	call := npPallet.CallSetClaimPermission{Permission: permission}
	return NewTransaction(this.client, call.ToPayload())
}

// `origin` can claim payouts on some pool member `other`'s behalf.
//
// Pool member `other` must have a `PermissionlessAll` or `PermissionlessWithdraw` in order
// for this call to be successful.
func (this *NominationPoolsTx) ClaimPayoutOther(other metadata.AccountId) Transaction {
	call := npPallet.CallClaimPayoutOther{Other: other}
	return NewTransaction(this.client, call.ToPayload())
}

// Set the commission of a pool.
//
// Both a commission percentage and a commission payee must be provided in the `current`
// tuple. Where a `current` of `None` is provided, any current commission will be removed.
//
// - If a `None` is supplied to `new_commission`, existing commission will be removed.
func (this *NominationPoolsTx) SetCommission(poolId uint32, newCommission prim.Option[metadata.Tuple2[metadata.Perbill, metadata.AccountId]]) Transaction {
	call := npPallet.CallSetCommission{PoolId: poolId, NewCommission: newCommission}
	return NewTransaction(this.client, call.ToPayload())
}

// Set the maximum commission of a pool.
//
//   - Initial max can be set to any `Perbill`, and only smaller values thereafter.
//   - Current commission will be lowered in the event it is higher than a new max
//     commission.
func (this *NominationPoolsTx) SetCommissionMax(poolId uint32, maxCommission metadata.Perbill) Transaction {
	call := npPallet.CallSetCommissionMax{PoolId: poolId, MaxCommission: maxCommission}
	return NewTransaction(this.client, call.ToPayload())
}

// Set the commission change rate for a pool.
//
// Initial change rate is not bounded, whereas subsequent updates can only be more
// restrictive than the current.
func (this *NominationPoolsTx) SetCommissionChangeRate(poolId uint32, changeRate metadata.PoolCommissionChangeRate) Transaction {
	call := npPallet.CallSetCommissionChangeRate{PoolId: poolId, ChangeRate: changeRate}
	return NewTransaction(this.client, call.ToPayload())
}

// Claim pending commission.
//
// The dispatch origin of this call must be signed by the `root` role of the pool. Pending
// commission is paid out and added to total claimed commission`. Total pending commission
// is reset to zero. the current.
func (this *NominationPoolsTx) ClaimCommission(poolId uint32) Transaction {
	call := npPallet.CallClaimCommission{PoolId: poolId}
	return NewTransaction(this.client, call.ToPayload())
}

// Top up the deficit or withdraw the excess ED from the pool.
//
// When a pool is created, the pool depositor transfers ED to the reward account of the
// pool. ED is subject to change and over time, the deposit in the reward account may be
// insufficient to cover the ED deficit of the pool or vice-versa where there is excess
// deposit to the pool. This call allows anyone to adjust the ED deposit of the
// pool by either topping up the deficit or claiming the excess.
func (this *NominationPoolsTx) AdjustPoolDeposit(poolId uint32) Transaction {
	call := npPallet.CallAdjustPoolDeposit{PoolId: poolId}
	return NewTransaction(this.client, call.ToPayload())
}

// Set or remove a pool's commission claim permission.
//
// Determines who can claim the pool's pending commission. Only the `Root` role of the pool
// is able to conifigure commission claim permissions.
func (this *NominationPoolsTx) SetCommissionClaimPermission(poolId uint32, permission prim.Option[metadata.CommissionClaimPermission]) Transaction {
	call := npPallet.CallSetCommissionClaimPermission{PoolId: poolId, Permission: permission}
	return NewTransaction(this.client, call.ToPayload())
}

type SystemTx struct {
	client *Client
}

// Make some on-chain remark.
//
// Can be executed by every `origin`
func (this *SystemTx) Remark(remark []byte) Transaction {
	call := syPallet.CallRemark{Remark: remark}
	return NewTransaction(this.client, call.ToPayload())
}

// Make some on-chain remark and emit event
func (this *SystemTx) RemarkWithEvent(remark []byte) Transaction {
	call := syPallet.CallRemarkWithEvent{Remark: remark}
	return NewTransaction(this.client, call.ToPayload())
}

type VectorTx struct {
	client *Client
}

func (this *VectorTx) SendMessage(message metadata.VectorMessageKind, To prim.H256, domain uint32) Transaction {
	call := vcPallet.CallSendMessage{Message: message, To: To, Domain: domain}
	return NewTransaction(this.client, call.ToPayload())
}

type SudoTx struct {
	client *Client
}

// Authenticates the sudo key and dispatches a function call with `Root` origin.
func (this *SudoTx) Sudo(call prim.Call) Transaction {
	c := sdPallet.CallSudo{Call: call}
	return NewTransaction(this.client, c.ToPayload())
}

// Authenticates the sudo key and dispatches a function call with `Root` origin.
// This function does not check the weight of the call, and instead allows the
// Sudo user to specify the weight of the call.
//
// The dispatch origin for this call must be _Signed_.
func (this *SudoTx) SudoUncheckedWeight(call prim.Call) Transaction {
	c := sdPallet.CallSudoUncheckedWeight{Call: call}
	return NewTransaction(this.client, c.ToPayload())
}

// Authenticates the sudo key and dispatches a function call with `Signed` origin from
// a given account.
//
// The dispatch origin for this call must be _Signed_.
func (this *SudoTx) SudoAs(who prim.MultiAddress, call prim.Call) Transaction {
	c := sdPallet.CallSudoAs{Who: who, Call: call}
	return NewTransaction(this.client, c.ToPayload())
}
