package lib

import (
	"fmt"
	_ "net/http"
	"strings"

	"github.com/go-playground/validator"
)

type (
	ValidationError struct {
		Namespace string
		Field     string
		Tag       string
		Message   string
	}

	ValidationErrors []ValidationError

	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (ve ValidationErrors) Error() string {
	sErrs := make([]string, len(ve))
	for i, validationError := range ve {
		sErrs[i] = validationError.Message
	}
	return strings.Join(sErrs, "\n")
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errors ValidationErrors
		for _, e := range validationErrors {
			errors = append(errors, ValidationError{
				Namespace: e.Namespace(),
				Field:     e.Field(),
				Tag:       e.Tag(),
				Message:   fmt.Sprintf("'%s' failed on the '%s' condition", e.Field(), e.Tag()),
			})
		}
		return errors
	}
	return nil
}
