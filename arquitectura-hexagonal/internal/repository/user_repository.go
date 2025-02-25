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
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Error finding by email")
		}
		return nil, errors.New("Error getting user by email ")
	}
	return &user, nil

}

func (r *UserRepository) CreateUser(user *domain.User) error {
	query := "INSERT INTO users(email, name) values (?,?)"
	_, err := r.db.Exec(query, user.Email, user.Name)
	if err != nil {
		return fmt.Errorf("error creating user : %v", err)
	}
	return nil
}
