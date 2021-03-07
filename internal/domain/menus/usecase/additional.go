package usecase

import (
	"context"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/menus/spec"
	"food-order-api/internal/shared"
)

func (m Menus) FetchAdditionals(ctx context.Context) ([]restspec.FetchAdditionalsResponse, error) {
	categories, err := m.menusRepo.FetchAdditionals(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]restspec.FetchAdditionalsResponse, 0)
	for _, d := range categories {
		res = append(res, restspec.FetchAdditionalsResponse{
			ID:        d.ID,
			Name:      d.Name,
			Price:     d.Price,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	return res, nil
}

func (m Menus) CreateAdditional(ctx context.Context, arg spec.CreateAdditionals) error {
	if err := arg.Validate(); err != nil {
		return err
	}

	return m.menusRepo.CreateAdditional(ctx, arg.Name, arg.Price)
}

func (m Menus) UpdateAdditional(ctx context.Context, arg spec.UpdateAdditionals) error {
	if err := arg.Validate(); err != nil {
		return err
	}

	return m.menusRepo.UpdateAdditional(ctx, arg.Name, arg.Price, arg.ID)
}

func (m Menus) DeleteAdditional(ctx context.Context, id int) error {
	if id == 0 {
		return shared.ErrIDCannotBeNil
	}

	return m.menusRepo.DeleteAdditional(ctx, id)
}
