package types

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto/tmhash"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Oracle Errors
var (
	ErrInvalidExchangeRate   = sdkerrors.Register(ModuleName, 1, "invalid exchange rate")
	ErrNoPrevote             = sdkerrors.Register(ModuleName, 2, "no prevote")
	ErrNoVote                = sdkerrors.Register(ModuleName, 3, "no vote")
	ErrNoVotingPermission    = sdkerrors.Register(ModuleName, 4, "unauthorized voter")
	ErrInvalidHash           = sdkerrors.Register(ModuleName, 5, "invalid hash")
	ErrInvalidHashLength     = sdkerrors.Register(ModuleName, 6, fmt.Sprintf("invalid hash length; should equal %d", tmhash.TruncatedSize))
	ErrVerificationFailed    = sdkerrors.Register(ModuleName, 7, "hash verification failed")
	ErrRevealPeriodMissMatch = sdkerrors.Register(ModuleName, 8, "reveal period of submitted vote do not match with registered prevote")
	ErrInvalidSaltLength     = sdkerrors.Register(ModuleName, 9, "invalid salt length; must be 64")
	ErrInvalidSaltFormat     = sdkerrors.Register(ModuleName, 10, "invalid salt format")
	ErrNoAggregatePrevote    = sdkerrors.Register(ModuleName, 11, "no aggregate prevote")
	ErrNoAggregateVote       = sdkerrors.Register(ModuleName, 12, "no aggregate vote")
	ErrUnknownDenom          = sdkerrors.Register(ModuleName, 13, "unknown denom")
	ErrBallotNotSorted       = sdkerrors.Register(ModuleName, 14, "ballot not sorted")
)
