package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	"car-auction/x/carauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EndAuctions(goCtx context.Context, req *types.QueryEndAuctionsRequest) (*types.QueryEndAuctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var end []*types.EndAuction

	endAuctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.EndAuctionKey))

	pageRes, err := query.Paginate(
		endAuctionStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var endAuc types.EndAuction
			if err := k.cdc.Unmarshal(value, &endAuc); err != nil {
				return err
			}

			end = append(end, &endAuc)

			return nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryEndAuctionsResponse{
		EndAuction: end,
		Pagination: pageRes,
	}, nil
}
