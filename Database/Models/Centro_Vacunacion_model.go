package Models

import (
	"fmt"
	"pollitoBackendGolang/Database"

	"gorm.io/gorm"
)

type Centro_Vacunacion struct {
	gorm.Model
	Nombre string `json:"nombre"`
}

func (b *Centro_Vacunacion) TableName() string {
	return "centro_vacunacion"
}

func GetAllCentro_Vacunacion(b *[]Centro_Vacunacion) (err error) {
	if err = Database.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewCentro_Vacunacion(b *Centro_Vacunacion) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCentro_Vacunacion(b *Centro_Vacunacion, id string) (err error) {
	if err := Database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneCentro_Vacunacion(b *Centro_Vacunacion, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteCentro_Vacunacion(b *Centro_Vacunacion, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
