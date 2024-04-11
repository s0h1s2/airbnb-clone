package common

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Uid   uint   `json:"uid"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func GetUserClaimsFromContext(ctx *gin.Context) *UserClaims {
	user, err := ctx.Get("user")
	if !err {
		log.Fatal("User doesn't exist in context")
		return nil
	}

	userClaims := user.(UserClaims)
	return &userClaims
}
