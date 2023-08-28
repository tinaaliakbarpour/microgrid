package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func (k msgServer) UpdateDeviceStatus(goCtx context.Context, msg *types.MsgUpdateDeviceStatus) (*types.MsgUpdateDeviceStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	device, found := k.GetDevice(ctx, msg.GridId, msg.Addres)
	if !found {

	}

	device.Voltage = msg.Voltage
	device.Power = msg.Power
	device.Others = msg.Others

	return &types.MsgUpdateDeviceStatusResponse{}, nil
}
