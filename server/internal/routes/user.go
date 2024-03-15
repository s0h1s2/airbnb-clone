package routes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/s0h1s2/airbnb-clone/config"
	"github.com/s0h1s2/airbnb-clone/internal/db"
	"github.com/s0h1s2/airbnb-clone/internal/models"
	"github.com/s0h1s2/airbnb-clone/internal/util"
	"github.com/s0h1s2/airbnb-clone/internal/validation"
	"gorm.io/gorm"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func userRegisterRoute(ctx *gin.Context) {
	data, _ := ctx.Get("body")
	json, _ := data.(validation.CreateUserRequest)

	user := models.User{}
	result := db.Db.Where("email = ?", json.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// let's create the user
		hashedPassword := util.HashPassword(json.Password)

		newUser := models.User{Name: json.Name, Email: json.Email, Password: string(hashedPassword)}

		db.Db.Create(&newUser)

		ctx.JSON(http.StatusOK, gin.H{
			"data": newUser,
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": "User already exist",
	})
	return

}
func userAuthRoute(ctx *gin.Context) {
	data, _ := ctx.Get("body")
	json, _ := data.(validation.LoginUserRequest)

	user := models.User{} // 0x1
	result := db.Db.Where("email = ?", json.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email or password are invalid",
		})
		return
	}
	// Check if hashed password and json password match
	hashResult := util.VerifyPassword(user.Password, json.Password)
	if !hashResult {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email and password are invalid",
		})
		return
	}
	// create token
	expireTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email: json.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.Jwt.JwtSecretKey))
	if err != nil {
		log.Error().Err(err).Msg("")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	ctx.SetCookie("token", tokenString, int(expireTime.Unix()), "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"data": "User authenticated",
	})
}
