package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/auction module sentinel errors
var (
	AuctionEnded            = sdkerrors.Register(ModuleName, 1100, "The target auction has ended")
	AuctionNotFound         = sdkerrors.Register(ModuleName, 1200, "The target auction is not found")
	AuctionPriceInvalid     = sdkerrors.Register(ModuleName, 1300, "The starting price must larger than 0")
	AuctionDurationInvalid  = sdkerrors.Register(ModuleName, 1400, "The auction duration must be at least 100")
	AuctionDurationPassed   = sdkerrors.Register(ModuleName, 1500, "The auction duration passed, thus it is not accepting new bid")
	AuctionFinalizeTooEarly = sdkerrors.Register(ModuleName, 1600, "Please wait until auction duration passed to finalize the result")

	BidNotFound    = sdkerrors.Register(ModuleName, 1700, "The target bid is not found")
	BidPriceTooLow = sdkerrors.Register(ModuleName, 1800, "Please input a bid price higher than the auction starting price")

	NotAuctionOwner = sdkerrors.Register(ModuleName, 1900, "You are not the owner of this auction")

	InsufficientBalance = sdkerrors.Register(ModuleName, 2000, "Insufficient balance for the bid price")

	InternalError = sdkerrors.Register(ModuleName, 500, "Internal error")
)
