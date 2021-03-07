package usecase

import (
	"food-order-api/internal/domain/menus"
)

type Menus struct {
	menusRepo menus.Repository
}

func New(menusRepo menus.Repository) menus.Usecase {
	return &Menus{menusRepo: menusRepo}
}
