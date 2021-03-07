package handler

import (
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/healthcheck"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/check")

	usecase := healthcheck.MockUsecase{}
	usecase.On("Check", mock.Anything).Return(restspec.HealthCheckResponse{})

	err := Check(&usecase)(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
