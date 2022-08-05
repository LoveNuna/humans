package human_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "human/testutil/keeper"
	"human/testutil/nullify"
	"human/x/human"
	"human/x/human/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FeeBalanceList: []types.FeeBalance{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		KeysignVoteDataList: []types.KeysignVoteData{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ObserveVoteList: []types.ObserveVote{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		PoolBalanceList: []types.PoolBalance{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		PoolBalanapList: []types.PoolBalanap{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HumanKeeper(t)
	human.InitGenesis(ctx, *k, genesisState)
	got := human.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FeeBalanceList, got.FeeBalanceList)
	require.ElementsMatch(t, genesisState.KeysignVoteDataList, got.KeysignVoteDataList)
	require.ElementsMatch(t, genesisState.ObserveVoteList, got.ObserveVoteList)
	require.ElementsMatch(t, genesisState.PoolBalanceList, got.PoolBalanceList)
	require.ElementsMatch(t, genesisState.PoolBalanapList, got.PoolBalanapList)
	// this line is used by starport scaffolding # genesis/test/assert
}
