package repository

import (
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/models"
	"gorm.io/gorm"
)

type usuarioRepository struct {
	conn *gorm.DB
}

type UsuarioRepository interface {
	// GetUsuarios() ([]models.Usuario, error)
	CreateUsuario(usuario *models.Usuario) (*int64, error)
	// UpdateUsuario(usuario models.Usuario) (*models.Usuario, error)
	// DeleteUsuario(id int64) error
}

func NewUsuarioRepository(c *gorm.DB) UsuarioRepository {
	return &usuarioRepository{conn: c}
}

// CreateUsuario implements UsuarioInterface
func (*usuarioRepository) CreateUsuario(usuario *models.Usuario) (*int64, error) {
	panic("unimplemented")
}
