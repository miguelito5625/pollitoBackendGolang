package Models

import (
	"fmt"
	"pollitoBackendGolang/Database"

	"gorm.io/gorm"
)

type Rol struct {
	gorm.Model
	Rol         int    `gorm:"uniqueIndex", json:"rol"`
	Descripcion string `json:"descripcion"`
}

func (b *Rol) TableName() string {
	return "rol"
}

func GetAllRol(b *[]Rol) (err error) {
	if err = Database.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewRol(b *Rol) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneRol(b *Rol, id string) (err error) {
	if err := Database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneRol(b *Rol, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteRol(b *Rol, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
