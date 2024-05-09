package listing

import (
	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/airbnb-clone/internal/middleware"
)

func RegisterRoutes(route gin.IRouter) {
	route.GET("/listing", getListings)
	route.GET("/listing/properties", middleware.IsAuth(), getOwnerProperties)
	route.GET("/listing/trips", middleware.IsAuth(), getUserTrips)
	route.GET("/listing/reservations", middleware.IsAuth(), getUserReservations)
	route.GET("/listing/:id", getListingById)
	route.GET("/listing/favorites", middleware.IsAuth(), getUserFavorites)
	route.POST("/listing", middleware.IsAuth(), createNewListing)
	route.POST("/listing/:id/favorite", middleware.IsAuth(), favoriteListing)
	route.POST("/listing/:id/reserve", middleware.IsAuth(), reserveListing)
	route.DELETE("/listing/:id/cancel_reserve", middleware.IsAuth(), cancelReserve)
	route.DELETE("/listing/:id", middleware.IsAuth(), deleteProperty)

}
