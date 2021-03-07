package handler

import (
	"food-order-api/internal/domain/healthcheck"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Check(usecase healthcheck.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, usecase.Check(c.Request().Context()))
	}
}
