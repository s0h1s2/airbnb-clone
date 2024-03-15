package main

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/s0h1s2/airbnb-clone/config"
	"github.com/s0h1s2/airbnb-clone/internal/db"
	"github.com/s0h1s2/airbnb-clone/internal/routes"
)

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
	r := gin.Default()
	v1 := r.Group("/api/v1")
	routes.RegisterRoutes(v1)

	r.Run()
}
