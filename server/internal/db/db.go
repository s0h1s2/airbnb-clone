package db

import (
	"github.com/s0h1s2/airbnb-clone/config"
	"github.com/s0h1s2/airbnb-clone/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func InitDb() {
	dbDsn := config.Config.Db.String()
	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	Db = db
	err = Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
