package models

import (
	"os"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Secretkey string
	URLDB	string
	typedb string
)

func init() {
	typedb, _ := web.AppConfig.String("typedb")
	logs.Info("Tipo de base de datos a utilizar: ", typedb)
	Secretkey = os.Getenv("secretkey")
	if Secretkey == "" {
		Secretkey, _ = web.AppConfig.String("secretkey")
	}
	URLDB = os.Getenv("DATABASE_URL")

	// Migración de la base de datos
	db := GetDatabase()
	db.AutoMigrate(&Usuario{}, &Contacto{}, &Message{}, &Conversación{})
}

func GetDatabase() *gorm.DB {
	//	database := "postgres"
	var db *gorm.DB
	var err error
	// Controlamos el tipo de base de datos a utilizar
	if typedb == "postgres" {
		db, err = gorm.Open(postgres.Open(URLDB), &gorm.Config{})
	} else if typedb == "sqlite" || typedb == "" {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}
	if err != nil {
		logs.Error("DB Invalid database uri: ", err.Error())
		return nil
	}
	return db
}
