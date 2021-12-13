package Migrations

import (
	"pollitoBackendGolang/Database"
	"pollitoBackendGolang/Database/Models"
)

func MigrateAll() {
	Database.DB.AutoMigrate(&Models.Rol{})
	Database.DB.AutoMigrate(&Models.Usuario{})
}
