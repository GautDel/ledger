package utils

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Param   string
	Message string
}

func msgForTag(fe validator.FieldError) string {

	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Minimum of " + fe.Param() + " characters"
	case "max":
		return "Max of " + fe.Param() + " characters"
	case "email":
		return "Invalid email"
	}

	return fe.Error()
}

func ErrorHandler(err error) []ApiError {

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), msgForTag(fe)}
		}
		log.Println(out)
		return out
	}
    return nil
}
