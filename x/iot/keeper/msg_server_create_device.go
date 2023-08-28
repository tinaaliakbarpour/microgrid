package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func (k msgServer) CreateDevice(goCtx context.Context, msg *types.MsgCreateDevice) (*types.MsgCreateDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	device := types.Device{
		GridId:  msg.GridId,
		Address: msg.Creator,
		Lat:     msg.Lat,
		Lon:     msg.Lon,
		Power:   msg.Power,
		Voltage: msg.Voltage,
		Others:  msg.Others,
	}

	k.SetDevice(ctx, device)

	return &types.MsgCreateDeviceResponse{}, nil
}
