package db

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice float32
	UserId     uint
	ListingID  uint
}
