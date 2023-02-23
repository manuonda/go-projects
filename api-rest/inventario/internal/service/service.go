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
}

type service struct {
	repo repository.Repository
}

func New(repo *repository.Repository) Service {
	return &service{
		repo: *repo,
	}
}
