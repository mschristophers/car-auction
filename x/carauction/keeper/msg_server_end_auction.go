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

	if uint64(auction.CreatedAt)+auction.Duration > uint64(ctx.BlockHeight()) {
		return nil, types.AuctionEndsTooEarly
	}

	hammerBidPrice := "0"
	bidCreator := ""
	if auction.HighestBidPresent {
		bid, found := k.GetBid(ctx, auction.CurrentHighestBidID)
		if found {
			hammerBidPrice = bid.BidPrice
			bidCreator = bid.Creator
		}
	}

	end := types.EndAuction{
		Creator:     msg.Creator,
		AuctionID:   msg.AuctionID,
		BidID:       auction.CurrentHighestBidID,
		HammerPrice: hammerBidPrice,
		Winner:      bidCreator,
	}

	if id, hammerPrice, err := k.AppendEndAuction(ctx, end); err != nil {
		return nil, err
	} else {
		if err = k.FinishAuction(ctx, msg.AuctionID); err != nil {
			return nil, err
		}

		hammerPriceCoin, err := sdk.ParseCoinsNormalized(end.HammerPrice)
		if err != nil {
			return nil, err
		}

		receiver, err := sdk.AccAddressFromBech32(msg.Creator)

		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, hammerPriceCoin); err != nil {
			return nil, err
		}

		return &types.MsgEndAuctionResponse{
			Id:          id,
			HammerPrice: hammerPrice,
		}, nil
	}
}
