package main

import (
	"arquitectura-hexagonal/internal/core/services"
	"arquitectura-hexagonal/internal/handlers"
	"arquitectura-hexagonal/internal/repository"
	"arquitectura-hexagonal/internal/server"
	"log"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// }

func main() {

	db, err := repository.ConnectToMySQL()
	if err != nil {
		log.Fatalf("Error connection to database : %v ", err)
	}
	//repositories
	userRepostory := repository.NewUserRepository(db)
	//services
	userService := services.NewUserService(userRepostory)
	//handlers
	userHandler := handlers.NewUserHandler(userService)
	//server
	httpServe := server.NewServer(userHandler)

	httpServe.Initialize()

}
