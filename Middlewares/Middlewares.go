package Middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		// token := c.Query("token")
		authHeader := c.GetHeader("Authorization")

		log.Println("Esto es un middleware")
		log.Println("Token:", authHeader)

		// c.JSON(http.StatusUnauthorized, gin.H{
		// 	"code": http.StatusUnauthorized,
		// 	"msg":  "ERROR",
		// 	"data": "",
		// })
		// c.Abort()

		c.Next()
	}

}
