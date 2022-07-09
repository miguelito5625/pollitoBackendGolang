package Models

import (
	"fmt"
	"pollitoBackendGolang/Database"

	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Indicaciones string  `json:"indicaciones"`
	Persona_id   int     `json:"persona_id"`
	Persona      Persona `gorm:"foreignkey:Persona_id"`
}

func (b *Cliente) TableName() string {
	return "cliente"
}

func GetAllCliente(b *[]Cliente) (err error) {
	if err = Database.DB.Joins("Persona").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewCliente(b *Cliente) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCliente(b *Cliente, id string) (err error) {
	if err := Database.DB.Joins("Persona").Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneCliente(b *Cliente, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteCliente(b *Cliente, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
