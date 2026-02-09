//go:build wireinject
// +build wireinject

package main

import (
	"project_template/internal/config"
	"project_template/internal/delivery/http"
	"project_template/internal/delivery/http/handler"
	"project_template/internal/delivery/http/middleware"
	"project_template/internal/delivery/http/router"
	"project_template/internal/usecase"

	"github.com/google/wire"
	common_dto "github.com/gordan0410/common/dto"
	common_log "github.com/gordan0410/common/log"
)

type Application struct {
	//version    *common_dto.VersionInfo
	envConfig  *config.EnvConfig
	logger     *common_log.CommonLogger
	httpServer *http.HttpApiServer
}

func NewApplication(
	//version *common_dto.VersionInfo,
	envConfig *config.EnvConfig,
	logger *common_log.CommonLogger,
	httpServer *http.HttpApiServer,
) *Application {
	return &Application{
		//version:    version,
		envConfig:  envConfig,
		logger:     logger,
		httpServer: httpServer,
	}
}

func InitializeApplication(
	version common_dto.VersionInfo,
	logger *common_log.CommonLogger,
	envConfig *config.EnvConfig,
	masterDBConfig *config.DatabaseConfig,
	redisConfig *config.RedisConfig,
) *Application {
	wire.Build(
		NewApplication,
		http.ProvideSet,
		handler.ProvideSet,
		router.ProvideSet,
		usecase.ProvideSet,
		middleware.ProvideSet,
		//repository.ProvideSet,
		//api.ProvideSet,
	)
	return &Application{}
}
