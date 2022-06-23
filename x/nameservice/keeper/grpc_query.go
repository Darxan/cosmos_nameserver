package keeper

import (
	"github.com/darxan/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
