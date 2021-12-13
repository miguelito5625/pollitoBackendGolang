package Systemroutes

import (
	"pollitoBackendGolang/Controllers"
	"pollitoBackendGolang/Middlewares"

	"github.com/gin-gonic/gin"
)

func Usuario_Routes(route *gin.Engine) {
	usuarioRoutes := route.Group("/usuario")
	usuarioRoutes.Use(Middlewares.CheckAuth())
	{

		usuarioRoutes.GET("", Controllers.ListUsuario)
		usuarioRoutes.GET("/:id", Controllers.GetOneUsuario)
		usuarioRoutes.POST("", Controllers.AddNewUsuario)
		usuarioRoutes.PUT("/:id", Controllers.PutOneUsuario)
		usuarioRoutes.DELETE("/:id", Controllers.DeleteUsuario)

	}

}
