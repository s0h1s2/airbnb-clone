package listing

import (
	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/airbnb-clone/internal/middleware"
)

func RegisterRoutes(route gin.IRouter) {
	route.GET("/listing",  getListings)
	route.POST("/listing", middleware.IsAuth(), createNewListing)
}
