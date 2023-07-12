package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/strangelove-ventures/noble-router/x/router/types"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMintQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RouterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMint(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMintRequest
		response *types.QueryGetMintResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMintRequest{
				Attestation: "attestation1",
			},
			response: &types.QueryGetMintResponse{Mint: msgs[0].mint},
		},
		{
			desc: "Second",
			request: &types.QueryGetMintRequest{
				Attestation: "attestation2",
			},
			response: &types.QueryGetMintResponse{Mint: msgs[0].mint},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMintRequest{
				Attestation: "nothing",
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Mint(wctx, tc.request)
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

func TestMintQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RouterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMint(keeper, ctx, 5)
	Mint := make([]types.Mints, len(msgs))
	for i, msg := range msgs {
		Mint[i] = msg.mint
	}

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMintsRequest {
		return &types.QueryAllMintsRequest{
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
		for i := 0; i < len(Mint); i += step {
			resp, err := keeper.Mints(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Mints), step)
			require.Subset(t,
				nullify.Fill(Mint),
				nullify.Fill(resp.Mints),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(Mint); i += step {
			resp, err := keeper.Mints(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Mints), step)
			require.Subset(t,
				nullify.Fill(Mint),
				nullify.Fill(resp.Mints),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.Mints(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(Mint), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(Mint),
			nullify.Fill(resp.Mints),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.Mints(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
