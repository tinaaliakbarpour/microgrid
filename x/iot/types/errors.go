package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/iot module sentinel errors
var (
	ErrSample         = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrorDuplicate    = sdkerrors.Register(ModuleName, 1101, "duplicate key")
	ErrorInvalidInput = sdkerrors.Register(ModuleName, 1102, "invalid input")
)
