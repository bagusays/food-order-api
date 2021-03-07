package spec

import "food-order-api/internal/shared"

type UpdateEligibleAdditionalMenu struct {
	ID           int
	MenuID       int
	AdditionalID int
}

func (c *UpdateEligibleAdditionalMenu) Validate() error {
	if c.ID == 0 {
		return shared.ErrIDCannotBeNil
	}
	if c.MenuID == 0 {
		return shared.ErrMenuIDCannotBeNil
	}
	if c.AdditionalID == 0 {
		return shared.ErrAdditionalIDCannotBeNil
	}
	return nil
}
