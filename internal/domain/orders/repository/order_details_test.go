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

func TestOrders_FetchOrderDetails(t *testing.T) {
	now := time.Now()
	orderID := 1
	testCases := []struct {
		name     string
		mockRows []model.OrderDetails
		mockErr  error
		isError  bool
	}{
		{
			name: "should be succeed",
			mockRows: []model.OrderDetails{
				{
					ID:               1,
					OrderID:          1,
					MenuID:           1,
					MenuName:         "Latte",
					MenuCategoryName: "Signature",
					PriceMenu:        1000,
					CreatedAt:        &now,
					UpdatedAt:        &now,
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

			dbResult := sqlmock.NewRows([]string{"id", "order_id", "menu_id", "menu_name", "menu_category_name", "price_menu", "created_at", "updated_at"})
			for _, m := range tc.mockRows {
				dbResult.AddRow(m.ID, m.OrderID, m.MenuID, m.MenuName, m.MenuCategoryName, m.PriceMenu, m.CreatedAt, m.UpdatedAt)
			}

			mock.ExpectQuery("SELECT (.+) FROM order_details od INNER JOIN menus m ON m.id = od.menu_id INNER JOIN menu_categories mc ON mc.id = m.menu_category_id WHERE (.+)").WithArgs(orderID).WillReturnRows(dbResult).WillReturnError(tc.mockErr)

			repo := New(dbx)
			res, err := repo.FetchOrderDetails(context.Background(), orderID)
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
				assert.Equal(t, res[0].OrderID, m.OrderID)
				assert.Equal(t, res[0].MenuID, m.MenuID)
				assert.Equal(t, res[0].MenuName, m.MenuName)
				assert.Equal(t, res[0].MenuCategoryName, m.MenuCategoryName)
				assert.Equal(t, res[0].PriceMenu, m.PriceMenu)
			}
		})
	}
}
