package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

// SetDevice set a specific device in the store from its index
func (k Keeper) SetDevice(ctx sdk.Context, device types.Device) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKeyPrefix))
	b := k.cdc.MustMarshal(&device)
	store.Set(types.DeviceKey(
		device.GridId,
		device.Address,
	), b)
}

// GetDevice returns a device from its index
func (k Keeper) GetDevice(
	ctx sdk.Context,
	gridId uint64,
	address string,

) (val types.Device, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKeyPrefix))

	b := store.Get(types.DeviceKey(
		gridId,
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDevice removes a device from the store
func (k Keeper) RemoveDevice(
	ctx sdk.Context,
	gridId uint64,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKeyPrefix))
	store.Delete(types.DeviceKey(
		gridId,
		address,
	))
}

// GetAllDevice returns all device
func (k Keeper) GetAllDevice(ctx sdk.Context) (list []types.Device) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Device
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
