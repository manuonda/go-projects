package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/manuonda/go-projects/inventario/encryption"
	"github.com/manuonda/go-projects/inventario/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("User already exists")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

func (s *service) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	fmt.Println("user found :", u)
	if u != nil {
		fmt.Println("user exists")
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

	fmt.Println("Service encontrado1 : ", u.ID)

	//TODO: decrypt password
	bb, err := encryption.FromBase64(u.Password)

	if err != nil {
		return nil, err
	}

	fmt.Println("Service encontrado2 : ", u.ID)

	decryptedPassword, err := encryption.Decrypt(bb)

	fmt.Println("Service encontrado3 : ", u.ID)

	if err != nil {
		return nil, err
	}

	fmt.Println("Service encontrado4 : ", u.ID)

	if string(decryptedPassword) != password {
		fmt.Println("aqui ingreso no es el mismo password")
		return nil, ErrInvalidCredentials
	}
	fmt.Println("Service encontrado3 : ", u.ID)

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
