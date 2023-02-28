package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	entity "github.com/manuonda/go-projects/inventario/internal/entitys"
)

//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)

	SaveUserRole(ctx context.Context, userId, roleId int64) error
	RemoveUserRole(ctx context.Context, userId, roleId int64) error
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
