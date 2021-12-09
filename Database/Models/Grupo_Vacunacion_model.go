package Models

import (
	"fmt"
	"pollitoBackendGolang/Database"
	"time"

	"gorm.io/gorm"
)

type Grupo_Vacunacion struct {
	gorm.Model
	Fecha1               time.Time         `json:"fecha1"`
	Fecha2               time.Time         `json:"fecha2"`
	Fecha3               time.Time         `json:"fecha3"`
	Centro_Vacunacion_id int               `json:"centro_vacunacion_id"`
	Usuario_id           int               `json:"usuario_id"`
	Vacuna_id            int               `json:"vacuna_id"`
	Centro_Vacunacion    Centro_Vacunacion `gorm:"foreignkey:Centro_Vacunacion_id"`
	Usuario              Usuario           `gorm:"foreignkey:Usuario_id"`
	Vacuna               Vacuna            `gorm:"foreignkey:Vacuna_id"`
}

func (b *Grupo_Vacunacion) TableName() string {
	return "grupo_vacunacion"
}

func GetAllGrupo_Vacunacion(b *[]Grupo_Vacunacion) (err error) {
	if err = Database.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewGrupo_Vacunacion(b *Grupo_Vacunacion) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneGrupo_Vacunacion(b *Grupo_Vacunacion, id string) (err error) {
	if err := Database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneGrupo_Vacunacion(b *Grupo_Vacunacion, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteGrupo_Vacunacion(b *Grupo_Vacunacion, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
