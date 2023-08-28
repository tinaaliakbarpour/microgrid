package iot

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tinaaliakbarpour/microgrid/testutil/sample"
	iotsimulation "github.com/tinaaliakbarpour/microgrid/x/iot/simulation"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = iotsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateGrid = "op_weight_msg_create_grid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGrid int = 100

	opWeightMsgRegisterAdmin = "op_weight_msg_register_admin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterAdmin int = 100

	opWeightMsgDeleteGrid = "op_weight_msg_delete_grid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteGrid int = 100

	opWeightMsgCreateDevice = "op_weight_msg_create_device"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDevice int = 100

	opWeightMsgUpdateDeviceStatus = "op_weight_msg_update_device_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDeviceStatus int = 100

	opWeightMsgDeleteDevice = "op_weight_msg_delete_device"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDevice int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	iotGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&iotGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateGrid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateGrid, &weightMsgCreateGrid, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGrid = defaultWeightMsgCreateGrid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGrid,
		iotsimulation.SimulateMsgCreateGrid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterAdmin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterAdmin, &weightMsgRegisterAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterAdmin = defaultWeightMsgRegisterAdmin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterAdmin,
		iotsimulation.SimulateMsgRegisterAdmin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteGrid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteGrid, &weightMsgDeleteGrid, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteGrid = defaultWeightMsgDeleteGrid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteGrid,
		iotsimulation.SimulateMsgDeleteGrid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateDevice int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDevice, &weightMsgCreateDevice, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDevice = defaultWeightMsgCreateDevice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDevice,
		iotsimulation.SimulateMsgCreateDevice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDeviceStatus int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDeviceStatus, &weightMsgUpdateDeviceStatus, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDeviceStatus = defaultWeightMsgUpdateDeviceStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDeviceStatus,
		iotsimulation.SimulateMsgUpdateDeviceStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDevice int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDevice, &weightMsgDeleteDevice, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDevice = defaultWeightMsgDeleteDevice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDevice,
		iotsimulation.SimulateMsgDeleteDevice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateGrid,
			defaultWeightMsgCreateGrid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				iotsimulation.SimulateMsgCreateGrid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterAdmin,
			defaultWeightMsgRegisterAdmin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				iotsimulation.SimulateMsgRegisterAdmin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteGrid,
			defaultWeightMsgDeleteGrid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				iotsimulation.SimulateMsgDeleteGrid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateDevice,
			defaultWeightMsgCreateDevice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				iotsimulation.SimulateMsgCreateDevice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateDeviceStatus,
			defaultWeightMsgUpdateDeviceStatus,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				iotsimulation.SimulateMsgUpdateDeviceStatus(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteDevice,
			defaultWeightMsgDeleteDevice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				iotsimulation.SimulateMsgDeleteDevice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
