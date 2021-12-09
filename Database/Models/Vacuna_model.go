package Models

import (
	"fmt"
	"pollitoBackendGolang/Database"

	"gorm.io/gorm"
)

type Vacuna struct {
	gorm.Model
	Nombre string `json:"nombre"`
}

func (b *Vacuna) TableName() string {
	return "vacuna"
}

func GetAllVacuna(b *[]Vacuna) (err error) {
	if err = Database.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewVacuna(b *Vacuna) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneVacuna(b *Vacuna, id string) (err error) {
	if err := Database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneVacuna(b *Vacuna, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteVacuna(b *Vacuna, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
