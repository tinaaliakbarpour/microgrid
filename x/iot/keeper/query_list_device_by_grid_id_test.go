package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tinaaliakbarpour/microgrid/testutil/keeper"
	"github.com/tinaaliakbarpour/microgrid/testutil/nullify"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeviceByGridId(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDevice(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryListDeviceByGridIdRequest
		response *types.QueryListDeviceByGridIdResponse
		err      error
	}{
		{
			desc:    "First",
			request: &types.QueryListDeviceByGridIdRequest{GridId: 0},
			response: &types.QueryListDeviceByGridIdResponse{
				Device: []*types.Device{
					&msgs[0],
					&msgs[1],
				},
			},
			err: nil,
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryListDeviceByGridIdRequest{GridId: 3},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ListDeviceByGridId(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
