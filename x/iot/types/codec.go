package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateGrid{}, "iot/CreateGrid", nil)
	cdc.RegisterConcrete(&MsgRegisterAdmin{}, "iot/RegisterAdmin", nil)
	cdc.RegisterConcrete(&MsgDeleteGrid{}, "iot/DeleteGrid", nil)
	cdc.RegisterConcrete(&MsgCreateDevice{}, "iot/CreateDevice", nil)
	cdc.RegisterConcrete(&MsgUpdateDeviceStatus{}, "iot/UpdateDeviceStatus", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGrid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterAdmin{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteGrid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDevice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateDeviceStatus{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
