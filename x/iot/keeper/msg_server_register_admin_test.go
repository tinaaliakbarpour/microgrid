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

func TestRegisterAdmin(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	msgServer.CreateGrid(context, &types.MsgCreateGrid{
		Creator:   testutils.Bob,
		Name:      "first",
		CenterLat: 14683,
		CenterLon: 483912,
		Side:      1000,
	})
	tests := []struct {
		desc     string
		request  *types.MsgRegisterAdmin
		response *types.MsgRegisterAdminResponse
		err      error
	}{
		{
			desc: "Success",
			request: &types.MsgRegisterAdmin{
				Creator: testutils.Alice,
				Id:      0,
				Address: testutils.Carol,
			},
			response: &types.MsgRegisterAdminResponse{
				Grid: &types.Grid{
					Id:        0,
					Name:      "first",
					CenterLat: 14683,
					CenterLon: 483912,
					Side:      1000,
					Admins:    []string{testutils.Carol},
				},
			},
			err: nil,
		},

		{
			desc: "NotFound",
			request: &types.MsgRegisterAdmin{
				Creator: testutils.Alice,
				Id:      100,
				Address: testutils.Bob,
			},
			err: sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", 0)),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := msgServer.RegisterAdmin(context, tc.request)
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
