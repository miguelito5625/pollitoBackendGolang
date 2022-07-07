package Controllers

import (
	"log"
	"pollitoBackendGolang/ApiHelpers"
	"pollitoBackendGolang/Interfaces"

	"github.com/gin-gonic/gin"
)

func AddNewClient(c *gin.Context) {
	var datosCliente Interfaces.DatosRegistroCliente
	c.BindJSON(&datosCliente)
	log.Println("Datos Client:")
	log.Println(datosCliente)
	ApiHelpers.RespondLoginJSON(c, 200, datosCliente, "Client Creado")
}
