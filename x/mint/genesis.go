package mint

import (
	"human/x/mint/keeper"
	"human/x/mint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, ak types.AccountKeeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetMinter(ctx, genState.Minter)
	k.SetParams(ctx, genState.Params)
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return types.NewGenesisState(minter, params)
}
