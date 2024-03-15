package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/s0h1s2/airbnb-clone/config"
	"github.com/s0h1s2/airbnb-clone/internal/common"
)

func IsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		claims := common.UserClaims{}
		tk, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Config.Jwt.JwtSecretKey), nil
		})
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if !tk.Valid {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		ctx.Set("user", claims)
		ctx.Next()
	}
}
