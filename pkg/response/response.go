package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"-"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) JSON(ctx *gin.Context) {
	if r.Message == nil || r.Message == "" {
		r.Message = http.StatusText(r.Code)
	}
	if r.Code == 0 {
		r.Code = http.StatusOK
	}

	err, ok := r.Message.(error)
	if ok {
		r.Message = err.Error()
	}

	ctx.JSON(r.Code, r)
}
