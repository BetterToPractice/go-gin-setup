package response

import (
	"errors"
	appErrors "github.com/BetterToPractice/go-gin-setup/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

type Response struct {
	Code    int         `json:"-"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type BadRequest struct {
	Req     interface{}
	Message interface{} `json:"message"`
}

type PolicyResponse struct {
	Message error `json:"message"`
}

type NotFound struct {
	Message interface{} `json:"message"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors []ValidationError

func (r Response) JSON(ctx *gin.Context) {
	if r.Code == 0 {
		r.Code = http.StatusOK
	}

	if r.Message == nil || r.Message == "" {
		r.Message = http.StatusText(r.Code)
	}

	if err, ok := r.Message.(error); ok {
		if errors.Is(err, appErrors.DatabaseInternalError) {
			r.Code = http.StatusInternalServerError
		}
		r.Message = err.Error()
	}

	ctx.JSON(r.Code, r)
}

func (r BadRequest) JSON(ctx *gin.Context) {
	resp := Response{Code: http.StatusBadRequest, Message: r.Message}

	if err, ok := r.Message.(validator.ValidationErrors); ok && err != nil {
		var validationErrors []ValidationError
		v := reflect.TypeOf(r.Req)

		for _, e := range err {
			field, _ := v.FieldByName(e.Field())
			validationErrors = append(validationErrors, ValidationError{
				Field:   field.Tag.Get("json"),
				Message: e.Tag(),
			})
		}
		resp.Data = validationErrors
		resp.Message = http.StatusText(resp.Code)
	}

	resp.JSON(ctx)
}

func (r NotFound) JSON(ctx *gin.Context) {
	resp := Response{Code: http.StatusNotFound, Message: r.Message}
	resp.JSON(ctx)
}

func (r PolicyResponse) JSON(ctx *gin.Context) {
	resp := Response{Code: http.StatusUnauthorized, Message: r.Message}
	if errors.Is(r.Message, appErrors.Forbidden) {
		resp.Code = http.StatusForbidden
	}
	resp.JSON(ctx)
}
