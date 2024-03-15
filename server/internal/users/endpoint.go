package users

import (
	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/airbnb-clone/internal/middleware"
)

func RegisterRoutes(route gin.IRouter) {
	route.POST("/users", userRegisterRoute)
	route.POST("/users/auth", userAuthRoute)
	route.POST("/users/logout", middleware.IsAuth(), userLogoutRoute)
	route.GET("/users/me", middleware.IsAuth(), userInfoRoute)
}
