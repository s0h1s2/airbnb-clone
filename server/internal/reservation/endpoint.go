package reservation

import "github.com/gin-gonic/gin"

func RegisterRoutes(router gin.IRouter) {
	router.POST("/reserve/:listingid/")
}
