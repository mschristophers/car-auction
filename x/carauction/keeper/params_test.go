package keeper_test

import (
	"testing"

	testkeeper "car-auction/testutil/keeper"
	"car-auction/x/carauction/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CarauctionKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
