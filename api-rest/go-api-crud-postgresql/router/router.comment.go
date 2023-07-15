package router

import (
	"github.com/gin-gonic/gin"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/handlers"
)

func RegisterRouterComment(router *gin.Engine) {

	router.GET("/comments", handlers.GetAll)
	router.POST("/comments", handlers.CreateComent)
}
