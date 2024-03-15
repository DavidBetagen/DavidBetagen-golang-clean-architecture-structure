package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

type XValidator struct{}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

var validate = validator.New()

func ValidateErrors(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = strcase.ToSnake(err.StructField()) // Export struct field name
			elem.Tag = err.Tag()                                  // Export struct tag
			elem.Value = err.Value()                              // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func Validate(data interface{}) []string {
	if errs := ValidateErrors(data); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"%s field is %s",
				err.FailedField,
				err.Tag,
			))
		}

		return errMsgs
	}

	return nil
}
