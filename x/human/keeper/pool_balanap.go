package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"human/x/human/types"
)

// SetPoolBalanap set a specific poolBalanap in the store from its index
func (k Keeper) SetPoolBalanap(ctx sdk.Context, poolBalanap types.PoolBalanap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolBalanapKeyPrefix))
	b := k.cdc.MustMarshal(&poolBalanap)
	store.Set(types.PoolBalanapKey(
		poolBalanap.Index,
	), b)
}

// GetPoolBalanap returns a poolBalanap from its index
func (k Keeper) GetPoolBalanap(
	ctx sdk.Context,
	index string,

) (val types.PoolBalanap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolBalanapKeyPrefix))

	b := store.Get(types.PoolBalanapKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePoolBalanap removes a poolBalanap from the store
func (k Keeper) RemovePoolBalanap(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolBalanapKeyPrefix))
	store.Delete(types.PoolBalanapKey(
		index,
	))
}

// GetAllPoolBalanap returns all poolBalanap
func (k Keeper) GetAllPoolBalanap(ctx sdk.Context) (list []types.PoolBalanap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolBalanapKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PoolBalanap
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
