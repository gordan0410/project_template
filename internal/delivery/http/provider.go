package http

import "github.com/google/wire"

var ProvideSet = wire.NewSet(
	NewHttpApiServer,
)
