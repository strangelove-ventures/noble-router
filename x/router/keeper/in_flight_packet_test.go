package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble-router/x/router/keeper"
	"github.com/strangelove-ventures/noble-router/x/router/types"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNInFlightPacket(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.InFlightPacket {
	items := make([]types.InFlightPacket, n)
	for i := range items {
		items[i].SourceDomainSender = strconv.Itoa(i)
		items[i].Nonce = uint64(i)

		keeper.SetInFlightPacket(ctx, items[i])
	}
	return items
}

func TestInFlightPacketGet(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNInFlightPacket(routerKeeper, ctx, 10)
	for _, item := range items {
		rst, found := routerKeeper.GetInFlightPacket(
			ctx,
			item.SourceDomainSender,
			item.Nonce)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestInFlightPacketRemove(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNInFlightPacket(routerKeeper, ctx, 10)
	for _, item := range items {
		routerKeeper.DeleteInFlightPacket(
			ctx,
			item.SourceDomainSender,
			item.Nonce,
		)
		_, found := routerKeeper.GetInFlightPacket(
			ctx,
			item.SourceDomainSender,
			item.Nonce,
		)
		require.False(t, found)
	}
}

func TestInFlightPacketGetAll(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNInFlightPacket(routerKeeper, ctx, 10)
	inFlightPacket := make([]types.InFlightPacket, len(items))
	for i, item := range items {
		inFlightPacket[i] = item
	}
	require.ElementsMatch(t,
		nullify.Fill(inFlightPacket),
		nullify.Fill(routerKeeper.GetAllInFlightPackets(ctx)),
	)
}
