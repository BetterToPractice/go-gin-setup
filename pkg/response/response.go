package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	if err, ok := r.Message.(validator.ValidationErrors); ok {
		r.Message = extractValidationErrors(err)
	}
	err, ok := r.Message.(error)
	if ok {
		r.Message = err.Error()
	}

	ctx.JSON(r.Code, r)
}

type ValidationError struct {
	Field   string // Field name that failed validation
	Message string // Error message describing the validation failure
}

type ValidationErrors []ValidationError

func extractValidationErrors(err validator.ValidationErrors) ValidationErrors {
	validationErrors := ValidationErrors{}
	for _, fieldError := range err {
		validationErrors = append(validationErrors, ValidationError{
			Field:   fieldError.StructField(),
			Message: fieldError.Tag(),
		})
	}
	return validationErrors
}
