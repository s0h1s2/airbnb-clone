package db

import "gorm.io/gorm"

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
	ListingId  uint `gorm:"unique" json:"-"`
	UserId     uint `gorm:"unique" json:"userId"`
}

func (l *Listing) Favorite(user uint) {
	Db.Create(&ListingFavorite{ListingId: l.ID, UserId: user})
}
