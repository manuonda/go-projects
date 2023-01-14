package router

import (
	"database/sql"
	"fmt"
	controller "gin-gonic/gi_yt/controllers"
	"gin-gonic/gi_yt/repository"
	service "gin-gonic/gi_yt/services"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func NewRouter(rg *gin.RouterGroup, db *sql.DB) {
	NewUserRouter(rg, db)
	DB = db
	fmt.Println("aqui en newrouter DB ====> ", DB)
}

func NewUserRouter(rg *gin.RouterGroup, db *sql.DB) {
	fmt.Println("Database DB router : ", db)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	userController := controller.New(userService)
	route := rg.Group("/user")

	route.POST("/create", userController.CreateUser)
}
