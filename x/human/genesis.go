package human

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"human/x/human/keeper"
	"human/x/human/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the feeBalance
	for _, elem := range genState.FeeBalanceList {
		k.SetFeeBalance(ctx, elem)
	}
	// Set all the keysignVoteData
	for _, elem := range genState.KeysignVoteDataList {
		k.SetKeysignVoteData(ctx, elem)
	}
	// Set all the observeVote
	for _, elem := range genState.ObserveVoteList {
		k.SetObserveVote(ctx, elem)
	}
	// Set all the poolBalance
	for _, elem := range genState.PoolBalanceList {
		k.SetPoolBalance(ctx, elem)
	}
	// Set all the poolBalanap
	for _, elem := range genState.PoolBalanapList {
		k.SetPoolBalanap(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.FeeBalanceList = k.GetAllFeeBalance(ctx)
	genesis.KeysignVoteDataList = k.GetAllKeysignVoteData(ctx)
	genesis.ObserveVoteList = k.GetAllObserveVote(ctx)
	genesis.PoolBalanceList = k.GetAllPoolBalance(ctx)
	genesis.PoolBalanapList = k.GetAllPoolBalanap(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
