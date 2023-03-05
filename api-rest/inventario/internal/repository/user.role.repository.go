package repository

import (
	"context"

	"github.com/manuonda/go-projects/inventario/internal/entitys"
)

const (
	queryUserRoleInsert = (`
	   insert into USER_ROLES( user_id, role_id) values(:user_id,:role_id);
	`)

	queryUserRoleRemove = (`
	  delete from USER_ROLES where user_id = :user_id and role_id = :role_id
	`)
)

func (r *repo) SaveUserRole(ctx context.Context, userId, roleId int64) error {

	data := entitys.UserRole{
		UserID: userId,
		RoleID: roleId,
	}

	_, err := r.db.NamedExecContext(ctx, queryUserRoleInsert, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) RemoveUserRole(ctx context.Context, userId, roleId int64) error {
	data := entitys.UserRole{
		UserID: userId,
		RoleID: roleId,
	}

	_, err := r.db.NamedExecContext(ctx, queryUserRoleInsert, data)
	return err
}

func (r *repo) GetUserRolesByUserId(ctx context.Context, userId int64) ([]entitys.UserRole, error) {
	var roles = []entitys.UserRole{}

	err := r.db.SelectContext(ctx, &roles, "select user_id, role_id from USER_ROLES where user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
