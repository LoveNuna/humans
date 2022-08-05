package human

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"human/testutil/sample"
	humansimulation "human/x/human/simulation"
	"human/x/human/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = humansimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRequestTransaction = "op_weight_msg_request_transaction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestTransaction int = 100

	opWeightMsgObservationVote = "op_weight_msg_observation_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgObservationVote int = 100

	opWeightMsgUpdateBalance = "op_weight_msg_update_balance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBalance int = 100

	opWeightMsgKeysignVote = "op_weight_msg_keysign_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgKeysignVote int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	humanGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&humanGenesis)
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

	var weightMsgRequestTransaction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestTransaction, &weightMsgRequestTransaction, nil,
		func(_ *rand.Rand) {
			weightMsgRequestTransaction = defaultWeightMsgRequestTransaction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestTransaction,
		humansimulation.SimulateMsgRequestTransaction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgObservationVote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgObservationVote, &weightMsgObservationVote, nil,
		func(_ *rand.Rand) {
			weightMsgObservationVote = defaultWeightMsgObservationVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgObservationVote,
		humansimulation.SimulateMsgObservationVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBalance int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBalance, &weightMsgUpdateBalance, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBalance = defaultWeightMsgUpdateBalance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBalance,
		humansimulation.SimulateMsgUpdateBalance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgKeysignVote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgKeysignVote, &weightMsgKeysignVote, nil,
		func(_ *rand.Rand) {
			weightMsgKeysignVote = defaultWeightMsgKeysignVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgKeysignVote,
		humansimulation.SimulateMsgKeysignVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
