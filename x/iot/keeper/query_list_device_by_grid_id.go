package keeper

import (
	"context"
	"encoding/binary"

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

	// var devices []*types.Device
	ctx := sdk.UnwrapSDKContext(goCtx)

	var key []byte

	gridIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(gridIdBytes, req.GridId)
	key = append(key, types.KeyPrefix(types.DeviceKeyPrefix)...)
	key = append(key, gridIdBytes...)
	key = append(key, []byte("/")...)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), key)

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
