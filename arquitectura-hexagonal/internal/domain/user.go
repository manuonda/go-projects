package domain

//https://carlos.lat/blog/hexagonal-architecture-using-go-fiber/

type User struct {
	ID       int
	Email    string
	Password string
}

// func NewPerson(id int, email string, password string) {
// 	return &User{
// 		ID:       id,
// 		Email:    email,
// 		Password: password,
// 	}
// }

func (u *User) GetEmail() string {
	return u.Email
}
