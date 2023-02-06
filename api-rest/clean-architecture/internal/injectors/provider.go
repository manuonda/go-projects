package injectors

import "github.com/manuonda/go-projects/api-rest/clean-architecture/internal/controllers"

func GetControllerUsuario() *controllers.UsuarioController {
	return usuarioController
}
