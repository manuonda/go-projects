package encryption

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/manuonda/go-projects/inventario/internal/models"
)

func SignedLoginToken(u *models.User) (string, error) {

	fmt.Println("signedLoginToken")
	fmt.Println("user found models : ", u.Email, u.Name)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
	})

	return token.SignedString([]byte(key))
}
