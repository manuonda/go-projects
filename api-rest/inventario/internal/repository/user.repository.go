package repository

import (
	"context"

	"github.com/manuonda/go-projects/inventario/internal/entitys"
)

const (
	qryInsterUser = `
     insert into USERS(email, name, password)
      value(?,?,?);
	`

	qryGetUserByEmail = `
	  select 
	  id, 
	  email, 
	  name,
	  password 
	  from users
	  where email = ?
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsterUser, email, name, password)
	return err
}
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entitys.User, error) {
	u := entitys.User{}

	err := r.db.GetContext(ctx, &u, qryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
