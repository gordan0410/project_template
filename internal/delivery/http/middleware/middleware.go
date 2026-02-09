package middleware

import (
	"github.com/gin-gonic/gin"
)

type MiddlewareI interface {
	CorsMiddleware() gin.HandlerFunc
	RecoverMiddleware() gin.HandlerFunc
	AddTraceID() gin.HandlerFunc
}
type middleware struct {
}

func NewMiddleware() MiddlewareI {
	return &middleware{}
}
