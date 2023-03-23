package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manuonda/go-projects/inventario/api/dtos"
	"github.com/manuonda/go-projects/inventario/internal/service"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid Request"})
	}

	//Validator de los tags al bindear los parametros
	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "Error User Already Exists"})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Error internal server"})

	}
	return c.JSON(http.StatusCreated, nil)
}
