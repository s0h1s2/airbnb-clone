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
	"gorm.io/gorm"
)

var (
	userAlredyExist    = errors.New("User already exist")
	invalidCredentials = errors.New("Invalid credentials")
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email" `
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=5,max=72"`
}
type createUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email" `
	Password string `json:"password" binding:"required"`
}
type loginUserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func userRegisterRoute(ctx *gin.Context) {
	var json createUserRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, errorApiResponse{Errors: err, StatusCode: http.StatusBadRequest})

		return
	}
	user := models.User{}
	result := db.Db.Where("email = ?", json.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// let's create the user
		hashedPassword := util.HashPassword(json.Password)

		newUser := models.User{Name: json.Name, Email: json.Email, Password: string(hashedPassword)}

		db.Db.Create(&newUser)

		ctx.JSON(http.StatusCreated, okApiResponse{StatusCode: http.StatusCreated, Data: createUserResponse{Email: newUser.Email, Name: newUser.Name}})
		return
	}
	ctx.JSON(http.StatusBadRequest, errorApiResponse{StatusCode: http.StatusBadRequest, Errors: createUserResponse{Email: userAlredyExist.Error()}})

	return

}
func userAuthRoute(ctx *gin.Context) {
	var json loginUserRequest
	if err := ctx.BindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{}
	result := db.Db.Where("email = ?", json.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusUnauthorized, errorApiResponse{Errors: invalidCredentials.Error(), StatusCode: http.StatusUnauthorized})
		return
	}
	// Check if hashed password and json password match
	hashResult := util.VerifyPassword(user.Password, json.Password)
	if !hashResult {
		ctx.JSON(http.StatusUnauthorized, errorApiResponse{Errors: invalidCredentials.Error(), StatusCode: http.StatusUnauthorized})
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
	ctx.JSON(http.StatusOK, okApiResponse{Data: loginUserResponse{Email: user.Email, Name: user.Name}, StatusCode: http.StatusOK})
}
