package Controllers

import (
	"log"
	"pollitoBackendGolang/ApiHelpers"
	"pollitoBackendGolang/Database/Models"

	"github.com/gin-gonic/gin"
)

func ListRol(c *gin.Context) {
	var rol []Models.Rol
	err := Models.GetAllRol(&rol)
	if err != nil {
		log.Println("Error on list rols")
		ApiHelpers.RespondJSON(c, 404, rol, "Error")
	} else {
		log.Println("Success list rols")
		ApiHelpers.RespondJSON(c, 200, rol, "ok")
	}
}

func AddNewRol(c *gin.Context) {
	var rol Models.Rol
	c.BindJSON(&rol)
	err := Models.AddNewRol(&rol)
	if err != nil {
		log.Println("Error on insert rol:", rol)
		ApiHelpers.RespondJSON(c, 404, rol, "Error")
	} else {
		log.Println("Success inserted rol:", rol)
		ApiHelpers.RespondJSON(c, 200, rol, "ok")
	}
}

func GetOneRol(c *gin.Context) {
	id := c.Params.ByName("id")
	var rol Models.Rol
	err := Models.GetOneRol(&rol, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, rol, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, rol, "ok")
	}
}

func PutOneRol(c *gin.Context) {
	var rol Models.Rol
	id := c.Params.ByName("id")
	err := Models.GetOneRol(&rol, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, rol, "Error")
	}
	c.BindJSON(&rol)
	err = Models.PutOneRol(&rol, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, rol, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, rol, "ok")
	}
}

func DeleteRol(c *gin.Context) {
	var rol Models.Rol
	id := c.Params.ByName("id")
	err := Models.DeleteRol(&rol, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, rol, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, rol, "ok")
	}
}
