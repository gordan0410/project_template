package http

import (
	"project_template/internal/domain"
)

// HealthHandler handles HTTP requests for health
type HealthHandler struct {
	healthUsecase domain.HealthUsecase
}

// NewHealthHandler creates a new instance of HealthHandler
func NewHealthHandler(healthUsecase domain.HealthUsecase) *HealthHandler {
	return &HealthHandler{
		healthUsecase: healthUsecase,
	}
}
