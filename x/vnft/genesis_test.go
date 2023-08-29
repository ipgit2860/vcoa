package vnft_test

import (
	"testing"

	keepertest "github.com/ipgit2860/vcoa/testutil/keeper"
	"github.com/ipgit2860/vcoa/testutil/nullify"
	"github.com/ipgit2860/vcoa/x/vnft"
	"github.com/ipgit2860/vcoa/x/vnft/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VnftKeeper(t)
	vnft.InitGenesis(ctx, *k, genesisState)
	got := vnft.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
