package proxy

// Understanding Proxy Account:
//
// From a simplified perspective, proxy account acts as a family member that we trust.
//
// If we designate the name "Child" to our main account, then te proxy account will
// be designated with the name "Parent". The "Parent" can do anything in the name of "Child"
// while the "Child" cannot do anything in the name of "Parent".
//
// This means that the "Parent", if he wants, can take all the funds from "Child" if he thinks
// that the "Child" has misbehaved, or the "Parent" can force the "Child" to work as a validator
// in order to gather some funds.
//
// The "Child" can, if needed, break this bond and force the "Parent" to become childless.
// The "Parent" can, if needed, break this bond and force the "Child" to become fatherless (an orphan).
//
// There are different types of Proxy account and in not of all has the "Parent" the same control options.
//
// Pure Proxy accounts work in the opposite direction. If we designate the name "Parent" to our main account,
// then the proxy account will be designated with the name "Child". The "Child" cannot do anything on it's own
// and the "Parent" has full control over it. The "Parent" can forget that he even had the "Child" in the first
// place thus forcing the "Child" to become an orphan that no will ever see again.

const PalletIndex = 40
const PalletName = "Proxy"
