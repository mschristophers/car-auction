package keeper

import (
	"context"

	"car-auction/x/carauction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBid(goCtx context.Context, msg *types.MsgAddBid) (*types.MsgAddBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddBidResponse{}, nil
}
