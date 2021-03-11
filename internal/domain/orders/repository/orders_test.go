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

func TestOrders_FetchAllOrders(t *testing.T) {
	now := time.Now()
	userID := 1
	testCases := []struct {
		name     string
		mockRows []model.Orders
		mockErr  error
		isError  bool
	}{
		{
			name: "should be succeed",
			mockRows: []model.Orders{
				{
					ID:            1,
					UserID:        1,
					PaymentStatus: model.Paid,
					PaidBy:        "LINKAJA",
					PaidAt:        &now,
					TotalPrice:    1000,
					OrderStatus:   model.Completed,
					CreatedAt:     &now,
					UpdatedAt:     &now,
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

			dbResult := sqlmock.NewRows([]string{"id", "user_id", "payment_status", "paid_by", "paid_at", "total_price", "order_status", "created_at", "updated_at"})
			for _, m := range tc.mockRows {
				dbResult.AddRow(m.ID, m.UserID, m.PaymentStatus, m.PaidBy, m.PaidAt, m.TotalPrice, m.OrderStatus, m.CreatedAt, m.UpdatedAt)
			}

			mock.ExpectQuery("SELECT (.+) FROM orders WHERE (.+)").WithArgs(userID).WillReturnRows(dbResult).WillReturnError(tc.mockErr)

			repo := New(dbx)
			res, err := repo.FetchAllOrders(context.Background(), userID)
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
				assert.Equal(t, res[0].UserID, m.UserID)
				assert.Equal(t, res[0].PaymentStatus, m.PaymentStatus)
				assert.Equal(t, res[0].PaidBy, m.PaidBy)
				assert.Equal(t, res[0].PaidAt, m.PaidAt)
				assert.Equal(t, res[0].TotalPrice, m.TotalPrice)
				assert.Equal(t, res[0].OrderStatus, m.OrderStatus)
			}
		})
	}
}
