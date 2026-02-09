package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (m *middleware) CorsMiddleware() gin.HandlerFunc {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	config.AddAllowMethods("*")
	config.AllowWildcard = true
	config.AllowCredentials = true

	return cors.New(config)
}
