package common

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}
