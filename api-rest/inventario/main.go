package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/manuonda/go-projects/inventario/api"
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
			api.New,
			echo.New,
		),
		fx.Invoke(
			// func(ctx context.Context, service service.Service) {
			// 	err := service.RegisterUser(ctx, "manuonda@gmail.com", "manuonda", "12345678")
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// },
			setLifeCycle,
		),
	)

	app.Run()
}

/*
*
Establece el ciclo de vida
*
*/
func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, echo *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(echo, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
