package repository

import (
	"context"
	"food-order-api/internal/model"
)

func (o *Orders) FetchOrderDetails(ctx context.Context, orderID int) ([]model.OrderDetails, error) {
	query := `
	SELECT 
		od.id, 
		od.order_id, 
		od.menu_id, 
		m.name as menu_name, 
		mc.name as menu_category_name, 
		od.price_menu,
		od.created_at,
		od.updated_at
	FROM order_details od
	INNER JOIN menus m ON m.id = od.menu_id
	INNER JOIN menu_categories mc ON mc.id = m.menu_category_id
	WHERE od.order_id = ?`

	var res []model.OrderDetails
	err := o.db.SelectContext(ctx, &res, query, orderID)
	if err != nil {
		return nil, checkReadErr(err)
	}
	return res, nil
}
