package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

const cost = bcrypt.DefaultCost

func HashPassword(raw string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(raw), cost)

	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}

func VerifyPassword(hashedPassword string, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	if err != nil {
		return false
	}
	return true
}
