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

func TestDeleteGrid(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	msgServer.CreateGrid(context, &types.MsgCreateGrid{
		Creator:   testutils.Bob,
		Name:      "first",
		CenterLat: 452546,
		CenterLon: 284290,
		Side:      1999,
	})
	tests := []struct {
		desc     string
		request  *types.MsgDeleteGrid
		response *types.MsgDeleteGridResponse
		err      error
	}{
		{
			desc: "Success",
			request: &types.MsgDeleteGrid{
				Creator: testutils.Alice,
				Id:      0,
			},
			response: &types.MsgDeleteGridResponse{},
			err:      nil,
		},

		{
			desc: "NotFound",
			request: &types.MsgDeleteGrid{
				Creator: testutils.Alice,
				Id:      0,
			},
			err: sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d-%s doesn't exist", 0, testutils.Bob)),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := msgServer.DeleteGrid(context, tc.request)
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
