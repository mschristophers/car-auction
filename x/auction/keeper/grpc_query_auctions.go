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

func (k Keeper) Auctions(goCtx context.Context, req *types.QueryAuctionsRequest) (*types.QueryAuctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var auctions []*types.Auction

	ctx := sdk.UnwrapSDKContext(goCtx)

	auctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionKey))

	pageRes, err := query.Paginate(
		auctionStore,
		req.Pagination,
		func(key []byte, value []byte) error {
			var auction types.Auction
			if err := k.cdc.Unmarshal(value, &auction); err != nil {
				return err
			}

			auctions = append(auctions, &auction)

			return nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAuctionsResponse{
		Auction:    auctions,
		Pagination: pageRes,
	}, nil
}
