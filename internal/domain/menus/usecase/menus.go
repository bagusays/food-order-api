package usecase

import (
	"context"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/menus/spec"
	"food-order-api/internal/shared"
)

func (m *Menus) FetchMenus(ctx context.Context) ([]restspec.FetchMenusResponse, error) {
	menus, err := m.menusRepo.FetchMenus(ctx)
	if err != nil {
		return nil, err
	}

	eligibleAdditionalMenu, err := m.menusRepo.FetchEligibleAdditionalMenu(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]restspec.FetchMenusResponse, 0)
	for _, menu := range menus {
		var eligibleAddMenuTmp []restspec.FetchEligibleAdditionalMenuResponse
		for _, edm := range eligibleAdditionalMenu {
			if menu.ID == edm.MenuID {
				eligibleAddMenuTmp = append(eligibleAddMenuTmp, restspec.FetchEligibleAdditionalMenuResponse{
					ID:              edm.ID,
					MenuID:          edm.MenuID,
					AdditionalName:  edm.AdditionalName,
					AdditionalPrice: edm.AdditionalPrice,
					CreatedAt:       edm.CreatedAt,
					UpdatedAt:       edm.UpdatedAt,
				})
			}
		}

		res = append(res, restspec.FetchMenusResponse{
			ID:                     menu.ID,
			Name:                   menu.Name,
			Description:            menu.Description,
			Price:                  menu.Price,
			CategoryName:           menu.CategoryName,
			EligibleAdditionalMenu: eligibleAddMenuTmp,
			CreatedAt:              menu.CreatedAt,
			UpdatedAt:              menu.UpdatedAt,
		})
	}

	return res, nil
}

func (m *Menus) CreateMenu(ctx context.Context, arg spec.CreateMenu) error {
	if err := arg.Validate(); err != nil {
		return err
	}

	return m.menusRepo.CreateMenu(ctx, arg)
}

func (m *Menus) UpdateMenu(ctx context.Context, arg spec.UpdateMenu) error {
	if err := arg.Validate(); err != nil {
		return err
	}

	return m.menusRepo.UpdateMenu(ctx, arg)
}

func (m *Menus) DeleteMenu(ctx context.Context, id int) error {
	if id == 0 {
		return shared.ErrIDCannotBeNil
	}

	return m.menusRepo.DeleteMenu(ctx, id)
}
