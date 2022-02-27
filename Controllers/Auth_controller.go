package Controllers

import (
	"log"
	"pollitoBackendGolang/ApiHelpers"
	"pollitoBackendGolang/Database/Models"
	"pollitoBackendGolang/Interfaces"
	"pollitoBackendGolang/Services"

	"github.com/gin-gonic/gin"
)

func LoginUsuario(c *gin.Context) {
	var credenciales Interfaces.Credenciales
	c.BindJSON(&credenciales)

	log.Println("INICIANDO SESION")
	log.Println("CREDENCIALES:")
	log.Println(credenciales)

	var usuario Models.Usuario
	err := Models.SearchUserForLogin(&usuario, credenciales.Username)
	if err != nil {
		// log.Println(err)
		log.Println("Error autenticacion usuario no existe")
		ApiHelpers.RespondLoginJSON(c, 404, nil, "Usuario no existe")
		return
	}

	hashClave := usuario.Password

	claveCorrecta := Services.CheckPasswordHash(credenciales.Password, hashClave)

	if !claveCorrecta {
		log.Println("Error autenticacion clave incorrecta")
		ApiHelpers.RespondLoginJSON(c, 404, nil, "Contraseña incorrecta")
		return
	}

	var loginData Interfaces.TokenUsuarioLoginData

	loginData.Id = int(usuario.ID)
	loginData.Cui = usuario.Persona.Cui
	loginData.Nombres = usuario.Persona.Nombres
	loginData.Apellidos = usuario.Persona.Apellidos
	loginData.Username = usuario.Username
	loginData.Rol = usuario.Rol.Rol

	token, _ := Services.CreateToken(loginData)
	ApiHelpers.RespondLoginJSON(c, 200, token, "Sesion iniciada")

}
