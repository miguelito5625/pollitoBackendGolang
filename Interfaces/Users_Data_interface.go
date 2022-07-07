package Interfaces

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DatosRegistroUser struct {
	Cui        string `json:"cui"`
	Nombres    string `json:"nombres"`
	Apellidos  string `json:"apellidos"`
	Nacimiento string `json:"nacimiento"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Rol_id     int    `json:"rol_id"`
}

type TokenUserLoginData struct {
	Id        int
	Cui       string
	Nombres   string
	Apellidos string
	Username  string
	Rol       int
}
