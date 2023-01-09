package service

import (
	"database/sql"
	"fmt"
	"gin-gonic/gi_yt/models"
)

type UserServiceMysqlImpl struct {
	database *sql.DB
}

func NewUserServiceMysqlImpl(db *sql.DB) *UserServiceMysqlImpl {
	return &UserServiceMysqlImpl{
		database: db,
	}
}

func (u *UserServiceMysqlImpl) CreateUser(user *models.User) string {
	fmt.Println("user parameter impl ", user)
	return user.Name
}
