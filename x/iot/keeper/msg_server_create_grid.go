package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	id := k.AppendGrid(
		ctx,
		grid,
	)

	return &types.MsgCreateGridResponse{Id: id}, nil
}

//TODO : may be we need some authentication on creator account
