package Systemroutes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	Rol_Routes(r)
	Usuario_Routes(r)
	Vacuna_Routes(r)
	Centro_Vacunacion_Routes(r)
	Grupo_Vacunacion_Routes(r)
	Autenticacion_Routes(r)

	return r
}
