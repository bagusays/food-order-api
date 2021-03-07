package middleware

import (
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/shared"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed || err == nil {
		return
	}

	resp := restspec.ErrorResponse{Message: err.Error()}

	statusCode, ok := shared.MapErrToStatusCode(err)
	if ok {
		resp.ErrorCode = statusCode
		_ = c.JSON(http.StatusBadRequest, resp)
		return
	}

	//if he, ok := err.(*echo.HTTPError); ok {
	//	resp.Message = he.Error()
	//	if he.Code == http.StatusBadGateway {
	//		_ = c.NoContent(he.Code)
	//		return
	//	}
	//}

	resp.ErrorCode = "-1"
	_ = c.JSON(http.StatusInternalServerError, resp)
}
