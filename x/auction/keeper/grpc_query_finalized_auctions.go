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

func (k Keeper) FinalizedAuctions(goCtx context.Context, req *types.QueryFinalizedAuctionsRequest) (*types.QueryFinalizedAuctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var finalized []*types.FinalizeAuction

	finalizedAuctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.FinalizeAuctionKey))

	pageRes, err := query.Paginate(
		finalizedAuctionStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var finalizedAuc types.FinalizeAuction
			if err := k.cdc.Unmarshal(value, &finalizedAuc); err != nil {
				return err
			}

			finalized = append(finalized, &finalizedAuc)

			return nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryFinalizedAuctionsResponse{
		FinalizedAuction: finalized,
		Pagination:       pageRes,
	}, nil
}
