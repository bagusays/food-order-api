package spec

import (
	"food-order-api/internal/shared"
)

type CreateMenuCategory struct {
	Name string
}

func (c *CreateMenuCategory) Validate() error {
	if c.Name == "" {
		return shared.ErrNameCannotBeNil
	}
	return nil
}
