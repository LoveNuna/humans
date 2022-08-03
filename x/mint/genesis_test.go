package hmint_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "human/testutil/keeper"
	"human/testutil/nullify"
	"human/x/hmint"
	"human/x/hmint/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HmintKeeper(t)
	hmint.InitGenesis(ctx, *k, genesisState)
	got := hmint.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
