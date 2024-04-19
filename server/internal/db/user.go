package db

import (
	"github.com/s0h1s2/airbnb-clone/internal/util"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string `json:"omitempty"`
	Listings     []Listing
	Reservations []Reservation
	Favorites    []ListingFavorite
}

func (User) TableName() string {
	return "users"
}

func (user *User) HashPassword() {
	hashedPassword := util.HashPassword(user.Password)
	user.Password = hashedPassword
}
