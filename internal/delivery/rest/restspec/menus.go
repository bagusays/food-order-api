package restspec

import (
	"time"
)

type FetchMenuCategoriesResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateMenuCategoryRequest struct {
	Name string `json:"name,omitempty"`
}

type UpdateMenuCategoryRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DeleteMenuCategoryRequest struct {
	ID int `json:"id"`
}

type FetchAdditionalsResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAdditionalRequest struct {
	Name  string `json:"name,omitempty"`
	Price int64  `json:"price,omitempty"`
}

type UpdateAdditionalRequest struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int64  `json:"price,omitempty"`
}

type DeleteAdditionalRequest struct {
	ID int `json:"id"`
}

type FetchMenusResponse struct {
	ID                     int                                   `json:"id"`
	Name                   string                                `json:"name"`
	Description            string                                `json:"description"`
	Price                  float64                               `json:"price"`
	CategoryName           string                                `json:"category_name"`
	EligibleAdditionalMenu []FetchEligibleAdditionalMenuResponse `json:"eligible_additional_menu"`
	CreatedAt              time.Time                             `json:"created_at"`
	UpdatedAt              time.Time                             `json:"updated_at"`
}

type CreateMenuRequest struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	MenuCategoryID int     `json:"menu_category_id"`
}

type UpdateMenuRequest struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	MenuCategoryID int     `json:"menu_category_id"`
}

type DeleteMenuRequest struct {
	ID int `json:"id"`
}

type FetchEligibleAdditionalMenuResponse struct {
	ID              int       `json:"id"`
	MenuID          int       `json:"menu_id"`
	AdditionalName  string    `json:"additional_name"`
	AdditionalPrice float64   `json:"additional_price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateEligibleAdditionalMenuRequest struct {
	MenuID       int `json:"menu_id"`
	AdditionalID int `json:"additional_id"`
}

type UpdateEligibleAdditionalMenuRequest struct {
	ID           int `json:"id"`
	MenuID       int `json:"menu_id"`
	AdditionalID int `json:"additional_id"`
}

type DeleteEligibleAdditionalMenuRequest struct {
	ID int `json:"id"`
}
