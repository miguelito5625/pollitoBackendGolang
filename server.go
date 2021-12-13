package main

import (
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

	//No utilizar en app engine//////////////////////////////////////////////////////////////////////////////////////////

	// logfile, errFile := os.OpenFile("server.log", os.O_RDWR|os.O_APPEND, 0660)
	// if errFile != nil {
	// 	logfile, _ = os.Create("server.log")
	// 	// log.Println("Error al abrir el archivo server.log")
	// }

	// logfile, _ := os.Create("server.log")
	// gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
	// log.SetOutput(gin.DefaultWriter)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	log.Println("")
	log.Println("Servidor iniciado")

	r := Systemroutes.SetupRouter()
	r.GET("/", func(c *gin.Context) {
		log.Println("Welcome to API GO")
		c.JSON(200, gin.H{
			"message": "Hello!!! " + os.Getenv("TEST_MESSAGE"),
		})
	})
	// User_routes.Routes(r)
	// Product_routes.Routes(r)

	//Para subir ha app engine hay que eliminar el puerto que se usa localmente
	// r.Run()

	r.Run(":3000")
}
