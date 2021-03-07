package spec

import "food-order-api/internal/shared"

type CreateEligibleAdditionalMenu struct {
	MenuID       int
	AdditionalID int
}

func (c *CreateEligibleAdditionalMenu) Validate() error {
	if c.MenuID == 0 {
		return shared.ErrMenuIDCannotBeNil
	}
	if c.AdditionalID == 0 {
		return shared.ErrAdditionalIDCannotBeNil
	}
	return nil
}
