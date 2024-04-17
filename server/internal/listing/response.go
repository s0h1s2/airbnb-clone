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
	Id            uint              `json:"id"`
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	ImageSrc      string            `json:"imageSrc"`
	Category      string            `json:"category"`
	Roomcount     int               `json:"roomCount"`
	BathroomCount int               `json:"bathroomCount"`
	GuestCount    int               `json:"guestCount"`
	Location      locationResponse  `json:"location"`
	Country       string            `json:"country"`
	Price         float32           `json:"price"`
	Favorites     []ListingFavorite `json:"favorites" gorm:"unique"`
}
type listingResponse struct {
	Id          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	ImageSrc    string           `json:"imageSrc"`
	Name        string           `json:"username"`
	Country     string           `json:"country"`
	Location    locationResponse `json:"location"`
	Email       string           `json:"email"`
}
type listingWithOwnerInfo struct {
	Listing
	Name  string
	Email string
}

func (l *listingResponse) Response(data listingWithOwnerInfo) listingResponse {
	lat, lng, _ := strings.Cut(data.Location, ",")
	return listingResponse{
		Title:       data.Title,
		ImageSrc:    data.Imagesrc,
		Name:        data.Name,
		Email:       data.Email,
		Country:     data.Country,
		Location:    locationResponse{Lat: lat, Lng: lng},
		Description: data.Description,
	}
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
			Country:       listing.Country,
			Location:      locationResponse{Lat: lat, Lng: lng},
			Favorites:     listing.Favorites,
		}

		response = append(response, listingResponse)
	}
	return response
}
