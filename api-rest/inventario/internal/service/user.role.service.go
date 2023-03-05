package service

import (
	"context"
	"errors"
)

type UserRoleService interface {
}

var (
	ErrRoleAlreadyAdded = errors.New("Role was already added for this user")
	ErrRoleNotFound     = errors.New("Role not found for this user")
)

// AddUserRole implements Service
func (s *service) AddUserRole(ctx context.Context, UserId int64, RoleId int64) error {
	roles, err := s.repo.GetUserRolesByUserId(ctx, UserId)
	if err != nil {
		return err
	}

	for _, rol := range roles {
		if rol.RoleID == RoleId {
			return ErrRoleAlreadyAdded
		}
	}
	return s.repo.SaveUserRole(ctx, UserId, RoleId)
}

// RemoveUserRole implements Service
func (s *service) RemoveUserRole(ctx context.Context, UserId int64, RoleId int64) error {
	roles, err := s.repo.GetUserRolesByUserId(ctx, UserId)
	if err != nil {
		return err
	}

	var roleFound = false
	for _, rol := range roles {
		if rol.RoleID == RoleId {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repo.RemoveUserRole(ctx, UserId, RoleId)
}
