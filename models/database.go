package models

import (
	"os"

	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Secretkey string
	URLDB	string
)

func init() {
	Secretkey = os.Getenv("secretkey")
	URLDB = os.Getenv("Database")

	// Migración de la base de datos
	db := GetDatabase()
	db.AutoMigrate(&Usuario{})
	logs.Info("Migración completada")
}

func GetDatabase() *gorm.DB {
	//	database := "postgres"
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		logs.Error("DB Invalid database uri")
		return nil
	}

	logs.Info("Conexion exitosa a database")

	return db
}
