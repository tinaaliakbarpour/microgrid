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

func TestCreateDevice(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	tests := []struct {
		desc     string
		request  *types.MsgCreateDevice
		response *types.MsgCreateDeviceResponse
		err      error
	}{
		{
			desc: "Success",
			request: &types.MsgCreateDevice{
				Creator: testutils.Alice,
				GridId:  0,
				Lat:     4693,
				Lon:     45234,
				Voltage: 100,
				Power:   1000,
				Others:  "hey",
			},
			response: &types.MsgCreateDeviceResponse{},
			err:      nil,
		},

		{
			desc: "AlreadyExist",
			request: &types.MsgCreateDevice{
				Creator: testutils.Alice,
				GridId:  0,
				Lat:     4693,
				Lon:     45234,
				Voltage: 100,
				Power:   1000,
				Others:  "hey",
			},
			err: sdkerrors.Wrap(types.ErrorDuplicate, fmt.Sprintf("key %d-%s already exist", 0, testutils.Alice)),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := msgServer.CreateDevice(context, tc.request)
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
