package db

import (
	"github.com/s0h1s2/airbnb-clone/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func InitDb() {
	dbDsn := os.Getenv("DB_DSN")
	Db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	err = Db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

}
