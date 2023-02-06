package service

import (
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/models"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/repository"
)

type usuarioService struct {
	repository *repository.UsuarioRepository
}

type UsuarioService interface {
	CreateUsuario(usuario *models.Usuario) (int64, error)
}

func NewUsuarioService(r *repository.UsuarioRepository) UsuarioService {
	return &usuarioService{repository: r}
}

// CreateUsuario implements UsuarioService
func (*usuarioService) CreateUsuario(usuario *models.Usuario) (int64, error) {
	panic("unimplemented")
}
