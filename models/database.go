package models

import (
	"os"

	"github.com/beego/beego/v2/core/logs"
	"github.com/jinzhu/gorm"
)

var (
	Secretkey string = "Llavesecreta"
	URLDB	string
)

func init() {
	URLDB = os.Getenv("DATABASE_URL")
	// Migración de la base de datos
	connection := GetDatabase()
	defer CloseDatabase(connection)
	connection.AutoMigrate(Usuario{})
}

func GetDatabase() *gorm.DB {
	database := "postgres"

	connection, err := gorm.Open(database, URLDB)
	if err != nil {
		logs.Error("DB Invalid database uri")
		return nil
	}

	sqldb := connection.DB()
	defer sqldb.Close()
	err = sqldb.Ping()
	if err != nil {
		logs.Error("Error al hacer ping a db")
		return nil
	}

	logs.Info("Conexion exitosa a database")

	return connection
}

func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
