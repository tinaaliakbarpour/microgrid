package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func (k msgServer) UpdateDeviceStatus(goCtx context.Context, msg *types.MsgUpdateDeviceStatus) (*types.MsgUpdateDeviceStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	device, found := k.GetDevice(ctx, msg.GridId, msg.Addres)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d-%s doesn't exist", msg.GridId, msg.Addres))
	}

	device.Voltage = msg.Voltage
	device.Power = msg.Power
	device.Others = msg.Others

	k.SetDevice(ctx, device)

	return &types.MsgUpdateDeviceStatusResponse{}, nil
}

//also there sould some authentication for msg creator
