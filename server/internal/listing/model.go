package listing

import (
	"github.com/s0h1s2/airbnb-clone/internal/db"
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
	Country       string
	UserId        uint
	Reservations  []reservation.Reservation
	Favorites     []ListingFavorite
}

type ListingFavorite struct {
	gorm.Model
	ListingId uint `gorm:"uniqueIndex"`
	UserId    uint `gorm:"uniqueIndex"`
}

func (l *Listing) favorite(user uint) {
	db.Db.Create(&ListingFavorite{ListingId: l.ID, UserId: user})
}
