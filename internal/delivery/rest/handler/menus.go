package handler

import (
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/menus"
	"food-order-api/internal/domain/menus/spec"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchCategories(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := usecase.FetchCategories(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}
func CreateCategory(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.CreateMenuCategoryRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		arg := spec.CreateMenuCategory{Name: reqBody.Name}
		err = usecase.CreateMenuCategory(c.Request().Context(), arg)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}
func UpdateCategory(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.UpdateMenuCategoryRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		arg := spec.UpdateMenuCategories{
			ID:   reqBody.ID,
			Name: reqBody.Name,
		}
		err = usecase.UpdateMenuCategory(c.Request().Context(), arg)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
func DeleteCategory(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.DeleteMenuCategoryRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		err = usecase.DeleteMenuCategory(c.Request().Context(), reqBody.ID)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func FetchAdditionals(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := usecase.FetchAdditionals(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}
func CreateAdditional(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.CreateAdditionalRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		arg := spec.CreateAdditionals{Name: reqBody.Name, Price: reqBody.Price}
		err = usecase.CreateAdditional(c.Request().Context(), arg)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}
func UpdateAdditional(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.UpdateAdditionalRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		arg := spec.UpdateAdditionals{
			ID:    reqBody.ID,
			Name:  reqBody.Name,
			Price: reqBody.Price,
		}
		err = usecase.UpdateAdditional(c.Request().Context(), arg)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
func DeleteAdditional(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.DeleteAdditionalRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		err = usecase.DeleteAdditional(c.Request().Context(), reqBody.ID)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func FetchMenus(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := usecase.FetchMenus(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}

func CreateMenu(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.CreateMenuRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		err = usecase.CreateMenu(c.Request().Context(), spec.CreateMenu{
			Name:           reqBody.Name,
			Description:    reqBody.Description,
			Price:          reqBody.Price,
			MenuCategoryID: reqBody.MenuCategoryID,
		})
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}

func UpdateMenu(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.UpdateMenuRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		err = usecase.UpdateMenu(c.Request().Context(), spec.UpdateMenu{
			ID:             reqBody.ID,
			Name:           reqBody.Name,
			Description:    reqBody.Description,
			Price:          reqBody.Price,
			MenuCategoryID: reqBody.MenuCategoryID,
		})
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func DeleteMenu(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.DeleteMenuRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		err = usecase.DeleteMenu(c.Request().Context(), reqBody.ID)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func FetchEligibleAdditionalMenu(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := usecase.FetchEligibleAdditionalMenu(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}

func CreateEligibleAdditionalMenu(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.CreateEligibleAdditionalMenuRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		err = usecase.CreateEligibleAdditionalMenu(c.Request().Context(), spec.CreateEligibleAdditionalMenu{
			MenuID:       reqBody.MenuID,
			AdditionalID: reqBody.AdditionalID,
		})
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}

func UpdateEligibleAdditionalMenu(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.UpdateEligibleAdditionalMenuRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		err = usecase.UpdateEligibleAdditionalMenu(c.Request().Context(), spec.UpdateEligibleAdditionalMenu{
			ID:           reqBody.ID,
			MenuID:       reqBody.MenuID,
			AdditionalID: reqBody.AdditionalID,
		})
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func DeleteEligibleAdditionalMenu(usecase menus.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody restspec.DeleteEligibleAdditionalMenuRequest
		err := c.Bind(&reqBody)
		if err != nil {
			return err
		}

		err = usecase.DeleteEligibleAdditionalMenu(c.Request().Context(), reqBody.ID)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
