package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateEndpoint(rules interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data interface{}
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Set("body", data)
		ctx.Next()
		return
	}
}
