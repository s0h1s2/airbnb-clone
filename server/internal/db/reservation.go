package db

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model `json:"-"`
	StartDate  time.Time `json:"start"`
	EndDate    time.Time `json:"end"`
	TotalPrice float32
	UserId     uint
	ListingID  uint
}
