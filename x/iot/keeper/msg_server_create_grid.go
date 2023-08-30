package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func (k msgServer) CreateGrid(goCtx context.Context, msg *types.MsgCreateGrid) (*types.MsgCreateGridResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var grid = types.Grid{
		Name:      msg.Name,
		CenterLat: msg.CenterLat,
		CenterLon: msg.CenterLon,
		Side:      msg.Side,
	}

	//todo : it is a non efficient and simple validation, we have to optimize it
	grids := k.GetAllGrid(ctx)
	for _, g := range grids {
		if grid.Name == g.Name {
			return nil, sdkerrors.Wrap(types.ErrorDuplicate, fmt.Sprintf("key %s already exist", msg.Name))
		} else if grid.CenterLat == g.CenterLat && grid.CenterLon == g.CenterLon {
			return nil, sdkerrors.Wrap(types.ErrorDuplicate, fmt.Sprintf("key with this location: %d, %d already exist", msg.Name, msg.CenterLat, msg.CenterLon))
		}
	}
	id := k.AppendGrid(
		ctx,
		grid,
	)

	return &types.MsgCreateGridResponse{Id: id}, nil
}

//TODO : may be we need some authentication on creator account
//TODO: is there any mechanism for preventing storing duplicate values
