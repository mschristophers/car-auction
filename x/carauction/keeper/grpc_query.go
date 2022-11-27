package keeper

import (
	"car-auction/x/carauction/types"
)

var _ types.QueryServer = Keeper{}
