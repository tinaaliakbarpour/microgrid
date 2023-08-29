package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinaaliakbarpour/microgrid/x/iot/testutils"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestCreateGrid(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGrid(context, &types.MsgCreateGrid{
		Creator:   testutils.Alice,
		Name:      "first",
		CenterLat: 378989,
		CenterLon: 378987,
		Side:      1000,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGridResponse{
		Id: 0, // TODO: update with a proper value when updated
	}, *createResponse)
}
