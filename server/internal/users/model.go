package users

import (
	"time"

	"github.com/s0h1s2/airbnb-clone/internal/listing"
	"github.com/s0h1s2/airbnb-clone/internal/reservation"
	"github.com/s0h1s2/airbnb-clone/internal/util"
)

type User struct {
	ID           uint
	Name         string
	Email        string
	Password     string `json:"omitempty"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Listings     []listing.Listing
	Reservations []reservation.Reservation
}

func (User) TableName() string {
	return "users"
}

func (user *User) HashPassword() {
	hashedPassword := util.HashPassword(user.Password)
	user.Password = hashedPassword
}
