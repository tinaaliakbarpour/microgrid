package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tinaaliakbarpour/microgrid/x/iot/testutils"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestDeleteDevice(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	_, err := msgServer.DeleteDevice(context, &types.MsgDeleteDevice{
		Creator: testutils.Alice,
		GridId:  0,
		Address: testutils.Bob,
	})
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	msgServer.CreateDevice(context, &types.MsgCreateDevice{
		Creator: testutils.Alice,
		GridId:  2,
		Lat:     36787,
		Lon:     52464,
		Voltage: 100,
		Power:   1000,
		Others:  "hey",
	})

	_, err = msgServer.DeleteDevice(context, &types.MsgDeleteDevice{
		Creator: testutils.Alice,
		GridId:  2,
		Address: testutils.Alice,
	})

	require.Nil(t, err)
}
