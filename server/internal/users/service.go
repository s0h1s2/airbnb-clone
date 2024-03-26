package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/s0h1s2/airbnb-clone/config"
	"github.com/s0h1s2/airbnb-clone/internal/common"
	"github.com/s0h1s2/airbnb-clone/internal/db"
	"github.com/s0h1s2/airbnb-clone/internal/util"
	"gorm.io/gorm"
)

var (
	userAlredyExist    = errors.New("Email already exist")
	invalidCredentials = errors.New("Invalid credentials")
)

func userRegisterRoute(ctx *gin.Context) {
	var json createUserRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{Errors: err, StatusCode: http.StatusBadRequest})

		return
	}
	user := User{}
	result := db.Db.Where("email = ?", json.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// let's create the user
		newUser := User{Name: json.Name, Email: json.Email, Password: json.Password}

		newUser.HashPassword()

		db.Db.Create(&newUser)

		ctx.JSON(http.StatusCreated, common.OkApiResponse{StatusCode: http.StatusCreated, Data: createUserResponse{Email: newUser.Email, Name: newUser.Name}})
		return
	}
	ctx.JSON(http.StatusBadRequest, common.ErrorApiResponse{StatusCode: http.StatusBadRequest, Errors: createUserResponse{Email: userAlredyExist.Error()}})

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
	user := User{}
	result := db.Db.Where("email = ?", json.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusUnauthorized, common.ErrorApiResponse{Errors: invalidCredentials.Error(), StatusCode: http.StatusUnauthorized})
		return
	}
	// Check if hashed password and json password match
	hashResult := util.VerifyPassword(user.Password, json.Password)
	if !hashResult {
		ctx.JSON(http.StatusUnauthorized, common.ErrorApiResponse{Errors: invalidCredentials.Error(), StatusCode: http.StatusUnauthorized})
		return
	}
	// create token
	expireTime := time.Now().Add(1 * time.Hour)
	claims := &common.UserClaims{
		Email: user.Email,
		Name:  user.Name,
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
	ctx.JSON(http.StatusOK, common.OkApiResponse{Data: loginUserResponse{Email: user.Email, Name: user.Name}, StatusCode: http.StatusOK})
}
func userInfoRoute(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	u := user.(common.UserClaims)
	ctx.JSON(http.StatusOK, common.OkApiResponse{
		Data: userInfoResponse{
			User: userInfo{
				Name:  u.Name,
				Email: u.Email,
			},
		},
		StatusCode: http.StatusOK,
	})

}
func userLogoutRoute(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{})
}
