package injectors

import (
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/controllers"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/repository"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/service"
	"gorm.io/gorm"
)

var (
	Connection *gorm.DB

	// repository
	usuarioRepository repository.UsuarioRepository

	// service
	usuarioService service.UsuarioService

	//controller
	usuarioController *controllers.UsuarioController
)
