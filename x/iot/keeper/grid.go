package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

// GetGridCount get the total number of grid
func (k Keeper) GetGridCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.GridCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetGridCount set the total number of grid
func (k Keeper) SetGridCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.GridCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendGrid appends a grid in the store with a new id and update the count
func (k Keeper) AppendGrid(
	ctx sdk.Context,
	grid types.Grid,
) uint64 {
	// Create the grid
	count := k.GetGridCount(ctx)

	// Set the ID of the appended value
	grid.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GridKey))
	appendedValue := k.cdc.MustMarshal(&grid)
	store.Set(GetGridIDBytes(grid.Id), appendedValue)

	// Update grid count
	k.SetGridCount(ctx, count+1)

	return count
}

// SetGrid set a specific grid in the store
func (k Keeper) SetGrid(ctx sdk.Context, grid types.Grid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GridKey))
	b := k.cdc.MustMarshal(&grid)
	store.Set(GetGridIDBytes(grid.Id), b)
}

// GetGrid returns a grid from its id
func (k Keeper) GetGrid(ctx sdk.Context, id uint64) (val types.Grid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GridKey))
	b := store.Get(GetGridIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGrid removes a grid from the store
func (k Keeper) RemoveGrid(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GridKey))
	store.Delete(GetGridIDBytes(id))
}

// GetAllGrid returns all grid
func (k Keeper) GetAllGrid(ctx sdk.Context) (list []types.Grid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GridKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Grid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetGridIDBytes returns the byte representation of the ID
func GetGridIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetGridIDFromBytes returns ID in uint64 format from a byte array
func GetGridIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
