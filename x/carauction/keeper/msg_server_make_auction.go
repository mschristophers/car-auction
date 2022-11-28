package keeper

import (
	"context"

	"car-auction/x/carauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MakeAuction(goCtx context.Context, msg *types.MsgMakeAuction) (*types.MsgMakeAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Duration < 100 {
		return nil, types.AuctionDurationInvalid
	}

	if msg.InitialPrice <= 0 {
		return nil, types.AuctionPriceMustBePositive
	}

	// Seller makes an auction
	auction := types.Auction{
		Creator:      msg.Creator,
		Name:         msg.Name,
		InitialPrice: msg.InitialPrice,
		Duration:     msg.Duration,
		CreatedAt:    ctx.BlockHeight(),
		Ended:        false,
	}

	id := k.AppendAuction(ctx, auction)

	return &types.MsgMakeAuctionResponse{Id: id}, nil
}
