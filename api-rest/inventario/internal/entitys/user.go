package entitys

type User struct {
	ID       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"-"`
}
