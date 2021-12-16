package Database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	DB_IP := os.Getenv("DB_IP")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_IP + ":3306)/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DATABASE CONNECT FAILED, status: ", err)
	}

	DB = db

}
