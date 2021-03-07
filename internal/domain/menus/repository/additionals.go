package repository

import (
	"context"
	"food-order-api/internal/model"
)

func (m Menus) FetchAdditionals(ctx context.Context) ([]model.Additionals, error) {
	query := "SELECT * FROM additionals"
	var res []model.Additionals
	err := m.db.SelectContext(ctx, &res, query)
	if err != nil {
		return res, checkReadErr(err)
	}
	return res, nil
}

func (m Menus) CreateAdditional(ctx context.Context, name string, price int64) error {
	query := "INSERT INTO additionals (name, price) VALUES (?, ?)"
	_, err := m.db.ExecContext(ctx, query, name, price)
	if err != nil {
		return checkInsertErr(err)
	}
	return nil
}

func (m Menus) UpdateAdditional(ctx context.Context, name string, price int64, id int) error {
	query := "UPDATE additionals SET name = ?, price = ?, updated_at = now() WHERE id = ?"
	_, err := m.db.ExecContext(ctx, query, name, price, id)
	if err != nil {
		return err
	}
	return nil
}

func (m Menus) DeleteAdditional(ctx context.Context, id int) error {
	query := "DELETE FROM additionals WHERE id = ?"
	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
