package menus

import (
	"context"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/menus/spec"
	"food-order-api/internal/model"
)

type Usecase interface {
	FetchCategories(ctx context.Context) ([]restspec.FetchMenuCategoriesResponse, error)
	CreateMenuCategory(ctx context.Context, arg spec.CreateMenuCategory) error
	UpdateMenuCategory(ctx context.Context, arg spec.UpdateMenuCategories) error
	DeleteMenuCategory(ctx context.Context, id int) error

	FetchAdditionals(ctx context.Context) ([]restspec.FetchAdditionalsResponse, error)
	CreateAdditional(ctx context.Context, arg spec.CreateAdditionals) error
	UpdateAdditional(ctx context.Context, arg spec.UpdateAdditionals) error
	DeleteAdditional(ctx context.Context, id int) error

	FetchMenus(ctx context.Context) ([]restspec.FetchMenusResponse, error)
	CreateMenu(ctx context.Context, arg spec.CreateMenu) error
	UpdateMenu(ctx context.Context, arg spec.UpdateMenu) error
	DeleteMenu(ctx context.Context, id int) error

	FetchEligibleAdditionalMenu(ctx context.Context) ([]restspec.FetchEligibleAdditionalMenuResponse, error)
	CreateEligibleAdditionalMenu(ctx context.Context, arg spec.CreateEligibleAdditionalMenu) error
	UpdateEligibleAdditionalMenu(ctx context.Context, arg spec.UpdateEligibleAdditionalMenu) error
	DeleteEligibleAdditionalMenu(ctx context.Context, id int) error
}

type Repository interface {
	FetchMenuCategory(ctx context.Context) ([]model.MenuCategories, error)
	CreateMenuCategory(ctx context.Context, name string) error
	UpdateMenuCategory(ctx context.Context, name string, id int) error
	DeleteMenuCategory(ctx context.Context, id int) error

	FetchAdditionals(ctx context.Context) ([]model.Additionals, error)
	CreateAdditional(ctx context.Context, name string, price int64) error
	UpdateAdditional(ctx context.Context, name string, price int64, id int) error
	DeleteAdditional(ctx context.Context, id int) error

	FetchMenus(ctx context.Context) ([]model.Menus, error)
	CreateMenu(ctx context.Context, arg spec.CreateMenu) error
	UpdateMenu(ctx context.Context, arg spec.UpdateMenu) error
	DeleteMenu(ctx context.Context, id int) error

	FetchEligibleAdditionalMenu(ctx context.Context) ([]model.EligibleAdditionalMenu, error)
	CreateEligibleAdditionalMenu(ctx context.Context, menuID, additionalID int) error
	UpdateEligibleAdditionalMenu(ctx context.Context, menuID, additionalID, id int) error
	DeleteEligibleAdditionalMenu(ctx context.Context, id int) error
}
