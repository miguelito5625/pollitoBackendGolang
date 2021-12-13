package Models

import (
	"fmt"
	"pollitoBackendGolang/Database"

	"gorm.io/gorm"
)

type Persona struct {
	gorm.Model
	Cui        string `json:"cui"`
	Nombres    string `json:"nombres"`
	Apeliidos  string `json:"apellidos"`
	Nacimiento string `json:"nacimiento"`
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
