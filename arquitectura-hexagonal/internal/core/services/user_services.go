package services

import (
	"arquitectura-hexagonal/internal/core/domain"
	port "arquitectura-hexagonal/internal/core/ports"
	"errors"
	"fmt"
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
	fmt.Println("\n Ingresando a create user service")
	fmt.Printf("email : %v , %v", user.Email, user.Name)
	if user.Email == "" {
		return errors.New("el email no puede estar vacio")
	}

	fmt.Println("paso por aca")
	//Validar que el email no exista ya en el repositorio
	existingUserEmail, err := s.userRepository.GetUserByEmail(user.Email)
	fmt.Printf("Error  : %v", err)
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
