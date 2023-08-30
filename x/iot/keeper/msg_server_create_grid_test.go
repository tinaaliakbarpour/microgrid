package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tinaaliakbarpour/microgrid/testutil/nullify"
	"github.com/tinaaliakbarpour/microgrid/x/iot/testutils"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestCreateGrid(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	tests := []struct {
		desc     string
		request  *types.MsgCreateGrid
		response *types.MsgCreateGridResponse
		err      error
	}{
		{
			desc: "Success",
			request: &types.MsgCreateGrid{
				Creator:   testutils.Alice,
				Name:      "first",
				CenterLat: 3679734,
				CenterLon: 3842005,
				Side:      10000,
			},
			response: &types.MsgCreateGridResponse{
				Id: 0,
			},
			err: nil,
		},

		{
			desc: "AlreadyExist",
			request: &types.MsgCreateGrid{
				Creator:   testutils.Alice,
				Name:      "first",
				CenterLat: 3679734,
				CenterLon: 3842005,
				Side:      10000,
			},
			err: sdkerrors.Wrap(types.ErrorDuplicate, "key first already exist"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := msgServer.CreateGrid(context, tc.request)
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
