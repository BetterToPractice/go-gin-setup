package dto

import "github.com/BetterToPractice/go-gin-setup/models"

type PostQueryParam struct {
	PaginationParam
}

type PostRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

type PostUpdateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostResponse struct {
	ID    uint             `json:"id"`
	Title string           `json:"title"`
	Body  string           `json:"body"`
	User  PostUserResponse `json:"user"`
}

type PostUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type PostPaginationResponse struct {
	List       []PostResponse `json:"list"`
	Pagination *Pagination    `json:"pagination"`
}

func (r *PostResponse) Serializer(post *models.Post) {
	r.ID = post.ID
	r.Title = post.Title
	r.Body = post.Body
	r.User = PostUserResponse{
		ID:       post.UserID,
		Username: post.User.Username,
	}
}

func (r *PostPaginationResponse) Serializer(posts *models.Posts) {
	var list []PostResponse
	for _, post := range *posts {
		p := PostResponse{}
		p.Serializer(&post)
		list = append(list, p)
	}
	r.List = list
}
