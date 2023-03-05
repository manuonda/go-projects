package entitys

type UserRole struct {
	UserID int64 `db:"user_id" json:"user_id"`
	RoleID int64 `db:"role_id"  json:"role_id"`
}
