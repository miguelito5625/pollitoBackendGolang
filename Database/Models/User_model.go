package Models

import (
	"fmt"
	"pollitoBackendGolang/Database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Persona_id int     `json:"persona_id"`
	Persona    Persona `gorm:"foreignkey:Persona_id"`
	Rol_id     int     `json:"rol_id"`
	Rol        Rol     `gorm:"foreignkey:Rol_id"`
}

func (b *User) TableName() string {
	return "user"
}

func GetAllUser(b *[]User) (err error) {
	if err = Database.DB.Joins("Rol").Joins("Persona").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUser(b *User) (err error) {
	if err = Database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(b *User, id string) (err error) {
	if err := Database.DB.Joins("Rol").Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func SearchUserForLogin(b *User, username string) (err error) {
	if err := Database.DB.Joins("Rol").Joins("Persona").Where("username = ?", username).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(b *User, id string) (err error) {
	fmt.Println(b)
	Database.DB.Save(b)
	return nil
}

func DeleteUser(b *User, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(b)
	return nil
}
