package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tinaaliakbarpour/microgrid/testutil/keeper"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeviceByGridId(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDevice(keeper, ctx, 2)
	// _ = msgs

	tests := []struct {
		desc     string
		request  *types.QueryListDeviceByGridIdRequest
		response *types.QueryListDeviceByGridIdResponse
		err      error
	}{

		// {
		// 	desc:    "KeyNotFound",
		// 	request: &types.QueryListDeviceByGridIdRequest{GridId: 100},
		// 	err:     status.Error(codes.NotFound, "not found"),
		// },
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
		{
			desc:    "First",
			request: &types.QueryListDeviceByGridIdRequest{GridId: 0},
			response: &types.QueryListDeviceByGridIdResponse{
				Device: []*types.Device{
					&msgs[0],
				},
			},
			err: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ListDeviceByGridId(wctx, tc.request)
			fmt.Println(err)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response.Device, response.Device)
			}
		})
	}
}
