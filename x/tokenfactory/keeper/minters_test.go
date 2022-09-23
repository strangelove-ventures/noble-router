package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "noble/testutil/keeper"
	"noble/testutil/nullify"
	"noble/x/tokenfactory/keeper"
	"noble/x/tokenfactory/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMinters(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Minters {
	items := make([]types.Minters, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetMinters(ctx, items[i])
	}
	return items
}

func TestMintersGet(t *testing.T) {
	keeper, ctx := keepertest.TokenfactoryKeeper(t)
	items := createNMinters(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMinters(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMintersRemove(t *testing.T) {
	keeper, ctx := keepertest.TokenfactoryKeeper(t)
	items := createNMinters(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMinters(ctx,
			item.Address,
		)
		_, found := keeper.GetMinters(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestMintersGetAll(t *testing.T) {
	keeper, ctx := keepertest.TokenfactoryKeeper(t)
	items := createNMinters(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMinters(ctx)),
	)
}