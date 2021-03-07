package infrastructure

import (
	"context"
	"fmt"
	"food-order-api/internal/shared/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQL(ctx context.Context, cfg config.MySQL) (*sqlx.DB, error) {
	connString := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sqlx.ConnectContext(ctx, "mysql", connString)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(15)
	db.SetMaxOpenConns(30)
	db.SetConnMaxIdleTime(3600 * time.Second)
	db.SetConnMaxLifetime(3600 * time.Second)

	return db, nil
}
