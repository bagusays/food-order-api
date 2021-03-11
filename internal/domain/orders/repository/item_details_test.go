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

func TestOrders_FetchItemDetails(t *testing.T) {
	now := time.Now()
	orderDetailID := 1
	testCases := []struct {
		name     string
		mockRows []model.ItemDetails
		mockErr  error
		isError  bool
	}{
		{
			name: "should be succeed",
			mockRows: []model.ItemDetails{
				{
					ID:              1,
					OrderDetailID:   1,
					AdditionalID:    1,
					AdditionalName:  "Espresso +1",
					AdditionalPrice: 500,
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

			dbResult := sqlmock.NewRows([]string{"id", "order_detail_id", "additional_id", "additional_name", "additional_price", "created_at", "updated_at"})
			for _, m := range tc.mockRows {
				dbResult.AddRow(m.ID, m.OrderDetailID, m.AdditionalID, m.AdditionalName, m.AdditionalPrice, m.CreatedAt, m.UpdatedAt)
			}

			mock.ExpectQuery("SELECT (.+) FROM item_details ids INNER JOIN additionals a ON a.id = ids.additional_id WHERE (.+)").WithArgs(orderDetailID).WillReturnRows(dbResult).WillReturnError(tc.mockErr)

			repo := New(dbx)
			res, err := repo.FetchItemDetails(context.Background(), orderDetailID)
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
				assert.Equal(t, res[0].OrderDetailID, m.OrderDetailID)
				assert.Equal(t, res[0].AdditionalID, m.AdditionalID)
				assert.Equal(t, res[0].AdditionalPrice, m.AdditionalPrice)
				assert.Equal(t, res[0].AdditionalName, m.AdditionalName)
			}
		})
	}
}
