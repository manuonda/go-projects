package router

import (
	"github.com/gin-gonic/gin"
	"github.com/manuonda/go-projects/api-rest/clean-architecture/internal/injectors"
)

var groupUsuario *gin.RouterGroup

func UsuarioRouter(router *gin.Engine) {
	Controller := injectors.GetControllerUsuario()
	groupUsuario = router.Group("/usuario")
	groupUsuario.GET("/list", Controller.CreateUsuario)
}
