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
		Name:         msg.Name,
		InitialPrice: msg.InitialPrice,
		MinIncrement: msg.MinIncrement,
	}

	id := k.AppendAuction(ctx, auction)

	return &types.MsgMakeAuctionResponse{Id: id}, nil
}
