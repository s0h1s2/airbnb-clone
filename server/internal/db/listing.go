package db

import (
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
	User          User
	Reservations  []Reservation
	Favorites     []ListingFavorite
}

type ListingFavorite struct {
	gorm.Model `json:"-"`
	ListingID  uint `json:"-"`
	Listing    Listing
	UserId     uint `json:"userId"`
}

func (l *Listing) Favorite(user uint) {
	Db.Create(&ListingFavorite{ListingID: l.ID, UserId: user})
}
