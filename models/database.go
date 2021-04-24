package models

import (
	"net/url"

	"github.com/beego/beego/v2/core/logs"
	"github.com/jinzhu/gorm"


	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	Secretkey string = "Llavesecreta"
	URLDB	string
	dsn url.URL
)

func init() {
	dsn = url.URL{
		User: url.UserPassword("hdxxpvdpyekofa", "ff761b47ba1979b829199b70b4d4c59298c2337d6c2161082aa69f67be407543"),
		Scheme: "postgres",
		Host: "ec2-52-21-153-207.compute-1.amazonaws.com:"+"5432",
		Path: "d5c452ol1qfuha",
	}
	URLDB = "postgres://hdxxpvdpyekofa:ff761b47ba1979b829199b70b4d4c59298c2337d6c2161082aa69f67be407543@ec2-52-21-153-207.compute-1.amazonaws.com:5432/d5c452ol1qfuha"

}

func InitMigration() {
	// Migración de la base de datos
	connection := GetDatabase()
	connection.AutoMigrate(Usuario{})
	defer CloseDatabase(connection)
	logs.Info("Migración completada")
}

func GetDatabase() *gorm.DB {
	database := "postgres"

	connection, err := gorm.Open(database, dsn.String())
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
	}else {
		logs.Info("El ping se funciona db")
	}

	logs.Info("Conexion exitosa a database")

	return connection
}

func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
