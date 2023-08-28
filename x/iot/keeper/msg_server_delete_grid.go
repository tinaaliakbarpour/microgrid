package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func (k msgServer) DeleteGrid(goCtx context.Context, msg *types.MsgDeleteGrid) (*types.MsgDeleteGridResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.GetGrid(ctx, msg.Id); !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.RemoveGrid(ctx, msg.Id)

	return &types.MsgDeleteGridResponse{}, nil
}
