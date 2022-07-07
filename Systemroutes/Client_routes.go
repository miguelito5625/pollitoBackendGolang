package Systemroutes

import (
	"pollitoBackendGolang/Controllers"
	"pollitoBackendGolang/Middlewares"

	"github.com/gin-gonic/gin"
)

func Client_Routes(route *gin.Engine) {
	autenticacionRoutes := route.Group("/client")
	autenticacionRoutes.Use(Middlewares.CheckAuth())
	{
		autenticacionRoutes.POST("/create", Controllers.AddNewClient)
	}

}
