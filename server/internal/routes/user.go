package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s0h1s2/airbnb-clone/internal/db"
	"github.com/s0h1s2/airbnb-clone/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type createUser struct {
	Email    string `json:"email" binding:"required,email" `
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func userRegisterRoute(ctx *gin.Context) {
	var json createUser
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := model.User{}
	result := db.Db.Where("email = ?", json.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// let's create the user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.Error(err)
			return
		}
		newUser := model.User{Name: json.Name, Email: json.Email, Password: string(hashedPassword)}
		result := db.Db.Create(&newUser)
		if result.Error != nil {
			log.Fatal(result.Error)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": newUser,
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": "User already exist",
	})

}
