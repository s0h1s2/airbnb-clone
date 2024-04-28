package listing

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/airbnb-clone/internal/common"
	"github.com/s0h1s2/airbnb-clone/internal/db"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	lisitngNotFoundErr        = errors.New("Listing not found.")
	acceptNumberIdOnlyErr     = errors.New("Only accept number only.")
	reservationExistErr       = errors.New("Reservation exist.")
	reservationDoesntExistErr = errors.New("Reservation doesn't exist.")
	favoriteExistErr          = errors.New("Favorite exist.")
)

func getListings(ctx *gin.Context) {
	clauses := make([]clause.Expression, 0)

	category, exists := ctx.GetQuery("category")
	if exists {
		clauses = append(clauses, clause.Like{Column: "category", Value: category})
	}
	listings := []db.Listing{}
	db.Db.Clauses(clauses...).Preload(clause.Associations).Order("ID desc").Find(&listings)
	response := &listingsResponse{}
	ctx.JSON(http.StatusOK, common.OkApiResponse{StatusCode: http.StatusOK, Data: response.Response(listings)})
	return
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
	queryResult := db.Db.Preload("User").Preload("Reservations").First(&result, "id=?", listingId)
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
	db.Db.Create(&db.ListingFavorite{ListingId: uint(listingId), UserId: user.Uid})
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
	user := common.GetUserClaimsFromContext(ctx)
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
	var listings []tripsResponse
	user := common.GetUserClaimsFromContext(ctx)
	db.Db.Raw("SELECT *,Reservations.id as reservation_id FROM Reservations INNER JOIN Listings ON Reservations.listing_id=Listings.id WHERE Reservations.user_id=? ORDER BY Reservations.created_at DESC", user.Uid).Find(&listings)
	ctx.JSON(http.StatusOK, gin.H{"data": listings})
}
func cancelReserve(ctx *gin.Context) {
	reservationId, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	userId := common.GetUserClaimsFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: acceptNumberIdOnlyErr.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	var reservation db.Reservation
	result := db.Db.First(&reservation, reservationId)
	println(result.Error.Error())
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: reservationDoesntExistErr.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	if reservation.UserId != userId.Uid {
		ctx.JSON(http.StatusUnauthorized, common.ErrorApiResponse{Errors: nil, StatusCode: http.StatusUnauthorized})
		return
	}
	db.Db.Unscoped().Delete(&reservation)
	ctx.JSON(http.StatusOK, common.OkApiResponse{StatusCode: http.StatusOK})
}
