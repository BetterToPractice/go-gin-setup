package services

import (
	"errors"
	"fmt"
	"github.com/BetterToPractice/go-gin-setup/api/dto"
	"github.com/BetterToPractice/go-gin-setup/api/mails"
	"github.com/BetterToPractice/go-gin-setup/constants"
	appErrors "github.com/BetterToPractice/go-gin-setup/errors"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
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
	signingKey := fmt.Sprintf("jwt:%s", config.Secret)
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

func (s AuthService) GenerateToken(user *models.User) (*dto.JWTResponse, error) {
	now := time.Now()
	resp := &dto.JWTResponse{}

	claims := &dto.JWTClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(s.config.Auth.TokenExpired) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	access, err := token.SignedString(s.opts.signingKey)
	if err != nil {
		return nil, err
	}

	resp.Serializer(access)
	return resp, nil
}

func (s AuthService) ParseToken(tokenString string) (*dto.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.JWTClaims{}, s.opts.keyfunc)
	if err != nil {
		return nil, err
	}

	if token != nil {
		if claims, ok := token.Claims.(*dto.JWTClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("invalid token")
}

func (s AuthService) Register(register *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	user, err := s.userService.Register(register)
	if err != nil {
		return nil, err
	}

	s.authMail.Register(user)

	resp := &dto.RegisterResponse{}
	resp.Serializer(user)

	return resp, nil
}

func (s AuthService) Login(login *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userService.Verify(login.Username, login.Password)
	if err != nil {
		return nil, appErrors.Unauthorized
	}

	access, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	resp := &dto.LoginResponse{}
	resp.Serializer(access.Access)

	return resp, nil
}

func (s AuthService) Authenticate(ctx *gin.Context) (*models.User, error) {
	claims, _ := ctx.Get(constants.CurrentUser)
	jwtClaims, _ := claims.(*dto.JWTClaims)
	if jwtClaims == nil {
		return nil, errors.New("unauthorized")
	}

	if user, err := s.userService.GetByUsername(jwtClaims.Username); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
