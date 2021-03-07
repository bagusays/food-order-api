package usecase

import (
	"context"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/menus/spec"
	"food-order-api/internal/shared"
)

func (m Menus) FetchCategories(ctx context.Context) ([]restspec.FetchMenuCategoriesResponse, error) {
	categories, err := m.menusRepo.FetchMenuCategory(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]restspec.FetchMenuCategoriesResponse, 0)
	for _, d := range categories {
		res = append(res, restspec.FetchMenuCategoriesResponse{
			ID:        d.ID,
			Name:      d.Name,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	return res, nil
}

func (m Menus) CreateMenuCategory(ctx context.Context, arg spec.CreateMenuCategory) error {
	if err := arg.Validate(); err != nil {
		return err
	}

	return m.menusRepo.CreateMenuCategory(ctx, arg.Name)
}

func (m Menus) UpdateMenuCategory(ctx context.Context, arg spec.UpdateMenuCategories) error {
	if err := arg.Validate(); err != nil {
		return err
	}

	return m.menusRepo.UpdateMenuCategory(ctx, arg.Name, arg.ID)
}

func (m Menus) DeleteMenuCategory(ctx context.Context, id int) error {
	if id == 0 {
		return shared.ErrIDCannotBeNil
	}

	return m.menusRepo.DeleteMenuCategory(ctx, id)
}
