package repository

import (
	"database/sql"
	"fmt"
	"gin-gonic/gi_yt/models"
	"log"
)

type IUserRepository interface {
	Save(u *models.User) (*models.User, error)
	FindAll() ([]*models.User, error)
}
type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	fmt.Println("que aque in new userrepostiroy : ", db)
	return &UserRepository{db}
}

// FindAll implements IUserRepository
func (*UserRepository) FindAll() ([]*models.User, error) {
	panic("unimplemented")
}

// Save implements IUserRepository
func (ur *UserRepository) Save(u *models.User) (*models.User, error) {
	fmt.Println("Save in Repoistory {} ", u)

	name := u.Name
	age := u.Age
	fmt.Print(ur.DB)
	stmt, err := ur.DB.Prepare("INSERT INTO users (name,age) VALUES (?, ?)")
	if err != nil {
		log.Fatalf("impossible insert user: %s", err)
		return nil, err
	}
	_, err = stmt.Exec(name, age)
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
		return nil, err
	}
	log.Printf("inserted id")
	return u, nil
}
