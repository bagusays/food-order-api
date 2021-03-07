package spec

import (
	"food-order-api/internal/shared"
)

type CreateAdditionals struct {
	Name  string
	Price int64
}

func (c *CreateAdditionals) Validate() error {
	if c.Name == "" {
		return shared.ErrNameCannotBeNil
	}
	if c.Price == 0 {
		return shared.ErrPriceCannotBeNil
	}
	return nil
}
