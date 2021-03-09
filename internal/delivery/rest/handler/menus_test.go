package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"food-order-api/internal/delivery/rest/restspec"
	"food-order-api/internal/domain/menus"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchCategories(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name       string
		mockReturn []restspec.FetchMenuCategoriesResponse
		mockErr    error
		isError    bool
	}{
		{
			name: "should be succeed",
			mockReturn: []restspec.FetchMenuCategoriesResponse{
				{
					ID:        1,
					Name:      "test",
					CreatedAt: &now,
					UpdatedAt: &now,
				},
			},
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("FetchCategories", mock.Anything).Return(tc.mockReturn, tc.mockErr)

			err := FetchCategories(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			var resBody []restspec.FetchMenuCategoriesResponse
			err = json.Unmarshal(rec.Body.Bytes(), &resBody)
			assert.NoError(t, err)

			assert.Greater(t, len(resBody), 0)
			assert.Equal(t, resBody[0].ID, tc.mockReturn[0].ID)
			assert.Equal(t, resBody[0].Name, tc.mockReturn[0].Name)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestCreateCategory(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("CreateMenuCategory", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := CreateCategory(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, rec.Code)
		})
	}
}

func TestEditCategory(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPatch, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("UpdateMenuCategory", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := UpdateCategory(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("DeleteMenuCategory", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := DeleteCategory(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestFetchAdditionals(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name       string
		mockReturn []restspec.FetchAdditionalsResponse
		mockErr    error
		isError    bool
	}{
		{
			name: "should be succeed",
			mockReturn: []restspec.FetchAdditionalsResponse{
				{
					ID:        1,
					Name:      "test",
					Price:     1000,
					CreatedAt: &now,
					UpdatedAt: &now,
				},
			},
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("FetchAdditionals", mock.Anything).Return(tc.mockReturn, tc.mockErr)

			err := FetchAdditionals(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			var resBody []restspec.FetchAdditionalsResponse
			err = json.Unmarshal(rec.Body.Bytes(), &resBody)
			assert.NoError(t, err)

			assert.Greater(t, len(resBody), 0)
			assert.Equal(t, resBody[0].ID, tc.mockReturn[0].ID)
			assert.Equal(t, resBody[0].Name, tc.mockReturn[0].Name)
			assert.Equal(t, resBody[0].Price, tc.mockReturn[0].Price)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestCreateAdditional(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("CreateAdditional", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := CreateAdditional(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, rec.Code)
		})
	}
}

func TestUpdateAdditional(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPatch, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("UpdateAdditional", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := UpdateAdditional(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestDeleteAdditional(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("DeleteAdditional", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := DeleteAdditional(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestFetchMenus(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name       string
		mockReturn []restspec.FetchMenusResponse
		mockErr    error
		isError    bool
	}{
		{
			name: "should be succeed",
			mockReturn: []restspec.FetchMenusResponse{
				{
					ID:           1,
					Name:         "test",
					Description:  "description",
					Price:        1000,
					CategoryName: "signature",
					CreatedAt:    &now,
					UpdatedAt:    &now,
				},
			},
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("FetchMenus", mock.Anything).Return(tc.mockReturn, tc.mockErr)

			err := FetchMenus(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			var resBody []restspec.FetchMenusResponse
			err = json.Unmarshal(rec.Body.Bytes(), &resBody)
			assert.NoError(t, err)

			assert.Greater(t, len(resBody), 0)
			assert.Equal(t, resBody[0].ID, tc.mockReturn[0].ID)
			assert.Equal(t, resBody[0].Name, tc.mockReturn[0].Name)
			assert.Equal(t, resBody[0].Description, tc.mockReturn[0].Description)
			assert.Equal(t, resBody[0].Price, tc.mockReturn[0].Price)
			assert.Equal(t, resBody[0].CategoryName, tc.mockReturn[0].CategoryName)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestCreateMenu(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("CreateMenu", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := CreateMenu(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, rec.Code)
		})
	}
}

func TestUpdateMenu(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPatch, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("UpdateMenu", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := UpdateMenu(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestDeleteMenu(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("DeleteMenu", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := DeleteMenu(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestFetchEligibleAdditionalMenu(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name       string
		mockReturn []restspec.FetchEligibleAdditionalMenuResponse
		mockErr    error
		isError    bool
	}{
		{
			name: "should be succeed",
			mockReturn: []restspec.FetchEligibleAdditionalMenuResponse{
				{
					ID:              1,
					MenuID:          1,
					AdditionalName:  "additional",
					AdditionalPrice: 100,
					CreatedAt:       &now,
					UpdatedAt:       &now,
				},
			},
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("FetchEligibleAdditionalMenu", mock.Anything).Return(tc.mockReturn, tc.mockErr)

			err := FetchEligibleAdditionalMenu(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			var resBody []restspec.FetchEligibleAdditionalMenuResponse
			err = json.Unmarshal(rec.Body.Bytes(), &resBody)
			assert.NoError(t, err)

			assert.Greater(t, len(resBody), 0)
			assert.Equal(t, resBody[0].ID, tc.mockReturn[0].ID)
			assert.Equal(t, resBody[0].MenuID, tc.mockReturn[0].MenuID)
			assert.Equal(t, resBody[0].AdditionalName, tc.mockReturn[0].AdditionalName)
			assert.Equal(t, resBody[0].AdditionalPrice, tc.mockReturn[0].AdditionalPrice)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestCreateEligibleAdditionalMenu(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("CreateEligibleAdditionalMenu", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := CreateEligibleAdditionalMenu(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, rec.Code)
		})
	}
}

func TestUpdateEligibleAdditionalMenu(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPatch, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("UpdateEligibleAdditionalMenu", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := UpdateEligibleAdditionalMenu(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func TestDeleteEligibleAdditionalMenu(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		mockErr error
		isError bool
	}{
		{
			name:    "should be succeed",
			payload: `{}`,
		},
		{
			name:    "should be failed while receive unexpected response",
			payload: "unexpected payload",
			isError: true,
		},
		{
			name:    "should be failed",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer([]byte(tc.payload)))
			req.Header = http.Header{echo.HeaderContentType: []string{echo.MIMEApplicationJSON}}
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			usecase := menus.MockUsecase{}
			usecase.On("DeleteEligibleAdditionalMenu", mock.Anything, mock.Anything).Return(tc.mockErr)

			err := DeleteEligibleAdditionalMenu(&usecase)(c)
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}
