package keeper

import (
	"context"

	"car-auction/x/carauction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MakeAuction(goCtx context.Context, msg *types.MsgMakeAuction) (*types.MsgMakeAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMakeAuctionResponse{}, nil
}
