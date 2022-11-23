package keeper_test

import (
	"testing"

	testkeeper "NICOTEEN/testutil/keeper"
	"NICOTEEN/x/nicoteen/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NicoteenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
