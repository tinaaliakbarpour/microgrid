package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func (k msgServer) CreateDevice(goCtx context.Context, msg *types.MsgCreateDevice) (*types.MsgCreateDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.GetDevice(ctx, msg.GridId, msg.Creator); found {
		return nil, sdkerrors.Wrap(types.ErrorDuplicate, fmt.Sprintf("key %d-%s already exist", msg.GridId, msg.Creator))
	}

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
