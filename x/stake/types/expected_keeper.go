package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	curatingtypes "github.com/public-awesome/stakebird/x/curating/types"
)

/*
When a module wishes to interact with another module, it is good practice to define what it will use
as an interface so the module cannot use things that are not permitted.
*/

// CurationKeeper defines the expected interface for the curation module
type CurationKeeper interface {
	GetPostZ(ctx sdk.Context, vendorID uint32, postIDBz []byte) (post curatingtypes.Post, found bool, err error)
}

// StakingKeeper expected staking keeper
type StakingKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	Delegate(ctx sdk.Context, delAddr sdk.AccAddress, bondAmt sdk.Int, tokenSrc stakingtypes.BondStatus,
		validator stakingtypes.Validator, subtractAccount bool) (newShares sdk.Dec, err error)
	Undelegate(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, sharesAmount sdk.Dec) (time.Time, error)
}