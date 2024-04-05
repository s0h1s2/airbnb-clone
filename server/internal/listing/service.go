package listing

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/airbnb-clone/internal/common"
	"github.com/s0h1s2/airbnb-clone/internal/db"
)

func createNewListing(ctx *gin.Context) {
	var json createListingRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}
	user, err := ctx.Get("user")
	if !err {
		log.Fatal("User doesn't exist in context")
		return
	}
	userClaims := user.(common.UserClaims)
	listing := Listing{Title: json.Title, Price: json.Price, Category: json.Category, Imagesrc: json.ImageSrc, Location: json.Location, Roomcount: json.RoomCount, Description: json.Description, GuestCount: json.GuestCount, BathroomCount: json.BathroomCount, UserId: userClaims.Uid}
	result := db.Db.Create(&listing)
	if result.Error != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusCreated, common.OkApiResponse{Data: createListingResponse{}, StatusCode: http.StatusCreated})
}
