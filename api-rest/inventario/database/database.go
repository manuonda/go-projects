package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/manuonda/go-projects/inventario/settings"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	fmt.Println(s.DB.User)
	fmt.Println(s.DB.Password)
	fmt.Println(s.DB.Host)
	fmt.Println(s.DB.Port)

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)

	fmt.Println("inforamcion connectionString")
	fmt.Println(connectionString)
	fmt.Println("end conneciton string")

	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}
