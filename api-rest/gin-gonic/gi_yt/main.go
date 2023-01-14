package main

import (
	"context"
	"database/sql"
	"fmt"
	"gin-gonic/gi_yt/database"
	"gin-gonic/gi_yt/router"
	service "gin-gonic/gi_yt/services"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
	us     service.UserService
	ctx    context.Context
	err    error
	DB     *sql.DB
)

func init() {

	//us = &service.UserServiceMysqlImpl{database: db}
	//uc = controller.UserControllerMysql(us)

	fmt.Println("db connected")
	fmt.Println("--- db init -----")
	fmt.Println(" salida db: ", DB)
}

func main() {
	fmt.Println("Connect to Database")
	DB, err := database.Connect()
	if err != nil {
		fmt.Print("falta error connection")
	}
	fmt.Println("-----------------")
	fmt.Println("--- main onta DB", DB)
	var server = gin.Default()
	basepath := server.Group("/v1")
	router.NewRouter(basepath, DB)
	log.Fatal(server.Run(":9091"))
}
