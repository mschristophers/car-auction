package keeper

import (
	"context"

	"car-auction/x/carauction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MakeAuction(goCtx context.Context, msg *types.MsgMakeAuction) (*types.MsgMakeAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Seller makes an auction
	auction := types.Auction{
		Creator:      msg.Creator,
		InitialPrice: msg.InitialPrice,
		MinIncrment:  msg.MinIncrment,
	}

	id := k.AppendAuction(ctx, Auction)

	return &types.MsgMakeAuctionResponse{}, nil
}
