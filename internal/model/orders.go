package model

import "time"

type PaymentStatus string
type OrderStatus string

const (
	Paid    PaymentStatus = "PAID"
	Pending               = "PENDING"
)

const (
	Completed         OrderStatus = "COMPLETED"
	WaitingForPayment             = "WAITING_FOR_PAYMENT"
	Canceled                      = "CANCELED"
)

type Orders struct {
	ID            int           `db:"id"`
	UserID        int           `db:"user_id"`
	PaymentStatus PaymentStatus `db:"payment_status"`
	PaidBy        string        `db:"paid_by"`
	PaidAt        *time.Time    `db:"paid_at"`
	TotalPrice    int64         `db:"total_price"`
	OrderStatus   OrderStatus   `db:"order_status"`
	CreatedAt     *time.Time    `db:"created_at"`
	UpdatedAt     *time.Time    `db:"updated_at"`
}

type OrderDetails struct {
	ID               int        `db:"id"`
	OrderID          int        `db:"order_id"`
	MenuID           int        `db:"menu_id"`
	MenuName         string     `db:"menu_name"`
	MenuCategoryName string     `db:"menu_category_name"`
	PriceMenu        int64      `db:"price_menu"`
	CreatedAt        *time.Time `db:"created_at"`
	UpdatedAt        *time.Time `db:"updated_at"`
}

type ItemDetails struct {
	ID              int        `db:"id"`
	OrderDetailID   int        `db:"order_detail_id"`
	AdditionalID    int        `db:"additional_id"`
	AdditionalName  string     `db:"additional_name"`
	AdditionalPrice int64      `db:"additional_price"`
	CreatedAt       *time.Time `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
}
