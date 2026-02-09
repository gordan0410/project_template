package middleware

import "github.com/google/wire"

var ProvideSet = wire.NewSet(
	NewMiddleware,
)
