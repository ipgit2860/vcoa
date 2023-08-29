package keeper_test

import (
	"testing"

	testkeeper "github.com/ipgit2860/vcoa/testutil/keeper"
	"github.com/ipgit2860/vcoa/x/vnft/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VnftKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
