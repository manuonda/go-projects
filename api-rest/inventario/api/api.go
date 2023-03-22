package api

import (
	"github.com/labstack/echo/v4"
	"github.com/manuonda/go-projects/inventario/internal/service"
)

type API struct {
	serv service.Service
}

func New(serv service.Service) *API {
	return &API{
		serv: serv,
	}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRouter(e)
	return e.Start(address)
}
