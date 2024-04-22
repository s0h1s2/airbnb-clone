package reservation

import "time"

type createReserveRequest struct {
	StartDate time.Time `json:"startDate" binding:"required,datetime"`
	EndDate   time.Time `json:"endDate" binding:"required,datetime"`
}
