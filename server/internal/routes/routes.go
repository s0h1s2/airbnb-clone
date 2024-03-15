package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/airbnb-clone/internal/middleware"
	"github.com/s0h1s2/airbnb-clone/internal/validation"
)

func RegisterRoutes(route gin.IRouter) {
	// users
	route.POST("/users", middleware.ValidateEndpoint(validation.CreateUserRequest{}), userRegisterRoute)
	route.POST("/users/auth", middleware.ValidateEndpoint(validation.LoginUserRequest{}), userAuthRoute)
}
