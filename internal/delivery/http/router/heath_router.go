package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (r *Router) InitHeathRoute(engine *gin.Engine) {
	r.apiGroup.GET("/health", r.heartbeatHandler.HealthCheck)
	r.apiGroup.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
