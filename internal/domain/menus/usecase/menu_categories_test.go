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

func TestMenus_FetchCategories(t *testing.T) {
	testCases := []struct {
		name           string
		mockRepoReturn []model.MenuCategories
		mockRepoErr    error
		isError        bool
	}{
		{
			name: "should be succeed",
			mockRepoReturn: []model.MenuCategories{
				{
					ID:   1,
					Name: "additionals",
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
			mockRepo.On("FetchMenuCategory", mock.Anything).Return(tc.mockRepoReturn, tc.mockRepoErr)

			usecase := New(&mockRepo)
			res, err := usecase.FetchCategories(ctx)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, res[0].ID, tc.mockRepoReturn[0].ID)
			assert.Equal(t, res[0].Name, tc.mockRepoReturn[0].Name)
		})
	}
}

func TestMenus_CreateAdditional(t *testing.T) {
	testCases := []struct {
		name        string
		arg         spec.CreateAdditionals
		mockRepoErr error
		isError     bool
	}{
		{
			name: "should be succeed",
			arg: spec.CreateAdditionals{
				Name:  "name",
				Price: 1,
			},
		},
		{
			name:        "should be failed because any error in repository",
			mockRepoErr: errors.New("unexpected error"),
			arg: spec.CreateAdditionals{
				Name:  "name",
				Price: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because name is nil",
			arg: spec.CreateAdditionals{
				Name:  "",
				Price: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because price is nil",
			arg: spec.CreateAdditionals{
				Name:  "name",
				Price: 0,
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("CreateAdditional", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.CreateAdditional(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMenus_UpdateAdditional(t *testing.T) {
	testCases := []struct {
		name        string
		arg         spec.UpdateAdditionals
		mockRepoErr error
		isError     bool
	}{
		{
			name: "should be succeed",
			arg: spec.UpdateAdditionals{
				ID:    1,
				Name:  "name",
				Price: 1,
			},
		},
		{
			name:        "should be failed because any error in repository",
			mockRepoErr: errors.New("unexpected error"),
			arg: spec.UpdateAdditionals{
				ID:    1,
				Name:  "name",
				Price: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because id is nil",
			arg: spec.UpdateAdditionals{
				ID:    0,
				Name:  "name",
				Price: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because name is nil",
			arg: spec.UpdateAdditionals{
				ID:    1,
				Name:  "",
				Price: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because price is nil",
			arg: spec.UpdateAdditionals{
				ID:    1,
				Name:  "name",
				Price: 0,
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("UpdateAdditional", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.UpdateAdditional(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMenus_DeleteAdditional(t *testing.T) {
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
			mockRepo.On("DeleteAdditional", mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.DeleteAdditional(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
