package services

import (
	"errors"
	"fmt"
	"github.com/BetterToPractice/go-gin-setup/api/mails"
	"github.com/BetterToPractice/go-gin-setup/constants"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
	"github.com/BetterToPractice/go-gin-setup/models/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	config      lib.Config
	opts        *options
	userService UserService
	authMail    mails.AuthMail
	db          lib.Database
}

type options struct {
	issuer        string
	signingMethod jwt.SigningMethod
	signingKey    interface{}
	keyfunc       jwt.Keyfunc
	expired       int
}

func NewAuthService(config lib.Config, userService UserService, authMail mails.AuthMail, db lib.Database) AuthService {
	signingKey := fmt.Sprintf("jwt:%s", config.Name)
	opts := &options{
		issuer:        config.Name,
		expired:       config.Auth.TokenExpired,
		signingMethod: jwt.SigningMethodHS512,
		signingKey:    []byte(signingKey),
		keyfunc: func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid Token")
			}
			return []byte(signingKey), nil
		},
	}

	return AuthService{
		config:      config,
		opts:        opts,
		userService: userService,
		authMail:    authMail,
		db:          db,
	}
}

func (s AuthService) GenerateToken(user *models.User) (string, error) {
	now := time.Now()
	claims := &dto.JwtClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(s.config.Auth.TokenExpired) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(s.opts.signingKey)
}

func (s AuthService) ParseToken(tokenString string) (*dto.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.JwtClaims{}, s.opts.keyfunc)
	if err != nil {
		return nil, err
	}

	if token != nil {
		if claims, ok := token.Claims.(*dto.JwtClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("invalid token")
}

func (s AuthService) Register(username, password, email string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: models.HashPassword(password),
		Email:    email,
	}
	if err := s.db.ORM.Create(&user).Error; err != nil {
		return nil, err
	}

	s.authMail.Register(user)

	return user, nil
}

func (s AuthService) Login(login *dto.Login) (*dto.LoginResponse, error) {
	user, err := s.userService.Verify(login.Username, login.Password)
	if err != nil {
		return nil, err
	}

	access, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Access: access}, nil
}

func (s AuthService) Authenticate(ctx *gin.Context) (*models.User, error) {
	claims, _ := ctx.Get(constants.CurrentUser)
	jwtClaims, _ := claims.(*dto.JwtClaims)
	if jwtClaims == nil {
		return nil, errors.New("unauthorized")
	}

	if user, err := s.userService.GetByUsername(jwtClaims.Username); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
