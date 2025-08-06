package staking

import (
	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"

	"github.com/itering/scale.go/utiles/uint128"
)

// Take the origin account as a stash and lock up `value` of its balance. `controller` will
// be the account that controls it.
type CallBond struct {
	Value metadata.Balance `scale:"compact"`
	Payee metadata.RewardDestination
}

func (cb CallBond) PalletIndex() uint8 {
	return PalletIndex
}

func (cb CallBond) PalletName() string {
	return PalletName
}

func (cb CallBond) CallIndex() uint8 {
	return 0
}

func (cb CallBond) CallName() string {
	return "bond"
}

// Add some extra amount that have appeared in the stash `free_balance` into the balance up
// for staking.
type CallBondExtra struct {
	MaxAdditional metadata.Balance `scale:"compact"`
}

func (cbe CallBondExtra) PalletIndex() uint8 {
	return PalletIndex
}

func (cbe CallBondExtra) PalletName() string {
	return PalletName
}

func (cbe CallBondExtra) CallIndex() uint8 {
	return 1
}

func (cbe CallBondExtra) CallName() string {
	return "bond_extra"
}

// Schedule a portion of the stash to be unlocked ready for transfer out after the bond
// period ends. If this leaves an amount actively bonded less than
// T::Currency::minimum_balance(), then it is increased to the full amount.
type CallUnbond struct {
	Value metadata.Balance `scale:"compact"`
}

func (cu CallUnbond) PalletIndex() uint8 {
	return PalletIndex
}

func (cu CallUnbond) PalletName() string {
	return PalletName
}

func (cu CallUnbond) CallIndex() uint8 {
	return 2
}

func (cu CallUnbond) CallName() string {
	return "unbond"
}

// Remove any unlocked chunks from the `unlocking` queue from our management.
//
// This essentially frees up that balance to be used by the stash account to do whatever
// it wants.
type CallWithdrawUnbonded struct {
	NumSlashingSpans uint32
}

func (cwu CallWithdrawUnbonded) PalletIndex() uint8 {
	return PalletIndex
}

func (cwu CallWithdrawUnbonded) PalletName() string {
	return PalletName
}

func (cwu CallWithdrawUnbonded) CallIndex() uint8 {
	return 3
}

func (cwu CallWithdrawUnbonded) CallName() string {
	return "withdraw_unbonded"
}

// Declare the desire to validate for the origin controller.
//
// Effects will be felt at the beginning of the next era.
type CallValidate struct {
	Prefs metadata.ValidatorPrefs
}

func (cv CallValidate) PalletIndex() uint8 {
	return PalletIndex
}

func (cv CallValidate) PalletName() string {
	return PalletName
}

func (cv CallValidate) CallIndex() uint8 {
	return 4
}

func (cv CallValidate) CallName() string {
	return "validate"
}

// Declare the desire to nominate `targets` for the origin controller.
//
// Effects will be felt at the beginning of the next era.
type CallNominate struct {
	Targets []prim.MultiAddress
}

func (cn CallNominate) PalletIndex() uint8 {
	return PalletIndex
}

func (cn CallNominate) PalletName() string {
	return PalletName
}

func (cn CallNominate) CallIndex() uint8 {
	return 5
}

func (cn CallNominate) CallName() string {
	return "nominate"
}

// Declare no desire to either validate or nominate.
//
// Effects will be felt at the beginning of the next era.
type CallChill struct{}

func (cc CallChill) PalletIndex() uint8 {
	return PalletIndex
}

func (cc CallChill) PalletName() string {
	return PalletName
}

func (cc CallChill) CallIndex() uint8 {
	return 6
}

func (cc CallChill) CallName() string {
	return "chill"
}

// (Re-)set the payment target for a controller.
//
// Effects will be felt instantly (as soon as this function is completed successfully).
type CallSetPayee struct {
	Payee metadata.RewardDestination
}

func (csp CallSetPayee) PalletIndex() uint8 {
	return PalletIndex
}

func (csp CallSetPayee) PalletName() string {
	return PalletName
}

func (csp CallSetPayee) CallIndex() uint8 {
	return 7
}

func (csp CallSetPayee) CallName() string {
	return "set_payee"
}

// (Re-)sets the controller of a stash to the stash itself. This function previously
// accepted a `controller` argument to set the controller to an account other than the
// stash itself. This functionality has now been removed, now only setting the controller
// to the stash, if it is not already.
//
// Effects will be felt instantly (as soon as this function is completed successfully).
type CallSetController struct{}

func (csc CallSetController) PalletIndex() uint8 {
	return PalletIndex
}

func (csc CallSetController) PalletName() string {
	return PalletName
}

func (csc CallSetController) CallIndex() uint8 {
	return 8
}

func (csc CallSetController) CallName() string {
	return "set_controller"
}

// Pay out next page of the stakers behind a validator for the given era.
//
// - `validator_stash` is the stash account of the validator.
// - `era` may be any era between `[current_era - history_depth; current_era]`.
type CallPayoutStakers struct {
	ValidatorStash prim.AccountId
	Era            uint32
}

func (cps CallPayoutStakers) PalletIndex() uint8 {
	return PalletIndex
}

func (cps CallPayoutStakers) PalletName() string {
	return PalletName
}

func (cps CallPayoutStakers) CallIndex() uint8 {
	return 18
}

func (cps CallPayoutStakers) CallName() string {
	return "payout_stakers"
}

// Rebond a portion of the stash scheduled to be unlocked.
type CallRebond struct {
	Value uint128.Uint128 `scale:"compact"`
}

func (cr CallRebond) PalletIndex() uint8 {
	return PalletIndex
}

func (cr CallRebond) PalletName() string {
	return PalletName
}

func (cr CallRebond) CallIndex() uint8 {
	return 19
}

func (cr CallRebond) CallName() string {
	return "rebond"
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
type CallReapStash struct {
	Stash            prim.AccountId
	NumSlashingSpans uint32
}

func (crs CallReapStash) PalletIndex() uint8 {
	return PalletIndex
}

func (crs CallReapStash) PalletName() string {
	return PalletName
}

func (crs CallReapStash) CallIndex() uint8 {
	return 20
}

func (crs CallReapStash) CallName() string {
	return "reap_stash"
}

// Remove the given nominations from the calling validator.
//
// Effects will be felt at the beginning of the next era.
type CallKick struct {
	Who []prim.MultiAddress
}

func (ck CallKick) PalletIndex() uint8 {
	return PalletIndex
}

func (ck CallKick) PalletName() string {
	return PalletName
}

func (ck CallKick) CallIndex() uint8 {
	return 21
}

func (ck CallKick) CallName() string {
	return "kick"
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
type CallChillOther struct {
	Stash prim.AccountId
}

func (cco CallChillOther) PalletIndex() uint8 {
	return PalletIndex
}

func (cco CallChillOther) PalletName() string {
	return PalletName
}

func (cco CallChillOther) CallIndex() uint8 {
	return 23
}

func (cco CallChillOther) CallName() string {
	return "chill_other"
}

// Force a validator to have at least the minimum commission. This will not affect a
// validator who already has a commission greater than or equal to the minimum. Any account
// can call this.
type CallForceApplyMinCommission struct {
	ValidatorStash prim.AccountId
}

func (cfamc CallForceApplyMinCommission) PalletIndex() uint8 {
	return PalletIndex
}

func (cfamc CallForceApplyMinCommission) PalletName() string {
	return PalletName
}

func (cfamc CallForceApplyMinCommission) CallIndex() uint8 {
	return 24
}

func (cfamc CallForceApplyMinCommission) CallName() string {
	return "force_apply_min_commission"
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
type CallPayoutStakersByPage struct {
	ValidatorStash prim.AccountId
	Era            uint32
	Page           uint32
}

func (cpsbp CallPayoutStakersByPage) PalletIndex() uint8 {
	return PalletIndex
}

func (cpsbp CallPayoutStakersByPage) PalletName() string {
	return PalletName
}

func (cpsbp CallPayoutStakersByPage) CallIndex() uint8 {
	return 26
}

func (cpsbp CallPayoutStakersByPage) CallName() string {
	return "payout_stakers_by_page"
}

// Migrates an account's `RewardDestination::Controller` to
// `RewardDestination::Account(controller)`.
//
// Effects will be felt instantly (as soon as this function is completed successfully).
//
// This will waive the transaction fee if the `payee` is successfully migrated.
type CallUpdatePayee struct {
	Controller prim.AccountId
}

func (cup CallUpdatePayee) PalletIndex() uint8 {
	return PalletIndex
}

func (cup CallUpdatePayee) PalletName() string {
	return PalletName
}

func (cup CallUpdatePayee) CallIndex() uint8 {
	return 27
}

func (cup CallUpdatePayee) CallName() string {
	return "update_payee"
}
