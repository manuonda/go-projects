package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/api/router"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/injectors"
)

func main() {

	err := injectors.LoadConfig()
	if err != nil {
		fmt.Println("Error connection database")
	}

	server := gin.Default()
	router.UsuarioRouter(server)
	server.Run("localhost:9003")
	fmt.Println("Database Connect")
}
