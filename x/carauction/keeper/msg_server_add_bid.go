package keeper

import (
	"context"

	"car-auction/x/carauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBid(goCtx context.Context, msg *types.MsgAddBid) (*types.MsgAddBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction, found := k.GetAuction(ctx, msg.AuctionID)
	if !found {
		return nil, types.AuctionNotFound
	}

	if auction.Ended {
		return nil, types.AuctionEnded
	}

	if msg.BidPrice < 0 {
		return nil, types.AuctionPriceMustBePositive
	}

	if uint64(auction.CreatedAt)+auction.Duration <= uint64(ctx.BlockHeight()) {
		if err := k.FinishAuction(ctx, auction.Id); err != nil {
			return nil, err
		}
		return nil, types.AuctionEnded
	}

	if auction.InitialPrice > msg.BidPrice {
		return nil, types.BidPriceTooLow
	}

	// Buyer makes a bid
	bid := types.Bid{
		Creator:   msg.Creator,
		AuctionID: msg.AuctionID,
		BidPrice:  msg.BidPrice,
	}

	if id, err := k.AppendBid(ctx, bid); err != nil {
		return nil, err
	} else {
		return &types.MsgAddBidResponse{Id: id}, nil
	}
}
