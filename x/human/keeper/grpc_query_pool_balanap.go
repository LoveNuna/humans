package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"human/x/human/types"
)

func (k Keeper) PoolBalanapAll(c context.Context, req *types.QueryAllPoolBalanapRequest) (*types.QueryAllPoolBalanapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var poolBalanaps []types.PoolBalanap
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	poolBalanapStore := prefix.NewStore(store, types.KeyPrefix(types.PoolBalanapKeyPrefix))

	pageRes, err := query.Paginate(poolBalanapStore, req.Pagination, func(key []byte, value []byte) error {
		var poolBalanap types.PoolBalanap
		if err := k.cdc.Unmarshal(value, &poolBalanap); err != nil {
			return err
		}

		poolBalanaps = append(poolBalanaps, poolBalanap)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPoolBalanapResponse{PoolBalanap: poolBalanaps, Pagination: pageRes}, nil
}

func (k Keeper) PoolBalanap(c context.Context, req *types.QueryGetPoolBalanapRequest) (*types.QueryGetPoolBalanapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPoolBalanap(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPoolBalanapResponse{PoolBalanap: val}, nil
}
