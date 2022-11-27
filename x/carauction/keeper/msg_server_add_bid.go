package keeper

import (
	"context"

	"car-auction/x/carauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBid(goCtx context.Context, msg *types.MsgAddBid) (*types.MsgAddBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Buyer makes a bid
	bid := types.Bid{
		Creator:   msg.Creator,
		AuctionID: msg.AuctionID,
		BidPrice:  msg.BidPrice,
	}

	id := k.AppendBid(ctx, bid)

	return &types.MsgAddBidResponse{Id: id}, nil
}
