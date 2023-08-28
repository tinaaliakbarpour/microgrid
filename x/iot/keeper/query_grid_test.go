package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/tinaaliakbarpour/microgrid/testutil/keeper"
	"github.com/tinaaliakbarpour/microgrid/testutil/nullify"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestGridQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGrid(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetGridRequest
		response *types.QueryGetGridResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetGridRequest{Id: msgs[0].Id},
			response: &types.QueryGetGridResponse{Grid: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetGridRequest{Id: msgs[1].Id},
			response: &types.QueryGetGridResponse{Grid: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetGridRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Grid(wctx, tc.request)
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

func TestGridQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IotKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGrid(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGridRequest {
		return &types.QueryAllGridRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GridAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Grid), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Grid),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GridAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Grid), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Grid),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GridAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Grid),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GridAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
