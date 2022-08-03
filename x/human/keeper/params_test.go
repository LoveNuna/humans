package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "human/testutil/keeper"
	"human/x/human/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HumanKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
