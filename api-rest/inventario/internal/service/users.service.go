package service

import (
	"context"
	"errors"

	"github.com/manuonda/go-projects/inventario/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("User already exists")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

func (s *service) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}
	//TODO: hasspassword
	return s.repo.SaveUser(ctx, email, name, password)
}

func (s *service) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//TODO: decrypt password
	if u.Password != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
