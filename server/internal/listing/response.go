package listing

import (
	"strings"

	"github.com/s0h1s2/airbnb-clone/internal/db"
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
	Id            uint                 `json:"id"`
	Title         string               `json:"title"`
	Description   string               `json:"description"`
	ImageSrc      string               `json:"imageSrc"`
	Category      string               `json:"category"`
	Roomcount     int                  `json:"roomCount"`
	BathroomCount int                  `json:"bathroomCount"`
	GuestCount    int                  `json:"guestCount"`
	Location      locationResponse     `json:"location"`
	Country       string               `json:"country"`
	Price         float32              `json:"price"`
	Favorites     []db.ListingFavorite `json:"favorites" gorm:"unique"`
}
type userResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type listingResponse struct {
	Id            uint             `json:"id"`
	Title         string           `json:"title"`
	Description   string           `json:"description"`
	ImageSrc      string           `json:"imageSrc"`
	Country       string           `json:"country"`
	Location      locationResponse `json:"location"`
	GuestCount    int              `json:"guestCount"`
	RoomCount     int              `json:"roomCount"`
	BathroomCount int              `json:"bathroomCount"`
	Category      string           `json:"category"`
	Price         float32          `json:"price"`
	Reservations  []db.Reservation `json:"reservations"`
	User          userResponse     `json:"user"`
}
type tripsResponse struct {
	ID         uint    `json:"listingId"`
	StartDate  string  `json:"startDate"`
	EndDate    string  `json:"endDate"`
	Title      string  `json:"title"`
	TotalPrice float32 `json:"price"`
	Country    string  `json:"country"`
	Imagesrc   string  `json:"imageSrc"`
	Category   string  `json:"category"`
}

func (l *listingResponse) Response(data db.Listing) listingResponse {
	lat, lng, _ := strings.Cut(data.Location, ",")
	return listingResponse{
		Id:            data.ID,
		Title:         data.Title,
		ImageSrc:      data.Imagesrc,
		Country:       data.Country,
		Location:      locationResponse{Lat: lat, Lng: lng},
		BathroomCount: data.BathroomCount,
		GuestCount:    data.GuestCount,
		RoomCount:     data.Roomcount,
		Description:   data.Description,
		Category:      data.Category,
		Reservations:  data.Reservations,
		Price:         data.Price,
		User:          userResponse{Name: data.User.Name, Email: data.User.Email},
	}
}
func (l *listingsResponse) Response(listings []db.Listing) (response []listingsResponse) {
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
