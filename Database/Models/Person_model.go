package Models

import (
	"fmt"
	"pollitoBackendGolang/Database"

	"gorm.io/gorm"
)

type Persona struct {
	gorm.Model
	Dpi             string `gorm:"unique", json:"dpi"`
	PrimerNombre    string `json:"primer_nombre"`
	SegundoNombre   string `json:"segundo_nombre"`
	PrimerApellido  string `json:"primer_apellido"`
	SegundoApellido string `json:"segundo_apellido"`
	Departamento    string `json:"departamento"`
	Municipio       string `json:"municipio"`
	Direccion       string `json:"direccion"`
}

func (b *Persona) TableName() string {
	return "persona"
}

func GetAllPersona(b *[]Persona) (err error) {
	if err = Database.DB.Joins("Rol").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewPersona(b *Persona) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOnePersona(b *Persona, id string) (err error) {
	if err := Database.DB.Joins("Rol").Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func LoginPersona(b *Persona, cui string) (err error) {
	if err := Database.DB.Joins("Rol").Where("cui = ?", cui).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOnePersona(b *Persona, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeletePersona(b *Persona, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
