package port

// Los puertos definen mediante interfaces los métodos que
// debe contener una estructura para ser un
// Servicio o Repositorio.

import (
	"arquitectura-hexagonal/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
}

type UserService interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
}

type UserHandlers interface {
	CreateUser(c *fiber.Ctx) error
	GetUserByEmail(c *fiber.Ctx) error
}
