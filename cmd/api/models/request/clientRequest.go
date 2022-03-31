package request

import "github.com/rhodeon/moviescreen/internal/validator"

// ClientRequest is the generic interface which all request models have to satisfy.
type ClientRequest interface {
	// Validate handles data validation in the request against specified rules,
	// and returns all errors found.
	// required holds the JSON names of the mandatory fields.
	Validate(required []string) *validator.Validator
}
