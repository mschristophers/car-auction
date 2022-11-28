package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Bids(goCtx context.Context, req *types.QueryBidsRequest) (*types.QueryBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var bids []*types.Bid

	bidStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BidKey))

	pageRes, err := query.Paginate(
		bidStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var bid types.Bid
			if err := k.cdc.Unmarshal(value, &bid); err != nil {
				return err
			}

			bids = append(bids, &bid)

			return nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryBidsResponse{
		Bid:        bids,
		Pagination: pageRes,
	}, nil
}
