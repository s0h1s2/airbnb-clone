package listing

import (
	"strings"
)

type createListingResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type locationResponse struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
type listingsResponse struct {
	Id            uint             `json:"id"`
	Title         string           `json:"title"`
	Description   string           `json:"description"`
	ImageSrc      string           `json:"imageSrc"`
	Category      string           `json:"category"`
	Roomcount     int              `json:"roomCount"`
	BathroomCount int              `json:"bathroomCount"`
	GuestCount    int              `json:"guestCount"`
	Location      locationResponse `json:"location"`
	Price         float32          `json:"price"`
}

func (l *listingsResponse) Response(listings []Listing) (response []listingsResponse) {
	for _, listing := range listings {
		lat, lng, _ := strings.Cut(listing.Location, ",")
		listingResponse := listingsResponse{
			Id:            listing.ID,
			Title:         listing.Title,
			Description:   listing.Description,
			ImageSrc:      listing.Imagesrc,
			Price:         listing.Price,
			Category:      listing.Category,
			Roomcount:     listing.Roomcount,
			BathroomCount: listing.BathroomCount,
			GuestCount:    listing.GuestCount,
			Location:      locationResponse{Lat: lat, Lng: lng},
		}

		response = append(response, listingResponse)
	}
	return response
}
