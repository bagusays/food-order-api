package model

import (
	"time"
)

type Additionals struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	Price     float64    `db:"price"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type MenuCategories struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type Menus struct {
	ID           int        `db:"id"`
	Name         string     `db:"name"`
	Description  string     `db:"description"`
	Price        float64    `db:"price"`
	CategoryName string     `db:"category_name"`
	CreatedAt    *time.Time `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}

type EligibleAdditionalMenu struct {
	ID              int        `db:"id"`
	MenuID          int        `db:"menu_id"`
	AdditionalName  string     `db:"additional_name"`
	AdditionalPrice float64    `db:"additional_price"`
	CreatedAt       *time.Time `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
}
