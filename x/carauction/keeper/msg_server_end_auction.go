package keeper

import (
	"context"

	"car-auction/x/carauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EndAuction(goCtx context.Context, msg *types.MsgEndAuction) (*types.MsgEndAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	end := types.EndAuction{
		Creator:   msg.Creator,
		AuctionID: msg.AuctionID,
		BidID:     msg.BidID,
	}

	if id, finalPrice, err := k.AppendEndAuction(ctx, end); err != nil {
		return nil, err
	} else {
		return &types.MsgEndAuctionResponse{
			Id:         id,
			FinalPrice: finalPrice,
		}, nil
	}
}
