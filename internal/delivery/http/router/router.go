package router

import (
	"project_template/internal/delivery/http/handler"
	"project_template/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
)

type RouterI interface {
	SetPrefix(string)
	InitGroup(engine *gin.Engine)
	InitHeathRoute(*gin.Engine)
}

type Router struct {
	prefix   string // 依照不同微服務給不同
	apiGroup *gin.RouterGroup

	middleware       middleware.MiddlewareI
	heartbeatHandler handler.HeartbeatHandlerI
}

func NewRouter(
	middleware middleware.MiddlewareI,
	heartbeatHandler handler.HeartbeatHandlerI,
) RouterI {
	return &Router{
		middleware:       middleware,
		heartbeatHandler: heartbeatHandler,
	}
}

func (r *Router) SetPrefix(prefix string) {
	r.prefix = prefix
}

func (r *Router) InitGroup(engine *gin.Engine) {
	apiGroup := engine.Group(r.prefix)
	r.apiGroup = apiGroup
}
