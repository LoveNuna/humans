package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "human/testutil/keeper"
	"human/testutil/nullify"
	"human/x/human/keeper"
	"human/x/human/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFeeBalance(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FeeBalance {
	items := make([]types.FeeBalance, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetFeeBalance(ctx, items[i])
	}
	return items
}

func TestFeeBalanceGet(t *testing.T) {
	keeper, ctx := keepertest.HumanKeeper(t)
	items := createNFeeBalance(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFeeBalance(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFeeBalanceRemove(t *testing.T) {
	keeper, ctx := keepertest.HumanKeeper(t)
	items := createNFeeBalance(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFeeBalance(ctx,
			item.Index,
		)
		_, found := keeper.GetFeeBalance(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestFeeBalanceGetAll(t *testing.T) {
	keeper, ctx := keepertest.HumanKeeper(t)
	items := createNFeeBalance(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFeeBalance(ctx)),
	)
}
