package iot_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tinaaliakbarpour/microgrid/testutil/keeper"
	"github.com/tinaaliakbarpour/microgrid/testutil/nullify"
	"github.com/tinaaliakbarpour/microgrid/x/iot"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		GridList: []types.Grid{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		GridCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IotKeeper(t)
	iot.InitGenesis(ctx, *k, genesisState)
	got := iot.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.GridList, got.GridList)
	require.Equal(t, genesisState.GridCount, got.GridCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
