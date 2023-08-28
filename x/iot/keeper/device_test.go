package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tinaaliakbarpour/microgrid/testutil/keeper"
	"github.com/tinaaliakbarpour/microgrid/testutil/nullify"
	"github.com/tinaaliakbarpour/microgrid/x/iot/keeper"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDevice(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Device {
	items := make([]types.Device, n)
	for i := range items {
		items[i].GridId = uint64(i)

		keeper.SetDevice(ctx, items[i])
	}
	return items
}

func TestDeviceGet(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	items := createNDevice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDevice(ctx,
			item.GridId,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDeviceRemove(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	items := createNDevice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDevice(ctx,
			item.GridId,
			item.Address,
		)
		_, found := keeper.GetDevice(ctx,
			item.GridId,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestDeviceGetAll(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	items := createNDevice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDevice(ctx)),
	)
}
