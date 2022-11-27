package carauction

import (
	"math/rand"

	"car-auction/testutil/sample"
	carauctionsimulation "car-auction/x/carauction/simulation"
	"car-auction/x/carauction/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = carauctionsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgMakeAuction = "op_weight_msg_make_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMakeAuction int = 100

	opWeightMsgAddBid = "op_weight_msg_add_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddBid int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	carauctionGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&carauctionGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgMakeAuction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMakeAuction, &weightMsgMakeAuction, nil,
		func(_ *rand.Rand) {
			weightMsgMakeAuction = defaultWeightMsgMakeAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMakeAuction,
		carauctionsimulation.SimulateMsgMakeAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddBid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddBid, &weightMsgAddBid, nil,
		func(_ *rand.Rand) {
			weightMsgAddBid = defaultWeightMsgAddBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddBid,
		carauctionsimulation.SimulateMsgAddBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
