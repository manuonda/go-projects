package controller

import (
	"fmt"
	"gin-gonic/gi_yt/models"
	service "gin-gonic/gi_yt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.IUserService
}

func New(userService service.IUserService) *UserController {
	return &UserController{
		service: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	fmt.Println("ingreso create user ")
	var user *models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := uc.service.Save(user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &user)
}

func createUser(user *models.User) *models.User {
	fmt.Println(user)
	return user
}
