package service

import (
	"context"
	"errors"

	"github.com/manuonda/go-projects/inventario/encryption"
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
	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	pass := encryption.ToBase64(bb)
	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *service) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//TODO: decrypt password
	bb, err := encryption.FromBase64(u.Password)

	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)

	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
