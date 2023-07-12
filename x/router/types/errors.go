package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO cleanup pass
// x/router module sentinel errors
var (
	ErrUnauthorized            = sdkerrors.Register(ModuleName, 2, "unauthorized")
	ErrUserNotFound            = sdkerrors.Register(ModuleName, 3, "user not found")
	ErrMint                    = sdkerrors.Register(ModuleName, 4, "tokens can not be minted")
	ErrDenomAlreadySet         = sdkerrors.Register(ModuleName, 5, "denom already set")
	ErrAuthorityNotSet         = sdkerrors.Register(ModuleName, 15, "authority not set")
	ErrMalformedField          = sdkerrors.Register(ModuleName, 16, "field cannot be empty or nil")
	ErrIbcForwardAlreadyQueued = sdkerrors.Register(ModuleName, 17, "the ibc forward cannot be processed because there is an existing one queued")
	ErrNoOp                    = sdkerrors.Register(ModuleName, 18, "unable to parse message into a supported message type")
)
