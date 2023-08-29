package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tinaaliakbarpour/microgrid/x/iot/testutils"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestRegisterAdmin(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	_, err := msgServer.RegisterAdmin(context, &types.MsgRegisterAdmin{
		Creator: testutils.Alice,
		Id:      1,
		Address: testutils.Alice,
	})
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	msgServer.CreateGrid(context, &types.MsgCreateGrid{
		Creator:   testutils.Alice,
		Name:      "a",
		CenterLat: 458493,
		CenterLon: 458359,
		Side:      1000,
	})

	_, err = msgServer.RegisterAdmin(context, &types.MsgRegisterAdmin{
		Creator: testutils.Alice,
		Id:      0,
		Address: testutils.Alice,
	})

	require.Nil(t, err)
}
