package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/EmilGeorgiev/leaderboard/testutil/keeper"
	"github.com/EmilGeorgiev/leaderboard/x/leaderboard/keeper"
	"github.com/EmilGeorgiev/leaderboard/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LeaderboardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
