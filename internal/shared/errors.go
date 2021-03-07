package shared

import "errors"

var (
	ErrIDCannotBeNil             = errors.New("id cannot be nil/zero")
	ErrNameCannotBeNil           = errors.New("name cannot be nil/empty")
	ErrPriceCannotBeNil          = errors.New("price cannot be nil/zero")
	ErrDescriptionCannotBeNil    = errors.New("description cannot be nil/zero")
	ErrMenuCategoryIDCannotBeNil = errors.New("menu_category_id cannot be nil/zero")
	ErrMenuIDCannotBeNil         = errors.New("menu_id cannot be nil/zero")
	ErrAdditionalIDCannotBeNil   = errors.New("additional_id cannot be nil/zero")
	ErrMenuCategoryIDNotFound    = errors.New("menu_category_id is not found. please insert before use")
)

var mapErrToStatusCode = map[error]string{
	ErrIDCannotBeNil:             StatusErrIDCannotBeZero,
	ErrNameCannotBeNil:           StatusErrNameCannotBeEmpty,
	ErrPriceCannotBeNil:          StatusErrPriceCannotBeZero,
	ErrDescriptionCannotBeNil:    StatusErrDescriptionCannotBeEmpty,
	ErrMenuCategoryIDCannotBeNil: StatusErrMenuCategoryIDCannotBeZero,
	ErrMenuIDCannotBeNil:         StatusErrMenuIDCannotBeZero,
	ErrAdditionalIDCannotBeNil:   StatusErrAdditionalIDCannotBeZero,
	ErrMenuCategoryIDNotFound:    StatusCategoryIDNotFound,
}

func MapErrToStatusCode(err error) (string, bool) {
	s, ok := mapErrToStatusCode[err]
	return s, ok
}
