package iot

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tinaaliakbarpour/microgrid/x/iot/keeper"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the grid
	for _, elem := range genState.GridList {
		k.SetGrid(ctx, elem)
	}

	// Set grid count
	k.SetGridCount(ctx, genState.GridCount)
	// Set all the device
	for _, elem := range genState.DeviceList {
		k.SetDevice(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.GridList = k.GetAllGrid(ctx)
	genesis.GridCount = k.GetGridCount(ctx)
	genesis.DeviceList = k.GetAllDevice(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
