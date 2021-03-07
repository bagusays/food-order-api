package middleware

import (
	"encoding/json"
	"errors"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHTTPErrorHandler(t *testing.T) {
	server := echo.New()
	defer func() { _ = server.Close() }()

	table := []struct {
		name string
		err  error
		msg  string
	}{
		{
			"id cannot be nil",
			shared.ErrIDCannotBeNil,
			shared.StatusErrIDCannotBeZero,
		},
		{
			"unexpected error",
			errors.New("unexpected error"),
			"-1",
		},
	}

	for _, tab := range table {
		t.Run(tab.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/dev/api/twallet/v1/users", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := server.NewContext(req, rec)

			HTTPErrorHandler(tab.err, c)

			var res restspec.ErrorResponse

			err := json.NewDecoder(rec.Body).Decode(&res)
			assert.NoError(t, err)
			assert.Equal(t, res.ErrorCode, tab.msg, "statusCode error")
			assert.Equal(t, res.Message, tab.err.Error(), "message error")
		})
	}
}
