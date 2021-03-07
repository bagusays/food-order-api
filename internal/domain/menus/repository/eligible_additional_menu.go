package repository

import (
	"context"
	"food-order-api/internal/model"
)

func (m Menus) FetchEligibleAdditionalMenu(ctx context.Context) ([]model.EligibleAdditionalMenu, error) {
	query := `
	SELECT edm.id, edm.menu_id, a.name as additional_name, a.price as additional_price, edm.created_at, edm.updated_at 
	FROM eligible_additional_menu edm
	INNER JOIN additionals a ON a.id = edm.additional_id
	ORDER BY a.name`

	var res []model.EligibleAdditionalMenu
	err := m.db.SelectContext(ctx, &res, query)
	if err != nil {
		return res, checkReadErr(err)
	}
	return res, nil
}

func (m Menus) CreateEligibleAdditionalMenu(ctx context.Context, menuID, additionalID int) error {
	query := "INSERT INTO eligible_additional_menu (menu_id, additional_id) VALUES (?, ?)"
	_, err := m.db.ExecContext(ctx, query, menuID, additionalID)
	if err != nil {
		return checkInsertErr(err)
	}
	return nil
}

func (m Menus) UpdateEligibleAdditionalMenu(ctx context.Context, menuID, additionalID, id int) error {
	query := "UPDATE eligible_additional_menu SET menu_id = ?, additional_id = ?, updated_at = now() WHERE id = ?"
	_, err := m.db.ExecContext(ctx, query, menuID, additionalID, id)
	if err != nil {
		return err
	}
	return nil
}

func (m Menus) DeleteEligibleAdditionalMenu(ctx context.Context, id int) error {
	query := "DELETE FROM eligible_additional_menu WHERE id = ?"
	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
