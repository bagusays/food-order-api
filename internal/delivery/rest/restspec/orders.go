package restspec

import (
	"food-order-api/internal/model"
	"time"
)

type FetchOrders struct {
	ID            int                 `json:"id"`
	UserID        int                 `json:"user_id"`
	PaymentStatus model.PaymentStatus `json:"payment_status"`
	PaidBy        string              `json:"paid_by"`
	PaidAt        *time.Time          `json:"paid_at"`
	TotalPrice    int64               `json:"total_price"`
	OrderStatus   model.OrderStatus   `json:"order_status"`
	OrderDetails  []OrderDetails      `json:"order_details"`
	CreatedAt     *time.Time          `json:"created_at"`
	UpdatedAt     *time.Time          `json:"updated_at"`
}

type OrderDetails struct {
	ID               int           `json:"id"`
	OrderID          int           `json:"order_id"`
	MenuID           int           `json:"menu_id"`
	PriceMenu        int64         `json:"price_menu"`
	MenuName         string        `json:"menu_name"`
	MenuCategoryName string        `json:"menu_category_name"`
	ItemDetails      []ItemDetails `json:"item_details"`
	CreatedAt        *time.Time    `json:"created_at"`
	UpdatedAt        *time.Time    `json:"updated_at"`
}

type ItemDetails struct {
	ID              int        `json:"id"`
	OrderDetailID   int        `json:"order_detail_id"`
	AdditionalID    int        `json:"additional_id"`
	AdditionalName  string     `json:"additional_name"`
	AdditionalPrice int64      `json:"additional_price"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
