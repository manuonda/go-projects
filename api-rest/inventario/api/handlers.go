package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manuonda/go-projects/inventario/api/dtos"
	"github.com/manuonda/go-projects/inventario/internal/service"
	"github.com/pkg/errors"
)

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, err)
		}
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error"))

	}
	return c.JSON(http.StatusCreated, nil)
}
