package dto

import "github.com/golang-jwt/jwt/v5"

type JwtResponse struct {
	Access string `form:"access" json:"access" binding:"required"`
}

type JwtClaims struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}
