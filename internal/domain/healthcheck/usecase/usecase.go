package usecase

import (
	"context"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/healthcheck"

	"github.com/jmoiron/sqlx"
)

type HealthCheck struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) healthcheck.Usecase {
	return &HealthCheck{db: db}
}

func (h *HealthCheck) Check(ctx context.Context) restspec.HealthCheckResponse {
	status := "OK"
	message := ""

	if err := h.db.PingContext(ctx); err != nil {
		status = "ERROR"
		message = err.Error()
	}

	return restspec.HealthCheckResponse{
		DB: restspec.HealthCheckDetail{
			Status:  status,
			Message: message,
		},
	}
}
