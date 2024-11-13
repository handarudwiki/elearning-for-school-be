package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequest(req interface{}) []ValidationError {
	var errors []ValidationError
	err := validate.Struct(req)
	fmt.Println(err)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			// Store each validation error with the field as the key
			errors = append(errors, ValidationError{
				Field:   e.Field(),
				Message: GetErrorMessage(e),
			})
		}
	}
	return errors
}

func GetErrorMessage(e validator.FieldError) (message string) {
	switch e.Tag() {
	case "required":
		message = e.Field() + " is required"
	case "email":
		message = e.Field() + " is not valid email"
	case "min":
		message = e.Field() + " min " + e.Param()
	case "max":
		message = e.Field() + " max " + e.Param()
	default:
		message = e.Field() + " is not valid"
	}

	return
}
