package router

import "github.com/google/wire"

var ProvideSet = wire.NewSet(
	NewRouter,
)
