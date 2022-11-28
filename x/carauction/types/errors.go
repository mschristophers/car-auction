package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/carauction module sentinel errors
var (
	AuctionEnded    = sdkerrors.Register(ModuleName, 1100, "The target auction has ended.")
	AuctionNotFound = sdkerrors.Register(ModuleName, 1200, "The target auction is not found.")

	BidNotFound            = sdkerrors.Register(ModuleName, 1300, "The target bid is not found.")
	BidPriceMustBePositive = sdkerrors.Register(ModuleName, 1400, "The bid price must be a positive number.")
	BidPriceTooLow         = sdkerrors.Register(ModuleName, 1500, "The bid price is lower than the initialprice.")
	BidPriceStepTooLow     = sdkerrors.Register(ModuleName, 1600, "The bid price increment is lower than the minimum increment.")

	NotAuctionOwner = sdkerrors.Register(ModuleName, 1700, "You are not the owner of this auction")
)
