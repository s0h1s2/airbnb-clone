package db

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model `json:"-"`
	StartDate  datatypes.Date `json:"start" gorm:"unique"`
	EndDate    datatypes.Date `json:"end" gorm:"unique"`
	TotalPrice float32
	UserId     uint
	ListingID  uint
}
