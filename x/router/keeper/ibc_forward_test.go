package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble-router/x/router/types"
	keepertest "github.com/strangelove-ventures/noble/testutil/keeper"
	"github.com/strangelove-ventures/noble/testutil/nullify"
	"github.com/strangelove-ventures/noble/testutil/sample"
	"github.com/strangelove-ventures/noble/x/router/keeper"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

type ibcForwardWrapper struct {
	address    string
	ibcForward types.IBCForward
}

func createNIBCForward(keeper *keeper.Keeper, ctx sdk.Context, n int) []ibcForwardWrapper {
	items := make([]ibcForwardWrapper, n)
	for i := range items {
		items[i].address = sample.AccAddress()

		keeper.SetIBCForward(ctx, items[i].ibcForward)
	}
	return items
}

func TestIBCForwardGet(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNIBCForward(routerKeeper, ctx, 10)
	for _, item := range items {
		rst, found := routerKeeper.GetIBCForward(ctx,
			item.address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item.ibcForward),
			nullify.Fill(&rst),
		)
	}
}

func TestIBCForwardRemove(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNIBCForward(routerKeeper, ctx, 10)
	for _, item := range items {
		routerKeeper.DeleteIBCForward(ctx, item.address)
		_, found := routerKeeper.GetIBCForward(ctx, item.address)
		require.False(t, found)
	}
}

func TestIBCForwardGetAll(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)
	items := createNIBCForward(routerKeeper, ctx, 10)
	ibcForward := make([]types.IBCForward, len(items))
	for i, item := range items {
		ibcForward[i] = item.ibcForward
	}
	require.ElementsMatch(t,
		nullify.Fill(ibcForward),
		nullify.Fill(routerKeeper.GetAllIBCForwards(ctx)),
	)
}
