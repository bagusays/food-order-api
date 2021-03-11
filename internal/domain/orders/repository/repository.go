package repository

import (
	"food-order-api/internal/domain/orders"

	"github.com/jmoiron/sqlx"
)

type Orders struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) orders.Repository {
	return &Orders{db: db}
}
