package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/manuonda/go-projects/inventario/database"
	"github.com/manuonda/go-projects/inventario/internal/repository"
	"github.com/manuonda/go-projects/inventario/internal/service"
	"github.com/manuonda/go-projects/inventario/settings"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("SELECT * FROM USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
