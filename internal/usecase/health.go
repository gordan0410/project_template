package usecase

import (
	"context"
	"project_template/internal/domain"
)

type HealthUsecase struct {
}

func NewHealthUsecase() domain.HealthUsecase {
	return &HealthUsecase{}
}

func (h *HealthUsecase) CheckHealth(ctx context.Context) error {
	return nil
}
