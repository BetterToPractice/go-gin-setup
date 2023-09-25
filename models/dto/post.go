package dto

type PostRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type PostResponse struct {
	Title string
	Body  string
}
