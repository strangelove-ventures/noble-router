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

func TestInFlightPacketQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RouterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNInFlightPacket(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetInFlightPacketRequest
		response *types.QueryGetInFlightPacketResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetInFlightPacketRequest{
				Sender: msgs[0].address,
			},
			response: &types.QueryGetInFlightPacketResponse{IbcForward: msgs[0].ibcForward},
		},
		{
			desc: "Second",
			request: &types.QueryGetInFlightPacketRequest{
				Sender: msgs[1].address,
			},
			response: &types.QueryGetInFlightPacketResponse{IbcForward: msgs[1].ibcForward},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetInFlightPacketRequest{
				Sender: "nothing",
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.InFlightPacket(wctx, tc.request)
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

func TestInFlightPacketQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RouterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNInFlightPacket(keeper, ctx, 5)
	InFlightPacket := make([]types.InFlightPackets, len(msgs))
	for i, msg := range msgs {
		InFlightPacket[i] = msg.ibcForward
	}

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllInFlightPacketsRequest {
		return &types.QueryAllInFlightPacketsRequest{
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
		for i := 0; i < len(InFlightPacket); i += step {
			resp, err := keeper.InFlightPackets(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.IbcForwards), step)
			require.Subset(t,
				nullify.Fill(InFlightPacket),
				nullify.Fill(resp.IbcForwards),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(InFlightPacket); i += step {
			resp, err := keeper.InFlightPackets(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.IbcForwards), step)
			require.Subset(t,
				nullify.Fill(InFlightPacket),
				nullify.Fill(resp.IbcForwards),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.InFlightPackets(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(InFlightPacket), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(InFlightPacket),
			nullify.Fill(resp.IbcForwards),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.InFlightPackets(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
