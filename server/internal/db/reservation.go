package db

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model `json:"-"`
	StartDate  time.Time `json:"start" gorm:"unique"`
	EndDate    time.Time `json:"end" gorm:"unique"`
	TotalPrice float32
	UserId     uint
	ListingID  uint
}
