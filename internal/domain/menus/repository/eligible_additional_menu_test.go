package repository

import (
	"context"
	"database/sql"
	"errors"
	"food-order-api/internal/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestMenus_FetchEligibleAdditionalMenu(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name     string
		mockRows []model.EligibleAdditionalMenu
		mockErr  error
		isError  bool
	}{
		{
			name: "should be succeed",
			mockRows: []model.EligibleAdditionalMenu{
				{
					ID:              1,
					MenuID:          1,
					AdditionalName:  "additional",
					AdditionalPrice: 1000,
					CreatedAt:       &now,
					UpdatedAt:       &now,
				},
			},
		},
		{
			name:    "should be succeed if no rows return",
			mockErr: sql.ErrNoRows,
		},
		{
			name:    "should be failed because return err from repo",
			mockErr: errors.New("unexpected error"),
			isError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			dbx := sqlx.NewDb(db, "sqlmock")

			dbResult := sqlmock.NewRows([]string{"id", "additional_name", "additional_price", "created_at", "updated_at"})
			for _, m := range tc.mockRows {
				dbResult.AddRow(m.ID, m.AdditionalName, m.AdditionalPrice, m.CreatedAt, m.UpdatedAt)
			}

			query := `
			SELECT (.+) 
			FROM eligible_additional_menu edm
			INNER JOIN additionals a ON a.id = edm.additional_id
			ORDER BY a.name`

			mock.ExpectQuery(query).WillReturnRows(dbResult).WillReturnError(tc.mockErr)

			repo := New(dbx)
			res, err := repo.FetchEligibleAdditionalMenu(context.Background())
			if tc.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			if len(tc.mockRows) == 0 {
				return
			}

			for _, m := range tc.mockRows {
				assert.Equal(t, res[0].ID, m.ID)
				assert.Equal(t, res[0].AdditionalName, m.AdditionalName)
				assert.Equal(t, res[0].AdditionalPrice, m.AdditionalPrice)
			}
		})
	}
}

func TestMenus_CreateEligibleAdditionalMenu(t *testing.T) {
	query := "INSERT INTO eligible_additional_menu (.+) VALUES (.+)"
	t.Run("should be succeed", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs(1, 1).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := New(dbx)
		err := repo.CreateEligibleAdditionalMenu(context.Background(), 1, 1)
		assert.NoError(t, err)
	})

	t.Run("should be failed because return err from repo", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs("name").
			WillReturnResult(nil).
			WillReturnError(errors.New("unexpected error"))

		repo := New(dbx)
		err := repo.CreateEligibleAdditionalMenu(context.Background(), 1, 1)
		assert.Error(t, err)
	})
}

func TestMenus_UpdateEligibleAdditionalMenu(t *testing.T) {
	query := "UPDATE eligible_additional_menu SET (.+) WHERE (.+)"
	t.Run("should be succeed", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs(1, 1, 1).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := New(dbx)
		err := repo.UpdateEligibleAdditionalMenu(context.Background(), 1, 1, 1)
		assert.NoError(t, err)
	})

	t.Run("should be failed because return err from repo", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs(1, 1000).
			WillReturnResult(nil).
			WillReturnError(errors.New("unexpected error"))

		repo := New(dbx)
		err := repo.UpdateEligibleAdditionalMenu(context.Background(), 1, 1, 1)
		assert.Error(t, err)
	})
}

func TestMenus_DeleteEligibleAdditionalMenu(t *testing.T) {
	query := "DELETE FROM eligible_additional_menu WHERE (.+)"
	t.Run("should be succeed", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := New(dbx)
		err := repo.DeleteEligibleAdditionalMenu(context.Background(), 1)
		assert.NoError(t, err)
	})

	t.Run("should be failed because return err from repo", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs("name", 1000, 1).
			WillReturnResult(nil).
			WillReturnError(errors.New("unexpected error"))

		repo := New(dbx)
		err := repo.DeleteEligibleAdditionalMenu(context.Background(), 1)
		assert.Error(t, err)
	})
}
