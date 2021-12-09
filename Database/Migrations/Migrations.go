package Migrations

import (
	"pollitoBackendGolang/Database"
	"pollitoBackendGolang/Database/Models"
)

func MigrateAll() {
	Database.DB.AutoMigrate(&Models.Rol{})
	Database.DB.AutoMigrate(&Models.Usuario{})
	// Database.DB.AutoMigrate(&Models.Vacuna{})
	// Database.DB.AutoMigrate(&Models.Centro_Vacunacion{})
	// Database.DB.AutoMigrate(&Models.Grupo_Vacunacion{})
}
