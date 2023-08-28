package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func (k msgServer) RegisterAdmin(goCtx context.Context, msg *types.MsgRegisterAdmin) (*types.MsgRegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetGrid(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	val.Admins = append(val.Admins, msg.Address)
	k.SetGrid(ctx, val)

	return &types.MsgRegisterAdminResponse{
		Grid: &val,
	}, nil
}

//TODO : may be we need some authentication on creator account and also the account that is going to
//be added to admins
