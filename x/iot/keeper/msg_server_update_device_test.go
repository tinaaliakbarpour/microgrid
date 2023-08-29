package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tinaaliakbarpour/microgrid/x/iot/testutils"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestUpdateDevice(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	_, err := msgServer.UpdateDeviceStatus(context, &types.MsgUpdateDeviceStatus{
		Creator: testutils.Carol,
		Voltage: 100,
		Power:   1000,
		Others:  "a",
		Addres:  testutils.Alice,
		GridId:  1,
	})
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
	msgServer.CreateDevice(context, &types.MsgCreateDevice{
		Creator: testutils.Alice,
		GridId:  0,
	})

	_, err = msgServer.UpdateDeviceStatus(context, &types.MsgUpdateDeviceStatus{
		Creator: testutils.Bob,
		Voltage: 100,
		Power:   1000,
		Addres:  testutils.Alice,
		GridId:  0,
	})

	require.Nil(t, err)
}
