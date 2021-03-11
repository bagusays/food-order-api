package usecase

import (
	"context"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/model"
)

func (o *Orders) FetchAllOrders(ctx context.Context, userID int) ([]restspec.FetchOrders, error) {
	orders, err := o.orderRepo.FetchAllOrders(ctx, userID)
	if err != nil {
		return nil, err
	}

	var res []restspec.FetchOrders
	for _, order := range orders {
		orderDetails, err := o.mappingOrderDetails(ctx, &order)
		if err != nil {
			return nil, err
		}

		prepareOrder := restspec.FetchOrders{
			ID:            order.ID,
			UserID:        order.UserID,
			PaymentStatus: order.PaymentStatus,
			PaidBy:        order.PaidBy,
			PaidAt:        order.PaidAt,
			TotalPrice:    order.TotalPrice,
			OrderStatus:   order.OrderStatus,
			OrderDetails:  orderDetails,
			CreatedAt:     order.CreatedAt,
			UpdatedAt:     order.UpdatedAt,
		}

		res = append(res, prepareOrder)
	}

	return res, nil
}

func (o *Orders) FetchOrder(ctx context.Context, userID, orderID int) (*restspec.FetchOrders, error) {
	order, err := o.orderRepo.FetchOrder(ctx, userID, orderID)
	if err != nil {
		return nil, err
	}

	orderDetails, err := o.mappingOrderDetails(ctx, order)
	if err != nil {
		return nil, err
	}

	resOrder := &restspec.FetchOrders{
		ID:            order.ID,
		UserID:        order.UserID,
		PaymentStatus: order.PaymentStatus,
		PaidBy:        order.PaidBy,
		PaidAt:        order.PaidAt,
		TotalPrice:    order.TotalPrice,
		OrderStatus:   order.OrderStatus,
		OrderDetails:  orderDetails,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}

	return resOrder, nil
}

func (o *Orders) mappingOrderDetails(ctx context.Context, order *model.Orders) ([]restspec.OrderDetails, error) {
	orderDetails, err := o.orderRepo.FetchOrderDetails(ctx, order.ID)
	if err != nil {
		return nil, err
	}

	var orderDetailResult []restspec.OrderDetails
	for _, orderDetail := range orderDetails {
		itemDetails, err := o.mappingItemDetails(ctx, &orderDetail)
		if err != nil {
			return nil, err
		}

		prepareOrderDetail := restspec.OrderDetails{
			ID:               orderDetail.ID,
			OrderID:          orderDetail.OrderID,
			MenuID:           orderDetail.MenuID,
			MenuName:         orderDetail.MenuName,
			MenuCategoryName: orderDetail.MenuCategoryName,
			PriceMenu:        orderDetail.PriceMenu,
			ItemDetails:      itemDetails,
			CreatedAt:        orderDetail.CreatedAt,
			UpdatedAt:        orderDetail.UpdatedAt,
		}

		orderDetailResult = append(orderDetailResult, prepareOrderDetail)
	}

	return orderDetailResult, nil
}

func (o *Orders) mappingItemDetails(ctx context.Context, orderDetail *model.OrderDetails) ([]restspec.ItemDetails, error) {
	itemDetails, err := o.orderRepo.FetchItemDetails(ctx, orderDetail.ID)
	if err != nil {
		return nil, err
	}

	var itemDetailsResult []restspec.ItemDetails
	for _, itemDetail := range itemDetails {
		itemDetailsResult = append(itemDetailsResult, restspec.ItemDetails{
			ID:              itemDetail.ID,
			OrderDetailID:   itemDetail.OrderDetailID,
			AdditionalID:    itemDetail.AdditionalID,
			AdditionalName:  itemDetail.AdditionalName,
			AdditionalPrice: itemDetail.AdditionalPrice,
			CreatedAt:       itemDetail.CreatedAt,
			UpdatedAt:       itemDetail.UpdatedAt,
		})
	}
	return itemDetailsResult, nil
}
