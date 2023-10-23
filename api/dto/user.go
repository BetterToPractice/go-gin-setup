package dto

import "github.com/BetterToPractice/go-gin-setup/models"

type UserQueryParam struct {
	PaginationParam
}

type ProfileResponse struct {
	PhoneNumber string `json:"phone_number,omitempty"`
	Gender      string `json:"gender,omitempty"`
}

type UserResponse struct {
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Profile  ProfileResponse `json:"profile"`
}

type UserPaginationResponse struct {
	List       []UserResponse `json:"list"`
	Pagination *Pagination    `json:"pagination"`
}

func (r *UserPaginationResponse) Serializer(users *models.Users) {
	var list []UserResponse
	for _, user := range *users {
		u := UserResponse{}
		u.Serializer(&user)
		list = append(list, u)
	}
	r.List = list
}

func (r *UserResponse) Serializer(user *models.User) {
	r.Username = user.Username
	r.Email = user.Email
	r.Profile = ProfileResponse{
		PhoneNumber: user.Profile.PhoneNumber,
		Gender:      user.Profile.Gender,
	}
}
