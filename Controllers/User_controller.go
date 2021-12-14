package Controllers

import (
	"log"
	"pollitoBackendGolang/ApiHelpers"
	"pollitoBackendGolang/Database/Models"
	"pollitoBackendGolang/Interfaces"
	"pollitoBackendGolang/Services"

	"github.com/gin-gonic/gin"
)

func ListUsuario(c *gin.Context) {
	var usuario []Models.Usuario
	err := Models.GetAllUsuario(&usuario)
	if err != nil {
		log.Println("Error on list usuarios")
		ApiHelpers.RespondJSON(c, 404, usuario, "Error")
	} else {
		log.Println("Success list usuarios")
		ApiHelpers.RespondJSON(c, 200, usuario, "ok")
	}
}

func AddNewUsuario(c *gin.Context) {
	var datosRegistroUsuario Interfaces.DatosRegistroUsuario
	c.BindJSON(&datosRegistroUsuario)
	log.Println("Registrando usuario: ", datosRegistroUsuario)

	persona := Models.Persona{
		Cui:        datosRegistroUsuario.Cui,
		Nombres:    datosRegistroUsuario.Nombres,
		Apeliidos:  datosRegistroUsuario.Apellidos,
		Nacimiento: datosRegistroUsuario.Nacimiento,
	}

	hashedPassword, _ := Services.HashPassword(datosRegistroUsuario.Password)

	usuario := Models.Usuario{
		Username: datosRegistroUsuario.Username,
		Password: hashedPassword,
		Persona:  persona,
		Rol_id:   datosRegistroUsuario.Rol_id,
	}

	err := Models.AddNewUsuario(&usuario)
	if err != nil {
		log.Println("Error on insert usuario:", usuario)
		ApiHelpers.RespondJSON(c, 500, usuario, "Error")
	}

	log.Println("Usuario creado: ", usuario)

	ApiHelpers.RespondJSON(c, 200, nil, "ok")

	// usuario.Password, _ = Services.HashPassword(usuario.Password)
	// err := Models.AddNewUsuario(&usuario)
	// if err != nil {
	// 	log.Println("Error on insert usuario:", usuario)
	// 	ApiHelpers.RespondJSON(c, 404, usuario, "Error")
	// } else {
	// 	log.Println("Success inserted usuario:", usuario)
	// 	ApiHelpers.RespondJSON(c, 200, usuario, "ok")
	// }
}

func GetOneUsuario(c *gin.Context) {
	id := c.Params.ByName("id")
	var usuario Models.Usuario
	err := Models.GetOneUsuario(&usuario, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, nil, "Usuario no existe")
	} else {
		ApiHelpers.RespondJSON(c, 200, usuario, "ok")
	}
}

func PutOneUsuario(c *gin.Context) {
	var usuario Models.Usuario
	id := c.Params.ByName("id")
	err := Models.GetOneUsuario(&usuario, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usuario, "Error")
	}
	c.BindJSON(&usuario)
	err = Models.PutOneUsuario(&usuario, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usuario, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, usuario, "ok")
	}
}

func DeleteUsuario(c *gin.Context) {
	var usuario Models.Usuario
	id := c.Params.ByName("id")
	err := Models.DeleteUsuario(&usuario, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usuario, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, usuario, "ok")
	}
}
