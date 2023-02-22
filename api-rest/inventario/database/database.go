package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context) *sqlx.DB {
	connectionString := "root:123456@tcp(localhost:3306)/max_invetory?parseTime=true"

	sqlx.ConnectContext(ctx, "mysql", connectionString)
}
