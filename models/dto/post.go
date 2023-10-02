package dto

type PostRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type PostUpdateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostResponse struct {
	Title string
	Body  string
}
