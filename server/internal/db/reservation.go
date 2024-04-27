package db

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model `json:"-"`
	StartDate  datatypes.Date `json:"start"`
	EndDate    datatypes.Date `json:"end"`
	TotalPrice float32
	UserId     uint `json:"-"`
	ListingID  uint `json:"-"`
}
