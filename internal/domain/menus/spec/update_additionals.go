package spec

import (
	"food-order-api/internal/shared"
)

type UpdateAdditionals struct {
	ID    int
	Name  string
	Price int64
}

func (c *UpdateAdditionals) Validate() error {
	if c.ID == 0 {
		return shared.ErrIDCannotBeNil
	}
	if c.Name == "" {
		return shared.ErrNameCannotBeNil
	}
	if c.Price == 0 {
		return shared.ErrPriceCannotBeNil
	}
	return nil
}
