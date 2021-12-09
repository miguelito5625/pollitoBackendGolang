package Controllers

import (
	"log"
	"pollitoBackendGolang/ApiHelpers"
	"pollitoBackendGolang/Database/Models"
	"pollitoBackendGolang/Interfaces"
	"pollitoBackendGolang/Services"

	"github.com/gin-gonic/gin"
)

func IniciarSesion(c *gin.Context) {
	var credenciales Interfaces.Credenciales
	c.BindJSON(&credenciales)
	log.Println(credenciales)

	var usuario Models.Usuario
	err := Models.LoginUsuario(&usuario, credenciales.Cui)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, nil, "usuario no existe")
		return
	}

	hashClave := usuario.Clave
	log.Println("LA CLAVE ES: ", hashClave)

	claveCorrecta := Services.CheckPasswordHash(credenciales.Clave, hashClave)

	if !claveCorrecta {
		log.Println("Error autenticacion clave incorrecta")
		ApiHelpers.RespondJSON(c, 404, nil, "Error: clave incorrecta")
		return
	}

	var loginData Interfaces.UserLoginData

	loginData.Id = int(usuario.ID)
	loginData.Cui = usuario.Cui
	loginData.Name = usuario.Nombres + " " + usuario.Apeliidos
	loginData.Rol = usuario.Rol.Rol

	token, _ := Services.CreateToken(loginData)
	ApiHelpers.RespondJSON(c, 200, token, "Sesión Iniciada")

}
