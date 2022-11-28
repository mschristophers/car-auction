package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/carauction module sentinel errors
var (
	AuctionEnded               = sdkerrors.Register(ModuleName, 1100, "The target auction has ended.")
	AuctionNotFound            = sdkerrors.Register(ModuleName, 1200, "The target auction is not found.")
	AuctionDurationInvalid     = sdkerrors.Register(ModuleName, 1300, "The auction duration must be at least 100 blocks")
	AuctionDurationPassed      = sdkerrors.Register(ModuleName, 1400, "The auction duration has passed, so we won't accept any new bids.")
	AuctionPriceMustBePositive = sdkerrors.Register(ModuleName, 1500, "The bid price must be a positive number.")
	AuctionEndsTooEarly        = sdkerrors.Register(ModuleName, 1600, "Please wait until the auction duration passed to end the auction.")

	BidNotFound    = sdkerrors.Register(ModuleName, 1700, "The target bid is not found.")
	BidPriceTooLow = sdkerrors.Register(ModuleName, 1800, "The bid price is lower than the initial price.")

	NotAuctionOwner     = sdkerrors.Register(ModuleName, 1900, "You are not the owner of this auction")
	InsufficientBalance = sdkerrors.Register(ModuleName, 2000, "You have insufficient balance for the bid price")

	InternalError = sdkerrors.Register(ModuleName, 500, "internal error")
)
