package keeper_test

import (
	"fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tinaaliakbarpour/microgrid/testutil/nullify"
	"github.com/tinaaliakbarpour/microgrid/x/iot/testutils"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestDeleteDevice(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	msgServer.CreateDevice(context, &types.MsgCreateDevice{
		Creator: testutils.Bob,
		GridId:  0,
		Lat:     42579,
		Lon:     54540,
	})
	tests := []struct {
		desc     string
		request  *types.MsgDeleteDevice
		response *types.MsgDeleteDeviceResponse
		err      error
	}{
		{
			desc: "Success",
			request: &types.MsgDeleteDevice{
				Creator: testutils.Alice,
				GridId:  0,
				Address: testutils.Bob,
			},
			response: &types.MsgDeleteDeviceResponse{},
			err:      nil,
		},

		{
			desc: "NotFound",
			request: &types.MsgDeleteDevice{
				Creator: testutils.Alice,
				GridId:  0,
				Address: testutils.Bob,
			},
			err: sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d-%s doesn't exist", 0, testutils.Bob)),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := msgServer.DeleteDevice(context, tc.request)
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
