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
func (ur *UserRepository) FindAll() ([]*models.User, error) {
	fmt.Print("Find All repository")
	rows, err := ur.DB.Query("select * from ")
	if err != nil {
		fmt.Println("Error Query Select")
		return nil, err
	}

	defer rows.Close()

	var user models.User
	var users *[]models.User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Println("Error scan variables")
		}
		append(*users, *user)
	}
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
