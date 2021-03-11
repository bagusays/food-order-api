package repository

import (
	"context"
	"food-order-api/internal/model"
)

func (o *Orders) FetchItemDetails(ctx context.Context, orderDetailID int) ([]model.ItemDetails, error) {
	query := `
	SELECT 
		ids.id, 
		ids.order_detail_id, 
		ids.additional_id, 
		a.name as additional_name,
		ids.additional_price, 
		ids.created_at, 
		ids.updated_at
	FROM item_details ids
	INNER JOIN additionals a ON a.id = ids.additional_id
	WHERE ids.order_detail_id = ?`

	var res []model.ItemDetails
	err := o.db.SelectContext(ctx, &res, query, orderDetailID)
	if err != nil {
		return nil, checkReadErr(err)
	}
	return res, err
}
