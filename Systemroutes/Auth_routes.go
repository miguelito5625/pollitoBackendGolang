package Systemroutes

import (
	"pollitoBackendGolang/Controllers"
	"pollitoBackendGolang/Middlewares"

	"github.com/gin-gonic/gin"
)

func Autenticacion_Routes(route *gin.Engine) {
	autenticacionRoutes := route.Group("/autenticacion")
	autenticacionRoutes.Use(Middlewares.CheckAuth())
	{

		autenticacionRoutes.POST("/login", Controllers.LoginUsuario)

		//asd
	}

}
