package main

import (
	"io"
	"log"
	"os"
	"pollitoBackendGolang/Database"
	"pollitoBackendGolang/Database/Migrations"
	"pollitoBackendGolang/Systemroutes"

	"github.com/gin-gonic/gin"
)

func main() {

	Database.ConnectDataBase()
	Migrations.MigrateAll()

	// logfile, errFile := os.OpenFile("server.log", os.O_RDWR|os.O_APPEND, 0660)
	// if errFile != nil {
	// 	logfile, _ = os.Create("server.log")
	// 	// log.Println("Error al abrir el archivo server.log")
	// }

	logfile, _ := os.Create("server.log")

	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
	log.SetOutput(gin.DefaultWriter)
	log.Println("")
	log.Println("Servidor iniciado")

	r := Systemroutes.SetupRouter()
	r.GET("/", func(c *gin.Context) {
		log.Println("Welcome to API GO")
		c.JSON(200, gin.H{
			"message": "Hello!!!",
		})
	})
	// User_routes.Routes(r)
	// Product_routes.Routes(r)

	r.Run(":3000")
}
