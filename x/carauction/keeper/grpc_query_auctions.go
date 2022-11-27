package keeper

import (
	"context"

    "car-auction/x/carauction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Auctions(goCtx context.Context,  req *types.QueryAuctionsRequest) (*types.QueryAuctionsResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Process the query
    _ = ctx

	return &types.QueryAuctionsResponse{}, nil
}
