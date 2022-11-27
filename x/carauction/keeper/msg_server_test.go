package keeper_test

import (
	"context"
	"testing"

	keepertest "car-auction/testutil/keeper"
	"car-auction/x/carauction/keeper"
	"car-auction/x/carauction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CarauctionKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
