package request

import "github.com/go-playground/validator/v10"

type ClientRequest interface {
	// ValidationErrors returns the errors incurred during validation.
	ValidationErrors(errs validator.ValidationErrors) map[string]string

	// Validate handles complex and type-specific validations not covered by
	// the validator library.
	Validate(errMessages map[string]string)
}

const (
	tagRequired           = "required"
	tagEmail              = "email"
	tagUnique             = "unique"
	tagMaximum            = "max"
	tagMinimum            = "min"
	tagGreaterThan        = "gt"
	tagGreaterThanOrEqual = "gte"
	tagLessThan           = "lt"
	tagLessThanOrEqual    = "lte"
)
