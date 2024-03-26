package listing

import (
	"github.com/s0h1s2/airbnb-clone/internal/reservation"
	"gorm.io/gorm"
)

type Listing struct {
	gorm.Model
	Title         string
	Description   string
	Imagesrc      string
	Category      string
	Roomcount     int
	BathroomCount int
	GuestCount    int
	Location      string
	Price         float32
	UserId        uint
	Reservations  []reservation.Reservation
}
