package Interfaces

type Credenciales struct {
	Cui   string `json:"cui"`
	Clave string `json:"clave"`
}

type DatosRegistroUsuario struct {
	Cui        string `json:"cui"`
	Nombres    string `json:"nombres"`
	Apellidos  string `json:"apellidos"`
	Nacimiento string `json:"nacimiento"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Rol_id     int    `json:"rol_id"`
}

type TokenUsuarioLoginData struct {
	Id   int
	Cui  string
	Name string
	Rol  int
}
