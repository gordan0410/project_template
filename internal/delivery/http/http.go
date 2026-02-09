package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"project_template/internal/delivery/http/handler"
	"project_template/internal/delivery/http/middleware"
	"project_template/internal/delivery/http/router"

	"github.com/gin-gonic/gin"
	common_log "github.com/gordan0410/common/log"
)

type HttpApiServer struct {
	engine           *gin.Engine
	router           router.RouterI
	srv              *http.Server
	logger           *common_log.CommonLogger
	middleware       middleware.MiddlewareI
	heartbeatHandler handler.HeartbeatHandlerI
}

func NewHttpApiServer(
	logger *common_log.CommonLogger,
	middleware middleware.MiddlewareI,
	router router.RouterI,
) *HttpApiServer {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	return &HttpApiServer{
		engine:     engine,
		router:     router,
		middleware: middleware,
		logger:     &common_log.CommonLogger{Logger: logger.With().Str(common_log.LogField_Component, "http_server").Logger()},
	}
}

func (e *HttpApiServer) Run(ctx context.Context, addr ...string) {
	e.initGlobalMiddleware()
	e.initGroupRouter()
	e.initRouter()

	e.logger.WithCtx(ctx).Info().Msg(fmt.Sprintf("start http listening port %s", addr[0]))
	srv := &http.Server{
		Addr: addr[0],
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// http middleware entrance
			e.engine.ServeHTTP(w, r)
		}),
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      600 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.logger.WithCtx(ctx).Panic().Interface("err", err).Msg("backend start failed")
			log.Fatalf("listen: %s\n", err)
		}
	}()

	e.srv = srv
}

func (e *HttpApiServer) Shutdown(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	if err := e.srv.Shutdown(ctx); err != nil {
		e.logger.WithCtx(ctx).Error().Interface("err", err).Msg("http server shutdown error")
	} else {
		e.logger.WithCtx(ctx).Info().Msg("http server shutdown")
	}
}

func (e *HttpApiServer) initGlobalMiddleware() {
	e.engine.Use(e.middleware.RecoverMiddleware())
	e.engine.Use(e.middleware.AddTraceID())
	e.engine.Use(e.middleware.CorsMiddleware())
}

func (e *HttpApiServer) initGroupRouter() {
	e.router.InitGroup(e.engine)
}

func (e *HttpApiServer) initRouter() {
	e.router.InitHeathRoute(e.engine)
}
