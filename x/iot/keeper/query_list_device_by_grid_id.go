package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListDeviceByGridId(goCtx context.Context, req *types.QueryListDeviceByGridIdRequest) (*types.QueryListDeviceByGridIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	devices := k.GetDevicesByGridId(
		ctx,
		req.GridId,
	)
	if len(devices) == 0 {
		return nil, status.Error(codes.NotFound, "not found")
	}

	response := &types.QueryListDeviceByGridIdResponse{}

	for _, device := range devices {
		response.Device = append(response.Device, &types.Device{
			GridId:  device.GridId,
			Address: device.Address,
			Lat:     device.Lat,
			Lon:     device.Lon,
			Power:   device.Power,
			Voltage: device.Voltage,
			Others:  device.Others,
		})
	}

	return response, nil
}
