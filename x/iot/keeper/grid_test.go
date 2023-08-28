package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tinaaliakbarpour/microgrid/testutil/keeper"
	"github.com/tinaaliakbarpour/microgrid/testutil/nullify"
	"github.com/tinaaliakbarpour/microgrid/x/iot/keeper"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func createNGrid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Grid {
	items := make([]types.Grid, n)
	for i := range items {
		items[i].Id = keeper.AppendGrid(ctx, items[i])
	}
	return items
}

func TestGridGet(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	items := createNGrid(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetGrid(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestGridRemove(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	items := createNGrid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGrid(ctx, item.Id)
		_, found := keeper.GetGrid(ctx, item.Id)
		require.False(t, found)
	}
}

func TestGridGetAll(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	items := createNGrid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGrid(ctx)),
	)
}

func TestGridCount(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	items := createNGrid(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetGridCount(ctx))
}
