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
	AuctionPriceMustBePositive = sdkerrors.Register(ModuleName, 1400, "The bid price must be a positive number.")

	BidNotFound    = sdkerrors.Register(ModuleName, 1500, "The target bid is not found.")
	BidPriceTooLow = sdkerrors.Register(ModuleName, 1600, "The bid price is lower than the initialprice.")

	NotAuctionOwner = sdkerrors.Register(ModuleName, 1700, "You are not the owner of this auction")
)
