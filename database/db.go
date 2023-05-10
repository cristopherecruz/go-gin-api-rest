package database

import (
	"github.com/cristopherecruz/go-gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB
var Err error

func ConectarBancoDados() {

	connStr := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/SAO_PAULO"

	Db, Err = gorm.Open(postgres.Open(connStr))
	if Err != nil {
		log.Panic(Err.Error())
	}

	err := Db.AutoMigrate(&models.Aluno{})
	if err != nil {
		log.Panic(Err.Error())
	}

}
