package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tinaaliakbarpour/microgrid/x/iot/testutils"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestDeleteGrid(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	_, err := msgServer.DeleteGrid(context, &types.MsgDeleteGrid{
		Creator: testutils.Alice,
		Id:      1,
	})
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	msgServer.CreateGrid(context, &types.MsgCreateGrid{
		Creator:   testutils.Alice,
		Name:      "a",
		CenterLat: 458493,
		CenterLon: 458359,
		Side:      1000,
	})

	_, err = msgServer.DeleteGrid(context, &types.MsgDeleteGrid{
		Creator: testutils.Alice,
		Id:      0,
	})

	require.Nil(t, err)
}
