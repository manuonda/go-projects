package repository

import (
	"context"
	"fmt"

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
	  from USERS
	  where email = ?
	`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsterUser, email, name, password)
	return err
}
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entitys.User, error) {
	u := entitys.User{}
	fmt.Println("email found", email)

	err := r.db.GetContext(ctx, &u, qryGetUserByEmail, email)
	fmt.Println("error found getcontext : ", err)
	if err != nil {

		return nil, err
	}
	fmt.Println("user found repositorio : ", u)
	return &u, nil
}
