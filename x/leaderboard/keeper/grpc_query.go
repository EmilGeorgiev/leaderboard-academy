package keeper

import (
	"github.com/EmilGeorgiev/leaderboard/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
