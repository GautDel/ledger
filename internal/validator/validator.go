package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Validate(s interface{}) error {

	validate = validator.New()

	err := validate.Struct(s)
	if err != nil {
		return err
	}

    return nil
}
