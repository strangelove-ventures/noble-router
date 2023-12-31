package keeper_test

import (
	"testing"

	testkeeper "github.com/strangelove-ventures/noble-router/testutil/keeper"
	"github.com/strangelove-ventures/noble-router/x/router/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RouterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	got := k.GetParams(ctx)

	require.EqualValues(t, params, got)
}
