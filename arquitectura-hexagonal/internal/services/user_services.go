package services

import port "arquitectura-hexagonal/internal/ports"

type UserService struct {
	userRepository port.UserRepository
}

func NewUserService(repository port.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) CreateUser(user *) error {
	return s.userRepository.CreateUser(user)
}
