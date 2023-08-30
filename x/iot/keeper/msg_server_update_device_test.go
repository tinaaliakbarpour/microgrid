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

func TestUpdateDevice(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	msgServer.CreateDevice(context, &types.MsgCreateDevice{
		Creator: testutils.Alice,
		GridId:  0,
		Lat:     469325,
		Lon:     455424,
	})
	tests := []struct {
		desc     string
		request  *types.MsgUpdateDeviceStatus
		response *types.MsgUpdateDeviceStatusResponse
		err      error
	}{
		{
			desc: "Success",
			request: &types.MsgUpdateDeviceStatus{
				Creator: testutils.Alice,
				Voltage: 100,
				Power:   1000,
				Others:  "hello",
				GridId:  0,
			},
			response: &types.MsgUpdateDeviceStatusResponse{},
			err:      nil,
		},

		{
			desc: "NotFound",
			request: &types.MsgUpdateDeviceStatus{
				Creator: testutils.Alice,
				Voltage: 100,
				Power:   1000,
				Others:  "bye",
				GridId:  100,
			},
			err: sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d-%s doesn't exist", 100, testutils.Alice)),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := msgServer.UpdateDeviceStatus(context, tc.request)
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
