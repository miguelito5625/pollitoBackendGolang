package Models

import (
	"fmt"
	"pollitoBackendGolang/Database"

	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Persona_id int     `json:"persona_id"`
	Persona    Persona `gorm:"foreignkey:Persona_id"`
	Rol_id     int     `json:"rol_id"`
	Rol        Rol     `gorm:"foreignkey:Rol_id"`
}

func (b *Usuario) TableName() string {
	return "usuario"
}

func GetAllUsuario(b *[]Usuario) (err error) {
	if err = Database.DB.Joins("Rol").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUsuario(b *Usuario) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUsuario(b *Usuario, id string) (err error) {
	if err := Database.DB.Joins("Rol").Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func LoginUsuario(b *Usuario, cui string) (err error) {
	if err := Database.DB.Joins("Rol").Where("cui = ?", cui).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUsuario(b *Usuario, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteUsuario(b *Usuario, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
