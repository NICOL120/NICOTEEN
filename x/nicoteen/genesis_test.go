package nicoteen_test

import (
	"testing"

	keepertest "NICOTEEN/testutil/keeper"
	"NICOTEEN/testutil/nullify"
	"NICOTEEN/x/nicoteen"
	"NICOTEEN/x/nicoteen/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NicoteenKeeper(t)
	nicoteen.InitGenesis(ctx, *k, genesisState)
	got := nicoteen.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
