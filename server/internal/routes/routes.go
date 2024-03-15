package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(route gin.IRouter) {
	// users
	route.POST("/users", userRegisterRoute)
	route.POST("/users/auth", userAuthRoute)
}
