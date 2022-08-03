package keeper

import (
	"human/x/mint/types"
)

var _ types.QueryServer = Keeper{}
