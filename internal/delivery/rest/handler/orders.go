package handler

import (
	"food-order-api/internal/domain/orders"
	"food-order-api/internal/shared"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllOrders(usecase orders.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("userID"))
		if err != nil {
			return shared.ErrIntegerFormat
		}

		res, err := usecase.FetchAllOrders(c.Request().Context(), userID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}
