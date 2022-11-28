package keeper

import (
	"context"

	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FinalizeAuction(goCtx context.Context, msg *types.MsgFinalizeAuction) (*types.MsgFinalizeAuctionResponse, error) {
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
		return nil, types.AuctionFinalizeTooEarly
	}

	finalBidPrice := "0"
	bidCreator := ""
	if auction.HighestBidIsPresent {
		bid, found := k.GetBid(ctx, auction.CurrentHighestBidID)
		if found {
			finalBidPrice = bid.BidPrice
			bidCreator = bid.Creator
		}
	}

	finalize := types.FinalizeAuction{
		Creator:    msg.Creator,
		AuctionID:  msg.AuctionID,
		BidID:      auction.CurrentHighestBidID,
		FinalPrice: finalBidPrice,
		Winner:     bidCreator,
	}

	if id, finalPrice, err := k.AppendFinalizeAuction(ctx, finalize); err != nil {
		return nil, err
	} else {
		if err = k.EndAuction(ctx, msg.AuctionID); err != nil {
			return nil, err
		}

		finalPriceCoin, err := sdk.ParseCoinsNormalized(finalize.FinalPrice)
		if err != nil {
			return nil, err
		}

		receiver, err := sdk.AccAddressFromBech32(msg.Creator)

		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, finalPriceCoin); err != nil {
			return nil, err
		}

		return &types.MsgFinalizeAuctionResponse{
			Id:         id,
			FinalPrice: finalPrice,
		}, nil
	}
}
