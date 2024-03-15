package users

import (
	"time"

	"github.com/s0h1s2/airbnb-clone/internal/util"
)

type UserModel struct {
	ID        uint
	Name      string
	Email     string
	Password  string `json:"omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}

func (user *UserModel) HashPassword() {
	hashedPassword := util.HashPassword(user.Password)
	user.Password = hashedPassword
}
