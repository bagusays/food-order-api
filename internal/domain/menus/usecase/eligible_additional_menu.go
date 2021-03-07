package usecase

import (
	"context"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/menus/spec"
	"food-order-api/internal/shared"
)

func (m Menus) FetchEligibleAdditionalMenu(ctx context.Context) ([]restspec.FetchEligibleAdditionalMenuResponse, error) {
	eligibleAdditional, err := m.menusRepo.FetchEligibleAdditionalMenu(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]restspec.FetchEligibleAdditionalMenuResponse, 0)
	for _, d := range eligibleAdditional {
		res = append(res, restspec.FetchEligibleAdditionalMenuResponse{
			ID:              d.ID,
			MenuID:          d.MenuID,
			AdditionalName:  d.AdditionalName,
			AdditionalPrice: d.AdditionalPrice,
			CreatedAt:       d.CreatedAt,
			UpdatedAt:       d.UpdatedAt,
		})
	}

	return res, nil
}

func (m Menus) CreateEligibleAdditionalMenu(ctx context.Context, arg spec.CreateEligibleAdditionalMenu) error {
	if err := arg.Validate(); err != nil {
		return err
	}

	return m.menusRepo.CreateEligibleAdditionalMenu(ctx, arg.MenuID, arg.AdditionalID)
}

func (m Menus) UpdateEligibleAdditionalMenu(ctx context.Context, arg spec.UpdateEligibleAdditionalMenu) error {
	if err := arg.Validate(); err != nil {
		return err
	}

	return m.menusRepo.UpdateEligibleAdditionalMenu(ctx, arg.MenuID, arg.AdditionalID, arg.ID)
}

func (m Menus) DeleteEligibleAdditionalMenu(ctx context.Context, id int) error {
	if id == 0 {
		return shared.ErrIDCannotBeNil
	}

	return m.menusRepo.DeleteEligibleAdditionalMenu(ctx, id)
}
