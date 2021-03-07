package usecase

import (
	"context"
	"errors"
	"food-order-api/internal/domain/menus"
	"food-order-api/internal/domain/menus/spec"
	"food-order-api/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestMenus_FetchMenus(t *testing.T) {
	testCases := []struct {
		name                   string
		mockMenuRepoReturn     []model.Menus
		mockMenuRepoErr        error
		mockEligibleRepoReturn []model.EligibleAdditionalMenu
		mockEligibleRepoErr    error
		isError                bool
	}{
		{
			name: "should be succeed",
			mockMenuRepoReturn: []model.Menus{
				{
					ID:           1,
					Name:         "menu",
					Description:  "menu desc",
					Price:        1000,
					CategoryName: "catName",
				},
			},
			mockEligibleRepoReturn: []model.EligibleAdditionalMenu{
				{
					ID:              1,
					MenuID:          1,
					AdditionalName:  "additional",
					AdditionalPrice: 1000,
				},
				{
					ID:              2,
					MenuID:          2,
					AdditionalName:  "additional2",
					AdditionalPrice: 1000,
				},
			},
		},
		{
			name:                "should be failed because return any error while getting eligible additional menu",
			mockEligibleRepoErr: errors.New("unexpected error"),
			isError:             true,
		},
		{
			name:            "should be failed because return any error while getting menu",
			mockMenuRepoErr: errors.New("unexpected error"),
			isError:         true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("FetchMenus", mock.Anything).Return(tc.mockMenuRepoReturn, tc.mockMenuRepoErr)
			mockRepo.On("FetchEligibleAdditionalMenu", mock.Anything).Return(tc.mockEligibleRepoReturn, tc.mockEligibleRepoErr)

			usecase := New(&mockRepo)
			res, err := usecase.FetchMenus(ctx)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, res[0].ID, tc.mockMenuRepoReturn[0].ID)
			assert.Equal(t, res[0].Name, tc.mockMenuRepoReturn[0].Name)
			assert.Equal(t, res[0].Description, tc.mockMenuRepoReturn[0].Description)
			assert.Equal(t, res[0].Price, tc.mockMenuRepoReturn[0].Price)
			assert.Equal(t, res[0].CategoryName, tc.mockMenuRepoReturn[0].CategoryName)
			assert.Equal(t, 1, len(res[0].EligibleAdditionalMenu))
			assert.Equal(t, res[0].EligibleAdditionalMenu[0].ID, tc.mockEligibleRepoReturn[0].ID)
			assert.Equal(t, res[0].EligibleAdditionalMenu[0].MenuID, tc.mockEligibleRepoReturn[0].MenuID)
			assert.Equal(t, res[0].EligibleAdditionalMenu[0].AdditionalName, tc.mockEligibleRepoReturn[0].AdditionalName)
			assert.Equal(t, res[0].EligibleAdditionalMenu[0].AdditionalPrice, tc.mockEligibleRepoReturn[0].AdditionalPrice)
		})
	}
}

func TestMenus_CreateMenu(t *testing.T) {
	testCases := []struct {
		name        string
		arg         spec.CreateMenu
		mockRepoErr error
		isError     bool
	}{
		{
			name: "should be succeed",
			arg: spec.CreateMenu{
				Name:           "name",
				Description:    "desc",
				Price:          1000,
				MenuCategoryID: 1,
			},
		},
		{
			name:        "should be failed because any error in repository",
			mockRepoErr: errors.New("unexpected error"),
			arg: spec.CreateMenu{
				Name:           "name",
				Description:    "desc",
				Price:          1000,
				MenuCategoryID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because Name is nil",
			arg: spec.CreateMenu{
				Name:           "",
				Description:    "desc",
				Price:          1000,
				MenuCategoryID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because Description is nil",
			arg: spec.CreateMenu{
				Name:           "name",
				Description:    "",
				Price:          1000,
				MenuCategoryID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because Price is nil",
			arg: spec.CreateMenu{
				Name:           "name",
				Description:    "desc",
				Price:          0,
				MenuCategoryID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because MenuCategoryID is nil",
			arg: spec.CreateMenu{
				Name:           "name",
				Description:    "desc",
				Price:          1000,
				MenuCategoryID: 0,
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("CreateMenu", mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.CreateMenu(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMenus_UpdateMenu(t *testing.T) {
	testCases := []struct {
		name        string
		arg         spec.UpdateMenu
		mockRepoErr error
		isError     bool
	}{
		{
			name: "should be succeed",
			arg: spec.UpdateMenu{
				ID:             1,
				Name:           "name",
				Description:    "desc",
				Price:          1000,
				MenuCategoryID: 1,
			},
		},
		{
			name:        "should be failed because any error in repository",
			mockRepoErr: errors.New("unexpected error"),
			arg: spec.UpdateMenu{
				ID:             1,
				Name:           "name",
				Description:    "desc",
				Price:          1000,
				MenuCategoryID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because id is nil",
			arg: spec.UpdateMenu{
				ID:             0,
				Name:           "name",
				Description:    "desc",
				Price:          1000,
				MenuCategoryID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because Name is nil",
			arg: spec.UpdateMenu{
				ID:             1,
				Name:           "",
				Description:    "desc",
				Price:          1000,
				MenuCategoryID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because Description is nil",
			arg: spec.UpdateMenu{
				ID:             1,
				Name:           "name",
				Description:    "",
				Price:          1000,
				MenuCategoryID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because Price is nil",
			arg: spec.UpdateMenu{
				ID:             1,
				Name:           "name",
				Description:    "desc",
				Price:          0,
				MenuCategoryID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because MenuCategoryID is nil",
			arg: spec.UpdateMenu{
				ID:             1,
				Name:           "name",
				Description:    "desc",
				Price:          1000,
				MenuCategoryID: 0,
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("UpdateMenu", mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.UpdateMenu(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMenus_DeleteMenu(t *testing.T) {
	testCases := []struct {
		name        string
		arg         int
		mockRepoErr error
		isError     bool
	}{
		{
			name: "should be succeed",
			arg:  1,
		},
		{
			name:        "should be failed because any error in repository",
			mockRepoErr: errors.New("unexpected error"),
			arg:         1,
			isError:     true,
		},
		{
			name:    "should be failed because id is nil",
			arg:     0,
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("DeleteMenu", mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.DeleteMenu(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
