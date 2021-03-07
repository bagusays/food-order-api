package healthcheck

import (
	"context"
	"food-order-api/internal/delivery/rest/restspec"
)

type Usecase interface {
	Check(ctx context.Context) restspec.HealthCheckResponse
}
