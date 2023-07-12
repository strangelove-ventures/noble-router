package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble-router/x/router/keeper"
	"github.com/strangelove-ventures/noble-router/x/router/types"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/strangelove-ventures/noble/testutil/sample"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

type inFlightPacketWrapper struct {
	address        string
	inFlightPacket types.InFlightPacket
}

func createNInFlightPacket(keeper *keeper.Keeper, ctx sdk.Context, n int) []inFlightPacketWrapper {
	items := make([]inFlightPacketWrapper, n)
	for i := range items {
		items[i].address = sample.AccAddress()

		keeper.SetInFlightPacket(ctx, items[i].inFlightPacket)
	}
	return items
}

func TestInFlightPacketGet(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNInFlightPacket(routerKeeper, ctx, 10)
	for _, item := range items {
		rst, found := routerKeeper.GetInFlightPacket(
			ctx,
			item.inFlightPacket.SourceDomainSender,
			item.inFlightPacket.Nonce)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item.inFlightPacket),
			nullify.Fill(&rst),
		)
	}
}

func TestInFlightPacketRemove(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNInFlightPacket(routerKeeper, ctx, 10)
	for _, item := range items {
		routerKeeper.DeleteInFlightPacket(ctx, item.address)
		_, found := routerKeeper.GetInFlightPacket(ctx, item.address)
		require.False(t, found)
	}
}

func TestInFlightPacketGetAll(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNInFlightPacket(routerKeeper, ctx, 10)
	inFlightPacket := make([]types.InFlightPacket, len(items))
	for i, item := range items {
		inFlightPacket[i] = item.inFlightPacket
	}
	require.ElementsMatch(t,
		nullify.Fill(inFlightPacket),
		nullify.Fill(routerKeeper.GetAllInFlightPackets(ctx)),
	)
}
