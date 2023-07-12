package keeper_test

import (
	"fmt"
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

type mintWrapper struct {
	attestation string
	mint        types.Mint
}

func createNMint(keeper *keeper.Keeper, ctx sdk.Context, n int) []mintWrapper {
	items := make([]mintWrapper, n)
	for i := range items {
		items[i].attestation = fmt.Sprintf("attestation %d", i)

		keeper.SetMint(ctx, items[i].mint)
	}
	return items
}

func TestMintGet(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNMint(routerKeeper, ctx, 10)
	for _, item := range items {
		rst, found := routerKeeper.GetMint(ctx,
			item.attestation,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item.mint),
			nullify.Fill(&rst),
		)
	}
}

func TestMintRemove(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNMint(routerKeeper, ctx, 10)
	for _, item := range items {
		routerKeeper.DeleteMint(ctx, item.attestation)
		_, found := routerKeeper.GetMint(ctx, item.attestation)
		require.False(t, found)
	}
}

func TestMintGetAll(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNMint(routerKeeper, ctx, 10)
	mint := make([]types.Mint, len(items))
	for i, item := range items {
		mint[i] = item.mint
	}
	require.ElementsMatch(t,
		nullify.Fill(mint),
		nullify.Fill(routerKeeper.GetAllMints(ctx)),
	)
}
