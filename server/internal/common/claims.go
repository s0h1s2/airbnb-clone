package common

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Uid   uint   `json:"uid"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}
