package Systemroutes

import (
	"pollitoBackendGolang/Controllers"
	"pollitoBackendGolang/Middlewares"

	"github.com/gin-gonic/gin"
)

func Client_Routes(route *gin.Engine) {
	autenticacionRoutes := route.Group("/clients")
	autenticacionRoutes.Use(Middlewares.CheckAuth())
	{
		autenticacionRoutes.GET("", Controllers.ListClient)
		autenticacionRoutes.POST("", Controllers.AddNewClient)
	}

}
