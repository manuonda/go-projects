package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manuonda/go-projects/api-rest/go-api-mongodb/controller"
	"github.com/manuonda/go-projects/api-rest/go-api-mongodb/models"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/posts", controller.CreatePost)
	router.GET("/posts/:id", controller.GetProductoById)
	router.GET("/posts", controller.GetAll)
	router.POST("/posts/:id", controller.Update)
	router.Run("localhost:8080")

}
