package listing

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/airbnb-clone/internal/common"
	"github.com/s0h1s2/airbnb-clone/internal/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	lisitngNotFoundErr    = errors.New("Listing not found.")
	acceptNumberIdOnlyErr = errors.New("Only accept number only.")
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
	result := db.Listing{}
	queryResult := db.Db.Preload("User").Preload("Reservations").First(&result, ctx.Param("id"))
	if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, common.ErrorApiResponse{StatusCode: http.StatusNotFound, Errors: lisitngNotFoundErr.Error()})
		return
	}
	response := &listingResponse{}
	ctx.JSON(http.StatusOK, common.OkApiResponse{Data: response.Response(result)})
}
func favoriteListing(ctx *gin.Context) {
	listing := db.Listing{}
	result := db.Db.Where("id=?", ctx.Param("id")).First(&listing)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{StatusCode: http.StatusBadRequest, Errors: lisitngNotFoundErr.Error()})
		return
	}
	user := common.GetUserClaimsFromContext(ctx)

	listing.Favorite(user.Uid)
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
	reservation := db.Reservation{UserId: user.Uid, ListingID: uint(listingID), StartDate: json.StartDate, EndDate: json.EndDate}
	record := db.Db.Create(&reservation)
	if record.Error != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: record.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, common.OkApiResponse{Data: nil})

}
