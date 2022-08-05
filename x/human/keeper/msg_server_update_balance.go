package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"human/x/human/types"
)

func (k msgServer) UpdateBalance(goCtx context.Context, msg *types.MsgUpdateBalance) (*types.MsgUpdateBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateBalanceResponse{}, nil
}
