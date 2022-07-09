package Interfaces

type DatosRegistroCliente struct {
	Dpi             string `json:"dpi"`
	PrimerNombre    string `json:"primer_nombre"`
	SegundoNombre   string `json:"segundo_nombre"`
	PrimerApellido  string `json:"primer_apellido"`
	SegundoApellido string `json:"segundo_apellido"`
	Departamento    string `json:"departamento"`
	Municipio       string `json:"municipio"`
	Direccion       string `json:"direccion"`
	Indicaciones    string `json:"indicaciones"`
}
