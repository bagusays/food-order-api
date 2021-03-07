package delivery

import (
	"context"
	"food-order-api/internal/domain/healthcheck"
	healthCheckUsecase "food-order-api/internal/domain/healthcheck/usecase"
	"food-order-api/internal/domain/menus"
	menusRepository "food-order-api/internal/domain/menus/repository"
	menusUsecase "food-order-api/internal/domain/menus/usecase"
	"food-order-api/internal/shared/config"
	"food-order-api/internal/shared/infrastructure"

	"github.com/labstack/echo/v4"
)

type Container struct {
	EchoServer         *echo.Echo
	HealthCheckUsecase healthcheck.Usecase
	MenusUsecase       menus.Usecase
}

func NewContainer(ctx context.Context, e *echo.Echo) *Container {
	cfg := config.New("./config")

	db, err := infrastructure.NewMySQL(ctx, cfg.MySQL)
	if err != nil {
		panic(err.Error())
	}

	healthCheckUcase := healthCheckUsecase.New(db)

	menusRepo := menusRepository.New(db)
	menusUcase := menusUsecase.New(menusRepo)

	return &Container{
		EchoServer:         e,
		HealthCheckUsecase: healthCheckUcase,
		MenusUsecase:       menusUcase,
	}
}
