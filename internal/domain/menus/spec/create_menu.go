package spec

import "food-order-api/internal/shared"

type CreateMenu struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	MenuCategoryID int     `json:"menu_category_id"`
}

func (c *CreateMenu) Validate() error {
	if c.Name == "" {
		return shared.ErrNameCannotBeNil
	}
	if c.Description == "" {
		return shared.ErrDescriptionCannotBeNil
	}
	if c.Price == 0 {
		return shared.ErrPriceCannotBeNil
	}
	if c.MenuCategoryID == 0 {
		return shared.ErrMenuCategoryIDCannotBeNil
	}

	return nil
}

type UpdateMenu struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	MenuCategoryID int     `json:"menu_category_id"`
}

func (e *UpdateMenu) Validate() error {
	if e.ID == 0 {
		return shared.ErrIDCannotBeNil
	}
	if e.Name == "" {
		return shared.ErrNameCannotBeNil
	}
	if e.Description == "" {
		return shared.ErrDescriptionCannotBeNil
	}
	if e.Price == 0 {
		return shared.ErrPriceCannotBeNil
	}
	if e.MenuCategoryID == 0 {
		return shared.ErrMenuCategoryIDCannotBeNil
	}

	return nil
}
