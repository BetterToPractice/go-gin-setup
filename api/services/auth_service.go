package services

import (
	"github.com/BetterToPractice/go-gin-setup/models"
	"github.com/BetterToPractice/go-gin-setup/models/dto"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct{}

func NewAuthService() AuthService {
	return AuthService{}
}

func (c AuthService) GenerateToken(user *models.User) (string, error) {
	now := time.Now()
	claims := &dto.JwtClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte("example"))
}
