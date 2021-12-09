package Interfaces

type Credenciales struct {
	Cui   string `json:"cui"`
	Clave string `json:"clave"`
}

type UserLoginData struct {
	Id   int
	Cui  string
	Name string
	Rol  int
}
