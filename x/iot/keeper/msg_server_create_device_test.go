package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinaaliakbarpour/microgrid/x/iot/testutils"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestCreateDevice(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	_, err := msgServer.CreateDevice(context, &types.MsgCreateDevice{
		Creator: testutils.Alice,
	})
	require.Nil(t, err)
}
