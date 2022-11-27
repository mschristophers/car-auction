package carauction_test

import (
	"testing"

	keepertest "car-auction/testutil/keeper"
	"car-auction/testutil/nullify"
	"car-auction/x/carauction"
	"car-auction/x/carauction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CarauctionKeeper(t)
	carauction.InitGenesis(ctx, *k, genesisState)
	got := carauction.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
