package main

import (
	"context"
	"fmt"
	controller "gin-gonic/gi_yt/controllers"
	"gin-gonic/gi_yt/database"
	service "gin-gonic/gi_yt/services"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	us     service.UserService
	uc     controller.UserControllerMysql
	ctx    context.Context
	err    error
)

func init() {
	fmt.Println("Connect to Database")
	database.Connect()

	//us = &service.UserServiceMysqlImpl{database: db}
	//uc = controller.UserControllerMysql(us)

	fmt.Println("db connected")
}

func main() {
	var server = gin.Default()
	basepath := server.Group("/v1")
	uc.RegisterUserRoutesMysql(basepath)
	log.Fatal(server.Run(":9090"))
}
