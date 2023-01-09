package controller

import (
	"fmt"
	"gin-gonic/gi_yt/models"
	service "gin-gonic/gi_yt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerMysql struct {
	UserService service.UserService
}

func New(userService service.UserService) *UserControllerMysql {
	return &UserControllerMysql{
		UserService: userService,
	}
}

type Body struct {
	// json tag to de-serialize json body
	Name string `json:"name"`
}

func (uc *UserControllerMysql) CreateUser(ctx *gin.Context) {

	fmt.Println("ingreso create user ")
	user := models.User{Name: "david", Age: 23}
	//body := Body{}
	// if err := ctx.BindJSON(&user); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	user = models.User{Name: "informacion numeo 23", Age: 44}
	fmt.Println("prueba informnacon", user)
	fmt.Println("--------------")
	fmt.Println("User , ", user)
	fmt.Println("user2 : ", &user)
	user2 := models.User{Name: user.Name, Age: user.Age}

	createUser(&user2)

	err := uc.UserService.CreateUser(&user2)

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

func (uc *UserControllerMysql) RegisterUserRoutesMysql(rg *gin.RouterGroup) {
	route := rg.Group("/user")
	route.POST("/create", uc.CreateUser)
	/*route.GET("/get/:name", uc.GetUser)
	route.GET("/getAll", uc.GetAll)
	route.PATCH("/update", uc.UpdateUser)
	route.DELETE("/delete/:name", uc.DeleteUser)
	*/
}
