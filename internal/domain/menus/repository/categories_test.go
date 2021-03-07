package repository

import (
	"context"
	"database/sql"
	"errors"
	"food-order-api/internal/model"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMenus_FetchCategory(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name     string
		mockRows []model.MenuCategories
		mockErr  error
		isError  bool
	}{
		{
			name: "should be succeed",
			mockRows: []model.MenuCategories{
				{
					ID:        1,
					Name:      "name",
					CreatedAt: now,
					UpdatedAt: now,
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

			dbResult := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
			for _, m := range tc.mockRows {
				dbResult.AddRow(m.ID, m.Name, m.CreatedAt, m.UpdatedAt)
			}

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM menu_categories")).WillReturnRows(dbResult).WillReturnError(tc.mockErr)

			repo := New(dbx)
			res, err := repo.FetchMenuCategory(context.Background())
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
				assert.Equal(t, res[0].Name, m.Name)
			}
		})
	}
}

func TestMenus_CreateCategory(t *testing.T) {
	query := "INSERT INTO menu_categories (.+) VALUES (.+)"
	t.Run("should be succeed", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs("name").
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := New(dbx)
		err := repo.CreateMenuCategory(context.Background(), "name")
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
		err := repo.CreateMenuCategory(context.Background(), "name")
		assert.Error(t, err)
	})
}

func TestMenus_UpdateCategory(t *testing.T) {
	query := "UPDATE menu_categories SET (.+) WHERE (.+)"
	t.Run("should be succeed", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs("name", 1000).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := New(dbx)
		err := repo.UpdateMenuCategory(context.Background(), "name", 1000)
		assert.NoError(t, err)
	})

	t.Run("should be failed because return err from repo", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs("name", 1000).
			WillReturnResult(nil).
			WillReturnError(errors.New("unexpected error"))

		repo := New(dbx)
		err := repo.UpdateMenuCategory(context.Background(), "name", 1000)
		assert.Error(t, err)
	})
}

func TestMenus_DeleteCategory(t *testing.T) {
	query := "DELETE FROM menu_categories WHERE (.+)"
	t.Run("should be succeed", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := New(dbx)
		err := repo.DeleteMenuCategory(context.Background(), 1)
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
		err := repo.DeleteMenuCategory(context.Background(), 1)
		assert.Error(t, err)
	})
}
