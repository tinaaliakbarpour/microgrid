package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func (k msgServer) DeleteDevice(goCtx context.Context, msg *types.MsgDeleteDevice) (*types.MsgDeleteDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.GetDevice(ctx, msg.GridId, msg.Address); !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d-%ddoesn't exist", msg.GridId, msg.Address))
	}
	k.RemoveDevice(ctx, msg.GridId, msg.Address)
	return &types.MsgDeleteDeviceResponse{}, nil
}

//todo : maybe we need so sort of authentication on msg creator
