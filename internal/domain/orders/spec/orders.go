package spec

import (
	"food-order-api/internal/model"
	"time"
)

type CreateOrder struct {
	UserID          int
	TotalPrice      int64
	MenuID          int
	PriceMenu       int64
	AdditionalID    int
	AdditionalPrice int64
}

type UpdatePaymentStatus struct {
	OrderID       int
	PaidBy        string
	PaidAt        time.Time
	PaymentStatus model.PaymentStatus
}
