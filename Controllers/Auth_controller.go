package Controllers

import (
	"pollitoBackendGolang/ApiHelpers"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	// var credenciales Interfaces.Credentials
	// c.BindJSON(&credenciales)

	// log.Println("INICIANDO SESION")
	// log.Println("CREDENCIALES:")
	// log.Println(credenciales)

	// var user Models.User
	// err := Models.SearchUserForLogin(&user, credenciales.Username)
	// if err != nil {
	// 	// log.Println(err)
	// 	log.Println("Error autenticacion user no existe")
	// 	ApiHelpers.RespondLoginJSON(c, 404, nil, "User no existe")
	// 	return
	// }

	// hashClave := user.Password

	// claveCorrecta := Services.CheckPasswordHash(credenciales.Password, hashClave)

	// if !claveCorrecta {
	// 	log.Println("Error autenticacion clave incorrecta")
	// 	ApiHelpers.RespondLoginJSON(c, 404, nil, "Contraseña incorrecta")
	// 	return
	// }

	// var loginData Interfaces.TokenUserLoginData

	// loginData.Id = int(user.ID)
	// loginData.Cui = user.Persona.Cui
	// loginData.Nombres = user.Persona.Nombres
	// loginData.Apellidos = user.Persona.Apellidos
	// loginData.Username = user.Username
	// loginData.Rol = user.Rol.Rol

	// token, _ := Services.CreateToken(loginData)
	ApiHelpers.RespondLoginJSON(c, 200, "token", "Sesion iniciada")

}
