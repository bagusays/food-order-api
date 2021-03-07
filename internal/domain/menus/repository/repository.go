package repository

import (
	"food-order-api/internal/domain/menus"

	"github.com/jmoiron/sqlx"
)

type Menus struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) menus.Repository {
	return &Menus{db: db}
}
