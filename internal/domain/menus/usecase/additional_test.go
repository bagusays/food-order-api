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

func TestMenus_FetchAdditionals(t *testing.T) {
	testCases := []struct {
		name           string
		mockRepoReturn []model.Additionals
		mockRepoErr    error
		isError        bool
	}{
		{
			name: "should be succeed",
			mockRepoReturn: []model.Additionals{
				{
					ID:    1,
					Name:  "additionals",
					Price: 1000,
				},
			},
		},
		{
			name:        "should be failed",
			mockRepoErr: errors.New("unexpected error"),
			isError:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("FetchAdditionals", mock.Anything).Return(tc.mockRepoReturn, tc.mockRepoErr)

			usecase := New(&mockRepo)
			res, err := usecase.FetchAdditionals(ctx)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, res[0].ID, tc.mockRepoReturn[0].ID)
			assert.Equal(t, res[0].Name, tc.mockRepoReturn[0].Name)
			assert.Equal(t, res[0].Price, tc.mockRepoReturn[0].Price)
		})
	}
}

func TestMenus_CreateCategory(t *testing.T) {
	testCases := []struct {
		name        string
		arg         spec.CreateMenuCategory
		mockRepoErr error
		isError     bool
	}{
		{
			name: "should be succeed",
			arg: spec.CreateMenuCategory{
				Name: "name",
			},
		},
		{
			name:        "should be failed because any error in repository",
			mockRepoErr: errors.New("unexpected error"),
			arg: spec.CreateMenuCategory{
				Name: "name",
			},
			isError: true,
		},
		{
			name: "should be failed because name is nil",
			arg: spec.CreateMenuCategory{
				Name: "",
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("CreateMenuCategory", mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.CreateMenuCategory(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMenus_UpdateCategory(t *testing.T) {
	testCases := []struct {
		name        string
		arg         spec.UpdateMenuCategories
		mockRepoErr error
		isError     bool
	}{
		{
			name: "should be succeed",
			arg: spec.UpdateMenuCategories{
				ID:   1,
				Name: "name",
			},
		},
		{
			name:        "should be failed because any error in repository",
			mockRepoErr: errors.New("unexpected error"),
			arg: spec.UpdateMenuCategories{
				ID:   1,
				Name: "name",
			},
			isError: true,
		},
		{
			name: "should be failed because id is nil",
			arg: spec.UpdateMenuCategories{
				ID:   0,
				Name: "name",
			},
			isError: true,
		},
		{
			name: "should be failed because name is nil",
			arg: spec.UpdateMenuCategories{
				ID:   1,
				Name: "",
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("UpdateMenuCategory", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.UpdateMenuCategory(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMenus_DeleteCategory(t *testing.T) {
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
			mockRepo.On("DeleteMenuCategory", mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.DeleteMenuCategory(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
