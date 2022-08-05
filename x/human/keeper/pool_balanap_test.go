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

func createNPoolBalanap(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PoolBalanap {
	items := make([]types.PoolBalanap, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPoolBalanap(ctx, items[i])
	}
	return items
}

func TestPoolBalanapGet(t *testing.T) {
	keeper, ctx := keepertest.HumanKeeper(t)
	items := createNPoolBalanap(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPoolBalanap(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPoolBalanapRemove(t *testing.T) {
	keeper, ctx := keepertest.HumanKeeper(t)
	items := createNPoolBalanap(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePoolBalanap(ctx,
			item.Index,
		)
		_, found := keeper.GetPoolBalanap(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPoolBalanapGetAll(t *testing.T) {
	keeper, ctx := keepertest.HumanKeeper(t)
	items := createNPoolBalanap(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPoolBalanap(ctx)),
	)
}
