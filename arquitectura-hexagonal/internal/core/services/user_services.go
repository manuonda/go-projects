package services

import (
	"arquitectura-hexagonal/internal/core/domain"
	port "arquitectura-hexagonal/internal/core/ports"
	"errors"
)

// Los servicios son la implementacion
type UserService struct {
	userRepository port.UserRepository
}

func NewUserService(repository port.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) CreateUser(user *domain.User) error {
	if user.Email == "" {
		return errors.New("el email no puede estar vacio")
	}

	//Validar que el email no exista ya en el repositorio
	existingUserEmail, err := s.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUserEmail != nil {
		return errors.New("el email ya se encuentra registrado")
	}
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return s.userRepository.GetUserByEmail(email)

}
