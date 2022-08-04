package mint_test

import (
	"testing"

	"human/testutil/nullify"
	"human/x/mint"
	"human/x/mint/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.mintKeeper(t)
	mint.InitGenesis(ctx, *k, genesisState)
	got := mint.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
