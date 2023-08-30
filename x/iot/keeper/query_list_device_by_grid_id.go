package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListDeviceByGridId(goCtx context.Context, req *types.QueryListDeviceByGridIdRequest) (*types.QueryListDeviceByGridIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DevicesKeyByGridId(req.GridId))

	var devices []*types.Device

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var device types.Device
		if err := k.cdc.Unmarshal(value, &device); err != nil {
			return err
		}

		devices = append(devices, &device)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryListDeviceByGridIdResponse{
		Device:     devices,
		Pagination: pageRes,
	}, nil
}
