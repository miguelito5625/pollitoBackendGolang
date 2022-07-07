package Systemroutes

import (
	"pollitoBackendGolang/Controllers"
	"pollitoBackendGolang/Middlewares"

	"github.com/gin-gonic/gin"
)

func User_Routes(route *gin.Engine) {
	userRoutes := route.Group("/user")
	userRoutes.Use(Middlewares.CheckAuth())
	{

		userRoutes.GET("", Controllers.ListUser)
		userRoutes.GET("/:id", Controllers.GetOneUser)
		userRoutes.POST("", Controllers.AddNewUser)
		userRoutes.PUT("/:id", Controllers.PutOneUser)
		userRoutes.DELETE("/:id", Controllers.DeleteUser)

	}

}
