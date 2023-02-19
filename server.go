package main

import (
	"log"
	"os"
	"pollitoBackendGolang/Systemroutes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Database.ConnectDataBase()
	// Migrations.MigrateAll()

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

	location, err := time.LoadLocation("America/Guatemala")
	if err != nil {
		panic(err)
	}
	currentTime := time.Now().In(location)
	fmt.Println("ESTO ES UNA PRUEBA")

	dbIP, ok := os.LookupEnv("DB_IP")
	if !ok {
		// La variable de entorno no está definida
		dbIP = "error"
	}

	r.GET("/", func(c *gin.Context) {
		fmt.Println("ESTO ES UNA PRUEBA2")
		c.JSON(200, gin.H{
			"message": "Mike Hello!!! " + os.Getenv("TEST_MESSAGE"),
			"DB_IP":   dbIP,
			"DATE: ":  currentTime.Format("2006-01-02 15:04:05"),
		})
	})
	// Usuario_routes.Routes(r)
	// Product_routes.Routes(r)

	//Para subir ha app engine hay que eliminar el puerto que se usa localmente
	r.Run()

	// r.Run(":3000")
}
