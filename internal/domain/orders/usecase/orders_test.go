package usecase

import (
	"context"
	"errors"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/orders"
	"food-order-api/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrders_FetchAllOrders(t *testing.T) {
	testCases := []struct {
		name                  string
		mockErrFetchAllOrders error
		mockErrOrderDetails   error
		mockErrItemDetails    error
		isError               bool
	}{
		{
			name: "should be succeed",
		},
		{
			name:                  "should be error because return any error from fetchAllOrder",
			mockErrFetchAllOrders: errors.New("unexpected error"),
			isError:               true,
		},
		{
			name:                "should be error because return any error from orderDetails",
			mockErrOrderDetails: errors.New("unexpected error"),
			isError:             true,
		},
		{
			name:               "should be error because return any error from itemDetails",
			mockErrItemDetails: errors.New("unexpected error"),
			isError:            true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			ordersData := []model.Orders{{
				ID:            1,
				UserID:        1,
				PaymentStatus: model.Paid,
				PaidBy:        "LINKAJA",
				PaidAt:        &now,
				TotalPrice:    1500,
				OrderStatus:   model.Completed,
				CreatedAt:     &now,
				UpdatedAt:     &now,
			}}
			orderDetails := []model.OrderDetails{{
				ID:               1,
				OrderID:          1,
				MenuID:           1,
				PriceMenu:        1000,
				MenuName:         "Latte",
				MenuCategoryName: "Signature",
				CreatedAt:        &now,
				UpdatedAt:        &now,
			}}
			itemDetails := []model.ItemDetails{{
				ID:              1,
				OrderDetailID:   1,
				AdditionalID:    1,
				AdditionalName:  "Espresso +1",
				AdditionalPrice: 500,
				CreatedAt:       &now,
				UpdatedAt:       &now,
			}}
			expectedResult := []restspec.FetchOrders{{
				ID:            1,
				UserID:        1,
				PaymentStatus: model.Paid,
				PaidBy:        "LINKAJA",
				PaidAt:        &now,
				TotalPrice:    1500,
				OrderStatus:   model.Completed,
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
			}}

			ctx := context.Background()
			mockRepo := orders.MockRepository{}
			mockRepo.On("FetchAllOrders", mock.Anything, mock.Anything).Return(ordersData, tc.mockErrFetchAllOrders)
			mockRepo.On("FetchOrderDetails", mock.Anything, mock.Anything).Return(orderDetails, tc.mockErrOrderDetails)
			mockRepo.On("FetchItemDetails", mock.Anything, mock.Anything).Return(itemDetails, tc.mockErrItemDetails)

			usecase := New(&mockRepo)
			res, err := usecase.FetchAllOrders(ctx, 1)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, expectedResult, res)
		})
	}
}
