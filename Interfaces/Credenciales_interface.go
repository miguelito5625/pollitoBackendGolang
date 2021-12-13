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
	Rol        int    `json:"tol"`
}

type TokenUsuarioLoginData struct {
	Id   int
	Cui  string
	Name string
	Rol  int
}
