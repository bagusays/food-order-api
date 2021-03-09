package repository

import (
	"context"
	"database/sql"
	"food-order-api/internal/domain/menus/spec"
	"food-order-api/internal/model"
	"food-order-api/internal/shared"

	"github.com/go-sql-driver/mysql"
)

func (m Menus) FetchMenus(ctx context.Context) ([]model.Menus, error) {
	query := `
	SELECT m.id, m.name, m.description, m.price, mc.name as category_name, m.created_at, m.updated_at 
	FROM menus m
	INNER JOIN menu_categories mc ON mc.id = m.menu_category_id
	WHERE m.deleted_at is null
	ORDER BY mc.name`

	var res []model.Menus
	err := m.db.SelectContext(ctx, &res, query)
	if err != nil {
		return res, checkReadErr(err)
	}
	return res, nil
}

func (m Menus) CreateMenu(ctx context.Context, arg spec.CreateMenu) error {
	query := "INSERT INTO menus (name, description, price, menu_category_id) VALUES (?, ?, ?, ?)"
	_, err := m.db.ExecContext(ctx, query, arg.Name, arg.Description, arg.Price, arg.MenuCategoryID)
	if err != nil {
		return checkInsertErr(err)
	}
	return nil
}

func (m Menus) UpdateMenu(ctx context.Context, arg spec.UpdateMenu) error {
	query := "UPDATE menus SET name = ?, description = ?, price = ?, menu_category_id = ?, updated_at = now() WHERE id = ?"
	_, err := m.db.ExecContext(ctx, query, arg.Name, arg.Description, arg.Price, arg.MenuCategoryID, arg.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m Menus) DeleteMenu(ctx context.Context, id int) error {
	query := "UPDATE menus SET deleted_at = now() WHERE id = ?"
	_, err := m.db.ExecContext(ctx, query, id)
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

func checkInsertErr(err error) error {
	switch v := err.(type) {
	case *mysql.MySQLError:
		switch v.Number {
		case 1452:
			return shared.ErrMenuCategoryIDNotFound
		default:
			return err
		}
	default:
		return err
	}
}
