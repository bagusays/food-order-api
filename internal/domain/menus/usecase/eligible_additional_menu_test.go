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

func TestMenus_FetchEligibleAdditionalMenu(t *testing.T) {
	testCases := []struct {
		name           string
		mockRepoReturn []model.EligibleAdditionalMenu
		mockRepoErr    error
		isError        bool
	}{
		{
			name: "should be succeed",
			mockRepoReturn: []model.EligibleAdditionalMenu{
				{
					ID:              1,
					MenuID:          1,
					AdditionalName:  "additional",
					AdditionalPrice: 1000,
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
			mockRepo.On("FetchEligibleAdditionalMenu", mock.Anything).Return(tc.mockRepoReturn, tc.mockRepoErr)

			usecase := New(&mockRepo)
			res, err := usecase.FetchEligibleAdditionalMenu(ctx)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, res[0].ID, tc.mockRepoReturn[0].ID)
			assert.Equal(t, res[0].MenuID, tc.mockRepoReturn[0].MenuID)
			assert.Equal(t, res[0].AdditionalName, tc.mockRepoReturn[0].AdditionalName)
			assert.Equal(t, res[0].AdditionalPrice, tc.mockRepoReturn[0].AdditionalPrice)
		})
	}
}

func TestMenus_CreateEligibleAdditionalMenu(t *testing.T) {
	testCases := []struct {
		name        string
		arg         spec.CreateEligibleAdditionalMenu
		mockRepoErr error
		isError     bool
	}{
		{
			name: "should be succeed",
			arg: spec.CreateEligibleAdditionalMenu{
				MenuID:       1,
				AdditionalID: 1,
			},
		},
		{
			name:        "should be failed because any error in repository",
			mockRepoErr: errors.New("unexpected error"),
			arg: spec.CreateEligibleAdditionalMenu{
				MenuID:       1,
				AdditionalID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because MenuID is nil",
			arg: spec.CreateEligibleAdditionalMenu{
				MenuID:       0,
				AdditionalID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because AdditionalID is nil",
			arg: spec.CreateEligibleAdditionalMenu{
				MenuID:       1,
				AdditionalID: 0,
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("CreateEligibleAdditionalMenu", mock.Anything, mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.CreateEligibleAdditionalMenu(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMenus_UpdateEligibleAdditionalMenu(t *testing.T) {
	testCases := []struct {
		name        string
		arg         spec.UpdateEligibleAdditionalMenu
		mockRepoErr error
		isError     bool
	}{
		{
			name: "should be succeed",
			arg: spec.UpdateEligibleAdditionalMenu{
				ID:           1,
				MenuID:       1,
				AdditionalID: 1,
			},
		},
		{
			name:        "should be failed because any error in repository",
			mockRepoErr: errors.New("unexpected error"),
			arg: spec.UpdateEligibleAdditionalMenu{
				ID:           1,
				MenuID:       1,
				AdditionalID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because id is nil",
			arg: spec.UpdateEligibleAdditionalMenu{
				ID:           0,
				MenuID:       1,
				AdditionalID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because MenuID is nil",
			arg: spec.UpdateEligibleAdditionalMenu{
				ID:           1,
				MenuID:       0,
				AdditionalID: 1,
			},
			isError: true,
		},
		{
			name: "should be failed because AdditionalID is nil",
			arg: spec.UpdateEligibleAdditionalMenu{
				ID:           1,
				MenuID:       1,
				AdditionalID: 0,
			},
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := menus.MockRepository{}
			mockRepo.On("UpdateEligibleAdditionalMenu", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.UpdateEligibleAdditionalMenu(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestMenus_DeleteEligibleAdditionalMenu(t *testing.T) {
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
			mockRepo.On("DeleteEligibleAdditionalMenu", mock.Anything, mock.Anything).Return(tc.mockRepoErr)

			usecase := New(&mockRepo)
			err := usecase.DeleteEligibleAdditionalMenu(ctx, tc.arg)
			if tc.isError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
