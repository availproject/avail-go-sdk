package nomination_pools

import (
	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"

	"github.com/itering/scale.go/utiles/uint128"
)

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
type CallJoin struct {
	Amount metadata.Balance `scale:"compact"`
	PoolId uint32
}

func (this CallJoin) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallJoin) PalletName() string {
	return PalletName
}

func (this CallJoin) CallIndex() uint8 {
	return 0
}

func (this CallJoin) CallName() string {
	return "join"
}

// Bond `extra` more funds from `origin` into the pool to which they already belong.
//
// Additional funds can come from either the free balance of the account, of from the
// accumulated rewards, see [`BondExtra`].
//
// Bonding extra funds implies an automatic payout of all pending rewards as well.
// See `bond_extra_other` to bond pending rewards of `other` members.
type CallBondExtra struct {
	Extra metadata.PoolBondExtra
}

func (this CallBondExtra) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallBondExtra) PalletName() string {
	return PalletName
}

func (this CallBondExtra) CallIndex() uint8 {
	return 1
}

func (this CallBondExtra) CallName() string {
	return "bond_extra"
}

// A bonded member can use this to claim their payout based on the rewards that the pool
// has accumulated since their last claimed payout (OR since joining if this is their first
// time claiming rewards). The payout will be transferred to the member's account.
//
// The member will earn rewards pro rata based on the members stake vs the sum of the
// members in the pools stake. Rewards do not "expire".
//
// See `claim_payout_other` to caim rewards on bahalf of some `other` pool member.
type CallClaimPayout struct{}

func (this CallClaimPayout) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallClaimPayout) PalletName() string {
	return PalletName
}

func (this CallClaimPayout) CallIndex() uint8 {
	return 2
}

func (this CallClaimPayout) CallName() string {
	return "claim_payout"
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
type CallUnbond struct {
	MemberAccount   prim.MultiAddress
	UnbondingPoints uint128.Uint128 `scale:"compact"`
}

func (this CallUnbond) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallUnbond) PalletName() string {
	return PalletName
}

func (this CallUnbond) CallIndex() uint8 {
	return 3
}

func (this CallUnbond) CallName() string {
	return "unbond"
}

// Call `withdraw_unbonded` for the pools account. This call can be made by any account.
//
// This is useful if there are too many unlocking chunks to call `unbond`, and some
// can be cleared by withdrawing. In the case there are too many unlocking chunks, the user
// would probably see an error like `NoMoreChunks` emitted from the staking system when
// they attempt to unbond.
type CallPoolWithdrawUnbonded struct {
	PoolId           uint32
	NumSlashingSpans uint32
}

func (this CallPoolWithdrawUnbonded) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallPoolWithdrawUnbonded) PalletName() string {
	return PalletName
}

func (this CallPoolWithdrawUnbonded) CallIndex() uint8 {
	return 4
}

func (this CallPoolWithdrawUnbonded) CallName() string {
	return "pool_withdraw_unbonded"
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
type CallWithdrawUnbonded struct {
	MemberAccount    prim.MultiAddress
	NumSlashingSpans uint32
}

func (this CallWithdrawUnbonded) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallWithdrawUnbonded) PalletName() string {
	return PalletName
}

func (this CallWithdrawUnbonded) CallIndex() uint8 {
	return 5
}

func (this CallWithdrawUnbonded) CallName() string {
	return "withdraw_unbonded"
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
type CallCreate struct {
	Amount    metadata.Balance `scale:"compact"`
	Root      prim.MultiAddress
	Nominator prim.MultiAddress
	Bouncer   prim.MultiAddress
}

func (this CallCreate) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallCreate) PalletName() string {
	return PalletName
}

func (this CallCreate) CallIndex() uint8 {
	return 6
}

func (this CallCreate) CallName() string {
	return "create"
}

// Create a new delegation pool with a previously used pool id
//
// # Arguments
//
// same as `create` with the inclusion of
// * `pool_id` - `A valid PoolId.
type CallCreateWithPoolId struct {
	Amount    metadata.Balance `scale:"compact"`
	Root      prim.MultiAddress
	Nominator prim.MultiAddress
	Bouncer   prim.MultiAddress
	PoolId    uint32
}

func (this CallCreateWithPoolId) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallCreateWithPoolId) PalletName() string {
	return PalletName
}

func (this CallCreateWithPoolId) CallIndex() uint8 {
	return 7
}

func (this CallCreateWithPoolId) CallName() string {
	return "create_with_pool_id"
}

// Nominate on behalf of the pool.
//
// The dispatch origin of this call must be signed by the pool nominator or the pool
// root role.
//
// This directly forward the call to the staking pallet, on behalf of the pool bonded
// account.
type CallNominate struct {
	PoolId     uint32
	Validators []metadata.AccountId
}

func (this CallNominate) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallNominate) PalletName() string {
	return PalletName
}

func (this CallNominate) CallIndex() uint8 {
	return 8
}

func (this CallNominate) CallName() string {
	return "nominate"
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
type CallSetState struct {
	PoolId uint32
	State  metadata.PoolState
}

func (this CallSetState) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetState) PalletName() string {
	return PalletName
}

func (this CallSetState) CallIndex() uint8 {
	return 9
}

func (this CallSetState) CallName() string {
	return "set_state"
}

// Set a new metadata for the pool.
//
// The dispatch origin of this call must be signed by the bouncer, or the root role of the
// pool.
type CallSetMetadata struct {
	PoolId   uint32
	Metadata []byte
}

func (this CallSetMetadata) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetMetadata) PalletName() string {
	return PalletName
}

func (this CallSetMetadata) CallIndex() uint8 {
	return 10
}

func (this CallSetMetadata) CallName() string {
	return "set_metadata"
}

// Update the roles of the pool.
//
// The root is the only entity that can change any of the roles, including itself,
// excluding the depositor, who can never change.
//
// It emits an event, notifying UIs of the role change. This event is quite relevant to
// most pool members and they should be informed of changes to pool roles.
type CallUpdateRoles struct {
	PoolId       uint32
	NewRoot      metadata.PoolRoleConfig
	NewNominator metadata.PoolRoleConfig
	NewBouncer   metadata.PoolRoleConfig
}

func (this CallUpdateRoles) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallUpdateRoles) PalletName() string {
	return PalletName
}

func (this CallUpdateRoles) CallIndex() uint8 {
	return 12
}

func (this CallUpdateRoles) CallName() string {
	return "update_roles"
}

// Chill on behalf of the pool.
//
// The dispatch origin of this call must be signed by the pool nominator or the pool
// root role, same as [`Pallet::nominate`].
//
// This directly forward the call to the staking pallet, on behalf of the pool bonded
// account.
type CallChill struct {
	PoolId uint32
}

func (this CallChill) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallChill) PalletName() string {
	return PalletName
}

func (this CallChill) CallIndex() uint8 {
	return 13
}

func (this CallChill) CallName() string {
	return "chill"
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
type CallBondExtraOther struct {
	Member prim.MultiAddress
	Extra  metadata.PoolBondExtra
}

func (this CallBondExtraOther) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallBondExtraOther) PalletName() string {
	return PalletName
}

func (this CallBondExtraOther) CallIndex() uint8 {
	return 14
}

func (this CallBondExtraOther) CallName() string {
	return "bond_extra_other"
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
type CallSetClaimPermission struct {
	Permission metadata.PoolClaimPermission
}

func (this CallSetClaimPermission) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetClaimPermission) PalletName() string {
	return PalletName
}

func (this CallSetClaimPermission) CallIndex() uint8 {
	return 15
}

func (this CallSetClaimPermission) CallName() string {
	return "set_claim_permission"
}

// `origin` can claim payouts on some pool member `other`'s behalf.
//
// Pool member `other` must have a `PermissionlessAll` or `PermissionlessWithdraw` in order
// for this call to be successful.
type CallClaimPayoutOther struct {
	Other metadata.AccountId
}

func (this CallClaimPayoutOther) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallClaimPayoutOther) PalletName() string {
	return PalletName
}

func (this CallClaimPayoutOther) CallIndex() uint8 {
	return 16
}

func (this CallClaimPayoutOther) CallName() string {
	return "claim_payout_other"
}

// Set the commission of a pool.
//
// Both a commission percentage and a commission payee must be provided in the `current`
// tuple. Where a `current` of `None` is provided, any current commission will be removed.
//
// - If a `None` is supplied to `new_commission`, existing commission will be removed.
type CallSetCommission struct {
	PoolId        uint32
	NewCommission prim.Option[metadata.Tuple2[metadata.Perbill, metadata.AccountId]]
}

func (this CallSetCommission) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetCommission) PalletName() string {
	return PalletName
}

func (this CallSetCommission) CallIndex() uint8 {
	return 17
}

func (this CallSetCommission) CallName() string {
	return "set_commission"
}

// Set the maximum commission of a pool.
//
//   - Initial max can be set to any `Perbill`, and only smaller values thereafter.
//   - Current commission will be lowered in the event it is higher than a new max
//     commission.
type CallSetCommissionMax struct {
	PoolId        uint32
	MaxCommission metadata.Perbill
}

func (this CallSetCommissionMax) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetCommissionMax) PalletName() string {
	return PalletName
}

func (this CallSetCommissionMax) CallIndex() uint8 {
	return 18
}

func (this CallSetCommissionMax) CallName() string {
	return "set_commission_max"
}

// Set the commission change rate for a pool.
//
// Initial change rate is not bounded, whereas subsequent updates can only be more
// restrictive than the current.
type CallSetCommissionChangeRate struct {
	PoolId     uint32
	ChangeRate metadata.PoolCommissionChangeRate
}

func (this CallSetCommissionChangeRate) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetCommissionChangeRate) PalletName() string {
	return PalletName
}

func (this CallSetCommissionChangeRate) CallIndex() uint8 {
	return 19
}

func (this CallSetCommissionChangeRate) CallName() string {
	return "set_commission_change_rate"
}

// Claim pending commission.
//
// The dispatch origin of this call must be signed by the `root` role of the pool. Pending
// commission is paid out and added to total claimed commission`. Total pending commission
// is reset to zero. the current.
type CallClaimCommission struct {
	PoolId uint32
}

func (this CallClaimCommission) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallClaimCommission) PalletName() string {
	return PalletName
}

func (this CallClaimCommission) CallIndex() uint8 {
	return 20
}

func (this CallClaimCommission) CallName() string {
	return "claim_commission"
}

// Top up the deficit or withdraw the excess ED from the pool.
//
// When a pool is created, the pool depositor transfers ED to the reward account of the
// pool. ED is subject to change and over time, the deposit in the reward account may be
// insufficient to cover the ED deficit of the pool or vice-versa where there is excess
// deposit to the pool. This call allows anyone to adjust the ED deposit of the
// pool by either topping up the deficit or claiming the excess.
type CallAdjustPoolDeposit struct {
	PoolId uint32
}

func (this CallAdjustPoolDeposit) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallAdjustPoolDeposit) PalletName() string {
	return PalletName
}

func (this CallAdjustPoolDeposit) CallIndex() uint8 {
	return 21
}

func (this CallAdjustPoolDeposit) CallName() string {
	return "adjust_pool_deposit"
}

// Set or remove a pool's commission claim permission.
//
// Determines who can claim the pool's pending commission. Only the `Root` role of the pool
// is able to conifigure commission claim permissions.
type CallSetCommissionClaimPermission struct {
	PoolId     uint32
	Permission prim.Option[metadata.CommissionClaimPermission]
}

func (this CallSetCommissionClaimPermission) PalletIndex() uint8 {
	return PalletIndex
}

func (this CallSetCommissionClaimPermission) PalletName() string {
	return PalletName
}

func (this CallSetCommissionClaimPermission) CallIndex() uint8 {
	return 22
}

func (this CallSetCommissionClaimPermission) CallName() string {
	return "set_commission_claim_permission"
}
