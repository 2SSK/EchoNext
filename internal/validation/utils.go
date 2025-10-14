package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct and returns a formatted error
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
