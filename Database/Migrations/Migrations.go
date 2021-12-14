package Migrations

import (
	"log"
	"pollitoBackendGolang/Database"
	"pollitoBackendGolang/Database/Models"
)

func MigrateAll() {
	Database.DB.AutoMigrate(&Models.Rol{})
	Database.DB.AutoMigrate(&Models.Usuario{})
	insertData()
}

func insertData() {
	rol := Models.Rol{
		Rol:         1,
		Descripcion: "Admin",
	}

	err := Models.AddNewRol(&rol)
	if err != nil {
		log.Println("Error on insert rol:", err)
	} else {
		log.Println("Rol creado: ", rol)
	}

}
