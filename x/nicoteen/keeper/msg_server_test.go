package keeper_test

import (
	"context"
	"testing"

	keepertest "NICOTEEN/testutil/keeper"
	"NICOTEEN/x/nicoteen/keeper"
	"NICOTEEN/x/nicoteen/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NicoteenKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
