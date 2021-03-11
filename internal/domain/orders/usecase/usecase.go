package usecase

import (
	"food-order-api/internal/domain/orders"
)

type Orders struct {
	orderRepo orders.Repository
}

func New(orderRepo orders.Repository) orders.Usecase {
	return &Orders{orderRepo: orderRepo}
}
