package dto

import "github.com/BetterToPractice/go-gin-setup/models"

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Username string
	Email    string
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Access string `json:"access"`
}

func (r *RegisterResponse) Serializer(user *models.User) {
	r.Username = user.Username
	r.Email = user.Email
}

func (r *LoginResponse) Serializer(access string) {
	r.Access = access
}
