package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manuonda/go-projects/inventario/api/dtos"
	"github.com/manuonda/go-projects/inventario/encryption"
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

func (api *API) LonginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}
	fmt.Println("params : ", params.Email)
	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid Request"})
	}

	err = api.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	u, err := api.serv.LoginUser(ctx, params.Email, params.Password)

	if err != nil {
		if err == service.ErrInvalidCredentials {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	token, err := encryption.SignedLoginToken(u)
	if err != nil {
		log.Println("existe error que onda : ", err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"success": token})

}
