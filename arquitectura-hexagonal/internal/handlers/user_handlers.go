package handlers

import (
	"arquitectura-hexagonal/internal/core/domain"
	port "arquitectura-hexagonal/internal/core/ports"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

/**
Claro, en la arquitectura hexagonal (también conocida como arquitectura de puertos y adaptadores),
los handlers (manejadores) son componentes que actúan como
intermediarios entre el mundo exterior (como interfaces de usuario, APIs, etc.)
y la lógica de negocio de la aplicación.
*/

type UserHandlers struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) GetUserByEmail(c *fiber.Ctx) error {

	email := c.Params("email")
	user, err := h.userService.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(&user)
}

func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	fmt.Println("ingresando create user handler")
	user := &domain.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.userService.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(nil)
	}

	return c.JSON(&user)
}
