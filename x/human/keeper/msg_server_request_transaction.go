package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"human/x/human/types"
)

func (k msgServer) RequestTransaction(goCtx context.Context, msg *types.MsgRequestTransaction) (*types.MsgRequestTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRequestTransactionResponse{}, nil
}
