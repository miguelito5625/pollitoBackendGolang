package Controllers

import (
	"log"
	"pollitoBackendGolang/ApiHelpers"
	"pollitoBackendGolang/Database/Models"
	"pollitoBackendGolang/Interfaces"

	"github.com/gin-gonic/gin"
)

func AddNewClient(c *gin.Context) {

	var datosCliente Interfaces.DatosRegistroCliente
	c.BindJSON(&datosCliente)
	log.Println("DATOS DEL CLIENTE:")
	log.Println(datosCliente)

	persona := Models.Persona{
		Dpi:             datosCliente.Dpi,
		PrimerNombre:    datosCliente.PrimerNombre,
		SegundoNombre:   datosCliente.SegundoNombre,
		PrimerApellido:  datosCliente.PrimerApellido,
		SegundoApellido: datosCliente.SegundoApellido,
		Departamento:    datosCliente.Departamento,
		Municipio:       datosCliente.Municipio,
		Direccion:       datosCliente.Direccion,
	}

	cliente := Models.Cliente{
		Indicaciones: datosCliente.Indicaciones,
		Persona:      persona,
	}

	err := Models.AddNewCliente(&cliente)
	if err != nil {
		log.Println("Error on insert cliente:", err)
		ApiHelpers.RespondJSON(c, 500, cliente, "Error al intentar crear el cliente")
	} else {
		log.Println("Cliente creado: ", cliente)
		ApiHelpers.RespondJSON(c, 200, cliente, "Cliente creado")
	}

}

func ListClient(c *gin.Context) {
	var clientes []Models.Cliente
	err := Models.GetAllCliente(&clientes)
	if err != nil {
		log.Println("Error on list clientes")
		ApiHelpers.RespondJSON(c, 404, clientes, "Error")
	} else {
		log.Println("Success list clientes")
		ApiHelpers.RespondJSON(c, 200, clientes, "ok")
	}
}
