package service

import (
	"fmt"
	"gin-gonic/gi_yt/models"
	"gin-gonic/gi_yt/repository"
)

type IUserService interface {
	Save(user *models.User) (*models.User, error)
	FindAll() ([]*models.User, error)
}

type userService struct {
	UserRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &userService{UserRepo: userRepo}
}

// FindAll implements IUserService
func (*userService) FindAll() ([]*models.User, error) {
	panic("unimplemented")
}

// Save implements IUserService
func (us *userService) Save(user *models.User) (*models.User, error) {
	//panic("unimplemented")
	fmt.Println("Save User service implement {}", user)
	user, err := us.UserRepo.Save(user)
	if err != nil {
		fmt.Print("Error Save ")
	}
	fmt.Println("Saved user in capa service implement ", user)
	return user, nil
}
