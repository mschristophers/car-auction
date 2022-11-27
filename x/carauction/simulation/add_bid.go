package simulation

import (
	"math/rand"

	"car-auction/x/carauction/keeper"
	"car-auction/x/carauction/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddBid(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddBid{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddBid simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddBid simulation not implemented"), nil, nil
	}
}
