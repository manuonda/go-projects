package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/database"
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/router"
)

func main() {
	routerDefault := gin.Default()

	database.ConnectDatabase()

	// Router configuration
	router.RegisterRouterPost(routerDefault)
	router.RegisterRouterComment(routerDefault)

	routerDefault.Run("localhost:8083")

}
