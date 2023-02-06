package injectors

import (
	"fmt"

	"github.com/manuonda/go-projects/api-rest/clean-architecture/database"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/controllers"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/repository"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/service"
	"gorm.io/gorm"
)

func LoadConfig() error {
	connection, err := database.InitDb()
	if err != nil {
		fmt.Println("Error init db")
	}
	InitRepository(connection)
	InitServices()
	InitController()

	return err

}

func initRepository(connection *gorm.DB) {
	Connection = connection
}

func InitRepository(connection *gorm.DB) {
	usuarioRepository = repository.NewUsuarioRepository(connection)
}

func InitServices() {
	usuarioService = service.NewUsuarioService(&usuarioRepository)
}

func InitController() {
	usuarioController = controllers.NewUsuarioController(&usuarioService)
}
