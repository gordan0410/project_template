package handler

import (
	"project_template/internal/domain"

	"github.com/gin-gonic/gin"
	common_response "github.com/gordan0410/common/enum/response"
	common_log "github.com/gordan0410/common/log"
)

type HeartbeatHandlerI interface {
	HealthCheck(c *gin.Context)
}

type heartbeatHandler struct {
	logger        *common_log.CommonLogger
	healthUseCase domain.HealthUsecase
}

func NewHeartbeatHandler(
	logger *common_log.CommonLogger,
	healthUseCase domain.HealthUsecase,
) HeartbeatHandlerI {
	return &heartbeatHandler{
		logger:        &common_log.CommonLogger{Logger: logger.With().Str(common_log.LogField_Handler, common_log.LogType_Heartbeat).Logger()},
		healthUseCase: healthUseCase,
	}
}

func (handler *heartbeatHandler) HealthCheck(c *gin.Context) {
	handler.healthUseCase.CheckHealth(c)
	handler.logger.WithCtx(c).Info().Msg("Health check passed")
	c.JSON(common_response.GetSuccessResponse(c))
}
