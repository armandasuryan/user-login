package middleware

import (
	"github.com/go-playground/validator/v10"
)

func ValidateData(data interface{}) ([]string, error) {
	validate := validator.New()

	// Validate the data structure
	err := validate.Struct(data)
	if err == nil {
		return nil, nil
	}

	// collect validation errors
	var errorFields []string
	for _, vErr := range err.(validator.ValidationErrors) {
		errorFields = append(errorFields, vErr.Field())
	}

	return errorFields, err
}
