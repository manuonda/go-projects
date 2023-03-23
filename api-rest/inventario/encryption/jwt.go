package encryption

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/manuonda/go-projects/inventario/internal/models"
)

func SignedLoginToken(u *models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
	})

	return token.SignedString([]byte(key))
}
