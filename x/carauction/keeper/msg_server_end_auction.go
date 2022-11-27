package keeper

import (
	"context"

	"car-auction/x/carauction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EndAuction(goCtx context.Context, msg *types.MsgEndAuction) (*types.MsgEndAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEndAuctionResponse{}, nil
}
