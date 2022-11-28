package auction

import (
	"math/rand"

	"auction/testutil/sample"
	auctionsimulation "auction/x/auction/simulation"
	"auction/x/auction/types"
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
	_ = auctionsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAuction = "op_weight_msg_create_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAuction int = 100

	opWeightMsgPlaceBid = "op_weight_msg_place_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPlaceBid int = 100

	opWeightMsgFinalizeAuction = "op_weight_msg_finalize_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFinalizeAuction int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	auctionGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&auctionGenesis)
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

	var weightMsgCreateAuction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAuction, &weightMsgCreateAuction, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAuction = defaultWeightMsgCreateAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAuction,
		auctionsimulation.SimulateMsgCreateAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPlaceBid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPlaceBid, &weightMsgPlaceBid, nil,
		func(_ *rand.Rand) {
			weightMsgPlaceBid = defaultWeightMsgPlaceBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPlaceBid,
		auctionsimulation.SimulateMsgPlaceBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFinalizeAuction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFinalizeAuction, &weightMsgFinalizeAuction, nil,
		func(_ *rand.Rand) {
			weightMsgFinalizeAuction = defaultWeightMsgFinalizeAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFinalizeAuction,
		auctionsimulation.SimulateMsgFinalizeAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
