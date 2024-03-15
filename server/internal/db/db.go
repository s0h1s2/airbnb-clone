package db

import (
	"log"

	"github.com/s0h1s2/airbnb-clone/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	dbDsn := config.Config.Db.String()
	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{})
	Db = db
	if err != nil {
		log.Fatal(err.Error())
	}
}
