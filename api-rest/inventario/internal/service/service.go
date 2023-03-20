package service

import (
	"context"

	"github.com/manuonda/go-projects/inventario/internal/models"
	"github.com/manuonda/go-projects/inventario/internal/repository"
)

// Service is the business logic of the application

//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)

	AddUserRole(ctx context.Context, UserId, RoleId int64) error
	RemoveUserRole(ctx context.Context, UserId, RoleId int64) error

	// save products
	SaveProduct(ctx context.Context, product models.ProductDTO, email string) error
	GetProduct(ctx context.Context) ([]models.ProductDTO, error)
	GetByIdProduct(ctx context.Context, ID int64) (*models.ProductDTO, error)
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
