package main

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/s0h1s2/airbnb-clone/config"
	"github.com/s0h1s2/airbnb-clone/internal/db"
	"github.com/s0h1s2/airbnb-clone/internal/listing"
	"github.com/s0h1s2/airbnb-clone/internal/users"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Config.ForntendProxy)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func migrateTables() {
	db.Db.AutoMigrate(&db.User{})
	db.Db.AutoMigrate(&db.Listing{})
	db.Db.AutoMigrate(&db.ListingFavorite{})
	db.Db.AutoMigrate(&db.Reservation{})
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load envs from .env file.")
	}
	if err = env.Parse(&config.Config); err != nil {
		log.Fatal(err)
		return
	}
	db.InitDb()
	migrateTables()
	r := gin.Default()
	r.Use(CORSMiddleware())
	v1 := r.Group("/api/v1/")
	users.RegisterRoutes(v1)
	listing.RegisterRoutes(v1)
	r.Run()
}
