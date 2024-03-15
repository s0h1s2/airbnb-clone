package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}
	}
}
