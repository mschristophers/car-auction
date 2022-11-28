package keeper

import (
	"context"

	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuction(goCtx context.Context, msg *types.MsgCreateAuction) (*types.MsgCreateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	initialPriceCoins, err := sdk.ParseCoinsNormalized(msg.InitialPrice)
	if err != nil {
		return nil, err
	}

	zero := sdk.Coins{sdk.NewInt64Coin("token", 0)}

	if initialPriceCoins.IsAllLTE(zero) {
		return nil, types.AuctionPriceInvalid
	}

	if msg.Duration < 100 {
		return nil, types.AuctionDurationInvalid
	}

	auction := types.Auction{
		Creator:      msg.Creator,
		Name:         msg.Name,
		InitialPrice: msg.InitialPrice,
		Duration:     msg.Duration,
		Ended:        false,
		CreatedAt:    ctx.BlockHeight(),
	}

	id := k.AppendAuction(ctx, auction)

	return &types.MsgCreateAuctionResponse{Id: id}, nil
}
