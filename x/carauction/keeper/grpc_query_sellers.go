package keeper

import (
	"context"

    "car-auction/x/carauction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Sellers(goCtx context.Context,  req *types.QuerySellersRequest) (*types.QuerySellersResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Process the query
    _ = ctx

	return &types.QuerySellersResponse{}, nil
}
