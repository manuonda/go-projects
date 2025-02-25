package repository

import (
	"arquitectura-hexagonal/internal/core/domain"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := "SELECT email, name FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)
	var user domain.User

	err := row.Scan(&user.Email, &user.Name)
	fmt.Printf("Error repository getUserByEmail : ", err)
	if err != nil {
		fmt.Print("Aqui ingreso uno 1")
		if err == sql.ErrNoRows {
			fmt.Println("Aqui ingreso dos")
			return nil, nil
		}
		return nil, errors.New("error getting user by email")
	}
	return &user, nil

}

func (r *UserRepository) CreateUser(user *domain.User) error {
	query := "INSERT INTO users(email, name) values (?,?)"
	_, err := r.db.Exec(query, user.Email, user.Name)
	if err != nil {
		fmt.Printf("Erro create user : %v", err)
		return fmt.Errorf("error creating user : %v", err)
	}
	return nil
}
