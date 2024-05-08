package listing

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/airbnb-clone/internal/common"
	"github.com/s0h1s2/airbnb-clone/internal/db"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	lisitngNotFoundErr                       = errors.New("Listing not found.")
	acceptNumberIdOnlyErr                    = errors.New("Only accept number only.")
	reservationExistErr                      = errors.New("Reservation exist.")
	reservationDoesntExistErr                = errors.New("Reservation doesn't exist.")
	ownerCantMakeReservationOnOwnPropertyErr = errors.New("Owner can't make reservation on own property.")
	favoriteExistErr                         = errors.New("Favorite exist.")
	bindQueryErr                             = errors.New("Unable to bind query")
)

const listingLimit = 10

type listingsQuery struct {
	Page      int64           `form:"page,default=1"`
	StartDate *datatypes.Date `form:"startDate"`
	EndDate   *datatypes.Date `form:"endDate"`
	Category  string          `form:"category"`
	Country   string          `form:"country"`
	Guests    int             `form:"guests"`
	Rooms     int             `form:"rooms"`
	Bathrooms int             `form:"bathrooms"`
}

func getListings(ctx *gin.Context) {
	queryParams := listingsQuery{}
	if err := ctx.BindQuery(&queryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	listings := []db.Listing{}
	dbQuery := db.Db.Model(&db.Listing{}).Joins("LEFT JOIN Reservations ON Reservations.listing_id=listings.id").Preload("Favorites").Order("ID desc").Session(&gorm.Session{})

	if queryParams.Rooms > 0 {
		dbQuery = dbQuery.Where("listings.RoomCount = ?", queryParams.Rooms)
	}
	if queryParams.Bathrooms > 0 {
		dbQuery = dbQuery.Where("listings.bathroom_count = ?", queryParams.Bathrooms)
	}
	if queryParams.Guests > 0 {
		dbQuery = dbQuery.Where("listings.guest_count = ?", queryParams.Guests)
	}
	if queryParams.Country != "" {
		dbQuery = dbQuery.Where("listings.country= ?", queryParams.Country)
	}
	if queryParams.EndDate != nil {
		dbQuery = dbQuery.Where("NOT reservations.start_date>=? AND reservations.end_date<?", queryParams.StartDate, queryParams.EndDate)
	}
	if queryParams.Category != "" {
		dbQuery = dbQuery.Where("listings.category=?", queryParams.Category)
	}
	var recordCount int64
	dbQuery.Count(&recordCount)
	totalPages := int64(math.Ceil(float64(recordCount) / listingLimit))
	if queryParams.Page > totalPages {
		queryParams.Page = 1
	}

	offset := (queryParams.Page - 1) * listingLimit

	dbQuery.Limit(listingLimit).Offset(int(offset)).Find(&listings)

	response := &listingsResponse{}
	ctx.JSON(http.StatusOK, common.OkApiResponse{CurrentPage: queryParams.Page, TotalPages: totalPages, StatusCode: http.StatusOK, Data: response.Response(listings)})
}
func createNewListing(ctx *gin.Context) {
	var json createListingRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	userClaims := common.GetUserClaimsFromContext(ctx)
	listing := db.Listing{Title: json.Title, Price: json.Price, Category: json.Category, Imagesrc: json.ImageSrc, Location: json.Location.String(), Roomcount: json.RoomCount, Description: json.Description, GuestCount: json.GuestCount, BathroomCount: json.BathroomCount, UserId: userClaims.Uid, Country: json.Country}
	result := db.Db.Create(&listing)
	if result.Error != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusCreated, common.OkApiResponse{Data: createListingResponse{}, StatusCode: http.StatusCreated})
}
func getListingById(ctx *gin.Context) {
	listingId, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusNotFound, common.ErrorApiResponse{StatusCode: http.StatusNotFound, Errors: lisitngNotFoundErr.Error()})
		return
	}
	result := db.Listing{}
	queryResult := db.Db.Preload("User").Preload("Favorites").Preload("Reservations").First(&result, "id=?", listingId)
	if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, common.ErrorApiResponse{StatusCode: http.StatusNotFound, Errors: lisitngNotFoundErr.Error()})
		return
	}
	response := &listingResponse{}
	ctx.JSON(http.StatusOK, common.OkApiResponse{Data: response.Response(result)})
}
func favoriteListing(ctx *gin.Context) {
	listing := db.ListingFavorite{}
	user := common.GetUserClaimsFromContext(ctx)
	result := db.Db.Where("listing_id=? AND user_id=?", ctx.Param("id"), user.Uid).First(&listing)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{StatusCode: http.StatusBadRequest, Errors: favoriteExistErr.Error()})
		return
	}
	listingId, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	db.Db.Create(&db.ListingFavorite{ListingID: uint(listingId), UserId: user.Uid})
	ctx.JSON(http.StatusOK, common.OkApiResponse{})
}
func reserveListing(ctx *gin.Context) {
	var json createReserveRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	listingID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: acceptNumberIdOnlyErr.Error()})
		return
	}
	// Check if user is owner.
	user := common.GetUserClaimsFromContext(ctx)
	var isOwner bool
	db.Db.Model(&db.Listing{}).Where("id=? AND user_id=?", listingID, user.Uid).Select("count(*)>0").Find(&isOwner)
	if isOwner {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: ownerCantMakeReservationOnOwnPropertyErr.Error()})
		return
	}
	reservationExist := db.Reservation{}
	result := db.Db.Where("start_date=? AND end_date=?", json.StartDate.Format(time.DateOnly), json.EndDate.Format(time.DateOnly)).First(&reservationExist, "listing_id=?", uint(listingID))
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: reservationExistErr.Error()})
		return
	}
	reservation := db.Reservation{UserId: user.Uid, ListingID: uint(listingID), StartDate: datatypes.Date(json.StartDate), EndDate: datatypes.Date(json.EndDate)}
	record := db.Db.Create(&reservation)
	if record.Error != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: record.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, common.OkApiResponse{Data: nil})
}
func getUserTrips(ctx *gin.Context) {
	var trips []tripsResponse
	user := common.GetUserClaimsFromContext(ctx)
	db.Db.Raw("SELECT *,Reservations.id as reservation_id FROM Reservations INNER JOIN Listings ON Reservations.listing_id=Listings.id WHERE Reservations.user_id=? AND Reservations.start_date>=? ORDER BY Reservations.created_at DESC", user.Uid, datatypes.Date(time.Now())).Find(&trips)
	ctx.JSON(http.StatusOK, common.OkApiResponse{Data: trips, StatusCode: http.StatusOK})
}
func cancelReserve(ctx *gin.Context) {
	reservationId, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: acceptNumberIdOnlyErr.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	userId := common.GetUserClaimsFromContext(ctx)
	var reservation db.Reservation
	result := db.Db.First(&reservation, reservationId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: reservationDoesntExistErr.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	// check if current user is owner of the listing or not
	var listing db.Listing
	listingResult := db.Db.Find(&listing, "id=?", reservation.ListingID)
	if errors.Is(listingResult.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: lisitngNotFoundErr.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	if reservation.UserId != userId.Uid && listing.UserId != userId.Uid {
		ctx.JSON(http.StatusUnauthorized, common.ErrorApiResponse{Errors: nil, StatusCode: http.StatusUnauthorized})
		return
	}
	db.Db.Unscoped().Delete(&reservation)
	ctx.JSON(http.StatusOK, common.OkApiResponse{StatusCode: http.StatusOK})
}

// / Get current user reservation on properties.
func getUserReservations(ctx *gin.Context) {
	var reservations []tripsResponse
	user := common.GetUserClaimsFromContext(ctx)
	db.Db.Raw("SELECT *,Reservations.id as reservation_id  FROM Reservations INNER JOIN Listings ON Reservations.listing_id=Listings.id WHERE Listings.user_id=? AND Reservations.start_date>=?", user.Uid, datatypes.Date(time.Now())).Scan(&reservations)
	ctx.JSON(http.StatusOK, common.OkApiResponse{StatusCode: http.StatusOK, Data: reservations})
}
func getUserFavorites(ctx *gin.Context) {
	var listings []favoritesResponse
	user := common.GetUserClaimsFromContext(ctx)
	db.Db.Raw("SELECT * FROM listing_favorites INNER JOIN listings ON listing_favorites.listing_id=listings.id WHERE listing_favorites.user_id=?", user.Uid).Scan(&listings)
	ctx.JSON(http.StatusOK, common.OkApiResponse{Data: listings, StatusCode: http.StatusOK})
}
