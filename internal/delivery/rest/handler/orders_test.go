package handler

import (
	"encoding/json"
	"errors"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/orders"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchAllOrders(t *testing.T) {
	now := time.Now().UTC()
	testCases := []struct {
		name       string
		paramURL   string
		mockReturn []restspec.FetchOrders
		mockErr    error
		isError    bool
	}{
		{
			name:     "should be succeed",
			paramURL: "1",
			mockReturn: []restspec.FetchOrders{
				{
					ID:            1,
					UserID:        1,
					PaymentStatus: "PENDING",
					PaidBy:        "",
					PaidAt:        nil,
					TotalPrice:    1500,
					OrderStatus:   "WAITING_FOR_PAYMENT",
					OrderDetails: []restspec.OrderDetails{{
						ID:               1,
						OrderID:          1,
						MenuID:           1,
						PriceMenu:        1000,
						MenuName:         "Latte",
						MenuCategoryName: "Signature",
						ItemDetails: []restspec.ItemDetails{{
							ID:              1,
							OrderDetailID:   1,
							AdditionalID:    1,
							AdditionalName:  "Espresso +1",
							AdditionalPrice: 500,
							CreatedAt:       &now,
							UpdatedAt:       &now,
						}},
						CreatedAt: &now,
						UpdatedAt: &now,
					}},
					CreatedAt: &now,
					UpdatedAt: &now,
				},
			},
		},
		{
			name:     "should be failed because return any error from usecase",
			paramURL: "1",
			mockErr:  errors.New("unexpected error"),
			isError:  true,
		},
		{
			name:     "should be failed because parameter is not integer",
			paramURL: "this is not integer",
			mockErr:  errors.New("unexpected error"),
			isError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/orders/:userID")
			c.SetParamNames("userID")
			c.SetParamValues(tc.paramURL)

			usecase := orders.MockUsecase{}
			usecase.On("FetchAllOrders", mock.Anything, 1).Return(tc.mockReturn, tc.mockErr)

			err := FetchAllOrders(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			var resBody []restspec.FetchOrders
			err = json.Unmarshal(rec.Body.Bytes(), &resBody)
			assert.NoError(t, err)

			assert.Equal(t, resBody, tc.mockReturn)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestFetchOrder(t *testing.T) {
	now := time.Now().UTC()
	testCases := []struct {
		name       string
		userID     string
		orderID    string
		mockReturn *restspec.FetchOrders
		mockErr    error
		isError    bool
	}{
		{
			name:    "should be succeed",
			userID:  "1",
			orderID: "1",
			mockReturn: &restspec.FetchOrders{
				ID:            1,
				UserID:        1,
				PaymentStatus: "PENDING",
				PaidBy:        "",
				PaidAt:        nil,
				TotalPrice:    1500,
				OrderStatus:   "WAITING_FOR_PAYMENT",
				OrderDetails: []restspec.OrderDetails{{
					ID:               1,
					OrderID:          1,
					MenuID:           1,
					PriceMenu:        1000,
					MenuName:         "Latte",
					MenuCategoryName: "Signature",
					ItemDetails: []restspec.ItemDetails{{
						ID:              1,
						OrderDetailID:   1,
						AdditionalID:    1,
						AdditionalName:  "Espresso +1",
						AdditionalPrice: 500,
						CreatedAt:       &now,
						UpdatedAt:       &now,
					}},
					CreatedAt: &now,
					UpdatedAt: &now,
				}},
				CreatedAt: &now,
				UpdatedAt: &now,
			},
		},
		{
			name:    "should be failed because return any error from usecase",
			userID:  "1",
			orderID: "1",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
		{
			name:    "should be failed because userID is not integer",
			userID:  "this is not integer",
			orderID: "1",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
		{
			name:    "should be failed because orderID is not integer",
			orderID: "this is not integer",
			userID:  "1",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/orders/:userID/:orderID")
			c.SetParamNames("userID", "orderID")
			c.SetParamValues(tc.userID, tc.orderID)

			usecase := orders.MockUsecase{}
			usecase.On("FetchOrder", mock.Anything, 1, 1).Return(tc.mockReturn, tc.mockErr)

			err := FetchOrder(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			var resBody restspec.FetchOrders
			err = json.Unmarshal(rec.Body.Bytes(), &resBody)
			assert.NoError(t, err)

			assert.Equal(t, &resBody, tc.mockReturn)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}
