package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/metadata/pallets"
	baPallet "github.com/availproject/avail-go-sdk/metadata/pallets/balances"
	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	npPallet "github.com/availproject/avail-go-sdk/metadata/pallets/nomination_pools"
	pxPallet "github.com/availproject/avail-go-sdk/metadata/pallets/proxy"
	sePallet "github.com/availproject/avail-go-sdk/metadata/pallets/session"
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

func (dat *DataAvailabilityTx) SubmitData(data []byte) Transaction {
	call := daPallet.CallSubmitData{Data: data}
	return NewTransaction(dat.client, pallets.ToPayload(call))
}

func (dat *DataAvailabilityTx) CreateApplicationKey(key []byte) Transaction {
	call := daPallet.CallCreateApplicationKey{Key: key}
	return NewTransaction(dat.client, pallets.ToPayload(call))
}

type UtilityTx struct {
	client *Client
}

// Send a batch of dispatch calls.
//
// May be called from any origin except `None`.
func (ut *UtilityTx) Batch(calls []prim.Call) Transaction {
	call := utPallet.CallBatch{Calls: calls}
	return NewTransaction(ut.client, pallets.ToPayload(call))
}

// Send a batch of dispatch calls and atomically execute them.
// The whole transaction will rollback and fail if any of the calls failed.
//
// May be called from any origin except `None`.
func (ut *UtilityTx) BatchAll(calls []prim.Call) Transaction {
	call := utPallet.CallBatchAll{Calls: calls}
	return NewTransaction(ut.client, pallets.ToPayload(call))
}

// Send a batch of dispatch calls.
// Unlike `batch`, it allows errors and won't interrupt.
//
// May be called from any origin except `None`.
func (ut *UtilityTx) ForceBatch(calls []prim.Call) Transaction {
	call := utPallet.CallForceBatch{Calls: calls}
	return NewTransaction(ut.client, pallets.ToPayload(call))
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
func (ut *UtilityTx) AsDerivate(index uint16, call prim.Call) Transaction {
	c := utPallet.CallAsDerivate{Index: index, Call: call}
	return NewTransaction(ut.client, pallets.ToPayload(c))
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
func (bt *BalancesTx) TransferAllowDeath(dest prim.MultiAddress, amount metadata.Balance) Transaction {
	call := baPallet.CallTransferAlowDeath{Dest: dest, Value: amount.Value}
	return NewTransaction(bt.client, pallets.ToPayload(call))
}

// Exactly as `TransferAlowDeath`, except the origin must be root and the source account
// may be specified
func (bt *BalancesTx) ForceTransfer(dest prim.MultiAddress, amount metadata.Balance) Transaction {
	call := baPallet.CallForceTransfer{Dest: dest, Value: amount}
	return NewTransaction(bt.client, pallets.ToPayload(call))
}

// Transfer the entire transferable balance from the caller account.
//
// NOTE: This function only attempts to transfer _transferable_ balances. This means that
// any locked, reserved, or existential deposits (when `keep_alive` is `true`), will not be
// transferred by this function.
func (bt *BalancesTx) TransferAll(dest prim.MultiAddress, keepAlive bool) Transaction {
	call := baPallet.CallTransferAll{Dest: dest, KeepAlive: keepAlive}
	return NewTransaction(bt.client, pallets.ToPayload(call))
}

// Same as the `TransferAlowDeath` call, but with a check that the transfer will not
// kill the origin account.
func (bt *BalancesTx) TransferKeepAlive(dest prim.MultiAddress, amount metadata.Balance) Transaction {
	call := baPallet.CallTransferKeepAlive{Dest: dest, Value: amount}
	return NewTransaction(bt.client, pallets.ToPayload(call))
}

type StakingTx struct {
	client *Client
}

// Take the origin account as a stash and lock up `value` of its balance. `controller` will
// be the account that controls it.
func (st *StakingTx) Bond(value metadata.Balance, payee metadata.RewardDestination) Transaction {
	call := stPallet.CallBond{Value: value, Payee: payee}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Add some extra amount that have appeared in the stash `free_balance` into the balance up
// for staking.
func (st *StakingTx) BondExtra(maxAdditional metadata.Balance) Transaction {
	call := stPallet.CallBondExtra{MaxAdditional: maxAdditional}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Schedule a portion of the stash to be unlocked ready for transfer out after the bond
// period ends. If this leaves an amount actively bonded less than
// T::Currency::minimum_balance(), then it is increased to the full amount.
func (st *StakingTx) Unbond(value metadata.Balance) Transaction {
	call := stPallet.CallUnbond{Value: value}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Remove any unlocked chunks from the `unlocking` queue from our management.
//
// This essentially frees up that balance to be used by the stash account to do whatever
// it wants.
func (st *StakingTx) WithdrawUnbonded(numSlashingSpans uint32) Transaction {
	call := stPallet.CallWithdrawUnbonded{NumSlashingSpans: numSlashingSpans}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Declare the desire to validate for the origin controller.
//
// Effects will be felt at the beginning of the next era.
func (st *StakingTx) Validate(prefs metadata.ValidatorPrefs) Transaction {
	call := stPallet.CallValidate{Prefs: prefs}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Declare the desire to nominate `targets` for the origin controller.
//
// Effects will be felt at the beginning of the next era.
func (st *StakingTx) Nominate(targets []prim.MultiAddress) Transaction {
	call := stPallet.CallNominate{Targets: targets}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Declare no desire to either validate or nominate.
//
// Effects will be felt at the beginning of the next era.
func (st *StakingTx) Chill() Transaction {
	call := stPallet.CallChill{}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// (Re-)set the payment target for a controller.
//
// Effects will be felt instantly (as soon as this function is completed successfully).
func (st *StakingTx) SetPayee(payee metadata.RewardDestination) Transaction {
	call := stPallet.CallSetPayee{Payee: payee}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// (Re-)sets the controller of a stash to the stash itself. This function previously
// accepted a `controller` argument to set the controller to an account other than the
// stash itself. This functionality has now been removed, now only setting the controller
// to the stash, if it is not already.
//
// Effects will be felt instantly (as soon as this function is completed successfully).
func (st *StakingTx) SetController() Transaction {
	call := stPallet.CallSetController{}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Pay out next page of the stakers behind a validator for the given era.
//
// - `validator_stash` is the stash account of the validator.
// - `era` may be any era between `[current_era - history_depth; current_era]`.
func (st *StakingTx) PayoutStakers(validatorStash prim.AccountId, era uint32) Transaction {
	call := stPallet.CallPayoutStakers{ValidatorStash: validatorStash, Era: era}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Rebond a portion of the stash scheduled to be unlocked.
func (st *StakingTx) Rebond(value uint128.Uint128) Transaction {
	call := stPallet.CallRebond{Value: value}
	return NewTransaction(st.client, pallets.ToPayload(call))
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
func (st *StakingTx) ReapStash(stash prim.AccountId, numSlashingSpans uint32) Transaction {
	call := stPallet.CallReapStash{Stash: stash, NumSlashingSpans: numSlashingSpans}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Remove the given nominations from the calling validator.
//
// Effects will be felt at the beginning of the next era.
func (st *StakingTx) Kick(who []prim.MultiAddress) Transaction {
	call := stPallet.CallKick{Who: who}
	return NewTransaction(st.client, pallets.ToPayload(call))
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
func (st *StakingTx) ChillOther(stash prim.AccountId) Transaction {
	call := stPallet.CallChillOther{Stash: stash}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Force a validator to have at least the minimum commission. This will not affect a
// validator who already has a commission greater than or equal to the minimum. Any account
// can call s.
func (st *StakingTx) ForceApplyMinCommission(validatorStash prim.AccountId) Transaction {
	call := stPallet.CallForceApplyMinCommission{ValidatorStash: validatorStash}
	return NewTransaction(st.client, pallets.ToPayload(call))
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
func (st *StakingTx) PayoutStakersByPage(validatorStash prim.AccountId, era uint32, page uint32) Transaction {
	call := stPallet.CallPayoutStakersByPage{ValidatorStash: validatorStash, Era: era, Page: page}
	return NewTransaction(st.client, pallets.ToPayload(call))
}

// Migrates an account's `RewardDestination::Controller` to
// `RewardDestination::Account(controller)`.
//
// Effects will be felt instantly (as soon as this function is completed successfully).
//
// This will waive the transaction fee if the `payee` is successfully migrated.
func (st *StakingTx) UpdatePayee(controller prim.AccountId) Transaction {
	call := stPallet.CallUpdatePayee{Controller: controller}
	return NewTransaction(st.client, pallets.ToPayload(call))
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
func (npt *NominationPoolsTx) Join(amount metadata.Balance, poolId uint32) Transaction {
	call := npPallet.CallJoin{Amount: amount, PoolId: poolId}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Bond `extra` more funds from `origin` into the pool to which they already belong.
//
// Additional funds can come from either the free balance of the account, of from the
// accumulated rewards, see [`BondExtra`].
//
// Bonding extra funds implies an automatic payout of all pending rewards as well.
// See `bond_extra_other` to bond pending rewards of `other` members.
func (npt *NominationPoolsTx) BondExtra(extra metadata.PoolBondExtra) Transaction {
	call := npPallet.CallBondExtra{Extra: extra}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// A bonded member can use this to claim their payout based on the rewards that the pool
// has accumulated since their last claimed payout (OR since joining if this is their first
// time claiming rewards). The payout will be transferred to the member's account.
//
// The member will earn rewards pro rata based on the members stake vs the sum of the
// members in the pools stake. Rewards do not "expire".
//
// See `claim_payout_other` to caim rewards on bahalf of some `other` pool member.
func (npt *NominationPoolsTx) ClaimPayout() Transaction {
	call := npPallet.CallClaimPayout{}
	return NewTransaction(npt.client, pallets.ToPayload(call))
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
func (npt *NominationPoolsTx) Unbond(memberAccount prim.MultiAddress, unbondingPoints uint128.Uint128) Transaction {
	call := npPallet.CallUnbond{MemberAccount: memberAccount, UnbondingPoints: unbondingPoints}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Call `withdraw_unbonded` for the pools account. This call can be made by any account.
//
// This is useful if there are too many unlocking chunks to call `unbond`, and some
// can be cleared by withdrawing. In the case there are too many unlocking chunks, the user
// would probably see an error like `NoMoreChunks` emitted from the staking system when
// they attempt to unbond.
func (npt *NominationPoolsTx) PoolWithdrawUnbonded(poolId uint32, numSlashingSpans uint32) Transaction {
	call := npPallet.CallPoolWithdrawUnbonded{PoolId: poolId, NumSlashingSpans: numSlashingSpans}
	return NewTransaction(npt.client, pallets.ToPayload(call))
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
func (npt *NominationPoolsTx) WithdrawUnbonded(memberAccount prim.MultiAddress, numSlashingSpans uint32) Transaction {
	call := npPallet.CallWithdrawUnbonded{MemberAccount: memberAccount, NumSlashingSpans: numSlashingSpans}
	return NewTransaction(npt.client, pallets.ToPayload(call))
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
func (npt *NominationPoolsTx) Create(amount metadata.Balance, root prim.MultiAddress, nominator prim.MultiAddress, bouncer prim.MultiAddress) Transaction {
	call := npPallet.CallCreate{Amount: amount, Root: root, Nominator: nominator, Bouncer: bouncer}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Create a new delegation pool with a previously used pool id
//
// # Arguments
//
// same as `create` with the inclusion of
// * `pool_id` - `A valid PoolId.
func (npt *NominationPoolsTx) CreateWithPoolId(amount metadata.Balance, root prim.MultiAddress, nominator prim.MultiAddress, bouncer prim.MultiAddress, poolId uint32) Transaction {
	call := npPallet.CallCreateWithPoolId{Amount: amount, Root: root, Nominator: nominator, Bouncer: bouncer, PoolId: poolId}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Nominate on behalf of the pool.
//
// The dispatch origin of this call must be signed by the pool nominator or the pool
// root role.
//
// This directly forward the call to the staking pallet, on behalf of the pool bonded
// account.
func (npt *NominationPoolsTx) Nominate(poolId uint32, validators []prim.AccountId) Transaction {
	call := npPallet.CallNominate{PoolId: poolId, Validators: validators}
	return NewTransaction(npt.client, pallets.ToPayload(call))
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
func (npt *NominationPoolsTx) SetState(poolId uint32, state metadata.PoolState) Transaction {
	call := npPallet.CallSetState{PoolId: poolId, State: state}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Set a new metadata for the pool.
//
// The dispatch origin of this call must be signed by the bouncer, or the root role of the
// pool.
func (npt *NominationPoolsTx) SetMetadata(poolId uint32, metadata []byte) Transaction {
	call := npPallet.CallSetMetadata{PoolId: poolId, Metadata: metadata}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Update the roles of the pool.
//
// The root is the only entity that can change any of the roles, including itself,
// excluding the depositor, who can never change.
//
// It emits an event, notifying UIs of the role change. This event is quite relevant to
// most pool members and they should be informed of changes to pool roles.
func (npt *NominationPoolsTx) UpdateRoles(poolId uint32, newRoot metadata.PoolRoleConfig, newNominator metadata.PoolRoleConfig, newBouncer metadata.PoolRoleConfig) Transaction {
	call := npPallet.CallUpdateRoles{PoolId: poolId, NewRoot: newRoot, NewNominator: newNominator, NewBouncer: newBouncer}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Chill on behalf of the pool.
//
// The dispatch origin of this call must be signed by the pool nominator or the pool
// root role, same as [`Pallet::nominate`].
//
// This directly forward the call to the staking pallet, on behalf of the pool bonded
// account.
func (npt *NominationPoolsTx) Chill(poolId uint32) Transaction {
	call := npPallet.CallChill{PoolId: poolId}
	return NewTransaction(npt.client, pallets.ToPayload(call))
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
func (npt *NominationPoolsTx) BondExtraOther(member prim.MultiAddress, extra metadata.PoolBondExtra) Transaction {
	call := npPallet.CallBondExtraOther{Member: member, Extra: extra}
	return NewTransaction(npt.client, pallets.ToPayload(call))
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
func (npt *NominationPoolsTx) SetClaimPermission(permission metadata.PoolClaimPermission) Transaction {
	call := npPallet.CallSetClaimPermission{Permission: permission}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// `origin` can claim payouts on some pool member `other`'s behalf.
//
// Pool member `other` must have a `PermissionlessAll` or `PermissionlessWithdraw` in order
// for this call to be successful.
func (npt *NominationPoolsTx) ClaimPayoutOther(other prim.AccountId) Transaction {
	call := npPallet.CallClaimPayoutOther{Other: other}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Set the commission of a pool.
//
// Both a commission percentage and a commission payee must be provided in the `current`
// tuple. Where a `current` of `None` is provided, any current commission will be removed.
//
// - If a `None` is supplied to `new_commission`, existing commission will be removed.
func (npt *NominationPoolsTx) SetCommission(poolId uint32, newCommission prim.Option[metadata.Tuple2[metadata.Perbill, prim.AccountId]]) Transaction {
	call := npPallet.CallSetCommission{PoolId: poolId, NewCommission: newCommission}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Set the maximum commission of a pool.
//
//   - Initial max can be set to any `Perbill`, and only smaller values thereafter.
//   - Current commission will be lowered in the event it is higher than a new max
//     commission.
func (npt *NominationPoolsTx) SetCommissionMax(poolId uint32, maxCommission metadata.Perbill) Transaction {
	call := npPallet.CallSetCommissionMax{PoolId: poolId, MaxCommission: maxCommission}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Set the commission change rate for a pool.
//
// Initial change rate is not bounded, whereas subsequent updates can only be more
// restrictive than the current.
func (npt *NominationPoolsTx) SetCommissionChangeRate(poolId uint32, changeRate metadata.PoolCommissionChangeRate) Transaction {
	call := npPallet.CallSetCommissionChangeRate{PoolId: poolId, ChangeRate: changeRate}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Claim pending commission.
//
// The dispatch origin of this call must be signed by the `root` role of the pool. Pending
// commission is paid out and added to total claimed commission`. Total pending commission
// is reset to zero. the current.
func (npt *NominationPoolsTx) ClaimCommission(poolId uint32) Transaction {
	call := npPallet.CallClaimCommission{PoolId: poolId}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Top up the deficit or withdraw the excess ED from the pool.
//
// When a pool is created, the pool depositor transfers ED to the reward account of the
// pool. ED is subject to change and over time, the deposit in the reward account may be
// insufficient to cover the ED deficit of the pool or vice-versa where there is excess
// deposit to the pool. This call allows anyone to adjust the ED deposit of the
// pool by either topping up the deficit or claiming the excess.
func (npt *NominationPoolsTx) AdjustPoolDeposit(poolId uint32) Transaction {
	call := npPallet.CallAdjustPoolDeposit{PoolId: poolId}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

// Set or remove a pool's commission claim permission.
//
// Determines who can claim the pool's pending commission. Only the `Root` role of the pool
// is able to conifigure commission claim permissions.
func (npt *NominationPoolsTx) SetCommissionClaimPermission(poolId uint32, permission prim.Option[metadata.CommissionClaimPermission]) Transaction {
	call := npPallet.CallSetCommissionClaimPermission{PoolId: poolId, Permission: permission}
	return NewTransaction(npt.client, pallets.ToPayload(call))
}

type SystemTx struct {
	client *Client
}

// Make some on-chain remark.
//
// Can be executed by every `origin`
func (syt *SystemTx) Remark(remark []byte) Transaction {
	call := syPallet.CallRemark{Remark: remark}
	return NewTransaction(syt.client, pallets.ToPayload(call))
}

// Make some on-chain remark and emit event
func (syt *SystemTx) RemarkWithEvent(remark []byte) Transaction {
	call := syPallet.CallRemarkWithEvent{Remark: remark}
	return NewTransaction(syt.client, pallets.ToPayload(call))
}

type VectorTx struct {
	client *Client
}

func (vt *VectorTx) FulfillCall(functionId prim.H256, input []byte, output []byte, proof []byte, slot uint64) Transaction {
	call := vcPallet.CallFulfillCall{FunctionId: functionId, Input: input, Output: output, Proof: proof, Slot: slot}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) Execute(slot uint64, addrMessage metadata.VectorMessage, accountProof []byte, storageProof []byte) Transaction {
	call := vcPallet.CallExecute{Slot: slot, AddrMessage: addrMessage, AccountProof: accountProof, StorageProof: storageProof}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SourceChainFroze(sourceChainId uint32, frozen bool) Transaction {
	call := vcPallet.CallSourceChainFroze{SourceChainId: sourceChainId, Frozen: frozen}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SendMessage(message metadata.VectorMessage, To prim.H256, domain uint32) Transaction {
	call := vcPallet.CallSendMessage{Message: message, To: To, Domain: domain}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetPoseidonHash(period uint64, poseidonHash []byte) Transaction {
	call := vcPallet.CallSetPoseidonHash{Period: period, PoseidonHash: poseidonHash}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetBroadcaster(broadcasterDomain uint32, broadcaster prim.H256) Transaction {
	call := vcPallet.CallSetBroadcaster{BroadcasterDomain: broadcasterDomain, Broadcaster: broadcaster}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetWhitelistedDomains(value []uint32) Transaction {
	call := vcPallet.CallSetWhitelistedDomains{Value: value}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetConfiguration(value metadata.VectorConfiguration) Transaction {
	call := vcPallet.CallSetConfiguration{Value: value}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetFunctionsIds(value prim.Option[metadata.Tuple2[prim.H256, prim.H256]]) Transaction {
	call := vcPallet.CallSetFunctionsIds{Value: value}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetStepVerificationKey(value prim.Option[[]byte]) Transaction {
	call := vcPallet.CallSetStepVerificationKey{Value: value}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetRotateVerificationKey(value prim.Option[[]byte]) Transaction {
	call := vcPallet.CallSetRotateVerificationKey{Value: value}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) FailedSendMessageTxs(failedTxs []uint32) Transaction {
	call := vcPallet.CallFailedSendMessageTxs{FailedTxs: failedTxs}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetUpdater(updater prim.H256) Transaction {
	call := vcPallet.CallSetUpdater{Updater: updater}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) Fulfill(proof []byte, publicValues []byte) Transaction {
	call := vcPallet.CallFulfill{Proof: proof, PublicValues: publicValues}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetSp1VerificationKey(sp1Vk prim.H256) Transaction {
	call := vcPallet.CallSetSp1VerificationKey{Sp1Vk: sp1Vk}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) SetSyncCommitteeHash(period uint64, Hash prim.H256) Transaction {
	call := vcPallet.CallSetSyncCommitteeHash{Period: period, Hash: Hash}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) EnableMock(value bool) Transaction {
	call := vcPallet.CallEnableMock{Value: value}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

func (vt *VectorTx) MockFulfill(publicValues []byte) Transaction {
	call := vcPallet.CallMockFulfill{PublicValues: publicValues}
	return NewTransaction(vt.client, pallets.ToPayload(call))
}

type SudoTx struct {
	client *Client
}

// Authenticates the sudo key and dispatches a function call with `Root` origin.
func (sdt *SudoTx) Sudo(call prim.Call) Transaction {
	c := sdPallet.CallSudo{Call: call}
	return NewTransaction(sdt.client, pallets.ToPayload(c))
}

// Authenticates the sudo key and dispatches a function call with `Root` origin.
// This function does not check the weight of the call, and instead allows the
// Sudo user to specify the weight of the call.
//
// The dispatch origin for this call must be _Signed_.
func (sdt *SudoTx) SudoUncheckedWeight(call prim.Call) Transaction {
	c := sdPallet.CallSudoUncheckedWeight{Call: call}
	return NewTransaction(sdt.client, pallets.ToPayload(c))
}

// Authenticates the sudo key and dispatches a function call with `Signed` origin from
// a given account.
//
// The dispatch origin for this call must be _Signed_.
func (sdt *SudoTx) SudoAs(who prim.MultiAddress, call prim.Call) Transaction {
	c := sdPallet.CallSudoAs{Who: who, Call: call}
	return NewTransaction(sdt.client, pallets.ToPayload(c))
}

type SessionTx struct {
	client *Client
}

// Sets the session key(s) of the function caller to `keys`.
// Allows an account to set its session key prior to becoming a validator.
// This doesn't take effect until the next session.
func (set *SessionTx) SetKeys(keys metadata.SessionKeys, proof []byte) Transaction {
	call := sePallet.CallSetKeys{Keys: keys, Proof: proof}
	return NewTransaction(set.client, pallets.ToPayload(call))
}

// Removes any session key(s) of the function caller.
//
// This doesn't take effect until the next session.
//
// The dispatch origin of this function must be Signed and the account must be either be
// convertible to a validator ID using the chain's typical addressing system (this usually
// means being a controller account) or directly convertible into a validator ID (which
// usually means being a stash account).
func (set *SessionTx) PurgeKeys() Transaction {
	call := sePallet.CallPurgeKeys{}
	return NewTransaction(set.client, pallets.ToPayload(call))
}

type ProxyTx struct {
	client *Client
}

// Dispatch the given `call` from an account that the sender is authorised for through
// `add_proxy`.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Real`: The account that the proxy will make a call on behalf of.
// - `ForceProxyType`: Specify the exact proxy type to be used and checked for this call.
// - `Call`: The call to be made by the `real` account.
func (pt *ProxyTx) Proxy(real prim.MultiAddress, forceProxyType prim.Option[metadata.ProxyType], call prim.Call) Transaction {
	c := pxPallet.CallProxy{Real: real, ForceProxyType: forceProxyType, Call: call}
	return NewTransaction(pt.client, pallets.ToPayload(c))
}

// Register a proxy account for the sender that is able to make calls on its behalf.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Delegate`: The account that the `caller` would like to make a proxy.
// - `ProxyType`: The permissions allowed for this proxy account.
// - `Delay`: The announcement period required of the initial proxy. Will generally be
// zero.
func (pt *ProxyTx) AddProxy(delegate prim.MultiAddress, proxyType metadata.ProxyType, delay uint32) Transaction {
	call := pxPallet.CallAddProxy{Delegate: delegate, ProxyType: proxyType, Delay: delay}
	return NewTransaction(pt.client, pallets.ToPayload(call))
}

// Unregister a proxy account for the sender.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Delegate`: The account that the `caller` would like to remove as a proxy.
// - `ProxyType`: The permissions currently enabled for the removed proxy account.
// - `Delay`:  Will generally be zero.
func (pt *ProxyTx) RemoveProxy(delegate prim.MultiAddress, proxyType metadata.ProxyType, delay uint32) Transaction {
	call := pxPallet.CallRemoveProxy{Delegate: delegate, ProxyType: proxyType, Delay: delay}
	return NewTransaction(pt.client, pallets.ToPayload(call))
}

// Unregister all proxy accounts for the sender.
//
// The dispatch origin for this call must be _Signed_.
//
// WARNING: This may be called on accounts created by `pure`, however if done, then
// the unreserved fees will be inaccessible. **All access to this account will be lost.**
func (pt *ProxyTx) RemoveProxies() Transaction {
	call := pxPallet.CallRemoveProxies{}
	return NewTransaction(pt.client, pallets.ToPayload(call))
}

// Spawn a fresh new account that is guaranteed to be otherwise inaccessible, and
// initialize it with a proxy of `proxy_type` for `origin` sender.
//
// Requires a `Signed` origin.
//
// - `ProxyType`: The type of the proxy that the sender will be registered as over the
// new account. This will almost always be the most permissive `ProxyType` possible to
// allow for maximum flexibility.
// - `Index`: A disambiguation index, in case this is called multiple times in the same
// transaction (e.g. with `utility::batch`). Unless you're using `batch` you probably just
// want to use `0`.
// - `Delay`: The announcement period required of the initial proxy. Will generally be
// zero.
//
// Fails with `Duplicate` if this has already been called in this transaction, from the
// same sender, with the same parameters.
//
// Fails if there are insufficient funds to pay for deposit.
func (pt *ProxyTx) CreatePure(proxyType metadata.ProxyType, delay uint32, index uint16) Transaction {
	call := pxPallet.CallCreatePure{ProxyType: proxyType, Delay: delay, Index: index}
	return NewTransaction(pt.client, pallets.ToPayload(call))
}

// Removes a previously spawned pure proxy.
//
// WARNING: **All access to this account will be lost.** Any funds held in it will be
// inaccessible.
//
// Requires a `Signed` origin, and the sender account must have been created by a call to
// `pure` with corresponding parameters.
//
// - `Spawner`: The account that originally called `pure` to create this account.
// - `Index`: The disambiguation index originally passed to `pure`. Probably `0`.
// - `ProxyType`: The proxy type originally passed to `pure`.
// - `Height`: The height of the chain when the call to `pure` was processed.
// - `ExtIndex`: The extrinsic index in which the call to `pure` was processed.
//
// Fails with `NoPermission` in case the caller is not a previously created pure
// account whose `pure` call has corresponding parameters.
func (pt *ProxyTx) KillPure(spawner prim.MultiAddress, proxyType metadata.ProxyType, index uint16, height uint32, extIndex uint32) Transaction {
	call := pxPallet.CallKillPure{Spawner: spawner, ProxyType: proxyType, Index: index, Height: height, ExtIndex: extIndex}
	return NewTransaction(pt.client, pallets.ToPayload(call))
}

// Publish the hash of a proxy-call that will be made in the future.
//
// This must be called some number of blocks before the corresponding `proxy` is attempted
// if the delay associated with the proxy relationship is greater than zero.
//
// No more than `MaxPending` announcements may be made at any one time.
//
// This will take a deposit of `AnnouncementDepositFactor` as well as
// `AnnouncementDepositBase` if there are no other pending announcements.
//
// The dispatch origin for this call must be _Signed_ and a proxy of `real`.
//
// Parameters:
// - `Real`: The account that the proxy will make a call on behalf of.
// - `CallHash`: The hash of the call to be made by the `real` account.
func (pt *ProxyTx) Announce(real prim.MultiAddress, callHash prim.H256) Transaction {
	call := pxPallet.CallAnnounce{Real: real, CallHash: callHash}
	return NewTransaction(pt.client, pallets.ToPayload(call))
}

// Remove a given announcement.
//
// May be called by a proxy account to remove a call they previously announced and return
// the deposit.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Real`: The account that the proxy will make a call on behalf of.
// - `CallHash`: The hash of the call to be made by the `real` account.
func (pt *ProxyTx) RemoveAnnouncement(real prim.MultiAddress, callHash prim.H256) Transaction {
	call := pxPallet.CallRemoveAnnouncement{Real: real, CallHash: callHash}
	return NewTransaction(pt.client, pallets.ToPayload(call))
}

// Remove the given announcement of a delegate.
//
// May be called by a target (proxied) account to remove a call that one of their delegates
// (`delegate`) has announced they want to execute. The deposit is returned.
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Delegate`: The account that previously announced the call.
// - `CallHash`: The hash of the call to be made.
func (pt *ProxyTx) RejectAnnouncement(delegate prim.MultiAddress, callHash prim.H256) Transaction {
	call := pxPallet.CallRejectAnnouncement{Delegate: delegate, CallHash: callHash}
	return NewTransaction(pt.client, pallets.ToPayload(call))
}

// Dispatch the given `call` from an account that the sender is authorized for through
// `add_proxy`.
//
// Removes any corresponding announcement(s).
//
// The dispatch origin for this call must be _Signed_.
//
// Parameters:
// - `Real`: The account that the proxy will make a call on behalf of.
// - `ForceProxyType`: Specify the exact proxy type to be used and checked for this call.
// - `Call`: The call to be made by the `real` account.
func (pt *ProxyTx) ProxyAnnounced(delegate prim.MultiAddress, real prim.MultiAddress, forceProxyType prim.Option[metadata.ProxyType], call prim.Call) Transaction {
	c := pxPallet.CallProxyAnnounced{Delegate: delegate, Real: real, ForceProxyType: forceProxyType, Call: call}
	return NewTransaction(pt.client, pallets.ToPayload(c))
}
