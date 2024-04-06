package listing

import "fmt"

type location struct {
	lat float32
	lng float32
}

func (loc location) String() string {
	return fmt.Sprintf("%f,%f", loc.lat, loc.lng)
}

type createListingRequest struct {
	Category      string   `json:"category" binding:"required"`
	Title         string   `json:"title" binding:"required"`
	Description   string   `json:"description" binding:"required"`
	ImageSrc      string   `json:"imageSrc" binding:"required,url"`
	RoomCount     int      `json:"roomCount" binding:"required,number"`
	BathroomCount int      `json:"bathroomCount" binding:"required,number"`
	GuestCount    int      `json:"guestCount" binding:"required,number"`
	Price         float32  `json:"price" binding:"required,numeric"`
	Location      location `json:"location" binding:"required"`
}
