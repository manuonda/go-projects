package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/models"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/service"
)

type UsuarioController struct {
	service *service.UsuarioService
}

func NewUsuarioController(s *service.UsuarioService) *UsuarioController {
	return &UsuarioController{service: s}
}
func (u *UsuarioController) CreateUsuario(c *gin.Context) {

	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "encontrado"})
	return

}
