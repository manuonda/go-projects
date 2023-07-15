package router

import (
	"github.com/gin-gonic/gin"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/handlers"
)

func RegisterRouterPost(router *gin.Engine) {

	router.Group("/posts")
	router.GET("/posts", handlers.GetAllPost)
	router.POST("/posts", handlers.CreatePost)
}
