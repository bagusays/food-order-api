package repository

import (
	"context"
	"database/sql"
	"errors"
	"food-order-api/internal/domain/menus/spec"
	"food-order-api/internal/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestMenus_FetchMenus(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name     string
		mockRows []model.Menus
		mockErr  error
		isError  bool
	}{
		{
			name: "should be succeed",
			mockRows: []model.Menus{
				{
					ID:           1,
					Name:         "menu",
					Description:  "desc",
					Price:        1000,
					CategoryName: "catname",
					CreatedAt:    &now,
					UpdatedAt:    &now,
					DeletedAt:    nil,
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

			query := `
			SELECT (.+)
			FROM menus m
			INNER JOIN menu_categories mc ON mc.id = m.menu_category_id
			WHERE (.+)
			ORDER BY mc.name`

			dbResult := sqlmock.NewRows([]string{"id", "name", "description", "price", "category_name", "created_at", "updated_at", "deleted_at"})
			for _, m := range tc.mockRows {
				dbResult.AddRow(m.ID, m.Name, m.Description, m.Price, m.CategoryName, m.CreatedAt, m.UpdatedAt, m.DeletedAt)
			}

			mock.ExpectQuery(query).WillReturnRows(dbResult).WillReturnError(tc.mockErr)

			repo := New(dbx)
			res, err := repo.FetchMenus(context.Background())
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
				assert.Equal(t, res[0].Description, m.Description)
				assert.Equal(t, res[0].Price, m.Price)
				assert.Equal(t, res[0].CategoryName, m.CategoryName)
			}
		})
	}
}

func TestMenus_CreateMenu(t *testing.T) {
	query := "INSERT INTO menus (.+) VALUES (.+)"
	t.Run("should be succeed", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs("name", "desc", float64(1000), 1).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := New(dbx)
		err := repo.CreateMenu(context.Background(), spec.CreateMenu{
			Name:           "name",
			Description:    "desc",
			Price:          1000,
			MenuCategoryID: 1,
		})
		assert.NoError(t, err)
	})

	t.Run("should be failed because return err from repo", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs("name", "desc", float64(1000), 1).
			WillReturnResult(nil).
			WillReturnError(errors.New("unexpected error"))

		repo := New(dbx)
		err := repo.CreateMenu(context.Background(), spec.CreateMenu{
			Name:           "name",
			Description:    "desc",
			Price:          1000,
			MenuCategoryID: 1,
		})
		assert.Error(t, err)
	})
}

func TestMenus_UpdateMenu(t *testing.T) {
	query := "UPDATE menus SET (.+) WHERE (.+)"
	t.Run("should be succeed", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs("name", "desc", float64(1000), 1, 1).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := New(dbx)
		err := repo.UpdateMenu(context.Background(), spec.UpdateMenu{
			ID:             1,
			Name:           "name",
			Description:    "desc",
			Price:          1000,
			MenuCategoryID: 1,
		})
		assert.NoError(t, err)
	})

	t.Run("should be failed because return err from repo", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs("name", "desc", float64(1000), 1, 1).
			WillReturnResult(nil).
			WillReturnError(errors.New("unexpected error"))

		repo := New(dbx)
		err := repo.UpdateMenu(context.Background(), spec.UpdateMenu{
			ID:             1,
			Name:           "name",
			Description:    "desc",
			Price:          1000,
			MenuCategoryID: 1,
		})
		assert.Error(t, err)
	})
}

func TestMenus_DeleteMenu(t *testing.T) {
	query := "UPDATE menus SET (.+) WHERE (.+)"
	t.Run("should be succeed", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		dbx := sqlx.NewDb(db, "sqlmock")

		mock.ExpectExec(query).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := New(dbx)
		err := repo.DeleteMenu(context.Background(), 1)
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
		err := repo.DeleteMenu(context.Background(), 1)
		assert.Error(t, err)
	})
}
