package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GridAll(goCtx context.Context, req *types.QueryAllGridRequest) (*types.QueryAllGridResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var grids []types.Grid
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	gridStore := prefix.NewStore(store, types.KeyPrefix(types.GridKey))

	pageRes, err := query.Paginate(gridStore, req.Pagination, func(key []byte, value []byte) error {
		var grid types.Grid
		if err := k.cdc.Unmarshal(value, &grid); err != nil {
			return err
		}

		grids = append(grids, grid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGridResponse{Grid: grids, Pagination: pageRes}, nil
}

func (k Keeper) Grid(goCtx context.Context, req *types.QueryGetGridRequest) (*types.QueryGetGridResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	grid, found := k.GetGrid(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetGridResponse{Grid: grid}, nil
}
