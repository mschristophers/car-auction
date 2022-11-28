package keeper

import (
	"context"

	"car-auction/x/carauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EndAuction(goCtx context.Context, msg *types.MsgEndAuction) (*types.MsgEndAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction, found := k.GetAuction(ctx, msg.AuctionID)
	if !found {
		return nil, types.AuctionNotFound
	}

	if auction.Ended {
		return nil, types.AuctionEnded
	}

	if auction.Creator != msg.Creator {
		return nil, types.NotAuctionOwner
	}

	bid, found := k.GetBid(ctx, msg.BidID)
	if !found {
		return nil, types.BidNotFound
	}

	end := types.EndAuction{
		Creator:    msg.Creator,
		AuctionID:  msg.AuctionID,
		BidID:      msg.BidID,
		FinalPrice: bid.BidPrice,
	}

	if id, finalPrice, err := k.AppendEndAuction(ctx, end); err != nil {
		return nil, err
	} else {
		if err = k.FinishAuction(ctx, msg.AuctionID); err != nil {
			return nil, err
		}

		return &types.MsgEndAuctionResponse{
			Id:         id,
			FinalPrice: finalPrice,
		}, nil
	}
}
