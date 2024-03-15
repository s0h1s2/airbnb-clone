package places

import (
	"gorm.io/gorm"
)

type PlaceModel struct {
	gorm.Model
	Title       string
	Address     string
	Photos      []PlacePhoto
	Perks       []PlacePerk
	ExtraInfo   string
	CheckIn     uint
	CheckOut    uint
	MaxGuests   uint
	UserModelID int
}
type PlacePhoto struct {
	gorm.Model
	PlaceModelID int
	Url          string
}
type PlacePerk struct {
	gorm.Model
	Url          string
	PlaceModelID int
}
