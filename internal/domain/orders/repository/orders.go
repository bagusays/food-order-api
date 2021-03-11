package repository

import (
	"context"
	"database/sql"
	"food-order-api/internal/domain/orders/spec"
	"food-order-api/internal/model"
)

func (o *Orders) FetchAllOrders(ctx context.Context, userID int) ([]model.Orders, error) {
	var res []model.Orders
	err := o.db.SelectContext(ctx, &res, "SELECT * FROM orders WHERE user_id = ?", userID)
	if err != nil {
		return nil, checkReadErr(err)
	}
	return res, nil
}

func (o *Orders) FetchOrder(ctx context.Context, userID, orderID int) (*model.Orders, error) {
	var res model.Orders
	err := o.db.GetContext(ctx, &res, "SELECT * FROM orders WHERE user_id = ? AND id = ?", userID, orderID)
	if err != nil {
		return nil, checkReadErr(err)
	}
	return &res, nil
}

func (o *Orders) CreateOrder(ctx context.Context, arg spec.CreateOrder) error {
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, "INSERT INTO orders (user_id, payment_status, total_price, order_status) VALUES (?, ?, ?, ?)", arg.UserID, model.Pending, arg.TotalPrice, model.WaitingForPayment)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	orderID, _ := res.LastInsertId()
	res, err = tx.ExecContext(ctx, "INSERT INTO order_details (order_id, menu_id, price_menu) VALUES (?, ?, ?)", orderID, arg.MenuID, arg.PriceMenu)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	orderDetailID, _ := res.LastInsertId()
	res, err = tx.ExecContext(ctx, "INSERT INTO item_details (order_detail_id, additional_id, additional_price) VALUES (?, ?, ?)", orderDetailID, arg.AdditionalID, arg.AdditionalPrice)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()
	return nil
}

func (o *Orders) UpdatePaymentStatus(ctx context.Context, arg spec.UpdatePaymentStatus) error {
	_, err := o.db.ExecContext(ctx, "UPDATE orders SET payment_status = ?, paid_by = ?, paid_at = ?, updated_at = now() WHERE id = ?", arg.PaymentStatus, arg.PaidBy, arg.PaidAt, arg.OrderID)
	if err != nil {
		return err
	}
	return nil
}

func (o *Orders) UpdateOrderStatus(ctx context.Context, status model.OrderStatus, orderID int) error {
	_, err := o.db.ExecContext(ctx, "UPDATE orders SET order_status = ?, updated_at = now() WHERE id = ?", status, orderID)
	if err != nil {
		return err
	}
	return nil
}

func checkReadErr(err error) error {
	switch err {
	case sql.ErrNoRows:
		return nil
	default:
		return err
	}
}
