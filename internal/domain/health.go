package domain

import "context"

// HealthUsecase interface
type HealthUsecase interface {
	CheckHealth(ctx context.Context) error
}
