package Controllers

import (
	"log"
	"pollitoBackendGolang/ApiHelpers"
	"pollitoBackendGolang/Database/Models"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUser(&user)
	if err != nil {
		log.Println("Error on list users")
		ApiHelpers.RespondJSON(c, 404, user, "Error")
	} else {
		log.Println("Success list users")
		ApiHelpers.RespondJSON(c, 200, user, "ok")
	}
}

func AddNewUser(c *gin.Context) {
	// var datosRegistroUser Interfaces.DatosRegistroUser
	// c.BindJSON(&datosRegistroUser)

	// persona := Models.Persona{
	// 	Cui:        datosRegistroUser.Cui,
	// 	Nombres:    datosRegistroUser.Nombres,
	// 	Apellidos:  datosRegistroUser.Apellidos,
	// 	Nacimiento: datosRegistroUser.Nacimiento,
	// }

	// hashedPassword, _ := Services.HashPassword(datosRegistroUser.Password)

	// user := Models.User{
	// 	Username: datosRegistroUser.Username,
	// 	Password: hashedPassword,
	// 	Persona:  persona,
	// 	Rol_id:   datosRegistroUser.Rol_id,
	// }

	// err := Models.AddNewUser(&user)
	// if err != nil {
	// 	log.Println("Error on insert user:", err)
	// 	ApiHelpers.RespondJSON(c, 500, user, "Error al intentar crear el user")
	// } else {
	// 	log.Println("User creado: ", user)
	// 	ApiHelpers.RespondJSON(c, 200, user, "User creado")
	// }

	ApiHelpers.RespondJSON(c, 200, "user", "User creado")

}

func GetOneUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetOneUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, nil, "User no existe")
	} else {
		ApiHelpers.RespondJSON(c, 200, user, "ok")
	}
}

func PutOneUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetOneUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user, "Error")
	}
	c.BindJSON(&user)
	err = Models.PutOneUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, user, "ok")
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user, "Error")
	} else {
		ApiHelpers.RespondJSON(c, 200, user, "ok")
	}
}
