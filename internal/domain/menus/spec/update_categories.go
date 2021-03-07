package spec

import (
	"food-order-api/internal/shared"
)

type UpdateMenuCategories struct {
	ID   int
	Name string
}

func (c *UpdateMenuCategories) Validate() error {
	if c.ID == 0 {
		return shared.ErrIDCannotBeNil
	}
	if c.Name == "" {
		return shared.ErrNameCannotBeNil
	}
	return nil
}
