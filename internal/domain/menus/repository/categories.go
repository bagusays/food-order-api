package repository

import (
	"context"
	"food-order-api/internal/model"
)

func (m Menus) FetchMenuCategory(ctx context.Context) ([]model.MenuCategories, error) {
	query := "SELECT * FROM menu_categories"
	var res []model.MenuCategories
	err := m.db.SelectContext(ctx, &res, query)
	if err != nil {
		return res, checkReadErr(err)
	}
	return res, nil
}

func (m Menus) CreateMenuCategory(ctx context.Context, name string) error {
	query := "INSERT INTO menu_categories (name) VALUES (?)"
	_, err := m.db.ExecContext(ctx, query, name)
	if err != nil {
		return checkInsertErr(err)
	}
	return nil
}

func (m Menus) UpdateMenuCategory(ctx context.Context, name string, id int) error {
	query := "UPDATE menu_categories SET name = ?, updated_at = now() WHERE id = ?"
	_, err := m.db.ExecContext(ctx, query, name, id)
	if err != nil {
		return err
	}
	return nil
}

func (m Menus) DeleteMenuCategory(ctx context.Context, id int) error {
	query := "DELETE FROM menu_categories WHERE id = ?"
	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
