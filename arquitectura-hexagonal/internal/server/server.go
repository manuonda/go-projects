package server

import (
	port "arquitectura-hexagonal/internal/core/ports"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	//We will add every new hander here
	userHandler port.UserHandlers
	//middlewares ports.Middleware
	//paymentHandlers ports.paymetnHandlers
}

func NewServer(uHandlers port.UserHandlers) *Server {
	return &Server{
		userHandler: uHandlers,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("/user")
	userRoutes.Post("/create", s.userHandler.CreateUser)
	userRoutes.Get("/:email", s.userHandler.GetUserByEmail)

	err := app.Listen(":5000")
	if err != nil {
		log.Fatal("Error port : ", err)
	}
}
