package orders

import (
	"context"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/model"
)

type Usecase interface {
	FetchAllOrders(ctx context.Context, userID int) ([]restspec.FetchOrders, error)
	FetchOrder(ctx context.Context, userID, orderID int) (*restspec.FetchOrders, error)
}

type Repository interface {
	FetchItemDetails(ctx context.Context, orderDetailID int) ([]model.ItemDetails, error)

	FetchOrderDetails(ctx context.Context, orderID int) ([]model.OrderDetails, error)

	FetchAllOrders(ctx context.Context, userID int) ([]model.Orders, error)
	FetchOrder(ctx context.Context, userID, orderID int) (*model.Orders, error)
}
